package constraint

import (
	"context"

	"entgo.io/ent/dialect/sql"

	logger "github.com/NpoolPlatform/go-service-framework/pkg/logger"
	wlog "github.com/NpoolPlatform/go-service-framework/pkg/wlog"
	"github.com/NpoolPlatform/good-middleware/pkg/db"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent"
	types "github.com/NpoolPlatform/message/npool/basetypes/good/v1"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/app/good/topmost/good/constraint"

	"github.com/shopspring/decimal"
)

type queryHandler struct {
	*baseQueryHandler
	stmCount *ent.TopMostGoodConstraintSelect
	infos    []*npool.TopMostGoodConstraint
	total    uint32
}

func (h *queryHandler) queryJoin() {
	h.baseQueryHandler.queryJoin()
	if h.stmCount == nil {
		return
	}
	h.stmCount.Modify(func(s *sql.Selector) {
		if err := h.queryJoinTopMostGood(s); err != nil {
			logger.Sugar().Errorw("queryJoinTopMostGood", "Error", err)
		}
	})
}

func (h *queryHandler) scan(ctx context.Context) error {
	return h.stmSelect.Scan(ctx, &h.infos)
}

func (h *queryHandler) formalize() {
	for _, info := range h.infos {
		info.TopMostType = types.GoodTopMostType(types.GoodTopMostType_value[info.TopMostTypeStr])
		info.Constraint = types.GoodTopMostConstraint(types.GoodTopMostConstraint_value[info.ConstraintStr])
		info.TargetValue = func() string { amount, _ := decimal.NewFromString(info.TargetValue); return amount.String() }()
	}
}

func (h *Handler) GetConstraint(ctx context.Context) (*npool.TopMostGoodConstraint, error) {
	handler := &queryHandler{
		baseQueryHandler: &baseQueryHandler{
			Handler: h,
		},
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryTopMostGoodConstraint(cli); err != nil {
			return wlog.WrapError(err)
		}
		handler.queryJoin()
		return handler.scan(ctx)
	})
	if err != nil {
		return nil, err
	}
	if len(handler.infos) == 0 {
		return nil, nil
	}
	if len(handler.infos) > 1 {
		return nil, wlog.Errorf("too many records")
	}

	handler.formalize()

	return handler.infos[0], nil
}

func (h *Handler) GetConstraints(ctx context.Context) ([]*npool.TopMostGoodConstraint, uint32, error) {
	handler := &queryHandler{
		baseQueryHandler: &baseQueryHandler{
			Handler: h,
		},
	}

	var err error

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		handler.stmSelect, err = handler.queryTopMostGoodConstraints(cli)
		if err != nil {
			return wlog.WrapError(err)
		}
		handler.stmCount, err = handler.queryTopMostGoodConstraints(cli)
		if err != nil {
			return wlog.WrapError(err)
		}

		handler.queryJoin()

		total, err := handler.stmCount.Count(_ctx)
		if err != nil {
			return wlog.WrapError(err)
		}
		handler.total = uint32(total)

		handler.stmSelect.
			Offset(int(handler.Offset)).
			Limit(int(handler.Limit))
		return handler.scan(_ctx)
	})
	if err != nil {
		return nil, 0, err
	}

	handler.formalize()

	return handler.infos, handler.total, nil
}
