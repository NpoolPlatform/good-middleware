package powerrental

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/good-middleware/pkg/db"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent"
	entgoodbase "github.com/NpoolPlatform/good-middleware/pkg/db/ent/goodbase"
	entpowerrental "github.com/NpoolPlatform/good-middleware/pkg/db/ent/powerrental"
)

type powerRentalAppGoodQueryHandler struct {
	*Handler
	powerRental *ent.PowerRental
	goodBase    *ent.GoodBase
}

func (h *powerRentalAppGoodQueryHandler) _getPowerRentalGood(ctx context.Context, must bool) (err error) {
	if h.AppGoodBaseReq.GoodID == nil {
		return fmt.Errorf("invalid good id")
	}

	return db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		stm := cli.PowerRental.Query()
		if h.AppGoodBaseReq.GoodID != nil {
			stm.Where(entpowerrental.GoodID(*h.AppGoodBaseReq.GoodID))
		}
		if h.powerRental, err = stm.Only(_ctx); err != nil {
			if ent.IsNotFound(err) && !must {
				return nil
			}
			return err
		}
		if h.goodBase, err = cli.
			GoodBase.
			Query().
			Where(
				entgoodbase.EntID(h.powerRental.GoodID),
				entgoodbase.DeletedAt(0),
			).Only(_ctx); err != nil {
			if ent.IsNotFound(err) && !must {
				return nil
			}
			return err
		}
		return nil
	})
}

func (h *powerRentalAppGoodQueryHandler) getPowerRentalGood(ctx context.Context) error {
	return h._getPowerRentalGood(ctx, false)
}

func (h *powerRentalAppGoodQueryHandler) requirePowerRentalGood(ctx context.Context) error {
	return h._getPowerRentalGood(ctx, true)
}
