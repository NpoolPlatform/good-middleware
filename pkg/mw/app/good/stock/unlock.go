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

type unlockHandler struct {
	*stockAppGoodQuery
	lockOp *lockopHandler
}

func (h *unlockHandler) unlockStock(ctx context.Context, lock *ent.AppStockLock, tx *ent.Tx) error {
	spotQuantity := h.stock.SpotQuantity
	locked := h.stock.Locked
	appReserved := h.stock.AppReserved

	locked = locked.Sub(lock.Units)
	platformLocked := lock.Units.Sub(lock.AppSpotUnits)
	appReserved = lock.AppSpotUnits.Add(appReserved)
	if platformLocked.Cmp(decimal.NewFromInt(0)) < 0 {
		return fmt.Errorf("invalid appspotlocked")
	}
	spotQuantity = platformLocked.Add(spotQuantity)
	if locked.Cmp(decimal.NewFromInt(0)) < 0 {
		return fmt.Errorf("invalid locked")
	}
	if appReserved.Cmp(decimal.NewFromInt(0)) < 0 {
		return fmt.Errorf("invalid appreserved")
	}
	if spotQuantity.Cmp(decimal.NewFromInt(0)) < 0 {
		return fmt.Errorf("invalid stock")
	}
	if spotQuantity.Cmp(h.stock.Total) > 0 {
		return fmt.Errorf("invalid stock")
	}

	if locked.Add(h.stock.InService).
		Add(h.stock.WaitStart).
		Add(appReserved).
		Add(spotQuantity).
		Cmp(h.stock.Total) > 0 {
		return fmt.Errorf("invalid stock")
	}

	if _, err := stockcrud.UpdateSet(
		tx.Stock.UpdateOneID(h.stock.ID),
		&stockcrud.Req{
			SpotQuantity: &spotQuantity,
			Locked:       &locked,
			AppReserved:  &appReserved,
		},
	).Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *unlockHandler) unlockAppStock(ctx context.Context, lock *ent.AppStockLock, tx *ent.Tx) error {
	locked := h.appGoodStock.Locked
	spotQuantity := h.appGoodStock.SpotQuantity

	locked = locked.Sub(lock.Units)
	spotQuantity = lock.AppSpotUnits.Add(spotQuantity)
	if locked.Cmp(decimal.NewFromInt(0)) < 0 {
		return fmt.Errorf("invalid locked")
	}
	if spotQuantity.Cmp(h.appGoodStock.Reserved) > 0 {
		return fmt.Errorf("invalid spotquantity")
	}
	if spotQuantity.Cmp(decimal.NewFromInt(0)) < 0 {
		spotQuantity = decimal.NewFromInt(0)
	}

	if _, err := appstockcrud.UpdateSet(
		tx.AppStock.UpdateOneID(h.appGoodStock.ID),
		&appstockcrud.Req{
			SpotQuantity: &spotQuantity,
			Locked:       &locked,
		},
	).Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *Handler) UnlockStock(ctx context.Context) error {
	handler := &unlockHandler{
		stockAppGoodQuery: &stockAppGoodQuery{
			Handler: h,
		},
		lockOp: &lockopHandler{
			Handler: h,
			state:   types.AppStockLockState_AppStockRollback.Enum(),
		},
	}

	if err := handler.lockOp.getLocks(ctx); err != nil {
		if ent.IsNotFound(err) && h.Rollback != nil && *h.Rollback {
			return nil
		}
		return err
	}
	if h.Rollback == nil || !*h.Rollback {
		handler.lockOp.state = types.AppStockLockState_AppStockCanceled.Enum()
	}
	if err := handler.requireStockAppGood(ctx); err != nil {
		return err
	}

	return db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		for _, lock := range handler.lockOp.locks {
			if err := handler.unlockAppStock(ctx, lock, tx); err != nil {
				return err
			}
			if err := handler.unlockStock(ctx, lock, tx); err != nil {
				return err
			}
		}
		if err := handler.lockOp.updateLocks(ctx, tx); err != nil {
			return err
		}
		return nil
	})
}
