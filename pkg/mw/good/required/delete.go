package required

import (
	"context"
	"time"

	requiredcrud "github.com/NpoolPlatform/good-middleware/pkg/crud/good/required"
	"github.com/NpoolPlatform/good-middleware/pkg/db"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent"
)

func (h *Handler) DeleteRequired(ctx context.Context) error {
	info, err := h.GetRequired(ctx)
	if err != nil {
		return err
	}
	if info == nil {
		return nil
	}

	h.ID = &info.ID
	now := uint32(time.Now().Unix())

	return db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if _, err := requiredcrud.UpdateSet(
			cli.RequiredGood.UpdateOneID(*h.ID),
			&requiredcrud.Req{
				DeletedAt: &now,
			},
		).Save(_ctx); err != nil {
			return err
		}
		return nil
	})
}
