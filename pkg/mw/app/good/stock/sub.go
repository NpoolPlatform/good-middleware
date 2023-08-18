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

type subHandler struct {
	*Handler
}

func (h *subHandler) subStock(ctx context.Context, tx *ent.Tx) error {
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
		return fmt.Errorf("stock not found")
	}

	locked := info.Locked
	if h.Locked != nil {
		locked = locked.Sub(*h.Locked)
	}
	if locked.Cmp(decimal.NewFromInt(0)) < 0 {
		return fmt.Errorf("invalid locked")
	}
	inService := info.InService
	sold := info.Sold
	if h.InService != nil {
		inService = inService.Sub(*h.InService)
		sold = sold.Sub(*h.InService)
	}
	if inService.Cmp(decimal.NewFromInt(0)) < 0 {
		return fmt.Errorf("invalid inservice")
	}
	waitStart := info.WaitStart
	if h.WaitStart != nil {
		waitStart = waitStart.Sub(*h.WaitStart)
		sold = sold.Sub(*h.WaitStart)
	}
	if waitStart.Cmp(decimal.NewFromInt(0)) < 0 {
		return fmt.Errorf("invalid waitstart")
	}
	if sold.Cmp(decimal.NewFromInt(0)) < 0 {
		return fmt.Errorf("invalid sold")
	}

	if locked.Add(inService).
		Add(waitStart).
		Add(info.AppReserved).
		Cmp(info.Total) > 0 {
		return fmt.Errorf("invalid stock")
	}

	if _, err := stockcrud.UpdateSet(
		tx.Stock.UpdateOneID(info.ID),
		&stockcrud.Req{
			Locked:    &locked,
			InService: &inService,
			WaitStart: &waitStart,
			Sold:      &sold,
		},
	).Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *subHandler) subAppStock(ctx context.Context, tx *ent.Tx) error {
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
		return fmt.Errorf("stock not found")
	}
	h.GoodID = &info.GoodID
	spotQuantity := info.SpotQuantity

	locked := info.Locked
	if h.Locked != nil {
		locked = locked.Sub(*h.Locked)
		spotQuantity = spotQuantity.Add(*h.Locked)
	}
	if locked.Cmp(decimal.NewFromInt(0)) < 0 {
		return fmt.Errorf("invalid locked")
	}
	inService := info.InService
	sold := info.Sold
	if h.InService != nil {
		inService = inService.Sub(*h.InService)
		sold = sold.Sub(*h.InService)
		spotQuantity = spotQuantity.Add(*h.InService)
	}
	if inService.Cmp(decimal.NewFromInt(0)) < 0 {
		return fmt.Errorf("invalid inservice")
	}
	waitStart := info.WaitStart
	if h.WaitStart != nil {
		waitStart = waitStart.Sub(*h.WaitStart)
		sold = sold.Sub(*h.WaitStart)
		spotQuantity = spotQuantity.Add(*h.WaitStart)
	}
	if waitStart.Cmp(decimal.NewFromInt(0)) < 0 {
		return fmt.Errorf("invalid waitstart")
	}
	if sold.Cmp(decimal.NewFromInt(0)) < 0 {
		return fmt.Errorf("invalid sold")
	}
	if spotQuantity.Cmp(info.Reserved) > 0 {
		return fmt.Errorf("invalid spotquantity")
	}

	if locked.Add(inService).
		Add(waitStart).
		Cmp(info.SpotQuantity) > 0 {
		spotQuantity = decimal.NewFromInt(0)
	}

	if _, err := appstockcrud.UpdateSet(
		tx.AppStock.UpdateOneID(info.ID),
		&appstockcrud.Req{
			SpotQuantity: &spotQuantity,
			Locked:       &locked,
			InService:    &inService,
			WaitStart:    &waitStart,
			Sold:         &sold,
		},
	).Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *Handler) SubStock(ctx context.Context) (*npool.Stock, error) {
	handler := &subHandler{
		Handler: h,
	}

	err := db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		if err := handler.subStock(ctx, tx); err != nil {
			return err
		}
		if err := handler.subAppStock(ctx, tx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return h.GetStock(ctx)
}
