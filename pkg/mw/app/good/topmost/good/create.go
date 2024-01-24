package topmostgood

import (
	"context"
	"fmt"

	redis2 "github.com/NpoolPlatform/go-service-framework/pkg/redis"
	topmostgoodcrud "github.com/NpoolPlatform/good-middleware/pkg/crud/app/good/topmost/good"
	"github.com/NpoolPlatform/good-middleware/pkg/db"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/app/good/topmost/good"

	"github.com/google/uuid"
)

type createHandler struct {
	*Handler
}

func (h *createHandler) createTopMostGood(ctx context.Context, tx *ent.Tx) error {
	if _, err := topmostgoodcrud.CreateSet(
		tx.TopMostGood.Create(),
		&topmostgoodcrud.Req{
			EntID:        h.EntID,
			AppID:        h.AppID,
			GoodID:       h.GoodID,
			TopMostID:    h.TopMostID,
			AppGoodID:    h.AppGoodID,
			CoinTypeID:   h.CoinTypeID,
			DisplayIndex: h.DisplayIndex,
			Posters:      h.Posters,
			UnitPrice:    h.UnitPrice,
			PackagePrice: h.PackagePrice,
		},
	).Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *Handler) CreateTopMostGood(ctx context.Context) (*npool.TopMostGood, error) {
	if h.UnitPrice == nil && h.PackagePrice == nil {
		return nil, fmt.Errorf("invalid topmostgood price")
	}

	if err := h.GetAppGood(ctx); err != nil {
		return nil, err
	}

	key := fmt.Sprintf("%v:%v:%v:%v", basetypes.Prefix_PrefixCreateTopMostGood, *h.AppID, *h.TopMostID, *h.AppGoodID)
	if err := redis2.TryLock(key, 0); err != nil {
		return nil, err
	}
	defer func() {
		_ = redis2.Unlock(key)
	}()

	h.Conds = &topmostgoodcrud.Conds{
		AppID:     &cruder.Cond{Op: cruder.EQ, Val: *h.AppID},
		TopMostID: &cruder.Cond{Op: cruder.EQ, Val: *h.TopMostID},
		AppGoodID: &cruder.Cond{Op: cruder.EQ, Val: *h.AppGoodID},
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
		if err := handler.createTopMostGood(_ctx, tx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return h.GetTopMostGood(ctx)
}
