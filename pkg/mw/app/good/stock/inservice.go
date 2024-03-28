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

type inServiceHandler struct {
	*stockAppGoodQuery
	lockOp *lockopHandler
}

func (h *inServiceHandler) inServiceStock(ctx context.Context, lock *ent.AppStockLock, tx *ent.Tx) error {
	inService := h.stock.InService
	waitStart := h.stock.WaitStart

	inService = lock.Units.Add(inService)
	waitStart = waitStart.Sub(lock.Units)

	if waitStart.Cmp(decimal.NewFromInt(0)) < 0 {
		return fmt.Errorf("invalid stock")
	}

	if waitStart.Add(inService).
		Add(h.stock.Locked).
		Add(h.stock.AppReserved).
		Add(h.stock.SpotQuantity).
		Cmp(h.stock.Total) > 0 {
		return fmt.Errorf("stock exhausted")
	}

	if _, err := stockcrud.UpdateSet(
		tx.Stock.UpdateOneID(h.stock.ID),
		&stockcrud.Req{
			InService: &inService,
			WaitStart: &waitStart,
		},
	).Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *inServiceHandler) inServiceAppStock(ctx context.Context, lock *ent.AppStockLock, tx *ent.Tx) error {
	inService := h.appGoodStock.InService
	waitStart := h.appGoodStock.WaitStart

	inService = lock.Units.Add(inService)
	waitStart = waitStart.Sub(lock.Units)

	if waitStart.Cmp(decimal.NewFromInt(0)) < 0 {
		return fmt.Errorf("invalid stock")
	}

	if _, err := appstockcrud.UpdateSet(
		tx.AppStock.UpdateOneID(h.appGoodStock.ID),
		&appstockcrud.Req{
			InService: &inService,
			WaitStart: &waitStart,
		},
	).Save(ctx); err != nil {
		return err
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
		return err
	}
	if err := handler.requireStockAppGood(ctx); err != nil {
		return err
	}

	return db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		for _, lock := range handler.lockOp.locks {
			if err := handler.inServiceAppStock(ctx, lock, tx); err != nil {
				return err
			}
			if err := handler.inServiceStock(ctx, lock, tx); err != nil {
				return err
			}
		}
		if err := handler.lockOp.updateLocks(ctx, tx); err != nil {
			return err
		}
		return nil
	})
}
