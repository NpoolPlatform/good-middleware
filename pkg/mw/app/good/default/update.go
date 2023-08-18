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
)

type updateHandler struct {
	*Handler
}

func (h *updateHandler) updateDefault(ctx context.Context, tx *ent.Tx) error {
	if _, err := appdefaultgoodcrud.UpdateSet(
		tx.AppDefaultGood.UpdateOneID(*h.ID),
		&appdefaultgoodcrud.Req{
			GoodID:    h.GoodID,
			AppGoodID: h.AppGoodID,
		},
	).Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *Handler) UpdateDefault(ctx context.Context) (*npool.Default, error) {
	if h.AppGoodID != nil {
		key := fmt.Sprintf("%v:%v:%v:%v", basetypes.Prefix_PrefixCreateAppDefaultGood, *h.AppID, *h.GoodID, *h.AppGoodID)
		if err := redis2.TryLock(key, 0); err != nil {
			return nil, err
		}
		defer func() {
			_ = redis2.Unlock(key)
		}()

		h.Conds = &appdefaultgoodcrud.Conds{
			AppID:     &cruder.Cond{Op: cruder.EQ, Val: *h.AppID},
			GoodID:    &cruder.Cond{Op: cruder.EQ, Val: *h.GoodID},
			AppGoodID: &cruder.Cond{Op: cruder.EQ, Val: *h.AppGoodID},
		}
		exist, err := h.ExistDefaultConds(ctx)
		if err != nil {
			return nil, err
		}
		if exist {
			return nil, fmt.Errorf("already exists")
		}
	}

	handler := &updateHandler{
		Handler: h,
	}

	err := db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		if err := handler.updateDefault(_ctx, tx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return h.GetDefault(ctx)
}