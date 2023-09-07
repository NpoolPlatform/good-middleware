package appgood

import (
	"context"
	"fmt"

	appgoodcrud "github.com/NpoolPlatform/good-middleware/pkg/crud/app/good"
	"github.com/NpoolPlatform/good-middleware/pkg/db"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/app/good"

	"github.com/google/uuid"
)

type updateHandler struct {
	*Handler
}

func (h *updateHandler) updateAppGood(ctx context.Context, tx *ent.Tx) error {
	if _, err := appgoodcrud.UpdateSet(
		tx.AppGood.UpdateOneID(*h.ID),
		&appgoodcrud.Req{
			GoodName:               h.GoodName,
			Online:                 h.Online,
			Visible:                h.Visible,
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
			TechniqueFeeRatio:      h.TechniqueFeeRatio,
			ElectricityFeeRatio:    h.ElectricityFeeRatio,
		},
	).Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *Handler) UpdateGood(ctx context.Context) (*npool.Good, error) {
	info, err := h.GetGood(ctx)
	if err != nil {
		return nil, err
	}
	if info == nil {
		return nil, fmt.Errorf("invalid good")
	}

	goodID, err := uuid.Parse(info.GoodID)
	if err != nil {
		return nil, err
	}

	h.GoodID = &goodID

	if err := h.checkPrice(ctx); err != nil {
		return nil, err
	}

	handler := &updateHandler{
		Handler: h,
	}

	err = db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		if err := handler.updateAppGood(ctx, tx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return h.GetGood(ctx)
}
