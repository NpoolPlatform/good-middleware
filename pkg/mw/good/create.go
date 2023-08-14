package good

import (
	"context"

	redis2 "github.com/NpoolPlatform/go-service-framework/pkg/redis"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/good"
)

type createHandler struct {
	*Handler
}

func (h *createHandler) createExtraInfo(ctx context.Context, tx *ent.Tx) error {
	if _, err := extrainfocrud.CreateSet(
		tx.ExtraInfo.Create(),
		&extrainfocrud.Req{
			GoodID:  h.ID,
			Posters: h.Posters,
			Labels:  h.Labels,
		},
	).Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *createHandler) createReward(ctx context.Context, tx *ent.Tx) error {
	if _, err := rewardcrud.CreateSet(
		tx.GoodReward.Create(),
		&rewardcrud.Req{
			GoodID: h.ID,
		},
	).Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *createHandler) createStock(ctx context.Context, tx *ent.Tx) error {
	if _, err := stockcrud.CreateSet(
		tx.Stock.Create(),
		&stockcrud.Req{
			GoodID: h.ID,
			Total:  h.Total,
		},
	).Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *Handler) CreateGood(ctx context.Context) (*npool.Good, error) {
	key := h.lockKey()
	if err := redis2.TryLock(key, 0); err != nil {
		return nil, err
	}
	defer func() {
		_ = redis2.Unlock(key)
	}()

	h.Conds = &goodcrud.Conds{
		Title: &cruder.Cond{Op: cruder.EQ, Val: *h.Title},
	}
	exist, err := h.ExistGoodConds(ctx)
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

	err := db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		if err := handler.createExtraInfo(_ctx, tx); err != nil {
			return err
		}
		if err := handler.createReward(_ctx, tx); err != nil {
			return err
		}
		if err := handler.createStock(_ctx, tx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return h.GetGood(ctx)
}
