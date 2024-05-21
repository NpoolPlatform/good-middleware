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

type inServiceHandler struct {
	*stockAppGoodQuery
	lockOp *lockopHandler
}

func (h *inServiceHandler) inServiceStock(ctx context.Context, lock *ent.AppStockLock, tx *ent.Tx) (err error) {
	stock, ok := h.stocks[lock.AppGoodID]
	if !ok || stock.stock == nil {
		return wlog.Errorf("invalid stock")
	}

	inService := stock.stock.InService
	waitStart := stock.stock.WaitStart

	inService = lock.Units.Add(inService)
	waitStart = waitStart.Sub(lock.Units)

	if waitStart.Cmp(decimal.NewFromInt(0)) < 0 {
		return wlog.Errorf("invalid stock")
	}

	if waitStart.Add(inService).
		Add(stock.stock.Locked).
		Add(stock.stock.AppReserved).
		Add(stock.stock.SpotQuantity).
		Cmp(stock.stock.Total) > 0 {
		return wlog.Errorf("stock exhausted")
	}

	updatedStock := stock.stock
	if updatedStock, err = stockcrud.UpdateSet(
		tx.Stock.UpdateOneID(stock.stock.ID),
		&stockcrud.Req{
			InService: &inService,
			WaitStart: &waitStart,
		},
	).Save(ctx); err != nil {
		return wlog.WrapError(err)
	}
	*stock.stock = *updatedStock
	return nil
}

func (h *inServiceHandler) inServiceMiningGoodStock(ctx context.Context, lock *ent.AppStockLock, tx *ent.Tx) (err error) {
	stock, ok := h.stocks[lock.AppGoodID]
	if !ok || stock.miningGoodStock == nil {
		return wlog.Errorf("invalid stock")
	}

	inService := stock.miningGoodStock.InService
	waitStart := stock.miningGoodStock.WaitStart

	inService = lock.Units.Add(inService)
	waitStart = waitStart.Sub(lock.Units)

	if waitStart.Cmp(decimal.NewFromInt(0)) < 0 {
		return wlog.Errorf("invalid stock")
	}

	if waitStart.Add(inService).
		Add(stock.miningGoodStock.Locked).
		Add(stock.miningGoodStock.AppReserved).
		Add(stock.miningGoodStock.SpotQuantity).
		Cmp(stock.miningGoodStock.Total) > 0 {
		return wlog.Errorf("stock exhausted")
	}

	updatedStock := stock.miningGoodStock
	if updatedStock, err = mininggoodstockcrud.UpdateSet(
		tx.MiningGoodStock.UpdateOneID(stock.miningGoodStock.ID),
		&mininggoodstockcrud.Req{
			InService: &inService,
			WaitStart: &waitStart,
		},
	).Save(ctx); err != nil {
		return wlog.WrapError(err)
	}
	*stock.miningGoodStock = *updatedStock
	return nil
}

func (h *inServiceHandler) inServiceAppStock(ctx context.Context, lock *ent.AppStockLock, tx *ent.Tx) (err error) {
	stock, ok := h.stocks[lock.AppGoodID]
	if !ok || stock.appGoodStock == nil {
		return wlog.Errorf("invalid stock")
	}

	inService := stock.appGoodStock.InService
	waitStart := stock.appGoodStock.WaitStart

	inService = lock.Units.Add(inService)
	waitStart = waitStart.Sub(lock.Units)

	if waitStart.Cmp(decimal.NewFromInt(0)) < 0 {
		return wlog.Errorf("invalid stock")
	}

	if stock.appGoodStock, err = appstockcrud.UpdateSet(
		tx.AppStock.UpdateOneID(stock.appGoodStock.ID),
		&appstockcrud.Req{
			InService: &inService,
			WaitStart: &waitStart,
		},
	).Save(ctx); err != nil {
		return wlog.WrapError(err)
	}

	return nil
}

func (h *inServiceHandler) inServiceAppMiningGoodStock(ctx context.Context, lock *ent.AppStockLock, tx *ent.Tx) (err error) {
	stock, ok := h.stocks[lock.AppGoodID]
	if !ok || stock.appMiningGoodStock == nil {
		return wlog.Errorf("invalid stock")
	}

	inService := stock.appMiningGoodStock.InService
	waitStart := stock.appMiningGoodStock.WaitStart

	inService = lock.Units.Add(inService)
	waitStart = waitStart.Sub(lock.Units)

	if waitStart.Cmp(decimal.NewFromInt(0)) < 0 {
		return wlog.Errorf("invalid stock")
	}

	if stock.appMiningGoodStock, err = appmininggoodstockcrud.UpdateSet(
		tx.AppMiningGoodStock.UpdateOneID(stock.appMiningGoodStock.ID),
		&appmininggoodstockcrud.Req{
			InService: &inService,
			WaitStart: &waitStart,
		},
	).Save(ctx); err != nil {
		return wlog.WrapError(err)
	}

	return nil
}

func (h *Handler) InServiceStock(ctx context.Context) error {
	handler := &inServiceHandler{
		stockAppGoodQuery: &stockAppGoodQuery{
			Handler: h,
		},
		lockOp: &lockopHandler{
			Handler: h,
			state:   types.AppStockLockState_AppStockInService.Enum(),
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
			if err := handler.inServiceAppStock(ctx, lock, tx); err != nil {
				return wlog.WrapError(err)
			}
			if handler.stockByMiningPool(lock.AppGoodID) {
				if err := handler.inServiceAppMiningGoodStock(ctx, lock, tx); err != nil {
					return wlog.WrapError(err)
				}
			}
			if err := handler.inServiceStock(ctx, lock, tx); err != nil {
				return wlog.WrapError(err)
			}
			if handler.stockByMiningPool(lock.AppGoodID) {
				if err := handler.inServiceMiningGoodStock(ctx, lock, tx); err != nil {
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
