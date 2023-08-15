package recommend

import (
	"context"
	"time"

	recommendcrud "github.com/NpoolPlatform/good-middleware/pkg/crud/good/recommend"
	"github.com/NpoolPlatform/good-middleware/pkg/db"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/good/recommend"
)

type deleteHandler struct {
	*Handler
}

func (h *deleteHandler) deleteRecommend(ctx context.Context, tx *ent.Tx) error {
	now := uint32(time.Now().Unix())
	if _, err := recommendcrud.UpdateSet(
		tx.Recommend.UpdateOneID(*h.ID),
		&recommendcrud.Req{
			DeletedAt: &now,
		},
	).Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *Handler) DeleteRecommend(ctx context.Context) (*npool.Recommend, error) {
	info, err := h.GetRecommend(ctx)
	if err != nil {
		return nil, err
	}
	if info == nil {
		return nil, nil
	}

	handler := &deleteHandler{
		Handler: h,
	}

	err = db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		if err := handler.deleteRecommend(ctx, tx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}
