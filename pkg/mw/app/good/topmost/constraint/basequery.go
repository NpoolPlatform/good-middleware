package constraint

import (
	"entgo.io/ent/dialect/sql"

	logger "github.com/NpoolPlatform/go-service-framework/pkg/logger"
	wlog "github.com/NpoolPlatform/go-service-framework/pkg/wlog"
	constraintcrud "github.com/NpoolPlatform/good-middleware/pkg/crud/app/good/topmost/constraint"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent"
	enttopmost "github.com/NpoolPlatform/good-middleware/pkg/db/ent/topmost"
	enttopmostconstraint "github.com/NpoolPlatform/good-middleware/pkg/db/ent/topmostconstraint"
	types "github.com/NpoolPlatform/message/npool/basetypes/good/v1"

	"github.com/google/uuid"
)

type baseQueryHandler struct {
	*Handler
	stmSelect *ent.TopMostConstraintSelect
}

func (h *baseQueryHandler) selectTopMostConstraint(stm *ent.TopMostConstraintQuery) *ent.TopMostConstraintSelect {
	return stm.Select(enttopmostconstraint.FieldID)
}

func (h *baseQueryHandler) queryTopMostConstraint(cli *ent.Client) error {
	if h.ID == nil && h.EntID == nil {
		return wlog.Errorf("invalid id")
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

func (h *baseQueryHandler) queryTopMostConstraints(cli *ent.Client) (*ent.TopMostConstraintSelect, error) {
	stm, err := constraintcrud.SetQueryConds(cli.TopMostConstraint.Query(), h.ConstraintConds)
	if err != nil {
		return nil, err
	}
	return h.selectTopMostConstraint(stm), nil
}

func (h *baseQueryHandler) queryJoinMyself(s *sql.Selector) {
	t := sql.Table(enttopmostconstraint.Table)
	s.LeftJoin(t).
		On(
			s.C(enttopmostconstraint.FieldID),
			t.C(enttopmostconstraint.FieldID),
		).
		AppendSelect(
			t.C(enttopmostconstraint.FieldEntID),
			t.C(enttopmostconstraint.FieldTopMostID),
			t.C(enttopmostconstraint.FieldConstraint),
			t.C(enttopmostconstraint.FieldTargetValue),
			t.C(enttopmostconstraint.FieldIndex),
			t.C(enttopmostconstraint.FieldCreatedAt),
			t.C(enttopmostconstraint.FieldUpdatedAt),
		)
}

func (h *baseQueryHandler) queryJoinTopMost(s *sql.Selector) error {
	t := sql.Table(enttopmost.Table)
	s.Join(t).
		On(
			s.C(enttopmostconstraint.FieldTopMostID),
			t.C(enttopmost.FieldEntID),
		)
	if h.TopMostConds.AppID != nil {
		id, ok := h.TopMostConds.AppID.Val.(uuid.UUID)
		if !ok {
			return wlog.Errorf("invalid appid")
		}
		s.OnP(
			sql.EQ(t.C(enttopmost.FieldAppID), id),
		)
	}
	if h.TopMostConds.EntID != nil {
		id, ok := h.TopMostConds.EntID.Val.(uuid.UUID)
		if !ok {
			return wlog.Errorf("invalid entid")
		}
		s.OnP(
			sql.EQ(t.C(enttopmost.FieldEntID), id),
		)
	}
	if h.TopMostConds.TopMostType != nil {
		_type, ok := h.TopMostConds.TopMostType.Val.(types.GoodTopMostType)
		if !ok {
			return wlog.Errorf("invalid topmosttype")
		}
		s.OnP(
			sql.EQ(t.C(enttopmost.FieldTopMostType), _type.String()),
		)
	}
	s.AppendSelect(
		t.C(enttopmost.FieldAppID),
		t.C(enttopmost.FieldTopMostType),
		sql.As(t.C(enttopmost.FieldTitle), "top_most_title"),
		sql.As(t.C(enttopmost.FieldMessage), "top_most_message"),
		sql.As(t.C(enttopmost.FieldTargetURL), "top_most_target_url"),
	)
	return nil
}

func (h *baseQueryHandler) queryJoin() {
	h.stmSelect.Modify(func(s *sql.Selector) {
		h.queryJoinMyself(s)
		if err := h.queryJoinTopMost(s); err != nil {
			logger.Sugar().Errorw("queryJoinTopMost", "Error", err)
		}
	})
}
