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

type unlockHandler struct {
	*lockopHandler
}

func (h *unlockHandler) unlockStock(ctx context.Context, tx *ent.Tx) error {
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
		return fmt.Errorf("stock not found")
	}

	spotQuantity := info.SpotQuantity
	locked := info.Locked
	appReserved := info.AppReserved

	locked = locked.Sub(h.lock.Units)
	platformLocked := h.lock.Units.Sub(h.lock.AppSpotUnits)
	appReserved = h.lock.AppSpotUnits.Add(appReserved)
	if platformLocked.Cmp(decimal.NewFromInt(0)) < 0 {
		return fmt.Errorf("invalid appspotlocked")
	}
	spotQuantity = platformLocked.Add(spotQuantity)
	if locked.Cmp(decimal.NewFromInt(0)) < 0 {
		return fmt.Errorf("invalid locked")
	}
	if appReserved.Cmp(decimal.NewFromInt(0)) < 0 {
		return fmt.Errorf("invalid appreserved")
	}
	if spotQuantity.Cmp(decimal.NewFromInt(0)) < 0 {
		return fmt.Errorf("invalid stock")
	}
	if spotQuantity.Cmp(info.Total) > 0 {
		return fmt.Errorf("invalid stock")
	}

	if locked.Add(info.InService).
		Add(info.WaitStart).
		Add(appReserved).
		Add(spotQuantity).
		Cmp(info.Total) > 0 {
		return fmt.Errorf("invalid stock")
	}

	if _, err := stockcrud.UpdateSet(
		tx.Stock.UpdateOneID(info.ID),
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

func (h *unlockHandler) unlockAppStock(ctx context.Context, tx *ent.Tx) error {
	info, err := tx.
		AppStock.
		Query().
		Where(
			entappstock.ID(*h.ID),
			entappstock.DeletedAt(0),
		).
		ForUpdate().
		Only(ctx)
	if err != nil {
		return err
	}
	if info == nil {
		return fmt.Errorf("stock not found")
	}

	h.GoodID = &info.GoodID
	locked := info.Locked
	spotQuantity := info.SpotQuantity

	locked = locked.Sub(h.lock.Units)
	spotQuantity = h.lock.AppSpotUnits.Add(spotQuantity)
	if locked.Cmp(decimal.NewFromInt(0)) < 0 {
		return fmt.Errorf("invalid locked")
	}
	if spotQuantity.Cmp(info.Reserved) > 0 {
		return fmt.Errorf("invalid spotquantity")
	}
	if spotQuantity.Cmp(decimal.NewFromInt(0)) < 0 {
		spotQuantity = decimal.NewFromInt(0)
	}

	if _, err := appstockcrud.UpdateSet(
		tx.AppStock.UpdateOneID(info.ID),
		&appstockcrud.Req{
			SpotQuantity: &spotQuantity,
			Locked:       &locked,
		},
	).Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *Handler) UnlockStock(ctx context.Context) (*npool.Stock, error) {
	handler := &unlockHandler{
		lockopHandler: &lockopHandler{
			Handler: h,
			state:   types.AppStockLockState_AppStockRollback.Enum(),
		},
	}

	if err := handler.getLock(ctx); err != nil {
		if ent.IsNotFound(err) && h.Rollback != nil && *h.Rollback {
			return nil, nil
		}
		return nil, err
	}

	err := db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		if err := handler.unlockAppStock(ctx, tx); err != nil {
			return err
		}
		if err := handler.unlockStock(ctx, tx); err != nil {
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
