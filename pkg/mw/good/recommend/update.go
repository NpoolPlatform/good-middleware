package recommend

import (
	"context"
	"fmt"

	extrainfocrud "github.com/NpoolPlatform/good-middleware/pkg/crud/good/extrainfo"
	recommendcrud "github.com/NpoolPlatform/good-middleware/pkg/crud/good/recommend"
	"github.com/NpoolPlatform/good-middleware/pkg/db"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/good/recommend"

	"github.com/google/uuid"
)

type updateHandler struct {
	*Handler
}

func (h *updateHandler) updateRecommend(ctx context.Context, tx *ent.Tx) error {
	if _, err := recommendcrud.UpdateSet(
		tx.Recommend.UpdateOneID(*h.ID),
		&recommendcrud.Req{
			Message:        h.Message,
			RecommendIndex: h.RecommendIndex,
		},
	).Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *updateHandler) updateGoodRecommend(ctx context.Context, tx *ent.Tx) error {
	stm, err := extrainfocrud.SetQueryConds(tx.ExtraInfo.Query(), &extrainfocrud.Conds{
		GoodID: &cruder.Cond{Op: cruder.EQ, Val: *h.GoodID},
	})
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

func (h *Handler) UpdateRecommend(ctx context.Context) (*npool.Recommend, error) {
	info, err := h.GetRecommend(ctx)
	if err != nil {
		return nil, err
	}
	if info == nil {
		return nil, fmt.Errorf("invalid recommend")
	}

	goodID := uuid.MustParse(info.GoodID)
	h.GoodID = &goodID
	handler := &updateHandler{
		Handler: h,
	}

	err = db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		if err := handler.updateRecommend(ctx, tx); err != nil {
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
