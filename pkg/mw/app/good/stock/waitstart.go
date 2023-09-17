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

type waitStartHandler struct {
	*lockopHandler
}

func (h *waitStartHandler) waitStartStock(ctx context.Context, tx *ent.Tx) error {
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

	locked := info.Locked
	sold := info.Sold
	waitStart := info.WaitStart

	waitStart = h.lock.Units.Add(waitStart)
	locked = locked.Sub(h.lock.Units)
	sold = h.lock.Units.Add(sold)

	if locked.Cmp(decimal.NewFromInt(0)) < 0 {
		return fmt.Errorf("invalid stock")
	}
	if waitStart.Cmp(decimal.NewFromInt(0)) < 0 {
		return fmt.Errorf("invalid stock")
	}

	if locked.Add(info.InService).
		Add(info.WaitStart).
		Add(info.AppReserved).
		Add(info.SpotQuantity).
		Cmp(info.Total) > 0 {
		return fmt.Errorf("stock exhausted")
	}

	if _, err := stockcrud.UpdateSet(
		tx.Stock.UpdateOneID(info.ID),
		&stockcrud.Req{
			Locked:    &locked,
			WaitStart: &waitStart,
			Sold:      &sold,
		},
	).Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *waitStartHandler) waitStartAppStock(ctx context.Context, tx *ent.Tx) error {
	info, err := tx.
		AppStock.
		Query().
		Where(
			entappstock.ID(h.lock.AppStockID),
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
	locked := info.Locked
	sold := info.Sold
	waitStart := info.WaitStart

	waitStart = h.lock.Units.Add(waitStart)
	locked = locked.Sub(h.lock.Units)
	sold = h.lock.Units.Add(sold)

	if locked.Cmp(decimal.NewFromInt(0)) < 0 {
		return fmt.Errorf("invalid stock")
	}
	if waitStart.Cmp(decimal.NewFromInt(0)) < 0 {
		return fmt.Errorf("invalid stock")
	}

	if _, err := appstockcrud.UpdateSet(
		tx.AppStock.UpdateOneID(info.ID),
		&appstockcrud.Req{
			Locked:    &locked,
			WaitStart: &waitStart,
			Sold:      &sold,
		},
	).Save(ctx); err != nil {
		return err
	}

	return nil
}

func (h *Handler) WaitStartStock(ctx context.Context) (*npool.Stock, error) {
	handler := &waitStartHandler{
		lockopHandler: &lockopHandler{
			Handler: h,
			state:   types.AppStockLockState_AppStockWaitStart.Enum(),
		},
	}

	if err := handler.getLock(ctx); err != nil {
		return nil, err
	}

	err := db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		if err := handler.waitStartAppStock(ctx, tx); err != nil {
			return err
		}
		if err := handler.waitStartStock(ctx, tx); err != nil {
			return err
		}
		if err := handler.updateLock(ctx, tx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return h.GetStock(ctx)
}
