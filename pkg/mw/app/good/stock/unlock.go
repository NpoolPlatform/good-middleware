//nolint:dupl
package appstock

import (
	"context"

	wlog "github.com/NpoolPlatform/go-service-framework/pkg/wlog"
	appstockcrud "github.com/NpoolPlatform/good-middleware/pkg/crud/app/good/stock"
	appmininggoodstockcrud "github.com/NpoolPlatform/good-middleware/pkg/crud/app/good/stock/mining"
	stockcrud "github.com/NpoolPlatform/good-middleware/pkg/crud/good/stock"
	mininggoodstockcrud "github.com/NpoolPlatform/good-middleware/pkg/crud/good/stock/mining"
	"github.com/NpoolPlatform/good-middleware/pkg/db"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent"
	types "github.com/NpoolPlatform/message/npool/basetypes/good/v1"

	"github.com/shopspring/decimal"
)

type unlockHandler struct {
	*stockAppGoodQuery
	lockOp *lockopHandler
}

func (h *unlockHandler) unlockStock(ctx context.Context, lock *ent.AppStockLock, tx *ent.Tx) (err error) {
	stock, ok := h.stocks[lock.AppGoodID]
	if !ok || stock.stock == nil {
		return wlog.Errorf("invalid stock")
	}

	spotQuantity := stock.stock.SpotQuantity
	locked := stock.stock.Locked
	appReserved := stock.stock.AppReserved

	locked = locked.Sub(lock.Units)
	platformLocked := lock.Units.Sub(lock.AppSpotUnits)
	appReserved = lock.AppSpotUnits.Add(appReserved)
	if platformLocked.Cmp(decimal.NewFromInt(0)) < 0 {
		return wlog.Errorf("invalid appspotlocked")
	}
	spotQuantity = platformLocked.Add(spotQuantity)
	if locked.Cmp(decimal.NewFromInt(0)) < 0 {
		return wlog.Errorf("invalid locked")
	}
	if appReserved.Cmp(decimal.NewFromInt(0)) < 0 {
		return wlog.Errorf("invalid appreserved")
	}
	if spotQuantity.Cmp(decimal.NewFromInt(0)) < 0 {
		return wlog.Errorf("invalid stock")
	}
	if spotQuantity.Cmp(stock.stock.Total) > 0 {
		return wlog.Errorf("invalid stock")
	}

	if locked.Add(stock.stock.InService).
		Add(stock.stock.WaitStart).
		Add(appReserved).
		Add(spotQuantity).
		Cmp(stock.stock.Total) > 0 {
		return wlog.Errorf("invalid stock")
	}

	updatedStock := stock.stock
	if updatedStock, err = stockcrud.UpdateSet(
		tx.Stock.UpdateOneID(stock.stock.ID),
		&stockcrud.Req{
			SpotQuantity: &spotQuantity,
			Locked:       &locked,
			AppReserved:  &appReserved,
		},
	).Save(ctx); err != nil {
		return wlog.WrapError(err)
	}
	*stock.stock = *updatedStock
	return nil
}

func (h *unlockHandler) unlockMiningGoodStock(ctx context.Context, lock *ent.AppStockLock, tx *ent.Tx) (err error) {
	stock, ok := h.stocks[lock.AppGoodID]
	if !ok || stock.miningGoodStock == nil {
		return wlog.Errorf("invalid stock")
	}

	spotQuantity := stock.miningGoodStock.SpotQuantity
	locked := stock.miningGoodStock.Locked
	appReserved := stock.miningGoodStock.AppReserved

	locked = locked.Sub(lock.Units)
	platformLocked := lock.Units.Sub(lock.AppSpotUnits)
	appReserved = lock.AppSpotUnits.Add(appReserved)
	if platformLocked.Cmp(decimal.NewFromInt(0)) < 0 {
		return wlog.Errorf("invalid appspotlocked")
	}
	spotQuantity = platformLocked.Add(spotQuantity)
	if locked.Cmp(decimal.NewFromInt(0)) < 0 {
		return wlog.Errorf("invalid locked")
	}
	if appReserved.Cmp(decimal.NewFromInt(0)) < 0 {
		return wlog.Errorf("invalid appreserved")
	}
	if spotQuantity.Cmp(decimal.NewFromInt(0)) < 0 {
		return wlog.Errorf("invalid stock")
	}
	if spotQuantity.Cmp(stock.miningGoodStock.Total) > 0 {
		return wlog.Errorf("invalid stock")
	}

	if locked.Add(stock.miningGoodStock.InService).
		Add(stock.miningGoodStock.WaitStart).
		Add(appReserved).
		Add(spotQuantity).
		Cmp(stock.miningGoodStock.Total) > 0 {
		return wlog.Errorf("invalid stock")
	}

	updatedStock := stock.miningGoodStock
	if updatedStock, err = mininggoodstockcrud.UpdateSet(
		tx.MiningGoodStock.UpdateOneID(stock.miningGoodStock.ID),
		&mininggoodstockcrud.Req{
			SpotQuantity: &spotQuantity,
			Locked:       &locked,
			AppReserved:  &appReserved,
		},
	).Save(ctx); err != nil {
		return wlog.WrapError(err)
	}
	*stock.miningGoodStock = *updatedStock
	return nil
}

func (h *unlockHandler) unlockAppStock(ctx context.Context, lock *ent.AppStockLock, tx *ent.Tx) (err error) {
	stock, ok := h.stocks[lock.AppGoodID]
	if !ok || stock.appGoodStock == nil {
		return wlog.Errorf("invalid stock")
	}

	locked := stock.appGoodStock.Locked
	spotQuantity := stock.appGoodStock.SpotQuantity

	locked = locked.Sub(lock.Units)
	spotQuantity = lock.AppSpotUnits.Add(spotQuantity)
	if locked.Cmp(decimal.NewFromInt(0)) < 0 {
		return wlog.Errorf("invalid locked")
	}
	if spotQuantity.Cmp(stock.appGoodStock.Reserved) > 0 {
		return wlog.Errorf("invalid spotquantity")
	}
	if spotQuantity.Cmp(decimal.NewFromInt(0)) < 0 {
		spotQuantity = decimal.NewFromInt(0)
	}

	if stock.appGoodStock, err = appstockcrud.UpdateSet(
		tx.AppStock.UpdateOneID(stock.appGoodStock.ID),
		&appstockcrud.Req{
			SpotQuantity: &spotQuantity,
			Locked:       &locked,
		},
	).Save(ctx); err != nil {
		return wlog.WrapError(err)
	}
	return nil
}

func (h *unlockHandler) unlockAppMiningGoodStock(ctx context.Context, lock *ent.AppStockLock, tx *ent.Tx) (err error) {
	stock, ok := h.stocks[lock.AppGoodID]
	if !ok || stock.appMiningGoodStock == nil {
		return wlog.Errorf("invalid stock")
	}

	locked := stock.appMiningGoodStock.Locked
	spotQuantity := stock.appMiningGoodStock.SpotQuantity

	locked = locked.Sub(lock.Units)
	spotQuantity = lock.AppSpotUnits.Add(spotQuantity)
	if locked.Cmp(decimal.NewFromInt(0)) < 0 {
		return wlog.Errorf("invalid locked")
	}
	if spotQuantity.Cmp(stock.appMiningGoodStock.Reserved) > 0 {
		return wlog.Errorf("invalid spotquantity")
	}
	if spotQuantity.Cmp(decimal.NewFromInt(0)) < 0 {
		spotQuantity = decimal.NewFromInt(0)
	}

	if stock.appMiningGoodStock, err = appmininggoodstockcrud.UpdateSet(
		tx.AppMiningGoodStock.UpdateOneID(stock.appMiningGoodStock.ID),
		&appmininggoodstockcrud.Req{
			SpotQuantity: &spotQuantity,
			Locked:       &locked,
		},
	).Save(ctx); err != nil {
		return wlog.WrapError(err)
	}
	return nil
}

//nolint:gocyclo
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
		return wlog.WrapError(err)
	}
	if h.Rollback == nil || !*h.Rollback {
		handler.lockOp.state = types.AppStockLockState_AppStockCanceled.Enum()
	}
	h.Stocks = handler.lockOp.lock2Stocks()
	if err := handler.getStockAppGoods(ctx); err != nil {
		return wlog.WrapError(err)
	}

	return db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		for _, lock := range handler.lockOp.locks {
			if err := handler.unlockAppStock(ctx, lock, tx); err != nil {
				return wlog.WrapError(err)
			}
			if handler.stockByMiningPool(lock.AppGoodID) {
				if err := handler.unlockAppMiningGoodStock(ctx, lock, tx); err != nil {
					return wlog.WrapError(err)
				}
			}
			if err := handler.unlockStock(ctx, lock, tx); err != nil {
				return wlog.WrapError(err)
			}
			if handler.stockByMiningPool(lock.AppGoodID) {
				if err := handler.unlockMiningGoodStock(ctx, lock, tx); err != nil {
					return wlog.WrapError(err)
				}
			}
		}
		if err := handler.lockOp.updateLocks(ctx, tx); err != nil {
			return wlog.WrapError(err)
		}
		return nil
	})
}
