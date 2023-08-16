package stock

import (
	"context"
	"fmt"

	stockcrud "github.com/NpoolPlatform/good-middleware/pkg/crud/good/stock"
	"github.com/NpoolPlatform/good-middleware/pkg/db"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent"
	entstock "github.com/NpoolPlatform/good-middleware/pkg/db/ent/stock"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/good/stock"
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

	locked := info.Locked
	if h.Locked != nil {
		locked = h.Locked.Add(locked)
	}
	inService := info.InService
	if h.InService != nil {
		inService = h.InService.Add(inService)
	}
	waitStart := info.WaitStart
	sold := info.Sold
	if h.WaitStart != nil {
		waitStart = h.WaitStart.Add(waitStart)
		sold = h.WaitStart.Add(sold)
	}
	appLocked := info.AppLocked
	if h.AppLocked != nil {
		appLocked = h.AppLocked.Add(appLocked)
	}

	if locked.Add(inService).
		Add(waitStart).
		Add(appLocked).
		Cmp(info.Total) > 0 {
		return fmt.Errorf("stock exhausted")
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
