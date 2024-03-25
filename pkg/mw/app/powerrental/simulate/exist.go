package appsimulatepowerrental

import (
	"context"

	appsimulatepowerrentalcrud "github.com/NpoolPlatform/good-middleware/pkg/crud/app/powerrental/simulate"
	"github.com/NpoolPlatform/good-middleware/pkg/db"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent"
	entappsimulatepowerrental "github.com/NpoolPlatform/good-middleware/pkg/db/ent/appsimulatepowerrental"
)

func (h *Handler) ExistSimulate(ctx context.Context) (exist bool, err error) {
	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		exist, err = cli.
			AppSimulatePowerRental.
			Query().
			Where(
				entappsimulatepowerrental.EntID(*h.EntID),
				entappsimulatepowerrental.DeletedAt(0),
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

func (h *Handler) ExistSimulateConds(ctx context.Context) (exist bool, err error) {
	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		stm, err := appsimulatepowerrentalcrud.SetQueryConds(cli.AppSimulatePowerRental.Query(), h.AppSimulatePowerRentalConds)
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
