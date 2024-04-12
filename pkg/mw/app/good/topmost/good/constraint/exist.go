package constraint

import (
	"context"

	constraintcrud "github.com/NpoolPlatform/good-middleware/pkg/crud/app/good/topmost/good/constraint"
	"github.com/NpoolPlatform/good-middleware/pkg/db"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent"
	entconstraint "github.com/NpoolPlatform/good-middleware/pkg/db/ent/topmostgoodconstraint"
)

func (h *Handler) ExistConstraint(ctx context.Context) (exist bool, err error) {
	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		exist, err = cli.
			TopMostGoodConstraint.
			Query().
			Where(
				entconstraint.EntID(*h.EntID),
				entconstraint.DeletedAt(0),
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

func (h *Handler) ExistConstraintConds(ctx context.Context) (exist bool, err error) {
	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		stm, err := constraintcrud.SetQueryConds(cli.TopMostGoodConstraint.Query(), h.ConstraintConds)
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
