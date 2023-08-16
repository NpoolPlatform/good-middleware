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
		return fmt.Errorf("invalid stock")
	}

	locked := info.Locked
	if h.Locked != nil {
		locked = h.Locked.Sub(locked)
	}
	if locked.Cmp(decimal.NewFromInt(0)) < 0 {
		return fmt.Errorf("invalid locked")
	}
	inService := info.InService
	sold := info.Sold
	if h.InService != nil {
		inService = h.InService.Sub(inService)
		sold = h.WaitStart.Sub(sold)
	}
	if inService.Cmp(decimal.NewFromInt(0)) < 0 {
		return fmt.Errorf("invalid inservice")
	}
	waitStart := info.WaitStart
	if h.WaitStart != nil {
		waitStart = h.WaitStart.Sub(waitStart)
		sold = h.WaitStart.Sub(sold)
	}
	if waitStart.Cmp(decimal.NewFromInt(0)) < 0 {
		return fmt.Errorf("invalid waitstart")
	}
	if sold.Cmp(decimal.NewFromInt(0)) < 0 {
		return fmt.Errorf("invalid sold")
	}
	appLocked := info.AppLocked
	if h.AppLocked != nil {
		appLocked = h.AppLocked.Sub(appLocked)
	}
	if appLocked.Cmp(decimal.NewFromInt(0)) < 0 {
		return fmt.Errorf("invalid applocked")
	}

	if locked.Add(inService).
		Add(waitStart).
		Add(appLocked).
		Cmp(info.Total) > 0 {
		return fmt.Errorf("invalid stock")
	}

	if _, err := stockcrud.UpdateSet(
		tx.Stock.UpdateOneID(*h.ID),
		&stockcrud.Req{
			Locked:    &locked,
			InService: &inService,
			WaitStart: &waitStart,
			AppLocked: &appLocked,
			Sold:      &sold,
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
		return nil
	})
	if err != nil {
		return nil, err
	}

	return h.GetStock(ctx)
}
