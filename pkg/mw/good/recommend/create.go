package recommend

import (
	"context"
	"fmt"

	redis2 "github.com/NpoolPlatform/go-service-framework/pkg/redis"
	extrainfocrud "github.com/NpoolPlatform/good-middleware/pkg/crud/good/extrainfo"
	recommendcrud "github.com/NpoolPlatform/good-middleware/pkg/crud/good/recommend"
	"github.com/NpoolPlatform/good-middleware/pkg/db"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/good/recommend"

	"github.com/google/uuid"
)

type createHandler struct {
	*Handler
}

func (h *createHandler) createRecommend(ctx context.Context, tx *ent.Tx) error {
	if _, err := recommendcrud.CreateSet(
		tx.Recommend.Create(),
		&recommendcrud.Req{
			EntID:          h.EntID,
			AppID:          h.AppID,
			RecommenderID:  h.RecommenderID,
			GoodID:         h.GoodID,
			Message:        h.Message,
			RecommendIndex: h.RecommendIndex,
		},
	).Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *createHandler) updateGoodRecommend(ctx context.Context, tx *ent.Tx) error {
	stm, err := extrainfocrud.SetQueryConds(
		tx.ExtraInfo.Query(),
		&extrainfocrud.Conds{
			GoodID: &cruder.Cond{Op: cruder.EQ, Val: *h.GoodID},
		},
	)
	if err != nil {
		return err
	}
	info, err := stm.Only(ctx)
	if err != nil {
		return err
	}
	recommendCount := info.RecommendCount + 1
	if _, err := extrainfocrud.UpdateSet(
		info.Update(),
		&extrainfocrud.Req{
			RecommendCount: &recommendCount,
		},
	).Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *Handler) CreateRecommend(ctx context.Context) (*npool.Recommend, error) {
	key := fmt.Sprintf("%v:%v:%v:%v", basetypes.Prefix_PrefixRecommendGood, *h.AppID, *h.RecommenderID, *h.GoodID)
	if err := redis2.TryLock(key, 0); err != nil {
		return nil, err
	}
	defer func() {
		_ = redis2.Unlock(key)
	}()

	id := uuid.New()
	if h.EntID == nil {
		h.EntID = &id
	}

	handler := &createHandler{
		Handler: h,
	}

	err := db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		if err := handler.createRecommend(ctx, tx); err != nil {
			return err
		}
		if err := handler.updateGoodRecommend(ctx, tx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return h.GetRecommend(ctx)
}
