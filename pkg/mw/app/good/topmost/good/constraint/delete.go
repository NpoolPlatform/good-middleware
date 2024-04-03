package constraint

import (
	"context"
	"time"

	constraintcrud "github.com/NpoolPlatform/good-middleware/pkg/crud/app/good/topmost/good/constraint"
	"github.com/NpoolPlatform/good-middleware/pkg/db"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent"
)

type deleteHandler struct {
	*Handler
	now uint32
}

func (h *deleteHandler) deleteConstraint(ctx context.Context, tx *ent.Tx) error {
	if _, err := constraintcrud.UpdateSet(
		tx.TopMostGoodConstraint.UpdateOneID(*h.ID),
		&constraintcrud.Req{
			DeletedAt: &h.now,
		},
	).Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *Handler) DeleteConstraint(ctx context.Context) error {
	info, err := h.GetConstraint(ctx)
	if err != nil {
		return err
	}
	if info == nil {
		return nil
	}

	h.ID = &info.ID
	handler := &deleteHandler{
		Handler: h,
		now:     uint32(time.Now().Unix()),
	}

	return db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		return handler.deleteConstraint(_ctx, tx)
	})
}
