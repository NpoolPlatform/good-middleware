package topmostgood

import (
	"context"

	"github.com/NpoolPlatform/good-middleware/pkg/db"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent"
	entappgood "github.com/NpoolPlatform/good-middleware/pkg/db/ent/appgood"
)

func (h *Handler) GetAppGood(ctx context.Context) error {
	return db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		info, err := cli.
			AppGood.
			Query().
			Where(
				entappgood.ID(*h.AppGoodID),
				entappgood.DeletedAt(0),
			).
			Only(ctx)
		if err != nil {
			return err
		}

		h.AppID = &info.AppID
		h.GoodID = &info.GoodID
		return nil
	})
}
