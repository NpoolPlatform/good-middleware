package comment

import (
	"context"

	commentcrud "github.com/NpoolPlatform/good-middleware/pkg/crud/good/comment"
	"github.com/NpoolPlatform/good-middleware/pkg/db"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/good/comment"
)

type updateHandler struct {
	*Handler
}

func (h *updateHandler) updateComment(ctx context.Context, tx *ent.Tx) error {
	if _, err := commentcrud.UpdateSet(
		tx.Comment.UpdateOneID(*h.ID),
		&commentcrud.Req{
			Content: h.Content,
		},
	).Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *Handler) UpdateComment(ctx context.Context) (*npool.Comment, error) {
	handler := &updateHandler{
		Handler: h,
	}

	err := db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		if err := handler.updateComment(ctx, tx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return h.GetComment(ctx)
}
