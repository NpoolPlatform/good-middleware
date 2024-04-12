package powerrental

import (
	"context"
	"fmt"

	apppowerrentalcrud "github.com/NpoolPlatform/good-middleware/pkg/crud/app/powerrental"
	"github.com/NpoolPlatform/good-middleware/pkg/db"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent"
	entapppowerrental "github.com/NpoolPlatform/good-middleware/pkg/db/ent/apppowerrental"
)

func (h *Handler) ExistPowerRental(ctx context.Context) (exist bool, err error) {
	if h.ID == nil && h.EntID == nil && h.AppGoodID == nil {
		return false, fmt.Errorf("invalid appgoodid")
	}
	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		stm := cli.AppPowerRental.Query()
		if h.ID != nil {
			stm.Where(entapppowerrental.ID(*h.ID))
		}
		if h.EntID != nil {
			stm.Where(entapppowerrental.EntID(*h.EntID))
		}
		if h.AppGoodID != nil {
			stm.Where(entapppowerrental.AppGoodID(*h.AppGoodID))
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

func (h *Handler) ExistPowerRentalConds(ctx context.Context) (exist bool, err error) {
	if err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		stm, err := apppowerrentalcrud.SetQueryConds(cli.AppPowerRental.Query(), h.AppPowerRentalConds)
		if err != nil {
			return err
		}
		exist, err = stm.Exist(_ctx)
		return err
	}); err != nil {
		return false, err
	}
	return exist, nil
}
