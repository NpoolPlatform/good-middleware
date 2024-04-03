package constraint

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"

	constraintcrud "github.com/NpoolPlatform/good-middleware/pkg/crud/app/good/topmost/good/constraint"
	"github.com/NpoolPlatform/good-middleware/pkg/db"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent"
	entappgoodbase "github.com/NpoolPlatform/good-middleware/pkg/db/ent/appgoodbase"
	enttopmost "github.com/NpoolPlatform/good-middleware/pkg/db/ent/topmost"
	enttopmostgood "github.com/NpoolPlatform/good-middleware/pkg/db/ent/topmostgood"
	enttopmostgoodconstraint "github.com/NpoolPlatform/good-middleware/pkg/db/ent/topmostgoodconstraint"
	types "github.com/NpoolPlatform/message/npool/basetypes/good/v1"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/app/good/topmost/good/constraint"

	"github.com/shopspring/decimal"
)

type queryHandler struct {
	*Handler
	stmSelect *ent.TopMostGoodConstraintSelect
	stmCount  *ent.TopMostGoodConstraintSelect
	infos     []*npool.TopMostGoodConstraint
	total     uint32
}

func (h *queryHandler) selectTopMostGoodConstraint(stm *ent.TopMostGoodConstraintQuery) *ent.TopMostGoodConstraintSelect {
	return stm.Select(enttopmostgoodconstraint.FieldID)
}

func (h *queryHandler) queryTopMostGoodConstraint(cli *ent.Client) error {
	if h.ID == nil && h.EntID == nil {
		return fmt.Errorf("invalid id")
	}
	stm := cli.TopMostGoodConstraint.Query().Where(enttopmostgoodconstraint.DeletedAt(0))
	if h.ID != nil {
		stm.Where(enttopmostgoodconstraint.ID(*h.ID))
	}
	if h.EntID != nil {
		stm.Where(enttopmostgoodconstraint.EntID(*h.EntID))
	}
	h.stmSelect = h.selectTopMostGoodConstraint(stm)
	return nil
}

func (h *queryHandler) queryTopMostGoodConstraints(cli *ent.Client) (*ent.TopMostGoodConstraintSelect, error) {
	stm, err := constraintcrud.SetQueryConds(cli.TopMostGoodConstraint.Query(), h.ConstraintConds)
	if err != nil {
		return nil, err
	}
	return h.selectTopMostGoodConstraint(stm), nil
}

func (h *queryHandler) queryJoinMyself(s *sql.Selector) {
	t := sql.Table(enttopmostgoodconstraint.Table)
	s.LeftJoin(t).
		On(
			s.C(enttopmostgoodconstraint.FieldID),
			t.C(enttopmostgoodconstraint.FieldID),
		).
		AppendSelect(
			t.C(enttopmostgoodconstraint.FieldEntID),
			t.C(enttopmostgoodconstraint.FieldTopMostGoodID),
			t.C(enttopmostgoodconstraint.FieldConstraint),
			t.C(enttopmostgoodconstraint.FieldTargetValue),
			t.C(enttopmostgoodconstraint.FieldIndex),
			t.C(enttopmostgoodconstraint.FieldCreatedAt),
			t.C(enttopmostgoodconstraint.FieldUpdatedAt),
		)
}

func (h *queryHandler) queryJoinTopMostGood(s *sql.Selector) {
	t1 := sql.Table(enttopmostgood.Table)
	t2 := sql.Table(enttopmost.Table)
	t3 := sql.Table(entappgoodbase.Table)
	s.LeftJoin(t1).
		On(
			s.C(enttopmostgoodconstraint.FieldTopMostGoodID),
			t1.C(enttopmostgood.FieldEntID),
		).
		LeftJoin(t2).
		On(
			t1.C(enttopmostgood.FieldTopMostID),
			t2.C(enttopmost.FieldEntID),
		).
		LeftJoin(t3).
		On(
			t1.C(enttopmostgood.FieldAppGoodID),
			t3.C(entappgoodbase.FieldEntID),
		).
		AppendSelect(
			t1.C(enttopmostgood.FieldTopMostID),
			t1.C(enttopmostgood.FieldAppGoodID),
			sql.As(t3.C(entappgoodbase.FieldName), "app_good_name"),
			t2.C(enttopmost.FieldAppID),
			t2.C(enttopmost.FieldTopMostType),
			sql.As(t2.C(enttopmost.FieldTitle), "top_most_title"),
			sql.As(t2.C(enttopmost.FieldMessage), "top_most_message"),
			sql.As(t2.C(enttopmost.FieldTargetURL), "top_most_target_url"),
		)
}

func (h *queryHandler) queryJoin() {
	h.stmSelect.Modify(func(s *sql.Selector) {
		h.queryJoinMyself(s)
		h.queryJoinTopMostGood(s)
	})
	if h.stmCount == nil {
		return
	}
	h.stmCount.Modify(func(s *sql.Selector) {
		h.queryJoinTopMostGood(s)
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
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryTopMostGoodConstraint(cli); err != nil {
			return err
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
		return nil, fmt.Errorf("too many records")
	}

	handler.formalize()

	return handler.infos[0], nil
}

func (h *Handler) GetConstraints(ctx context.Context) ([]*npool.TopMostGoodConstraint, uint32, error) {
	handler := &queryHandler{
		Handler: h,
	}

	var err error

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		handler.stmSelect, err = handler.queryTopMostGoodConstraints(cli)
		if err != nil {
			return err
		}
		handler.stmCount, err = handler.queryTopMostGoodConstraints(cli)
		if err != nil {
			return err
		}

		handler.queryJoin()

		total, err := handler.stmCount.Count(_ctx)
		if err != nil {
			return err
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
