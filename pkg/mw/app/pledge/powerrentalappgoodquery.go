package pledge

import (
	"context"

	wlog "github.com/NpoolPlatform/go-service-framework/pkg/wlog"
	"github.com/NpoolPlatform/good-middleware/pkg/db"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent"
	entappgoodbase "github.com/NpoolPlatform/good-middleware/pkg/db/ent/appgoodbase"
	entapppledge "github.com/NpoolPlatform/good-middleware/pkg/db/ent/apppledge"
	entgoodbase "github.com/NpoolPlatform/good-middleware/pkg/db/ent/goodbase"
	entpledge "github.com/NpoolPlatform/good-middleware/pkg/db/ent/pledge"
)

type pledgeAppGoodQueryHandler struct {
	*Handler
	_ent pledge
}

func (h *pledgeAppGoodQueryHandler) getPledge(ctx context.Context, cli *ent.Client, must bool) (err error) {
	if h._ent.pledge, err = cli.
		Pledge.
		Query().
		Where(
			entpledge.GoodID(*h.AppGoodBaseReq.GoodID),
			entpledge.DeletedAt(0),
		).
		Only(ctx); err != nil {
		if ent.IsNotFound(err) && !must {
			return nil
		}
		return wlog.WrapError(err)
	}
	return nil
}

func (h *pledgeAppGoodQueryHandler) getGoodBase(ctx context.Context, cli *ent.Client, must bool) (err error) {
	if h._ent.goodBase, err = cli.
		GoodBase.
		Query().
		Where(
			entgoodbase.EntID(h._ent.pledge.GoodID),
			entgoodbase.DeletedAt(0),
		).Only(ctx); err != nil {
		if ent.IsNotFound(err) && !must {
			return nil
		}
		return wlog.WrapError(err)
	}
	return nil
}

func (h *pledgeAppGoodQueryHandler) _getPledgeGood(ctx context.Context, must bool) (err error) {
	if h.AppGoodBaseReq.GoodID == nil {
		return wlog.Errorf("invalid goodid")
	}

	return db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := h.getPledge(_ctx, cli, must); err != nil {
			return wlog.WrapError(err)
		}
		if h._ent.pledge == nil {
			return nil
		}
		if err := h.getGoodBase(_ctx, cli, must); err != nil {
			return wlog.WrapError(err)
		}
		return nil
	})
}

func (h *pledgeAppGoodQueryHandler) getAppPledge(ctx context.Context, cli *ent.Client, must bool) (err error) {
	stm := cli.AppPledge.Query().Where(entapppledge.DeletedAt(0))
	if h.ID != nil {
		stm.Where(entapppledge.ID(*h.ID))
	}
	if h.EntID != nil {
		stm.Where(entapppledge.EntID(*h.EntID))
	}
	if h.AppGoodID != nil {
		stm.Where(entapppledge.AppGoodID(*h.AppGoodID))
	}
	h._ent.appPledge, err = stm.Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) && !must {
			return nil
		}
		return wlog.WrapError(err)
	}
	return nil
}

func (h *pledgeAppGoodQueryHandler) getAppGoodBase(ctx context.Context, cli *ent.Client, must bool) (err error) {
	if h._ent.appGoodBase, err = cli.
		AppGoodBase.
		Query().
		Where(
			entappgoodbase.EntID(h._ent.appPledge.AppGoodID),
			entappgoodbase.DeletedAt(0),
		).Only(ctx); err != nil {
		if ent.IsNotFound(err) && !must {
			return nil
		}
		return wlog.WrapError(err)
	}
	return nil
}

func (h *pledgeAppGoodQueryHandler) _getAppPledgeAppGood(ctx context.Context, must bool) (err error) {
	if h.ID == nil && h.EntID == nil && h.AppGoodID == nil {
		return wlog.Errorf("invalid appgoodid")
	}

	if err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := h.getAppPledge(_ctx, cli, must); err != nil {
			return wlog.WrapError(err)
		}
		if h._ent.appPledge == nil {
			return nil
		}
		if err := h.getAppGoodBase(_ctx, cli, must); err != nil {
			return wlog.WrapError(err)
		}

		return nil
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
	return h._getPledgeGood(ctx, must)
}

//nolint:unused
func (h *pledgeAppGoodQueryHandler) getPledgeGood(ctx context.Context) error {
	return h._getPledgeGood(ctx, false)
}

func (h *pledgeAppGoodQueryHandler) requirePledgeGood(ctx context.Context) error {
	return h._getPledgeGood(ctx, true)
}

func (h *pledgeAppGoodQueryHandler) getAppPledgeAppGood(ctx context.Context) error {
	return h._getAppPledgeAppGood(ctx, false)
}

func (h *pledgeAppGoodQueryHandler) requireAppPledgeAppGood(ctx context.Context) error {
	return h._getAppPledgeAppGood(ctx, true)
}

func (h *Handler) QueryPledgeEnt(ctx context.Context) (Pledge, error) {
	handler := &pledgeAppGoodQueryHandler{
		Handler: h,
	}
	if err := handler.requireAppPledgeAppGood(ctx); err != nil {
		return nil, wlog.WrapError(err)
	}
	return &handler._ent, nil
}
