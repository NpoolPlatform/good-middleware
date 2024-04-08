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

	"github.com/shopspring/decimal"
)

type lockHandler struct {
	*stockAppGoodQuery
	lockOp *lockopHandler
}

func (h *lockHandler) lockStock(ctx context.Context, stock *LockStock, tx *ent.Tx) error {
	_stock, ok := h.stocks[*stock.AppGoodID]
	if !ok {
		return fmt.Errorf("invalid stock")
	}

	spotQuantity := _stock.stock.SpotQuantity
	locked := _stock.stock.Locked
	appReserved := _stock.stock.AppReserved

	locked = stock.Locked.Add(locked)
	platformLocked := *stock.Locked
	if stock.AppSpotLocked != nil {
		platformLocked = platformLocked.Sub(*stock.AppSpotLocked)
		appReserved = appReserved.Sub(*stock.AppSpotLocked)
	}
	if platformLocked.Cmp(decimal.NewFromInt(0)) < 0 {
		return fmt.Errorf("invalid appspotlocked")
	}
	if appReserved.Cmp(decimal.NewFromInt(0)) < 0 {
		return fmt.Errorf("invalid appreserved")
	}
	spotQuantity = spotQuantity.Sub(platformLocked)
	if spotQuantity.Cmp(decimal.NewFromInt(0)) < 0 {
		return fmt.Errorf("invalid stock")
	}
	if spotQuantity.Cmp(_stock.stock.Total) > 0 {
		return fmt.Errorf("invalid stock")
	}
	if locked.Cmp(decimal.NewFromInt(0)) < 0 {
		return fmt.Errorf("invalid stock")
	}

	if locked.Add(_stock.stock.InService).
		Add(_stock.stock.WaitStart).
		Add(appReserved).
		Add(spotQuantity).
		Cmp(_stock.stock.Total) > 0 {
		return fmt.Errorf("stock exhausted")
	}

	if _, err := stockcrud.UpdateSet(
		tx.Stock.UpdateOneID(_stock.stock.ID),
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

func (h *lockHandler) lockMiningGoodStock(ctx context.Context, stock *LockStock, tx *ent.Tx) error {
	_stock, ok := h.stocks[*stock.AppGoodID]
	if !ok {
		return fmt.Errorf("invalid stock")
	}

	spotQuantity := _stock.miningGoodStock.SpotQuantity
	locked := _stock.miningGoodStock.Locked
	appReserved := _stock.miningGoodStock.AppReserved

	locked = stock.Locked.Add(locked)
	platformLocked := *stock.Locked
	if stock.AppSpotLocked != nil {
		platformLocked = platformLocked.Sub(*stock.AppSpotLocked)
		appReserved = appReserved.Sub(*stock.AppSpotLocked)
	}
	if platformLocked.Cmp(decimal.NewFromInt(0)) < 0 {
		return fmt.Errorf("invalid appspotlocked")
	}
	if appReserved.Cmp(decimal.NewFromInt(0)) < 0 {
		return fmt.Errorf("invalid appreserved")
	}
	spotQuantity = spotQuantity.Sub(platformLocked)
	if spotQuantity.Cmp(decimal.NewFromInt(0)) < 0 {
		return fmt.Errorf("invalid stock")
	}
	if spotQuantity.Cmp(_stock.miningGoodStock.Total) > 0 {
		return fmt.Errorf("invalid stock")
	}
	if locked.Cmp(decimal.NewFromInt(0)) < 0 {
		return fmt.Errorf("invalid stock")
	}

	if locked.Add(_stock.miningGoodStock.InService).
		Add(_stock.miningGoodStock.WaitStart).
		Add(appReserved).
		Add(spotQuantity).
		Cmp(_stock.miningGoodStock.Total) > 0 {
		return fmt.Errorf("stock exhausted")
	}

	if _, err := mininggoodstockcrud.UpdateSet(
		tx.MiningGoodStock.UpdateOneID(_stock.miningGoodStock.ID),
		&mininggoodstockcrud.Req{
			SpotQuantity: &spotQuantity,
			Locked:       &locked,
			AppReserved:  &appReserved,
		},
	).Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *lockHandler) lockAppStock(ctx context.Context, stock *LockStock, tx *ent.Tx) error {
	_stock, ok := h.stocks[*stock.AppGoodID]
	if !ok {
		return fmt.Errorf("invalid stock")
	}

	spotQuantity := _stock.appGoodStock.SpotQuantity
	locked := _stock.appGoodStock.Locked

	locked = stock.Locked.Add(locked)
	if stock.AppSpotLocked != nil {
		spotQuantity = spotQuantity.Sub(*stock.AppSpotLocked)
	}
	if spotQuantity.Cmp(_stock.appGoodStock.Reserved) > 0 {
		return fmt.Errorf("invalid stock")
	}
	if spotQuantity.Cmp(decimal.NewFromInt(0)) < 0 {
		spotQuantity = decimal.NewFromInt(0)
	}
	if locked.Cmp(decimal.NewFromInt(0)) < 0 {
		return fmt.Errorf("invalid stock")
	}

	if _, err := appstockcrud.UpdateSet(
		tx.AppStock.UpdateOneID(_stock.appGoodStock.ID),
		&appstockcrud.Req{
			SpotQuantity: &spotQuantity,
			Locked:       &locked,
		},
	).Save(ctx); err != nil {
		return err
	}

	return nil
}

func (h *lockHandler) lockAppMiningGoodStock(ctx context.Context, stock *LockStock, tx *ent.Tx) error {
	_stock, ok := h.stocks[*stock.AppGoodID]
	if !ok {
		return fmt.Errorf("invalid stock")
	}

	spotQuantity := _stock.appMiningGoodStock.SpotQuantity
	locked := _stock.appMiningGoodStock.Locked

	locked = stock.Locked.Add(locked)
	if stock.AppSpotLocked != nil {
		spotQuantity = spotQuantity.Sub(*stock.AppSpotLocked)
	}
	if spotQuantity.Cmp(_stock.appMiningGoodStock.Reserved) > 0 {
		return fmt.Errorf("invalid stock")
	}
	if spotQuantity.Cmp(decimal.NewFromInt(0)) < 0 {
		spotQuantity = decimal.NewFromInt(0)
	}
	if locked.Cmp(decimal.NewFromInt(0)) < 0 {
		return fmt.Errorf("invalid stock")
	}

	if _, err := appmininggoodstockcrud.UpdateSet(
		tx.AppMiningGoodStock.UpdateOneID(_stock.appMiningGoodStock.ID),
		&appmininggoodstockcrud.Req{
			SpotQuantity: &spotQuantity,
			Locked:       &locked,
		},
	).Save(ctx); err != nil {
		return err
	}

	return nil
}

func (h *Handler) LockStock(ctx context.Context) error {
	handler := &lockHandler{
		stockAppGoodQuery: &stockAppGoodQuery{
			Handler: h,
		},
		lockOp: &lockopHandler{
			Handler: h,
		},
	}

	if err := handler.getStockAppGoods(ctx); err != nil {
		return err
	}

	return db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		stock := &LockStock{
			EntID:         h.EntID,
			AppGoodID:     h.AppGoodID,
			Locked:        h.Locked,
			AppSpotLocked: h.AppSpotLocked,
		}
		if err := handler.lockAppStock(ctx, stock, tx); err != nil {
			return err
		}
		if handler.stockByMiningPool(*h.AppGoodID) {
			if err := handler.lockAppMiningGoodStock(ctx, stock, tx); err != nil {
				return err
			}
		}
		if err := handler.lockStock(ctx, stock, tx); err != nil {
			return err
		}
		if handler.stockByMiningPool(*h.AppGoodID) {
			if err := handler.lockMiningGoodStock(ctx, stock, tx); err != nil {
				return err
			}
		}
		if err := handler.lockOp.createLocks(ctx, tx); err != nil {
			return err
		}
		return nil
	})
}

func (h *Handler) LockStocks(ctx context.Context) error {
	handler := &lockHandler{
		stockAppGoodQuery: &stockAppGoodQuery{
			Handler: h,
		},
		lockOp: &lockopHandler{
			Handler: h,
		},
	}

	if err := handler.getStockAppGoods(ctx); err != nil {
		return err
	}

	return db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		for _, stock := range h.Stocks {
			if err := handler.lockAppStock(ctx, stock, tx); err != nil {
				return err
			}
			if handler.stockByMiningPool(*stock.AppGoodID) {
				if err := handler.lockAppMiningGoodStock(ctx, stock, tx); err != nil {
					return err
				}
			}
			if err := handler.lockStock(ctx, stock, tx); err != nil {
				return err
			}
			if handler.stockByMiningPool(*stock.AppGoodID) {
				if err := handler.lockMiningGoodStock(ctx, stock, tx); err != nil {
					return err
				}
			}
		}
		if err := handler.lockOp.createLocks(ctx, tx); err != nil {
			return err
		}
		return nil
	})
}
