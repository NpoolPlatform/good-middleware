//nolint:dupl
package appstock

import (
	"context"
	"fmt"

	wlog "github.com/NpoolPlatform/go-service-framework/pkg/wlog"
	"github.com/NpoolPlatform/good-middleware/pkg/db"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent"
	types "github.com/NpoolPlatform/message/npool/basetypes/good/v1"
)

type expireHandler struct {
	*stockAppGoodQuery
	lockOp *lockopHandler
}

func (h *expireHandler) constructSQL(table string, lock *ent.AppStockLock, returnSpotQuantity bool, checkTotal bool, id uint32) (string, error) {
	sql := fmt.Sprintf(
		`update %v
		set
		  in_service = in_service - %v`,
		table,
		lock.Units,
	)
	if returnSpotQuantity {
		sql += fmt.Sprintf(
			`, spot_quantity = spot_quantity + %v`,
			lock.Units,
		)
	}
	sql += fmt.Sprintf(
		` where
		  id = %v
		and
		  deleted_at = 0
		and
		  in_service >= %v`,
		id,
		lock.Units,
	)
	if checkTotal {
		sql += ` and
		  in_service + wait_start + locked + app_reserved + spot_quantity = total`
	}
	return sql, nil
}

func (h *expireHandler) expireStock(ctx context.Context, lock *ent.AppStockLock, tx *ent.Tx) (err error) {
	stock, ok := h.stocks[lock.AppGoodID]
	if !ok || stock.stock == nil {
		return wlog.Errorf("invalid stock")
	}
	sql, err := h.constructSQL("stocks_v1", lock, true, true, stock.stock.ID)
	if err != nil {
		return wlog.WrapError(err)
	}
	return h.execSQL(ctx, tx, sql)
}

func (h *expireHandler) expireMiningGoodStock(ctx context.Context, lock *ent.AppStockLock, tx *ent.Tx) (err error) {
	stock, ok := h.stocks[lock.AppGoodID]
	if !ok || stock.miningGoodStock == nil {
		return wlog.Errorf("invalid stock")
	}
	sql, err := h.constructSQL("mining_good_stocks", lock, true, true, stock.miningGoodStock.ID)
	if err != nil {
		return wlog.WrapError(err)
	}
	return h.execSQL(ctx, tx, sql)
}

func (h *expireHandler) expireAppStock(ctx context.Context, lock *ent.AppStockLock, tx *ent.Tx) (err error) {
	stock, ok := h.stocks[lock.AppGoodID]
	if !ok || stock.appGoodStock == nil {
		return wlog.Errorf("invalid stock")
	}
	sql, err := h.constructSQL("app_stocks", lock, false, false, stock.appGoodStock.ID)
	if err != nil {
		return wlog.WrapError(err)
	}
	return h.execSQL(ctx, tx, sql)
}

func (h *expireHandler) expireAppMiningGoodStock(ctx context.Context, lock *ent.AppStockLock, tx *ent.Tx) (err error) {
	stock, ok := h.stocks[lock.AppGoodID]
	if !ok || stock.appMiningGoodStock == nil {
		return wlog.Errorf("invalid stock")
	}
	sql, err := h.constructSQL("app_mining_good_stocks", lock, false, false, stock.appMiningGoodStock.ID)
	if err != nil {
		return wlog.WrapError(err)
	}
	return h.execSQL(ctx, tx, sql)
}

func (h *Handler) ExpireStock(ctx context.Context) error {
	handler := &expireHandler{
		stockAppGoodQuery: &stockAppGoodQuery{
			Handler: h,
		},
		lockOp: &lockopHandler{
			Handler: h,
			state:   types.AppStockLockState_AppStockExpired.Enum(),
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
			if err := handler.expireAppStock(ctx, lock, tx); err != nil {
				return wlog.WrapError(err)
			}
			if handler.stockByMiningPool(lock.AppGoodID) {
				if err := handler.expireAppMiningGoodStock(ctx, lock, tx); err != nil {
					return wlog.WrapError(err)
				}
			}
			if err := handler.expireStock(ctx, lock, tx); err != nil {
				return wlog.WrapError(err)
			}
			if handler.stockByMiningPool(lock.AppGoodID) {
				if err := handler.expireMiningGoodStock(ctx, lock, tx); err != nil {
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
