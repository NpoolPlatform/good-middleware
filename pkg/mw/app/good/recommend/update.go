package recommend

import (
	"context"

	recommendcrud "github.com/NpoolPlatform/good-middleware/pkg/crud/app/good/recommend"
	"github.com/NpoolPlatform/good-middleware/pkg/db"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/app/good/recommend"
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

func (h *Handler) UpdateRecommend(ctx context.Context) (*npool.Recommend, error) {
	handler := &updateHandler{
		Handler: h,
	}

	err := db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		if err := handler.updateRecommend(ctx, tx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return h.GetRecommend(ctx)
}
