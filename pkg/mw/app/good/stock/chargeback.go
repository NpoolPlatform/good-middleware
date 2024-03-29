//nolint:dupl
package appstock

import (
	"context"
	"fmt"

	appstockcrud "github.com/NpoolPlatform/good-middleware/pkg/crud/app/good/stock"
	appmininggoodstockcrud "github.com/NpoolPlatform/good-middleware/pkg/crud/app/good/stock/mining"
	stockcrud "github.com/NpoolPlatform/good-middleware/pkg/crud/good/stock"
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

func (h *chargeBackHandler) chargeBackStock(ctx context.Context, lock *ent.AppStockLock, tx *ent.Tx) error {
	stock, ok := h.stocks[lock.AppGoodID]
	if !ok {
		return fmt.Errorf("invalid stock")
	}

	spotQuantity := stock.stock.SpotQuantity
	inService := stock.stock.InService
	waitStart := stock.stock.WaitStart
	sold := stock.stock.Sold

	switch lock.LockState {
	case types.AppStockLockState_AppStockInService.String():
		inService = inService.Sub(lock.Units)
	case types.AppStockLockState_AppStockWaitStart.String():
		waitStart = waitStart.Sub(lock.Units)
	}
	sold = sold.Sub(lock.Units)
	spotQuantity = spotQuantity.Add(lock.Units)

	if inService.Cmp(decimal.NewFromInt(0)) < 0 {
		return fmt.Errorf("invalid stock")
	}
	if waitStart.Cmp(decimal.NewFromInt(0)) < 0 {
		return fmt.Errorf("invalid stock")
	}
	if sold.Cmp(decimal.NewFromInt(0)) < 0 {
		return fmt.Errorf("invalid stock")
	}

	if inService.Add(waitStart).
		Add(stock.stock.Locked).
		Add(stock.stock.AppReserved).
		Add(spotQuantity).
		Cmp(stock.stock.Total) > 0 {
		return fmt.Errorf("stock exhausted")
	}

	if _, err := stockcrud.UpdateSet(
		tx.Stock.UpdateOneID(stock.stock.ID),
		&stockcrud.Req{
			SpotQuantity: &spotQuantity,
			InService:    &inService,
			WaitStart:    &waitStart,
			Sold:         &sold,
		},
	).Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *chargeBackHandler) chargeBackMiningGoodStock(ctx context.Context, lock *ent.AppStockLock, tx *ent.Tx) error {
	if !h.stockByMiningPool(lock.AppGoodID) {
		return nil
	}

	stock, ok := h.stocks[lock.AppGoodID]
	if !ok {
		return fmt.Errorf("invalid stock")
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
		return fmt.Errorf("invalid stock")
	}
	if waitStart.Cmp(decimal.NewFromInt(0)) < 0 {
		return fmt.Errorf("invalid stock")
	}
	if sold.Cmp(decimal.NewFromInt(0)) < 0 {
		return fmt.Errorf("invalid stock")
	}

	if inService.Add(waitStart).
		Add(stock.miningGoodStock.Locked).
		Add(stock.miningGoodStock.AppReserved).
		Add(spotQuantity).
		Cmp(stock.miningGoodStock.Total) > 0 {
		return fmt.Errorf("stock exhausted")
	}

	if _, err := mininggoodstockcrud.UpdateSet(
		tx.MiningGoodStock.UpdateOneID(stock.miningGoodStock.ID),
		&mininggoodstockcrud.Req{
			SpotQuantity: &spotQuantity,
			InService:    &inService,
			WaitStart:    &waitStart,
			Sold:         &sold,
		},
	).Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *chargeBackHandler) chargeBackAppStock(ctx context.Context, lock *ent.AppStockLock, tx *ent.Tx) error {
	stock, ok := h.stocks[lock.AppGoodID]
	if !ok {
		return fmt.Errorf("invalid stock")
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
		return fmt.Errorf("invalid stock")
	}
	if waitStart.Cmp(decimal.NewFromInt(0)) < 0 {
		return fmt.Errorf("invalid stock")
	}
	if sold.Cmp(decimal.NewFromInt(0)) < 0 {
		return fmt.Errorf("invalid stock")
	}

	if _, err := appstockcrud.UpdateSet(
		tx.AppStock.UpdateOneID(stock.appGoodStock.ID),
		&appstockcrud.Req{
			InService: &inService,
			WaitStart: &waitStart,
			Sold:      &sold,
		},
	).Save(ctx); err != nil {
		return err
	}

	return nil
}

func (h *chargeBackHandler) chargeBackAppMiningGoodStock(ctx context.Context, lock *ent.AppStockLock, tx *ent.Tx) error {
	if !h.stockByMiningPool(lock.AppGoodID) {
		return nil
	}

	stock, ok := h.stocks[lock.AppGoodID]
	if !ok {
		return fmt.Errorf("invalid stock")
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
		return fmt.Errorf("invalid stock")
	}
	if waitStart.Cmp(decimal.NewFromInt(0)) < 0 {
		return fmt.Errorf("invalid stock")
	}
	if sold.Cmp(decimal.NewFromInt(0)) < 0 {
		return fmt.Errorf("invalid stock")
	}

	if _, err := appmininggoodstockcrud.UpdateSet(
		tx.AppMiningGoodStock.UpdateOneID(stock.appMiningGoodStock.ID),
		&appmininggoodstockcrud.Req{
			InService: &inService,
			WaitStart: &waitStart,
			Sold:      &sold,
		},
	).Save(ctx); err != nil {
		return err
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
		return err
	}
	h.Stocks = handler.lockOp.lock2Stocks()
	if err := handler.getStockAppGoods(ctx); err != nil {
		return err
	}

	return db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		for _, lock := range handler.lockOp.locks {
			if err := handler.chargeBackAppStock(ctx, lock, tx); err != nil {
				return err
			}
			if err := handler.chargeBackStock(ctx, lock, tx); err != nil {
				return err
			}
			if err := handler.chargeBackMiningGoodStock(ctx, lock, tx); err != nil {
				return err
			}
			if err := handler.chargeBackAppMiningGoodStock(ctx, lock, tx); err != nil {
				return err
			}
		}
		if err := handler.lockOp.updateLocks(ctx, tx); err != nil {
			return err
		}
		return nil
	})
}
