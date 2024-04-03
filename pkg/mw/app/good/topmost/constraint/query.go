package constraint

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"

	constraintcrud "github.com/NpoolPlatform/good-middleware/pkg/crud/app/good/topmost/constraint"
	"github.com/NpoolPlatform/good-middleware/pkg/db"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent"
	enttopmost "github.com/NpoolPlatform/good-middleware/pkg/db/ent/topmost"
	enttopmostconstraint "github.com/NpoolPlatform/good-middleware/pkg/db/ent/topmostconstraint"
	types "github.com/NpoolPlatform/message/npool/basetypes/good/v1"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/app/good/topmost/constraint"

	"github.com/shopspring/decimal"
)

type queryHandler struct {
	*Handler
	stmSelect *ent.TopMostConstraintSelect
	stmCount  *ent.TopMostConstraintSelect
	infos     []*npool.TopMostConstraint
	total     uint32
}

func (h *queryHandler) selectTopMostConstraint(stm *ent.TopMostConstraintQuery) *ent.TopMostConstraintSelect {
	return stm.Select(enttopmostconstraint.FieldID)
}

func (h *queryHandler) queryTopMostConstraint(cli *ent.Client) error {
	if h.ID == nil && h.EntID == nil {
		return fmt.Errorf("invalid id")
	}
	stm := cli.TopMostConstraint.Query().Where(enttopmostconstraint.DeletedAt(0))
	if h.ID != nil {
		stm.Where(enttopmostconstraint.ID(*h.ID))
	}
	if h.EntID != nil {
		stm.Where(enttopmostconstraint.EntID(*h.EntID))
	}
	h.stmSelect = h.selectTopMostConstraint(stm)
	return nil
}

func (h *queryHandler) queryTopMostConstraints(cli *ent.Client) (*ent.TopMostConstraintSelect, error) {
	stm, err := constraintcrud.SetQueryConds(cli.TopMostConstraint.Query(), h.ConstraintConds)
	if err != nil {
		return nil, err
	}
	return h.selectTopMostConstraint(stm), nil
}

func (h *queryHandler) queryJoinMyself(s *sql.Selector) {
	t := sql.Table(enttopmostconstraint.Table)
	s.LeftJoin(t).
		On(
			s.C(enttopmostconstraint.FieldID),
			t.C(enttopmostconstraint.FieldID),
		).
		AppendSelect(
			t.C(enttopmostconstraint.FieldEntID),
			t.C(enttopmostconstraint.FieldTopMostID),
			t.C(enttopmostconstraint.FieldTargetValue),
			t.C(enttopmostconstraint.FieldIndex),
			t.C(enttopmostconstraint.FieldCreatedAt),
			t.C(enttopmostconstraint.FieldUpdatedAt),
		)
}

func (h *queryHandler) queryJoinTopMost(s *sql.Selector) {
	t := sql.Table(enttopmost.Table)
	s.LeftJoin(t).
		On(
			s.C(enttopmostconstraint.FieldTopMostID),
			t.C(enttopmost.FieldEntID),
		).
		AppendSelect(
			t.C(enttopmost.FieldTopMostType),
			t.C(enttopmost.FieldTitle),
			t.C(enttopmost.FieldMessage),
			t.C(enttopmost.FieldTargetURL),
		)
}

func (h *queryHandler) queryJoin() {
	h.stmSelect.Modify(func(s *sql.Selector) {
		h.queryJoinMyself(s)
		h.queryJoinTopMost(s)
	})
	if h.stmCount == nil {
		return
	}
	h.stmCount.Modify(func(s *sql.Selector) {
		h.queryJoinTopMost(s)
	})
}

func (h *queryHandler) scan(ctx context.Context) error {
	return h.stmSelect.Scan(ctx, &h.infos)
}

func (h *queryHandler) formalize() {
	for _, info := range h.infos {
		info.TopMostType = types.GoodTopMostType(types.GoodTopMostType_value[info.TopMostTypeStr])
		info.TargetValue = func() string { amount, _ := decimal.NewFromString(info.TargetValue); return amount.String() }()
	}
}

func (h *Handler) GetConstraint(ctx context.Context) (*npool.TopMostConstraint, error) {
	handler := &queryHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryTopMostConstraint(cli); err != nil {
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

func (h *Handler) GetConstraints(ctx context.Context) ([]*npool.TopMostConstraint, uint32, error) {
	handler := &queryHandler{
		Handler: h,
	}

	var err error

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		handler.stmSelect, err = handler.queryTopMostConstraints(cli)
		if err != nil {
			return err
		}
		handler.stmCount, err = handler.queryTopMostConstraints(cli)
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
