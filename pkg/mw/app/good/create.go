package appgood

import (
	"context"
	"fmt"

	redis2 "github.com/NpoolPlatform/go-service-framework/pkg/redis"
	appgoodcrud "github.com/NpoolPlatform/good-middleware/pkg/crud/app/good"
	appstockcrud "github.com/NpoolPlatform/good-middleware/pkg/crud/app/good/stock"
	"github.com/NpoolPlatform/good-middleware/pkg/db"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/app/good"

	"github.com/google/uuid"
)

type createHandler struct {
	*Handler
}

func (h *createHandler) createStock(ctx context.Context, tx *ent.Tx) error {
	if _, err := appstockcrud.CreateSet(
		tx.AppStock.Create(),
		&appstockcrud.Req{
			AppID:     h.AppID,
			GoodID:    h.GoodID,
			AppGoodID: h.ID,
		},
	).Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *createHandler) createAppGood(ctx context.Context, tx *ent.Tx) error {
	if _, err := appgoodcrud.CreateSet(
		tx.AppGood.Create(),
		&appgoodcrud.Req{
			ID:                     h.ID,
			AppID:                  h.AppID,
			GoodID:                 h.GoodID,
			GoodName:               h.GoodName,
			Price:                  h.Price,
			DisplayIndex:           h.DisplayIndex,
			PurchaseLimit:          h.PurchaseLimit,
			SaleStartAt:            h.SaleStartAt,
			SaleEndAt:              h.SaleEndAt,
			ServiceStartAt:         h.ServiceStartAt,
			Descriptions:           h.Descriptions,
			GoodBanner:             h.GoodBanner,
			DisplayNames:           h.DisplayNames,
			EnablePurchase:         h.EnablePurchase,
			EnableProductPage:      h.EnableProductPage,
			CancelMode:             h.CancelMode,
			UserPurchaseLimit:      h.UserPurchaseLimit,
			DisplayColors:          h.DisplayColors,
			CancellableBeforeStart: h.CancellableBeforeStart,
			ProductPage:            h.ProductPage,
			EnableSetCommission:    h.EnableSetCommission,
			Posters:                h.Posters,
		},
	).Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *Handler) CreateGood(ctx context.Context) (*npool.Good, error) {
	handler := &createHandler{
		Handler: h,
	}

	key := fmt.Sprintf("%v:%v:%v", basetypes.Prefix_PrefixCreateAppGood, *h.AppID, *h.GoodID)
	if err := redis2.TryLock(key, 0); err != nil {
		return nil, err
	}
	defer func() {
		_ = redis2.Unlock(key)
	}()

	id := uuid.New()
	if h.ID == nil {
		h.ID = &id
	}

	// Here we do not need to check if app-good exists due to one good can authorized to one app multiple times

	err := db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		if err := handler.createStock(ctx, tx); err != nil {
			return err
		}
		if err := handler.createAppGood(ctx, tx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return h.GetGood(ctx)
}
