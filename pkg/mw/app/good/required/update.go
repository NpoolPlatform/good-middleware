package required

import (
	"context"

	requiredcrud "github.com/NpoolPlatform/good-middleware/pkg/crud/app/good/required"
	"github.com/NpoolPlatform/good-middleware/pkg/db"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent"
)

func (h *Handler) UpdateRequired(ctx context.Context) error {
	return db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if _, err := requiredcrud.UpdateSet(
			cli.RequiredAppGood.UpdateOneID(*h.ID),
			&requiredcrud.Req{
				Must: h.Must,
			},
		).Save(_ctx); err != nil {
			return err
		}
		return nil
	})
}
