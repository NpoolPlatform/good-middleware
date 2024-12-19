package pledge

import (
	"context"

	wlog "github.com/NpoolPlatform/go-service-framework/pkg/wlog"
	"github.com/NpoolPlatform/good-middleware/pkg/db"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent"
	entgoodbase "github.com/NpoolPlatform/good-middleware/pkg/db/ent/goodbase"
	entgoodcoinreward "github.com/NpoolPlatform/good-middleware/pkg/db/ent/goodcoinreward"
	entgoodreward "github.com/NpoolPlatform/good-middleware/pkg/db/ent/goodreward"
	entpledge "github.com/NpoolPlatform/good-middleware/pkg/db/ent/pledge"
)

type pledgeGoodQueryHandler struct {
	*Handler
	pledge      *ent.Pledge
	goodBase    *ent.GoodBase
	goodReward  *ent.GoodReward
	coinRewards []*ent.GoodCoinReward
}

func (h *pledgeGoodQueryHandler) getPledge(ctx context.Context, cli *ent.Client, must bool) (err error) {
	stm := cli.Pledge.Query()
	if h.ID != nil {
		stm.Where(entpledge.ID(*h.ID))
	}
	if h.EntID != nil {
		stm.Where(entpledge.EntID(*h.EntID))
	}
	if h.GoodID != nil {
		stm.Where(entpledge.GoodID(*h.GoodID))
	}
	if h.pledge, err = stm.Only(ctx); err != nil {
		if ent.IsNotFound(err) && !must {
			return nil
		}
		return wlog.WrapError(err)
	}
	return nil
}

func (h *pledgeGoodQueryHandler) getGoodBase(ctx context.Context, cli *ent.Client, must bool) (err error) {
	if h.goodBase, err = cli.
		GoodBase.
		Query().
		Where(
			entgoodbase.EntID(h.pledge.GoodID),
			entgoodbase.DeletedAt(0),
		).Only(ctx); err != nil {
		if ent.IsNotFound(err) && !must {
			return nil
		}
		return wlog.WrapError(err)
	}
	return nil
}

func (h *pledgeGoodQueryHandler) getGoodReward(ctx context.Context, cli *ent.Client, must bool) (err error) {
	if h.goodReward, err = cli.
		GoodReward.
		Query().
		Where(
			entgoodreward.GoodID(h.pledge.GoodID),
			entgoodreward.DeletedAt(0),
		).Only(ctx); err != nil {
		if ent.IsNotFound(err) && !must {
			return nil
		}
		return wlog.WrapError(err)
	}
	return nil
}

func (h *pledgeGoodQueryHandler) getGoodCoinRewards(ctx context.Context, cli *ent.Client) (err error) {
	h.coinRewards, err = cli.
		GoodCoinReward.
		Query().
		Where(
			entgoodcoinreward.GoodID(h.pledge.GoodID),
			entgoodcoinreward.DeletedAt(0),
		).
		All(ctx)
	return wlog.WrapError(err)
}

func (h *pledgeGoodQueryHandler) _getPledgeGood(ctx context.Context, must bool) (err error) {
	if h.ID == nil && h.EntID == nil && h.GoodID == nil {
		return wlog.Errorf("invalid id")
	}

	return db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := h.getPledge(_ctx, cli, must); err != nil {
			return wlog.WrapError(err)
		}
		if h.pledge == nil {
			return nil
		}
		if err := h.getGoodBase(_ctx, cli, must); err != nil {
			return wlog.WrapError(err)
		}
		if err := h.getGoodReward(_ctx, cli, must); err != nil {
			return wlog.WrapError(err)
		}
		if err := h.getGoodCoinRewards(_ctx, cli); err != nil {
			return wlog.WrapError(err)
		}
		return nil
	})
}

func (h *pledgeGoodQueryHandler) getPledgeGood(ctx context.Context) error {
	return h._getPledgeGood(ctx, false)
}

func (h *pledgeGoodQueryHandler) requirePledgeGood(ctx context.Context) error {
	return h._getPledgeGood(ctx, true)
}
