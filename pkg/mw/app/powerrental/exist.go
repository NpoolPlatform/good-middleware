package powerrental

import (
	"context"
	"fmt"

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
