package appdefaultgood

import (
	"context"
	"fmt"

	redis2 "github.com/NpoolPlatform/go-service-framework/pkg/redis"
	appdefaultgoodcrud "github.com/NpoolPlatform/good-middleware/pkg/crud/app/good/default"
	"github.com/NpoolPlatform/good-middleware/pkg/db"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/app/good/default"

	"github.com/google/uuid"
)

type createHandler struct {
	*queryAppGoodHandler
}

func (h *createHandler) createDefault(ctx context.Context, tx *ent.Tx) error {
	if _, err := appdefaultgoodcrud.CreateSet(
		tx.AppDefaultGood.Create(),
		&appdefaultgoodcrud.Req{
			EntID:      h.EntID,
			AppID:      h.AppID,
			GoodID:     h.GoodID,
			AppGoodID:  h.AppGoodID,
			CoinTypeID: h.CoinTypeID,
			GoodType:   h.GoodType,
		},
	).Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *Handler) CreateDefault(ctx context.Context) (*npool.Default, error) {
	handler := &createHandler{
		queryAppGoodHandler: &queryAppGoodHandler{
			Handler: h,
		},
	}

	if err := handler.getAppGood(ctx); err != nil {
		return nil, err
	}

	key := fmt.Sprintf("%v:%v:%v", basetypes.Prefix_PrefixCreateAppDefaultGood, *h.AppID, *h.CoinTypeID)
	if err := redis2.TryLock(key, 0); err != nil {
		return nil, err
	}
	defer func() {
		_ = redis2.Unlock(key)
	}()

	h.Conds = &appdefaultgoodcrud.Conds{
		AppID:      &cruder.Cond{Op: cruder.EQ, Val: *h.AppID},
		CoinTypeID: &cruder.Cond{Op: cruder.EQ, Val: *h.CoinTypeID},
	}
	exist, err := h.ExistDefaultConds(ctx)
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
		if err := handler.createDefault(_ctx, tx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return h.GetDefault(ctx)
}
