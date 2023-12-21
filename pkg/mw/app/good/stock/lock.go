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
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/app/good/stock"

	"github.com/shopspring/decimal"
)

type lockHandler struct {
	*lockopHandler
}

func (h *lockHandler) lockStock(ctx context.Context, tx *ent.Tx) error {
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

	spotQuantity := info.SpotQuantity
	locked := info.Locked
	appReserved := info.AppReserved

	locked = h.Locked.Add(locked)
	platformLocked := *h.Locked
	if h.AppSpotLocked != nil {
		platformLocked = platformLocked.Sub(*h.AppSpotLocked)
		appReserved = appReserved.Sub(*h.AppSpotLocked)
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
	if spotQuantity.Cmp(info.Total) > 0 {
		return fmt.Errorf("invalid stock")
	}
	if locked.Cmp(decimal.NewFromInt(0)) < 0 {
		return fmt.Errorf("invalid stock")
	}

	if locked.Add(info.InService).
		Add(info.WaitStart).
		Add(appReserved).
		Add(spotQuantity).
		Cmp(info.Total) > 0 {
		return fmt.Errorf("stock exhausted")
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

func (h *lockHandler) lockAppStock(ctx context.Context, tx *ent.Tx) error {
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

	if *h.AppID != info.AppID {
		return fmt.Errorf("invalid appid")
	}
	if *h.GoodID != info.GoodID {
		return fmt.Errorf("invalid goodid")
	}
	if *h.AppGoodID != info.AppGoodID {
		return fmt.Errorf("invalid appgoodid")
	}
	spotQuantity := info.SpotQuantity
	locked := info.Locked

	locked = h.Locked.Add(locked)
	if h.AppSpotLocked != nil {
		spotQuantity = spotQuantity.Sub(*h.AppSpotLocked)
	}
	if spotQuantity.Cmp(info.Reserved) > 0 {
		return fmt.Errorf("invalid stock")
	}
	if spotQuantity.Cmp(decimal.NewFromInt(0)) < 0 {
		spotQuantity = decimal.NewFromInt(0)
	}
	if locked.Cmp(decimal.NewFromInt(0)) < 0 {
		return fmt.Errorf("invalid stock")
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

func (h *Handler) LockStock(ctx context.Context) (*npool.Stock, error) {
	handler := &lockHandler{
		lockopHandler: &lockopHandler{
			Handler: h,
		},
	}
	err := db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		if err := handler.lockAppStock(ctx, tx); err != nil {
			return err
		}
		if err := handler.lockStock(ctx, tx); err != nil {
			return err
		}
		if err := handler.createLock(ctx, tx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return h.GetStock(ctx)
}
