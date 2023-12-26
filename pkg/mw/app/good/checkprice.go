package appgood

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/good-middleware/pkg/db"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent"
	entgood "github.com/NpoolPlatform/good-middleware/pkg/db/ent/good"
)

func (h *Handler) checkUnitPrice(ctx context.Context) error {
	if h.UnitPrice == nil {
		return nil
	}
	return db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		info, err := cli.
			Good.
			Query().
			Where(
				entgood.EntID(*h.GoodID),
				entgood.DeletedAt(0),
			).
			Only(_ctx)
		if err != nil {
			return err
		}
		if h.UnitPrice.Cmp(info.UnitPrice) < 0 {
			return fmt.Errorf("invalid price")
		}
		return nil
	})
}

// TODO: we also need to check package price
