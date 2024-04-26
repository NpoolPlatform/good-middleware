package poster

import (
	"context"

	postercrud "github.com/NpoolPlatform/good-middleware/pkg/crud/device/poster"
	"github.com/NpoolPlatform/good-middleware/pkg/db"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent"
	entposter "github.com/NpoolPlatform/good-middleware/pkg/db/ent/deviceposter"
)

func (h *Handler) ExistPoster(ctx context.Context) (exist bool, err error) {
	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		exist, err = cli.
			DevicePoster.
			Query().
			Where(
				entposter.EntID(*h.EntID),
				entposter.DeletedAt(0),
			).
			Exist(_ctx)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return false, err
	}
	return exist, nil
}

func (h *Handler) ExistPosterConds(ctx context.Context) (exist bool, err error) {
	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		stm, err := postercrud.SetQueryConds(cli.DevicePoster.Query(), h.PosterConds)
		if err != nil {
			return err
		}
		exist, err = stm.Exist(_ctx)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return false, err
	}
	return exist, nil
}