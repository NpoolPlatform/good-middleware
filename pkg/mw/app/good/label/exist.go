package label

import (
	"context"

	labelcrud "github.com/NpoolPlatform/good-middleware/pkg/crud/app/good/label"
	"github.com/NpoolPlatform/good-middleware/pkg/db"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent"
	entlabel "github.com/NpoolPlatform/good-middleware/pkg/db/ent/appgoodlabel"
)

func (h *Handler) ExistLabel(ctx context.Context) (exist bool, err error) {
	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		exist, err = cli.
			AppGoodLabel.
			Query().
			Where(
				entlabel.EntID(*h.EntID),
				entlabel.DeletedAt(0),
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

func (h *Handler) ExistLabelConds(ctx context.Context) (exist bool, err error) {
	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		stm, err := labelcrud.SetQueryConds(cli.AppGoodLabel.Query(), h.LabelConds)
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
