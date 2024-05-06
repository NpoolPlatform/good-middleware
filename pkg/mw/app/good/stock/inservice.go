//nolint:dupl
package appstock

import (
	"context"

	wlog "github.com/NpoolPlatform/go-service-framework/pkg/wlog"
	appstockcrud "github.com/NpoolPlatform/good-middleware/pkg/crud/app/good/stock"
	stockcrud "github.com/NpoolPlatform/good-middleware/pkg/crud/good/stock"
	"github.com/NpoolPlatform/good-middleware/pkg/db"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent"
	types "github.com/NpoolPlatform/message/npool/basetypes/good/v1"

	"github.com/shopspring/decimal"
)

type inServiceHandler struct {
	*stockAppGoodQuery
	lockOp *lockopHandler
}

func (h *inServiceHandler) inServiceStock(ctx context.Context, lock *ent.AppStockLock, tx *ent.Tx) error {
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

	if _, err := stockcrud.UpdateSet(
		tx.Stock.UpdateOneID(stock.stock.ID),
		&stockcrud.Req{
			InService: &inService,
			WaitStart: &waitStart,
		},
	).Save(ctx); err != nil {
		return wlog.WrapError(err)
	}
	return nil
}

func (h *inServiceHandler) inServiceAppStock(ctx context.Context, lock *ent.AppStockLock, tx *ent.Tx) error {
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

	if _, err := appstockcrud.UpdateSet(
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
			if err := handler.inServiceStock(ctx, lock, tx); err != nil {
				return wlog.WrapError(err)
			}
		}
		if err := handler.lockOp.updateLocks(ctx, tx); err != nil {
			return wlog.WrapError(err)
		}
		return nil
	})
}
