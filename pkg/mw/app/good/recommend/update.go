package recommend

import (
	"context"
	"fmt"

	recommendcrud "github.com/NpoolPlatform/good-middleware/pkg/crud/app/good/recommend"
	"github.com/NpoolPlatform/good-middleware/pkg/db"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent"
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

func (h *Handler) UpdateRecommend(ctx context.Context) error {
	info, err := h.GetRecommend(ctx)
	if err != nil {
		return err
	}
	if info == nil {
		return fmt.Errorf("invalid recommend")
	}

	h.ID = &info.ID
	handler := &updateHandler{
		Handler: h,
	}

	return db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		return handler.updateRecommend(ctx, tx)
	})
}
