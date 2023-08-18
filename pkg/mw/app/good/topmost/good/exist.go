package topmostgood

import (
	"context"

	topmostgoodcrud "github.com/NpoolPlatform/good-middleware/pkg/crud/app/good/topmost/good"
	"github.com/NpoolPlatform/good-middleware/pkg/db"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent"
	enttopmostgood "github.com/NpoolPlatform/good-middleware/pkg/db/ent/topmostgood"
)

func (h *Handler) ExistTopMostGood(ctx context.Context) (bool, error) {
	exist := false
	var err error

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		exist, err = cli.
			TopMostGood.
			Query().
			Where(
				enttopmostgood.ID(*h.ID),
				enttopmostgood.DeletedAt(0),
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

func (h *Handler) ExistTopMostGoodConds(ctx context.Context) (bool, error) {
	exist := false

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		stm, err := topmostgoodcrud.SetQueryConds(cli.TopMostGood.Query(), h.Conds)
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
