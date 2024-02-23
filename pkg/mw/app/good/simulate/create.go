package appsimulategood

import (
	"context"
	"fmt"

	redis2 "github.com/NpoolPlatform/go-service-framework/pkg/redis"
	appsimulategoodcrud "github.com/NpoolPlatform/good-middleware/pkg/crud/app/good/simulate"
	"github.com/NpoolPlatform/good-middleware/pkg/db"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/app/good/simulate"

	"github.com/google/uuid"
)

type createHandler struct {
	*Handler
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
