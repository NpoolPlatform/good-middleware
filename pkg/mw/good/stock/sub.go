//nolint:dupl
package stock

import (
	"context"
	"fmt"

	stockcrud "github.com/NpoolPlatform/good-middleware/pkg/crud/good/stock"
	"github.com/NpoolPlatform/good-middleware/pkg/db"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent"
	entstock "github.com/NpoolPlatform/good-middleware/pkg/db/ent/stock"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/good/stock"

	"github.com/shopspring/decimal"
)

type subHandler struct {
	*Handler
}

//nolint:gocyclo
func (h *subHandler) subStock(ctx context.Context, tx *ent.Tx) error {
	info, err := tx.
		Stock.
		Query().
		Where(
			entstock.ID(*h.ID),
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

	spotQuantity := info.SpotQuantity
	locked := info.Locked
	inService := info.InService
	sold := info.Sold
	waitStart := info.WaitStart
	appReserved := info.AppReserved

	if h.Locked != nil {
		locked = locked.Sub(*h.Locked)
		spotQuantity = h.Locked.Add(spotQuantity)
	}
	if h.InService != nil {
		inService = inService.Sub(*h.InService)
		sold = sold.Sub(*h.InService)
		spotQuantity = h.InService.Add(spotQuantity)
	}
	if h.WaitStart != nil {
		waitStart = waitStart.Sub(*h.WaitStart)
		sold = sold.Sub(*h.WaitStart)
		spotQuantity = h.WaitStart.Add(spotQuantity)
	}
	if h.AppReserved != nil {
		appReserved = appReserved.Sub(*h.AppReserved)
		spotQuantity = h.AppReserved.Add(spotQuantity)
	}

	if spotQuantity.Cmp(decimal.NewFromInt(0)) < 0 {
		return fmt.Errorf("invalid stock")
	}
	if spotQuantity.Cmp(info.Total) > 0 {
		return fmt.Errorf("invalid stock")
	}
	if inService.Cmp(decimal.NewFromInt(0)) < 0 {
		return fmt.Errorf("invalid inservice")
	}
	if waitStart.Cmp(decimal.NewFromInt(0)) < 0 {
		return fmt.Errorf("invalid waitstart")
	}
	if sold.Cmp(decimal.NewFromInt(0)) < 0 {
		return fmt.Errorf("invalid sold")
	}
	if appReserved.Cmp(decimal.NewFromInt(0)) < 0 {
		return fmt.Errorf("invalid appreserved")
	}
	if locked.Cmp(decimal.NewFromInt(0)) < 0 {
		return fmt.Errorf("invalid locked")
	}

	if locked.Add(inService).
		Add(waitStart).
		Add(appReserved).
		Cmp(spotQuantity) > 0 {
		return fmt.Errorf("invalid stock")
	}

	if _, err := stockcrud.UpdateSet(
		tx.Stock.UpdateOneID(*h.ID),
		&stockcrud.Req{
			SpotQuantity: &spotQuantity,
			Locked:       &locked,
			InService:    &inService,
			WaitStart:    &waitStart,
			AppReserved:  &appReserved,
			Sold:         &sold,
		},
	).Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *Handler) SubStock(ctx context.Context) (*npool.Stock, error) {
	if err := h.validate(); err != nil {
		return nil, err
	}

	handler := &subHandler{
		Handler: h,
	}

	err := db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		if err := handler.subStock(ctx, tx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return h.GetStock(ctx)
}
