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

type expireHandler struct {
	*stockAppGoodQuery
	lockOp *lockopHandler
}

func (h *expireHandler) expireStock(ctx context.Context, lock *ent.AppStockLock, tx *ent.Tx) error {
	stock, ok := h.stocks[lock.AppGoodID]
	if !ok {
		return fmt.Errorf("invalid stock")
	}

	inService := stock.stock.InService
	spotQuantity := stock.stock.SpotQuantity

	inService = inService.Sub(lock.Units)
	spotQuantity = spotQuantity.Add(lock.Units)

	if inService.Cmp(decimal.NewFromInt(0)) < 0 {
		return fmt.Errorf("invalid stock")
	}

	if inService.Add(stock.stock.WaitStart).
		Add(stock.stock.Locked).
		Add(stock.stock.AppReserved).
		Add(stock.stock.SpotQuantity).
		Cmp(stock.stock.Total) > 0 {
		return fmt.Errorf("stock exhausted")
	}

	if _, err := stockcrud.UpdateSet(
		tx.Stock.UpdateOneID(stock.stock.ID),
		&stockcrud.Req{
			InService:    &inService,
			SpotQuantity: &spotQuantity,
		},
	).Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *expireHandler) expireAppStock(ctx context.Context, lock *ent.AppStockLock, tx *ent.Tx) error {
	stock, ok := h.stocks[lock.AppGoodID]
	if !ok {
		return fmt.Errorf("invalid stock")
	}

	inService := stock.appGoodStock.InService
	inService = inService.Sub(lock.Units)
	if inService.Cmp(decimal.NewFromInt(0)) < 0 {
		return fmt.Errorf("invalid stock")
	}

	if _, err := appstockcrud.UpdateSet(
		tx.AppStock.UpdateOneID(stock.appGoodStock.ID),
		&appstockcrud.Req{
			InService: &inService,
		},
	).Save(ctx); err != nil {
		return err
	}

	return nil
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
		return err
	}
	h.Stocks = handler.lockOp.lock2Stocks()
	if err := handler.getStockAppGoods(ctx); err != nil {
		return err
	}

	return db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		for _, lock := range handler.lockOp.locks {
			if err := handler.expireAppStock(ctx, lock, tx); err != nil {
				return err
			}
			if err := handler.expireStock(ctx, lock, tx); err != nil {
				return err
			}
		}
		if err := handler.lockOp.updateLocks(ctx, tx); err != nil {
			return err
		}
		return nil
	})
}
