package powerrental

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/good-middleware/pkg/db"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent"
	entgoodbase "github.com/NpoolPlatform/good-middleware/pkg/db/ent/goodbase"
	entpowerrental "github.com/NpoolPlatform/good-middleware/pkg/db/ent/powerrental"
)

type powerRentalGoodQueryHandler struct {
	*Handler
	powerRental *ent.PowerRental
	goodBase    *ent.GoodBase
}

func (h *powerRentalGoodQueryHandler) _getPowerRentalGood(ctx context.Context, must bool) (err error) {
	if h.ID == nil && h.EntID == nil && h.GoodID == nil {
		return fmt.Errorf("invalid id")
	}

	return db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		stm := cli.PowerRental.Query()
		if h.ID != nil {
			stm.Where(entpowerrental.ID(*h.ID))
		}
		if h.EntID != nil {
			stm.Where(entpowerrental.EntID(*h.EntID))
		}
		if h.GoodID != nil {
			stm.Where(entpowerrental.GoodID(*h.GoodID))
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

func (h *powerRentalGoodQueryHandler) getPowerRentalGood(ctx context.Context) error {
	return h._getPowerRentalGood(ctx, false)
}

func (h *powerRentalGoodQueryHandler) requirePowerRentalGood(ctx context.Context) error {
	return h._getPowerRentalGood(ctx, true)
}
