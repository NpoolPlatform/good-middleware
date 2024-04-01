package topmost

import (
	"context"

	topmostcrud "github.com/NpoolPlatform/good-middleware/pkg/crud/app/good/topmost"
	"github.com/NpoolPlatform/good-middleware/pkg/db"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/app/good/topmost"
)

type updateHandler struct {
	*Handler
}

func (h *updateHandler) updateTopMost(ctx context.Context, tx *ent.Tx) error {
	if _, err := topmostcrud.UpdateSet(
		tx.TopMost.UpdateOneID(*h.ID),
		&topmostcrud.Req{
			Title:   h.Title,
			Message: h.Message,
			StartAt: h.StartAt,
			EndAt:   h.EndAt,
		},
	).Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *Handler) UpdateTopMost(ctx context.Context) (*npool.TopMost, error) {
	if err := h.formalizeStartEnd(ctx); err != nil {
		return nil, err
	}
	if err := h.checkPromotion(ctx); err != nil {
		return nil, err
	}

	handler := &updateHandler{
		Handler: h,
	}

	err := db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		if err := handler.updateTopMost(_ctx, tx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return h.GetTopMost(ctx)
}
