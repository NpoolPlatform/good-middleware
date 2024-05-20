//nolint:dupl
package powerrental

import (
	"context"

	wlog "github.com/NpoolPlatform/go-service-framework/pkg/wlog"
	"github.com/NpoolPlatform/good-middleware/pkg/db"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent"
	entappgoodbase "github.com/NpoolPlatform/good-middleware/pkg/db/ent/appgoodbase"
	entappmininggoodstock "github.com/NpoolPlatform/good-middleware/pkg/db/ent/appmininggoodstock"
	entapppowerrental "github.com/NpoolPlatform/good-middleware/pkg/db/ent/apppowerrental"
	entappstock "github.com/NpoolPlatform/good-middleware/pkg/db/ent/appstock"
	entgoodbase "github.com/NpoolPlatform/good-middleware/pkg/db/ent/goodbase"
	entmininggoodstock "github.com/NpoolPlatform/good-middleware/pkg/db/ent/mininggoodstock"
	entpowerrental "github.com/NpoolPlatform/good-middleware/pkg/db/ent/powerrental"
	entstock "github.com/NpoolPlatform/good-middleware/pkg/db/ent/stock"
)

type powerRentalAppGoodQueryHandler struct {
	*Handler
	_ent powerRental
}

func (h *powerRentalAppGoodQueryHandler) getPowerRental(ctx context.Context, cli *ent.Client, must bool) (err error) {
	if h._ent.powerRental, err = cli.
		PowerRental.
		Query().
		Where(
			entpowerrental.GoodID(*h.AppGoodBaseReq.GoodID),
			entpowerrental.DeletedAt(0),
		).
		Only(ctx); err != nil {
		if ent.IsNotFound(err) && !must {
			return nil
		}
		return wlog.WrapError(err)
	}
	return nil
}

func (h *powerRentalAppGoodQueryHandler) getGoodBase(ctx context.Context, cli *ent.Client, must bool) (err error) {
	if h._ent.goodBase, err = cli.
		GoodBase.
		Query().
		Where(
			entgoodbase.EntID(h._ent.powerRental.GoodID),
			entgoodbase.DeletedAt(0),
		).Only(ctx); err != nil {
		if ent.IsNotFound(err) && !must {
			return nil
		}
		return wlog.WrapError(err)
	}
	return nil
}

func (h *powerRentalAppGoodQueryHandler) getGoodStock(ctx context.Context, cli *ent.Client, must bool) (err error) {
	if h._ent.stock, err = cli.
		Stock.
		Query().
		Where(
			entstock.GoodID(h._ent.powerRental.GoodID),
			entstock.DeletedAt(0),
		).Only(ctx); err != nil {
		if ent.IsNotFound(err) && !must {
			return nil
		}
		return wlog.WrapError(err)
	}
	return nil
}

func (h *powerRentalAppGoodQueryHandler) getMiningGoodStocks(ctx context.Context, cli *ent.Client) (err error) {
	if h._ent.miningGoodStocks, err = cli.
		MiningGoodStock.
		Query().
		Where(
			entmininggoodstock.GoodStockID(h._ent.stock.EntID),
			entmininggoodstock.DeletedAt(0),
		).All(ctx); err != nil {
		return wlog.WrapError(err)
	}
	return nil
}

func (h *powerRentalAppGoodQueryHandler) _getPowerRentalGood(ctx context.Context, must bool) (err error) {
	if h.AppGoodBaseReq.GoodID == nil {
		return wlog.Errorf("invalid goodid")
	}

	return db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := h.getPowerRental(_ctx, cli, must); err != nil {
			return wlog.WrapError(err)
		}
		if h._ent.powerRental == nil {
			return nil
		}
		if err := h.getGoodBase(_ctx, cli, must); err != nil {
			return wlog.WrapError(err)
		}
		if err := h.getGoodStock(_ctx, cli, must); err != nil {
			return wlog.WrapError(err)
		}
		if h._ent.stock == nil {
			return nil
		}
		return h.getMiningGoodStocks(_ctx, cli)
	})
}

func (h *powerRentalAppGoodQueryHandler) getAppPowerRental(ctx context.Context, cli *ent.Client, must bool) (err error) {
	stm := cli.AppPowerRental.Query().Where(entapppowerrental.DeletedAt(0))
	if h.ID != nil {
		stm.Where(entapppowerrental.ID(*h.ID))
	}
	if h.EntID != nil {
		stm.Where(entapppowerrental.EntID(*h.EntID))
	}
	if h.AppGoodID != nil {
		stm.Where(entapppowerrental.AppGoodID(*h.AppGoodID))
	}
	h._ent.appPowerRental, err = stm.Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) && !must {
			return nil
		}
		return wlog.WrapError(err)
	}
	return nil
}

func (h *powerRentalAppGoodQueryHandler) getAppGoodBase(ctx context.Context, cli *ent.Client, must bool) (err error) {
	if h._ent.appGoodBase, err = cli.
		AppGoodBase.
		Query().
		Where(
			entappgoodbase.EntID(h._ent.appPowerRental.AppGoodID),
			entappgoodbase.DeletedAt(0),
		).Only(ctx); err != nil {
		if ent.IsNotFound(err) && !must {
			return nil
		}
		return wlog.WrapError(err)
	}
	return nil
}

func (h *powerRentalAppGoodQueryHandler) getAppGoodStock(ctx context.Context, cli *ent.Client, must bool) (err error) {
	if h._ent.appGoodStock, err = cli.
		AppStock.
		Query().
		Where(
			entappstock.AppGoodID(h._ent.appPowerRental.AppGoodID),
			entappstock.DeletedAt(0),
		).Only(ctx); err != nil {
		if ent.IsNotFound(err) && !must {
			return nil
		}
		return wlog.WrapError(err)
	}
	return nil
}

func (h *powerRentalAppGoodQueryHandler) getAppMiningGoodStocks(ctx context.Context, cli *ent.Client) (err error) {
	if h._ent.appMiningGoodStocks, err = cli.
		AppMiningGoodStock.
		Query().
		Where(
			entappmininggoodstock.AppGoodStockID(h._ent.appGoodStock.EntID),
			entappmininggoodstock.DeletedAt(0),
		).All(ctx); err != nil {
		return wlog.WrapError(err)
	}
	return nil
}

func (h *powerRentalAppGoodQueryHandler) _getAppPowerRentalAppGood(ctx context.Context, must bool) (err error) {
	if h.ID == nil && h.EntID == nil && h.AppGoodID == nil {
		return wlog.Errorf("invalid appgoodid")
	}

	if err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := h.getAppPowerRental(_ctx, cli, must); err != nil {
			return wlog.WrapError(err)
		}
		if h._ent.appPowerRental == nil {
			return nil
		}
		if err := h.getAppGoodBase(_ctx, cli, must); err != nil {
			return wlog.WrapError(err)
		}
		if err := h.getAppGoodStock(_ctx, cli, must); err != nil {
			return wlog.WrapError(err)
		}
		if h._ent.appGoodStock == nil {
			return nil
		}
		return h.getAppMiningGoodStocks(_ctx, cli)
	}); err != nil {
		return wlog.WrapError(err)
	}

	if h._ent.appGoodBase == nil {
		if !must {
			return nil
		}
		return wlog.Errorf("invalid appgoodbase")
	}

	h.AppGoodBaseReq.GoodID = &h._ent.appGoodBase.GoodID
	return h._getPowerRentalGood(ctx, must)
}

//nolint:unused
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

func (h *Handler) QueryPowerRentalEnt(ctx context.Context) (PowerRental, error) {
	handler := &powerRentalAppGoodQueryHandler{
		Handler: h,
	}
	if err := handler.requireAppPowerRentalAppGood(ctx); err != nil {
		return nil, err
	}
	return &handler._ent, nil
}
