package constraint

import (
	"context"
	"fmt"

	constraintcrud "github.com/NpoolPlatform/good-middleware/pkg/crud/app/good/topmost/constraint"
	"github.com/NpoolPlatform/good-middleware/pkg/db"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent"
)

type updateHandler struct {
	*Handler
}

func (h *updateHandler) updateConstraint(ctx context.Context, tx *ent.Tx) error {
	if _, err := constraintcrud.UpdateSet(
		tx.TopMostConstraint.UpdateOneID(*h.ID),
		&constraintcrud.Req{
			TargetValue: h.TargetValue,
			Index:       h.Index,
		},
	).Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *Handler) UpdateConstraint(ctx context.Context) error {
	info, err := h.GetConstraint(ctx)
	if err != nil {
		return err
	}
	if info == nil {
		return fmt.Errorf("invalid constraint")
	}

	h.ID = &info.ID
	handler := &updateHandler{
		Handler: h,
	}
	return db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		return handler.updateConstraint(_ctx, tx)
	})
}
