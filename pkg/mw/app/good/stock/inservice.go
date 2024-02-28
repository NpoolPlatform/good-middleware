//nolint:dupl
package appstock

import (
	"context"
	"fmt"

	appstockcrud "github.com/NpoolPlatform/good-middleware/pkg/crud/app/good/stock"
	stockcrud "github.com/NpoolPlatform/good-middleware/pkg/crud/good/stock"
	"github.com/NpoolPlatform/good-middleware/pkg/db"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent"
	entappstock "github.com/NpoolPlatform/good-middleware/pkg/db/ent/appstock"
	entstock "github.com/NpoolPlatform/good-middleware/pkg/db/ent/stock"
	types "github.com/NpoolPlatform/message/npool/basetypes/good/v1"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/app/good/stock"

	"github.com/shopspring/decimal"
)

type inServiceHandler struct {
	*lockopHandler
}

func (h *inServiceHandler) inServiceStock(ctx context.Context, lock *ent.AppStockLock, tx *ent.Tx) error {
	info, err := tx.
		Stock.
		Query().
		Where(
			entstock.GoodID(*h.GoodID),
			entstock.DeletedAt(0),
		).
		ForUpdate().
		Only(ctx)
	if err != nil {
		return err
	}
	if info == nil {
		return fmt.Errorf("invalid stock")
	}

	inService := info.InService
	waitStart := info.WaitStart

	inService = lock.Units.Add(inService)
	waitStart = waitStart.Sub(lock.Units)

	if waitStart.Cmp(decimal.NewFromInt(0)) < 0 {
		return fmt.Errorf("invalid stock")
	}

	if waitStart.Add(inService).
		Add(info.Locked).
		Add(info.AppReserved).
		Add(info.SpotQuantity).
		Cmp(info.Total) > 0 {
		return fmt.Errorf("stock exhausted")
	}

	if _, err := stockcrud.UpdateSet(
		tx.Stock.UpdateOneID(info.ID),
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
	info, err := tx.
		AppStock.
		Query().
		Where(
			entappstock.EntID(*h.EntID),
			entappstock.DeletedAt(0),
		).
		ForUpdate().
		Only(ctx)
	if err != nil {
		return err
	}
	if info == nil {
		return fmt.Errorf("invalid appstock")
	}

	h.GoodID = &info.GoodID
	inService := info.InService
	waitStart := info.WaitStart

	inService = lock.Units.Add(inService)
	waitStart = waitStart.Sub(lock.Units)

	if waitStart.Cmp(decimal.NewFromInt(0)) < 0 {
		return fmt.Errorf("invalid stock")
	}

	if _, err := appstockcrud.UpdateSet(
		tx.AppStock.UpdateOneID(info.ID),
		&appstockcrud.Req{
			InService: &inService,
			WaitStart: &waitStart,
		},
	).Save(ctx); err != nil {
		return err
	}

	return nil
}

func (h *Handler) InServiceStock(ctx context.Context) (*npool.Stock, error) {
	handler := &inServiceHandler{
		lockopHandler: &lockopHandler{
			Handler: h,
			state:   types.AppStockLockState_AppStockInService.Enum(),
		},
	}

	if err := handler.getLocks(ctx); err != nil {
		return nil, err
	}

	err := db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		for _, lock := range handler.lockopHandler.locks {
			if err := handler.inServiceAppStock(ctx, lock, tx); err != nil {
				return err
			}
			if err := handler.inServiceStock(ctx, lock, tx); err != nil {
				return err
			}
		}
		if err := handler.updateLocks(ctx, tx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return h.GetStock(ctx)
}
