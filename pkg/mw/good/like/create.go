package like

import (
	"context"
	"fmt"

	redis2 "github.com/NpoolPlatform/go-service-framework/pkg/redis"
	extrainfocrud "github.com/NpoolPlatform/good-middleware/pkg/crud/good/extrainfo"
	likecrud "github.com/NpoolPlatform/good-middleware/pkg/crud/good/like"
	"github.com/NpoolPlatform/good-middleware/pkg/db"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/good/like"

	"github.com/google/uuid"
)

type createHandler struct {
	*Handler
}

func (h *createHandler) createLike(ctx context.Context, tx *ent.Tx) error {
	if _, err := likecrud.CreateSet(
		tx.Like.Create(),
		&likecrud.Req{
			ID:        h.ID,
			AppID:     h.AppID,
			UserID:    h.UserID,
			AppGoodID: h.AppGoodID,
			Like:      h.Like,
		},
	).Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *createHandler) addGoodLike(ctx context.Context, tx *ent.Tx) error {
	appGood, err := tx.AppGood.Get(ctx, *h.AppGoodID)
	if err != nil {
		return err
	}
	if appGood == nil {
		return fmt.Errorf("app good not found %v", *h.AppGoodID)
	}

	stm, err := extrainfocrud.SetQueryConds(
		tx.ExtraInfo.Query(),
		&extrainfocrud.Conds{
			GoodID: &cruder.Cond{Op: cruder.EQ, Val: appGood.GoodID},
		},
	)
	if err != nil {
		return err
	}
	info, err := stm.Only(ctx)
	if err != nil {
		return err
	}
	if *h.Like {
		info.Likes += 1
	} else {
		info.Dislikes += 1
	}
	if _, err := extrainfocrud.UpdateSet(
		info.Update(),
		&extrainfocrud.Req{
			Likes:    &info.Likes,
			Dislikes: &info.Dislikes,
		},
	).Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *Handler) CreateLike(ctx context.Context) (*npool.Like, error) {
	key := fmt.Sprintf("%v:%v:%v:%v", basetypes.Prefix_PrefixLikeGood, *h.AppID, *h.UserID, *h.AppGoodID)
	if err := redis2.TryLock(key, 0); err != nil {
		return nil, err
	}
	defer func() {
		_ = redis2.Unlock(key)
	}()

	h.Conds = &likecrud.Conds{
		AppID:     &cruder.Cond{Op: cruder.EQ, Val: *h.AppID},
		UserID:    &cruder.Cond{Op: cruder.EQ, Val: *h.UserID},
		AppGoodID: &cruder.Cond{Op: cruder.EQ, Val: *h.AppGoodID},
	}
	exist, err := h.ExistLikeConds(ctx)
	if err != nil {
		return nil, err
	}
	if exist {
		return nil, fmt.Errorf("already exists")
	}

	id := uuid.New()
	if h.ID == nil {
		h.ID = &id
	}

	handler := &createHandler{
		Handler: h,
	}

	err = db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		if err := handler.createLike(ctx, tx); err != nil {
			return err
		}
		if err := handler.addGoodLike(ctx, tx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return h.GetLike(ctx)
}
