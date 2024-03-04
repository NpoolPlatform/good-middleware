package appsimulategood

import (
	"context"

	appsimulategoodcrud "github.com/NpoolPlatform/good-middleware/pkg/crud/app/good/simulate"
	"github.com/NpoolPlatform/good-middleware/pkg/db"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent"
	entappsimulategood "github.com/NpoolPlatform/good-middleware/pkg/db/ent/appsimulategood"
)

func (h *Handler) ExistSimulate(ctx context.Context) (bool, error) {
	exist := false
	var err error

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		exist, err = cli.
			AppSimulateGood.
			Query().
			Where(
				entappsimulategood.EntID(*h.EntID),
				entappsimulategood.DeletedAt(0),
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

func (h *Handler) ExistSimulateConds(ctx context.Context) (bool, error) {
	exist := false

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		stm, err := appsimulategoodcrud.SetQueryConds(cli.AppSimulateGood.Query(), h.Conds)
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
