package fee

import (
	"context"
	"fmt"

	appfeecrud "github.com/NpoolPlatform/good-middleware/pkg/crud/app/fee"
	"github.com/NpoolPlatform/good-middleware/pkg/db"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent"
	entappfee "github.com/NpoolPlatform/good-middleware/pkg/db/ent/appfee"
)

func (h *Handler) ExistFee(ctx context.Context) (exist bool, err error) {
	if h.ID == nil && h.EntID == nil && h.AppGoodID == nil {
		return false, fmt.Errorf("invalid appgoodid")
	}
	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		stm := cli.AppFee.Query()
		if h.ID != nil {
			stm.Where(entappfee.ID(*h.ID))
		}
		if h.EntID != nil {
			stm.Where(entappfee.EntID(*h.EntID))
		}
		if h.AppGoodID != nil {
			stm.Where(entappfee.AppGoodID(*h.AppGoodID))
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

func (h *Handler) ExistFeeConds(ctx context.Context) (exist bool, err error) {
	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		stm, err := appfeecrud.SetQueryConds(cli.AppFee.Query(), h.AppFeeConds)
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
