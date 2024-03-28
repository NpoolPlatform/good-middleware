//nolint:dupl
package appstock

import (
	"context"
	"fmt"

	appstockcrud "github.com/NpoolPlatform/good-middleware/pkg/crud/app/good/stock"
	stockcrud "github.com/NpoolPlatform/good-middleware/pkg/crud/good/stock"
	"github.com/NpoolPlatform/good-middleware/pkg/db"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent"
	types "github.com/NpoolPlatform/message/npool/basetypes/good/v1"

	"github.com/shopspring/decimal"
)

type waitStartHandler struct {
	*stockAppGoodQuery
	lockOp *lockopHandler
}

func (h *waitStartHandler) waitStartStock(ctx context.Context, lock *ent.AppStockLock, tx *ent.Tx) error {
	locked := h.stock.Locked
	sold := h.stock.Sold
	waitStart := h.stock.WaitStart

	waitStart = lock.Units.Add(waitStart)
	locked = locked.Sub(lock.Units)
	sold = lock.Units.Add(sold)

	if locked.Cmp(decimal.NewFromInt(0)) < 0 {
		return fmt.Errorf("invalid stock")
	}
	if waitStart.Cmp(decimal.NewFromInt(0)) < 0 {
		return fmt.Errorf("invalid stock")
	}

	if locked.Add(h.stock.InService).
		Add(h.stock.WaitStart).
		Add(h.stock.AppReserved).
		Add(h.stock.SpotQuantity).
		Cmp(h.stock.Total) > 0 {
		return fmt.Errorf("stock exhausted")
	}

	if _, err := stockcrud.UpdateSet(
		tx.Stock.UpdateOneID(h.stock.ID),
		&stockcrud.Req{
			Locked:    &locked,
			WaitStart: &waitStart,
			Sold:      &sold,
		},
	).Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *waitStartHandler) waitStartAppStock(ctx context.Context, lock *ent.AppStockLock, tx *ent.Tx) error {
	h.ID = &h.appGoodStock.ID
	locked := h.appGoodStock.Locked
	sold := h.appGoodStock.Sold
	waitStart := h.appGoodStock.WaitStart

	waitStart = lock.Units.Add(waitStart)
	locked = locked.Sub(lock.Units)
	sold = lock.Units.Add(sold)

	if locked.Cmp(decimal.NewFromInt(0)) < 0 {
		return fmt.Errorf("invalid stock")
	}
	if waitStart.Cmp(decimal.NewFromInt(0)) < 0 {
		return fmt.Errorf("invalid stock")
	}

	if _, err := appstockcrud.UpdateSet(
		tx.AppStock.UpdateOneID(h.appGoodStock.ID),
		&appstockcrud.Req{
			Locked:    &locked,
			WaitStart: &waitStart,
			Sold:      &sold,
		},
	).Save(ctx); err != nil {
		return err
	}

	return nil
}

func (h *Handler) WaitStartStock(ctx context.Context) error {
	handler := &waitStartHandler{
		stockAppGoodQuery: &stockAppGoodQuery{
			Handler: h,
		},
		lockOp: &lockopHandler{
			Handler: h,
			state:   types.AppStockLockState_AppStockWaitStart.Enum(),
		},
	}

	if err := handler.lockOp.getLocks(ctx); err != nil {
		return err
	}
	if err := handler.requireStockAppGood(ctx); err != nil {
		return err
	}

	return db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		for _, lock := range handler.lockOp.locks {
			if err := handler.waitStartAppStock(ctx, lock, tx); err != nil {
				return err
			}
			if err := handler.waitStartStock(ctx, lock, tx); err != nil {
				return err
			}
		}
		if err := handler.lockOp.updateLocks(ctx, tx); err != nil {
			return err
		}
		return nil
	})
}
