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

type addHandler struct {
	*Handler
}

func (h *addHandler) addStock(ctx context.Context, tx *ent.Tx) error {
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
		return fmt.Errorf("invalid stock")
	}

	spotQuantity := info.SpotQuantity
	locked := info.Locked
	if h.Locked != nil {
		locked = h.Locked.Add(locked)
		spotQuantity = spotQuantity.Sub(*h.Locked)
	}
	inService := info.InService
	sold := info.Sold
	if h.InService != nil {
		inService = h.InService.Add(inService)
		sold = h.InService.Add(sold)
		spotQuantity = spotQuantity.Sub(*h.InService)
	}
	waitStart := info.WaitStart
	if h.WaitStart != nil {
		waitStart = h.WaitStart.Add(waitStart)
		sold = h.WaitStart.Add(sold)
		spotQuantity = spotQuantity.Sub(*h.WaitStart)
	}
	appReserved := info.AppReserved
	if h.AppReserved != nil {
		appReserved = h.AppReserved.Add(appReserved)
		spotQuantity = spotQuantity.Sub(*h.AppReserved)
	}

	if spotQuantity.Cmp(decimal.NewFromInt(0)) < 0 {
		return fmt.Errorf("stock exhausted")
	}
	if locked.Add(inService).
		Add(waitStart).
		Add(appReserved).
		Cmp(info.Total) > 0 {
		return fmt.Errorf("stock exhausted")
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

func (h *Handler) AddStock(ctx context.Context) (*npool.Stock, error) {
	handler := &addHandler{
		Handler: h,
	}

	err := db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		if err := handler.addStock(ctx, tx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return h.GetStock(ctx)
}
