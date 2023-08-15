package required

import (
	"context"
	"time"

	requiredcrud "github.com/NpoolPlatform/good-middleware/pkg/crud/good/required"
	"github.com/NpoolPlatform/good-middleware/pkg/db"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/good/required"
)

func (h *Handler) DeleteRequired(ctx context.Context) (*npool.Required, error) {
	info, err := h.GetRequired(ctx)
	if err != nil {
		return nil, err
	}
	if info == nil {
		return nil, nil
	}

	now := uint32(time.Now().Unix())
	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
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
	if err != nil {
		return nil, err
	}

	return info, nil
}
