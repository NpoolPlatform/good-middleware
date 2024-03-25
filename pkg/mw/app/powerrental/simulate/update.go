package appsimulategood

import (
	"context"
	"fmt"

	appsimulategoodcrud "github.com/NpoolPlatform/good-middleware/pkg/crud/app/powerrental/simulate"
	"github.com/NpoolPlatform/good-middleware/pkg/db"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent"
	appgoodmw "github.com/NpoolPlatform/good-middleware/pkg/mw/app/good"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/app/powerrental/simulate"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type updateHandler struct {
	*Handler
}

func (h *updateHandler) checkUnits(ctx context.Context) error {
	handler, err := appgoodmw.NewHandler(ctx)
	if err != nil {
		return err
	}
	handler.EntID = h.AppGoodID
	appgood, err := handler.GetGood(ctx)
	if err != nil {
		return err
	}
	if appgood == nil {
		return fmt.Errorf("invalid appgood")
	}
	maxOrderAmount, err := decimal.NewFromString(appgood.MaxOrderAmount)
	if err != nil {
		return err
	}
	minOrderAmount, err := decimal.NewFromString(appgood.MinOrderAmount)
	if err != nil {
		return err
	}
	if h.FixedOrderUnits != nil && h.FixedOrderUnits.Cmp(minOrderAmount) < 0 {
		return fmt.Errorf("units is less than minorderamount")
	}
	if h.FixedOrderUnits != nil && h.FixedOrderUnits.Cmp(maxOrderAmount) > 0 {
		return fmt.Errorf("units is more than maxorderamount")
	}
	if h.FixedOrderDuration == nil {
		return nil
	}
	if *h.FixedOrderDuration > appgood.MaxOrderDuration || *h.FixedOrderDuration < appgood.MinOrderDuration {
		return fmt.Errorf("invalid duration")
	}
	return nil
}

func (h *updateHandler) updateSimulate(ctx context.Context, tx *ent.Tx) error {
	if _, err := appsimulategoodcrud.UpdateSet(
		tx.AppSimulateGood.UpdateOneID(*h.ID),
		&appsimulategoodcrud.Req{
			FixedOrderUnits:    h.FixedOrderUnits,
			FixedOrderDuration: h.FixedOrderDuration,
		},
	).Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *Handler) UpdateSimulate(ctx context.Context) (*npool.Simulate, error) {
	info, err := h.GetSimulate(ctx)
	if err != nil {
		return nil, err
	}
	if h.FixedOrderUnits == nil && h.FixedOrderDuration == nil {
		return info, nil
	}

	appgoodID := uuid.MustParse(info.AppGoodID)
	h.AppGoodID = &appgoodID

	handler := &updateHandler{
		Handler: h,
	}

	if err := handler.checkUnits(ctx); err != nil {
		return nil, err
	}

	err = db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		if err := handler.updateSimulate(_ctx, tx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return h.GetSimulate(ctx)
}
