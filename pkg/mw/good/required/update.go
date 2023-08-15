package required

import (
	"context"

	requiredcrud "github.com/NpoolPlatform/good-middleware/pkg/crud/good/required"
	"github.com/NpoolPlatform/good-middleware/pkg/db"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/good/required"
)

func (h *Handler) UpdateRequired(ctx context.Context) (*npool.Required, error) {
	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if _, err := requiredcrud.UpdateSet(
			cli.RequiredGood.UpdateOneID(*h.ID),
			&requiredcrud.Req{
				Must:       h.Must,
				Commission: h.Commission,
			},
		).Save(_ctx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return h.GetRequired(ctx)
}
