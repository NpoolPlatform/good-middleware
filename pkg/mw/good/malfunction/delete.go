package goodmalfunction

import (
	"context"
	"time"

	malfunctioncrud "github.com/NpoolPlatform/good-middleware/pkg/crud/good/malfunction"
	"github.com/NpoolPlatform/good-middleware/pkg/db"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent"
)

func (h *Handler) DeleteMalfunction(ctx context.Context) error {
	info, err := h.GetMalfunction(ctx)
	if err != nil {
		return err
	}
	if info == nil {
		return nil
	}

	h.ID = &info.ID
	now := uint32(time.Now().Unix())

	return db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if _, err := malfunctioncrud.UpdateSet(
			cli.GoodMalfunction.UpdateOneID(*h.ID),
			&malfunctioncrud.Req{
				DeletedAt: &now,
			},
		).Save(_ctx); err != nil {
			return err
		}
		return nil
	})
}
