package comment

import (
	"context"

	commentcrud "github.com/NpoolPlatform/good-middleware/pkg/crud/app/good/comment"
	"github.com/NpoolPlatform/good-middleware/pkg/db"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent"
)

type updateHandler struct {
	*Handler
}

func (h *updateHandler) updateComment(ctx context.Context, tx *ent.Tx) error {
	if _, err := commentcrud.UpdateSet(
		tx.Comment.UpdateOneID(*h.ID),
		&commentcrud.Req{
			Content:    h.Content,
			Anonymous:  h.Anonymous,
			Hide:       h.Hide,
			HideReason: h.HideReason,
		},
	).Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *Handler) UpdateComment(ctx context.Context) error {
	handler := &updateHandler{
		Handler: h,
	}

	return db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		return handler.updateComment(ctx, tx)
	})
}
