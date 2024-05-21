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

type waitStartHandler struct {
	*stockAppGoodQuery
	lockOp *lockopHandler
}

func (h *waitStartHandler) waitStartStock(ctx context.Context, lock *ent.AppStockLock, tx *ent.Tx) (err error) {
	stock, ok := h.stocks[lock.AppGoodID]
	if !ok || stock.stock == nil {
		return wlog.Errorf("invalid stock")
	}

	locked := stock.stock.Locked
	sold := stock.stock.Sold
	waitStart := stock.stock.WaitStart

	waitStart = lock.Units.Add(waitStart)
	locked = locked.Sub(lock.Units)
	sold = lock.Units.Add(sold)

	if locked.Cmp(decimal.NewFromInt(0)) < 0 {
		return wlog.Errorf("invalid stock")
	}
	if waitStart.Cmp(decimal.NewFromInt(0)) < 0 {
		return wlog.Errorf("invalid stock")
	}

	if locked.Add(stock.stock.InService).
		Add(stock.stock.WaitStart).
		Add(stock.stock.AppReserved).
		Add(stock.stock.SpotQuantity).
		Cmp(stock.stock.Total) > 0 {
		return wlog.Errorf("stock exhausted")
	}

	updatedStock := stock.stock
	if updatedStock, err = stockcrud.UpdateSet(
		tx.Stock.UpdateOneID(stock.stock.ID),
		&stockcrud.Req{
			Locked:    &locked,
			WaitStart: &waitStart,
			Sold:      &sold,
		},
	).Save(ctx); err != nil {
		return wlog.WrapError(err)
	}
	*stock.stock = *updatedStock
	return nil
}

func (h *waitStartHandler) waitStartMiningGoodStock(ctx context.Context, lock *ent.AppStockLock, tx *ent.Tx) (err error) {
	stock, ok := h.stocks[lock.AppGoodID]
	if !ok || stock.miningGoodStock == nil {
		return wlog.Errorf("invalid stock")
	}

	locked := stock.miningGoodStock.Locked
	sold := stock.miningGoodStock.Sold
	waitStart := stock.miningGoodStock.WaitStart

	waitStart = lock.Units.Add(waitStart)
	locked = locked.Sub(lock.Units)
	sold = lock.Units.Add(sold)

	if locked.Cmp(decimal.NewFromInt(0)) < 0 {
		return wlog.Errorf("invalid stock")
	}
	if waitStart.Cmp(decimal.NewFromInt(0)) < 0 {
		return wlog.Errorf("invalid stock")
	}

	if locked.Add(stock.miningGoodStock.InService).
		Add(stock.miningGoodStock.WaitStart).
		Add(stock.miningGoodStock.AppReserved).
		Add(stock.miningGoodStock.SpotQuantity).
		Cmp(stock.miningGoodStock.Total) > 0 {
		return wlog.Errorf("stock exhausted")
	}

	updatedStock := stock.miningGoodStock
	if updatedStock, err = mininggoodstockcrud.UpdateSet(
		tx.MiningGoodStock.UpdateOneID(stock.miningGoodStock.ID),
		&mininggoodstockcrud.Req{
			Locked:    &locked,
			WaitStart: &waitStart,
			Sold:      &sold,
		},
	).Save(ctx); err != nil {
		return wlog.WrapError(err)
	}
	*stock.miningGoodStock = *updatedStock
	return nil
}

func (h *waitStartHandler) waitStartAppStock(ctx context.Context, lock *ent.AppStockLock, tx *ent.Tx) (err error) {
	stock, ok := h.stocks[lock.AppGoodID]
	if !ok || stock.appGoodStock == nil {
		return wlog.Errorf("invalid stock")
	}

	h.ID = &stock.appGoodStock.ID
	locked := stock.appGoodStock.Locked
	sold := stock.appGoodStock.Sold
	waitStart := stock.appGoodStock.WaitStart

	waitStart = lock.Units.Add(waitStart)
	locked = locked.Sub(lock.Units)
	sold = lock.Units.Add(sold)

	if locked.Cmp(decimal.NewFromInt(0)) < 0 {
		return wlog.Errorf("invalid stock")
	}
	if waitStart.Cmp(decimal.NewFromInt(0)) < 0 {
		return wlog.Errorf("invalid stock")
	}

	if stock.appGoodStock, err = appstockcrud.UpdateSet(
		tx.AppStock.UpdateOneID(stock.appGoodStock.ID),
		&appstockcrud.Req{
			Locked:    &locked,
			WaitStart: &waitStart,
			Sold:      &sold,
		},
	).Save(ctx); err != nil {
		return wlog.WrapError(err)
	}

	return nil
}

func (h *waitStartHandler) waitStartAppMiningGoodStock(ctx context.Context, lock *ent.AppStockLock, tx *ent.Tx) (err error) {
	stock, ok := h.stocks[lock.AppGoodID]
	if !ok || stock.appMiningGoodStock == nil {
		return wlog.Errorf("invalid stock")
	}

	h.ID = &stock.appMiningGoodStock.ID
	locked := stock.appMiningGoodStock.Locked
	sold := stock.appMiningGoodStock.Sold
	waitStart := stock.appMiningGoodStock.WaitStart

	waitStart = lock.Units.Add(waitStart)
	locked = locked.Sub(lock.Units)
	sold = lock.Units.Add(sold)

	if locked.Cmp(decimal.NewFromInt(0)) < 0 {
		return wlog.Errorf("invalid stock")
	}
	if waitStart.Cmp(decimal.NewFromInt(0)) < 0 {
		return wlog.Errorf("invalid stock")
	}

	if stock.appMiningGoodStock, err = appmininggoodstockcrud.UpdateSet(
		tx.AppMiningGoodStock.UpdateOneID(stock.appMiningGoodStock.ID),
		&appmininggoodstockcrud.Req{
			Locked:    &locked,
			WaitStart: &waitStart,
			Sold:      &sold,
		},
	).Save(ctx); err != nil {
		return wlog.WrapError(err)
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
		return wlog.WrapError(err)
	}
	h.Stocks = handler.lockOp.lock2Stocks()
	if err := handler.getStockAppGoods(ctx); err != nil {
		return wlog.WrapError(err)
	}

	return db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		for _, lock := range handler.lockOp.locks {
			if err := handler.waitStartAppStock(ctx, lock, tx); err != nil {
				return wlog.WrapError(err)
			}
			if handler.stockByMiningPool(lock.AppGoodID) {
				if err := handler.waitStartAppMiningGoodStock(ctx, lock, tx); err != nil {
					return wlog.WrapError(err)
				}
			}
			if err := handler.waitStartStock(ctx, lock, tx); err != nil {
				return wlog.WrapError(err)
			}
			if handler.stockByMiningPool(lock.AppGoodID) {
				if err := handler.waitStartMiningGoodStock(ctx, lock, tx); err != nil {
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
