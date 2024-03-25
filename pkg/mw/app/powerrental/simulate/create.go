package appsimulategood

import (
	"context"
	"fmt"

	redis2 "github.com/NpoolPlatform/go-service-framework/pkg/redis"
	appsimulategoodcrud "github.com/NpoolPlatform/good-middleware/pkg/crud/app/powerrental/simulate"
	"github.com/NpoolPlatform/good-middleware/pkg/db"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent"
	appgoodmw "github.com/NpoolPlatform/good-middleware/pkg/mw/app/good"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/app/powerrental/simulate"
	"github.com/shopspring/decimal"

	"github.com/google/uuid"
)

type createHandler struct {
	*Handler
}

func (h *createHandler) checkUnits(ctx context.Context) error {
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
	if h.FixedOrderUnits.Cmp(minOrderAmount) < 0 {
		return fmt.Errorf("units is less than minorderamount")
	}
	if h.FixedOrderUnits.Cmp(maxOrderAmount) > 0 {
		return fmt.Errorf("units is more than maxorderamount")
	}
	if *h.FixedOrderDuration > appgood.MaxOrderDuration || *h.FixedOrderDuration < appgood.MinOrderDuration {
		return fmt.Errorf("invalid duration")
	}
	return nil
}

func (h *createHandler) createSimulate(ctx context.Context, tx *ent.Tx) error {
	if _, err := appsimulategoodcrud.CreateSet(
		tx.AppSimulateGood.Create(),
		&appsimulategoodcrud.Req{
			EntID:              h.EntID,
			AppID:              h.AppID,
			GoodID:             h.GoodID,
			AppGoodID:          h.AppGoodID,
			CoinTypeID:         h.CoinTypeID,
			FixedOrderUnits:    h.FixedOrderUnits,
			FixedOrderDuration: h.FixedOrderDuration,
		},
	).Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *Handler) CreateSimulate(ctx context.Context) (*npool.Simulate, error) {
	if err := h.GetAppGood(ctx); err != nil {
		return nil, err
	}

	key := fmt.Sprintf("%v:%v:%v", basetypes.Prefix_PrefixCreateAppSimulateGood, *h.AppID, *h.AppGoodID)
	if err := redis2.TryLock(key, 0); err != nil {
		return nil, err
	}
	defer func() {
		_ = redis2.Unlock(key)
	}()

	h.Conds = &appsimulategoodcrud.Conds{
		AppID:     &cruder.Cond{Op: cruder.EQ, Val: *h.AppID},
		AppGoodID: &cruder.Cond{Op: cruder.EQ, Val: *h.AppGoodID},
	}
	exist, err := h.ExistSimulateConds(ctx)
	if err != nil {
		return nil, err
	}
	if exist {
		return nil, fmt.Errorf("already exists")
	}

	id := uuid.New()
	if h.EntID == nil {
		h.EntID = &id
	}

	handler := &createHandler{
		Handler: h,
	}

	if err := handler.checkUnits(ctx); err != nil {
		return nil, err
	}

	err = db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		if err := handler.createSimulate(_ctx, tx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return h.GetSimulate(ctx)
}
