package topmostgood

import (
	"context"
	"fmt"

	topmostgoodcrud "github.com/NpoolPlatform/good-middleware/pkg/crud/app/good/topmost/good"
	"github.com/NpoolPlatform/good-middleware/pkg/db"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent"
)

type updateHandler struct {
	*Handler
}

func (h *updateHandler) updateTopMostGood(ctx context.Context, tx *ent.Tx) error {
	if _, err := topmostgoodcrud.UpdateSet(
		tx.TopMostGood.UpdateOneID(*h.ID),
		&topmostgoodcrud.Req{
			DisplayIndex: h.DisplayIndex,
			UnitPrice:    h.UnitPrice,
		},
	).Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *Handler) UpdateTopMostGood(ctx context.Context) error {
	info, err := h.GetTopMostGood(ctx)
	if err != nil {
		return err
	}
	if info == nil {
		return fmt.Errorf("invalid topmostgood")
	}

	h.ID = &info.ID
	handler := &updateHandler{
		Handler: h,
	}
	return db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		return handler.updateTopMostGood(_ctx, tx)
	})
}
