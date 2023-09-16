//nolint:dupl
package appstock

import (
	"context"
	"fmt"

	appstockcrud "github.com/NpoolPlatform/good-middleware/pkg/crud/app/good/stock"
	stockcrud "github.com/NpoolPlatform/good-middleware/pkg/crud/good/stock"
	"github.com/NpoolPlatform/good-middleware/pkg/db"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent"
	entappstock "github.com/NpoolPlatform/good-middleware/pkg/db/ent/appstock"
	entstock "github.com/NpoolPlatform/good-middleware/pkg/db/ent/stock"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/app/good/stock"

	"github.com/shopspring/decimal"
)

type reserveHandler struct {
	*Handler
}

func (h *reserveHandler) reserveStock(ctx context.Context, tx *ent.Tx) error { //nolint:gocyclo
	info, err := tx.
		Stock.
		Query().
		Where(
			entstock.GoodID(*h.GoodID),
			entstock.DeletedAt(0),
		).
		ForUpdate().
		Only(ctx)
	if err != nil {
		return err
	}
	if info == nil {
		return fmt.Errorf("invalid stock")
	}

	spotQuantity := info.SpotQuantity
	appReserved := info.AppReserved

	appReserved = h.Reserved.Add(appReserved)
	spotQuantity = spotQuantity.Sub(*h.Reserved)

	if spotQuantity.Cmp(decimal.NewFromInt(0)) < 0 {
		return fmt.Errorf("invalid stock")
	}
	if spotQuantity.Cmp(info.Total) > 0 {
		return fmt.Errorf("invalid stock")
	}

	if appReserved.Add(info.Locked).
		Add(info.WaitStart).
		Add(info.InService).
		Add(spotQuantity).
		Cmp(info.Total) > 0 {
		return fmt.Errorf("stock exhausted")
	}

	if _, err := stockcrud.UpdateSet(
		tx.Stock.UpdateOneID(info.ID),
		&stockcrud.Req{
			SpotQuantity: &spotQuantity,
			AppReserved:  &appReserved,
		},
	).Save(ctx); err != nil {
		return err
	}
	return nil
}

//nolint:gocyclo,funlen
func (h *reserveHandler) reserveAppStock(ctx context.Context, tx *ent.Tx) error {
	info, err := tx.
		AppStock.
		Query().
		Where(
			entappstock.ID(*h.ID),
			entappstock.DeletedAt(0),
		).
		ForUpdate().
		Only(ctx)
	if err != nil {
		return err
	}
	if info == nil {
		return fmt.Errorf("invalid appstock")
	}

	h.GoodID = &info.GoodID
	spotQuantity := info.SpotQuantity
	reserved := info.Reserved

	spotQuantity = h.Reserved.Add(spotQuantity)
	reserved = h.Reserved.Add(reserved)

	if spotQuantity.Cmp(info.Reserved) > 0 {
		return fmt.Errorf("invalid stock")
	}

	if _, err := appstockcrud.UpdateSet(
		tx.AppStock.UpdateOneID(info.ID),
		&appstockcrud.Req{
			SpotQuantity: &spotQuantity,
			Reserved:     &reserved,
		},
	).Save(ctx); err != nil {
		return err
	}

	return nil
}

func (h *Handler) ReserveStock(ctx context.Context) (*npool.Stock, error) {
	handler := &reserveHandler{
		Handler: h,
	}
	err := db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		if err := handler.reserveAppStock(ctx, tx); err != nil {
			return err
		}
		if err := handler.reserveStock(ctx, tx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return h.GetStock(ctx)
}
