package topmost

import (
	"context"

	topmostcrud "github.com/NpoolPlatform/good-middleware/pkg/crud/app/good/topmost"
	"github.com/NpoolPlatform/good-middleware/pkg/db"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent"
)

type updateHandler struct {
	*Handler
}

func (h *updateHandler) updateTopMost(ctx context.Context, tx *ent.Tx) error {
	if _, err := topmostcrud.UpdateSet(
		tx.TopMost.UpdateOneID(*h.ID),
		&topmostcrud.Req{
			Title:     h.Title,
			Message:   h.Message,
			TargetURL: h.TargetURL,
			StartAt:   h.StartAt,
			EndAt:     h.EndAt,
		},
	).Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *Handler) UpdateTopMost(ctx context.Context) error {
	if err := h.formalizeStartEnd(ctx); err != nil {
		return err
	}
	if err := h.checkPromotion(ctx); err != nil {
		return err
	}

	handler := &updateHandler{
		Handler: h,
	}

	return db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		return handler.updateTopMost(_ctx, tx)
	})
}
