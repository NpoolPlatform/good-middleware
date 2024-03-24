package powerrental

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/good-middleware/pkg/db"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent"
	entappgoodbase "github.com/NpoolPlatform/good-middleware/pkg/db/ent/appgoodbase"
	entapppowerrental "github.com/NpoolPlatform/good-middleware/pkg/db/ent/apppowerrental"
	entgoodbase "github.com/NpoolPlatform/good-middleware/pkg/db/ent/goodbase"
	entpowerrental "github.com/NpoolPlatform/good-middleware/pkg/db/ent/powerrental"
)

type powerRentalAppGoodQueryHandler struct {
	*Handler
	powerRental    *ent.PowerRental
	goodBase       *ent.GoodBase
	appGoodBase    *ent.AppGoodBase
	appPowerRental *ent.AppPowerRental
}

func (h *powerRentalAppGoodQueryHandler) _getPowerRentalGood(ctx context.Context, must bool) (err error) {
	if h.AppGoodBaseReq.GoodID == nil {
		return fmt.Errorf("invalid goodid")
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

func (h *powerRentalAppGoodQueryHandler) _getAppPowerRentalAppGood(ctx context.Context, must bool) (err error) {
	if h.ID == nil && h.EntID == nil && h.AppGoodID == nil {
		return fmt.Errorf("invalid appgoodid")
	}

	if err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
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
		h.appPowerRental, err = stm.Only(_ctx)
		if err != nil {
			if ent.IsNotFound(err) && !must {
				return nil
			}
			return err
		}
		if h.appGoodBase, err = cli.
			AppGoodBase.
			Query().
			Where(
				entappgoodbase.EntID(h.appPowerRental.AppGoodID),
				entappgoodbase.DeletedAt(0),
			).Only(_ctx); err != nil {
			if ent.IsNotFound(err) && !must {
				return nil
			}
			return err
		}
		return nil
	}); err != nil {
		return err
	}

	h.AppGoodBaseReq.GoodID = &h.appGoodBase.GoodID
	return h._getPowerRentalGood(ctx, must)
}

func (h *powerRentalAppGoodQueryHandler) getPowerRentalGood(ctx context.Context) error {
	return h._getPowerRentalGood(ctx, false)
}

func (h *powerRentalAppGoodQueryHandler) requirePowerRentalGood(ctx context.Context) error {
	return h._getPowerRentalGood(ctx, true)
}

func (h *powerRentalAppGoodQueryHandler) getAppPowerRentalAppGood(ctx context.Context) error {
	return h._getAppPowerRentalAppGood(ctx, false)
}

func (h *powerRentalAppGoodQueryHandler) requireAppPowerRentalAppGood(ctx context.Context) error {
	return h._getAppPowerRentalAppGood(ctx, true)
}
