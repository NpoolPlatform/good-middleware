//nolint:dupl
package appstock

import (
	"context"
	"fmt"

	wlog "github.com/NpoolPlatform/go-service-framework/pkg/wlog"
	appstockcrud "github.com/NpoolPlatform/good-middleware/pkg/crud/app/good/stock"
	appmininggoodstockcrud "github.com/NpoolPlatform/good-middleware/pkg/crud/app/good/stock/mining"
	mininggoodstockcrud "github.com/NpoolPlatform/good-middleware/pkg/crud/good/stock/mining"
	"github.com/NpoolPlatform/good-middleware/pkg/db"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent"
	types "github.com/NpoolPlatform/message/npool/basetypes/good/v1"

	"github.com/shopspring/decimal"
)

type chargeBackHandler struct {
	*stockAppGoodQuery
	lockOp *lockopHandler
}

func (h *chargeBackHandler) constructSQL(table string, lock *ent.AppStockLock, returnSpotQuantity bool) (string, error) {
	sql := fmt.Sprintf(
		`update %v
		set
		  sold = sold - %v`,
		table,
		lock.Units,
	)
	if returnSpotQuantity {
		sql += fmt.Sprintf(
			`, spot_quantity = spot_quantity + %v`,
			lock.Units,
		)
	}
	switch lock.LockState {
	case types.AppStockLockState_AppStockInService.String():
		sql += fmt.Sprintf(
			`, in_service = in_service - %v
			where
			  in_service >= %v `,
			lock.Units,
			lock.Units,
		)
	case types.AppStockLockState_AppStockWaitStart.String():
		sql += fmt.Sprintf(
			`, wait_start = wait_start - %v
			where
			  wait_start >= %v `,
			lock.Units,
			lock.Units,
		)
	default:
		return "", wlog.Errorf("invalid lockstate")
	}
	sql += fmt.Sprintf(
		`and
		  sold >= %v
		and
		  in_service + wait_start + locked + app_reserved + spot_quantity = total`,
		lock.Units,
	)
	return sql, nil
}

func (h *chargeBackHandler) chargeBackStock(ctx context.Context, lock *ent.AppStockLock, tx *ent.Tx) error {
	sql, err := h.constructSQL("stocks_v1", lock, true)
	if err != nil {
		return wlog.WrapError(err)
	}
	return h.execSQL(ctx, tx, sql)
}

func (h *chargeBackHandler) chargeBackMiningGoodStock(ctx context.Context, lock *ent.AppStockLock, tx *ent.Tx) (err error) {
	stock, ok := h.stocks[lock.AppGoodID]
	if !ok || stock.miningGoodStock == nil {
		return wlog.Errorf("invalid stock")
	}

	spotQuantity := stock.miningGoodStock.SpotQuantity
	inService := stock.miningGoodStock.InService
	waitStart := stock.miningGoodStock.WaitStart
	sold := stock.miningGoodStock.Sold

	switch lock.LockState {
	case types.AppStockLockState_AppStockInService.String():
		inService = inService.Sub(lock.Units)
	case types.AppStockLockState_AppStockWaitStart.String():
		waitStart = waitStart.Sub(lock.Units)
	}
	sold = sold.Sub(lock.Units)
	spotQuantity = spotQuantity.Add(lock.Units)

	if inService.Cmp(decimal.NewFromInt(0)) < 0 {
		return wlog.Errorf("invalid stock")
	}
	if waitStart.Cmp(decimal.NewFromInt(0)) < 0 {
		return wlog.Errorf("invalid stock")
	}
	if sold.Cmp(decimal.NewFromInt(0)) < 0 {
		return wlog.Errorf("invalid stock")
	}

	if inService.Add(waitStart).
		Add(stock.miningGoodStock.Locked).
		Add(stock.miningGoodStock.AppReserved).
		Add(spotQuantity).
		Cmp(stock.miningGoodStock.Total) > 0 {
		return wlog.Errorf("stock exhausted")
	}

	updatedStock := stock.miningGoodStock
	if updatedStock, err = mininggoodstockcrud.UpdateSet(
		tx.MiningGoodStock.UpdateOneID(stock.miningGoodStock.ID),
		&mininggoodstockcrud.Req{
			SpotQuantity: &spotQuantity,
			InService:    &inService,
			WaitStart:    &waitStart,
			Sold:         &sold,
		},
	).Save(ctx); err != nil {
		return wlog.WrapError(err)
	}
	*stock.miningGoodStock = *updatedStock
	return nil
}

func (h *chargeBackHandler) chargeBackAppStock(ctx context.Context, lock *ent.AppStockLock, tx *ent.Tx) (err error) {
	stock, ok := h.stocks[lock.AppGoodID]
	if !ok || stock.appGoodStock == nil {
		return wlog.Errorf("invalid stock")
	}

	inService := stock.appGoodStock.InService
	waitStart := stock.appGoodStock.WaitStart
	sold := stock.appGoodStock.Sold

	switch lock.LockState {
	case types.AppStockLockState_AppStockInService.String():
		inService = inService.Sub(lock.Units)
	case types.AppStockLockState_AppStockWaitStart.String():
		waitStart = waitStart.Sub(lock.Units)
	}
	sold = sold.Sub(lock.Units)

	if inService.Cmp(decimal.NewFromInt(0)) < 0 {
		return wlog.Errorf("invalid stock")
	}
	if waitStart.Cmp(decimal.NewFromInt(0)) < 0 {
		return wlog.Errorf("invalid stock")
	}
	if sold.Cmp(decimal.NewFromInt(0)) < 0 {
		return wlog.Errorf("invalid stock")
	}

	if stock.appGoodStock, err = appstockcrud.UpdateSet(
		tx.AppStock.UpdateOneID(stock.appGoodStock.ID),
		&appstockcrud.Req{
			InService: &inService,
			WaitStart: &waitStart,
			Sold:      &sold,
		},
	).Save(ctx); err != nil {
		return wlog.WrapError(err)
	}

	return nil
}

func (h *chargeBackHandler) chargeBackAppMiningGoodStock(ctx context.Context, lock *ent.AppStockLock, tx *ent.Tx) (err error) {
	stock, ok := h.stocks[lock.AppGoodID]
	if !ok || stock.appMiningGoodStock == nil {
		return wlog.Errorf("invalid stock")
	}

	inService := stock.appMiningGoodStock.InService
	waitStart := stock.appMiningGoodStock.WaitStart
	sold := stock.appMiningGoodStock.Sold

	switch lock.LockState {
	case types.AppStockLockState_AppStockInService.String():
		inService = inService.Sub(lock.Units)
	case types.AppStockLockState_AppStockWaitStart.String():
		waitStart = waitStart.Sub(lock.Units)
	}
	sold = sold.Sub(lock.Units)

	if inService.Cmp(decimal.NewFromInt(0)) < 0 {
		return wlog.Errorf("invalid stock")
	}
	if waitStart.Cmp(decimal.NewFromInt(0)) < 0 {
		return wlog.Errorf("invalid stock")
	}
	if sold.Cmp(decimal.NewFromInt(0)) < 0 {
		return wlog.Errorf("invalid stock")
	}

	if stock.appMiningGoodStock, err = appmininggoodstockcrud.UpdateSet(
		tx.AppMiningGoodStock.UpdateOneID(stock.appMiningGoodStock.ID),
		&appmininggoodstockcrud.Req{
			InService: &inService,
			WaitStart: &waitStart,
			Sold:      &sold,
		},
	).Save(ctx); err != nil {
		return wlog.WrapError(err)
	}

	return nil
}

func (h *Handler) ChargeBackStock(ctx context.Context) error {
	handler := &chargeBackHandler{
		stockAppGoodQuery: &stockAppGoodQuery{
			Handler: h,
		},
		lockOp: &lockopHandler{
			Handler: h,
			state:   types.AppStockLockState_AppStockChargeBack.Enum(),
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
			if err := handler.chargeBackAppStock(ctx, lock, tx); err != nil {
				return wlog.WrapError(err)
			}
			if handler.stockByMiningPool(lock.AppGoodID) {
				if err := handler.chargeBackAppMiningGoodStock(ctx, lock, tx); err != nil {
					return wlog.WrapError(err)
				}
			}
			if err := handler.chargeBackStock(ctx, lock, tx); err != nil {
				return wlog.WrapError(err)
			}
			if handler.stockByMiningPool(lock.AppGoodID) {
				if err := handler.chargeBackMiningGoodStock(ctx, lock, tx); err != nil {
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
