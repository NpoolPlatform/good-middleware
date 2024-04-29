package topmostgood

import (
	"fmt"

	"entgo.io/ent/dialect/sql"

	topmostgoodcrud "github.com/NpoolPlatform/good-middleware/pkg/crud/app/good/topmost/good"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent"
	entappgoodbase "github.com/NpoolPlatform/good-middleware/pkg/db/ent/appgoodbase"
	entgoodbase "github.com/NpoolPlatform/good-middleware/pkg/db/ent/goodbase"
	enttopmost "github.com/NpoolPlatform/good-middleware/pkg/db/ent/topmost"
	enttopmostgood "github.com/NpoolPlatform/good-middleware/pkg/db/ent/topmostgood"
	types "github.com/NpoolPlatform/message/npool/basetypes/good/v1"
)

type baseQueryHandler struct {
	*Handler
	stmSelect *ent.TopMostGoodSelect
}

func (h *baseQueryHandler) selectTopMostGood(stm *ent.TopMostGoodQuery) *ent.TopMostGoodSelect {
	return stm.Select(enttopmostgood.FieldID)
}

func (h *baseQueryHandler) queryTopMostGood(cli *ent.Client) error {
	if h.ID == nil && h.EntID == nil {
		return fmt.Errorf("invalid id")
	}
	stm := cli.TopMostGood.Query().Where(enttopmostgood.DeletedAt(0))
	if h.ID != nil {
		stm.Where(enttopmostgood.ID(*h.ID))
	}
	if h.EntID != nil {
		stm.Where(enttopmostgood.EntID(*h.EntID))
	}
	h.stmSelect = h.selectTopMostGood(stm)
	return nil
}

func (h *baseQueryHandler) queryTopMostGoods(cli *ent.Client) (*ent.TopMostGoodSelect, error) {
	stm, err := topmostgoodcrud.SetQueryConds(cli.TopMostGood.Query(), h.Conds)
	if err != nil {
		return nil, err
	}
	return h.selectTopMostGood(stm), nil
}

func (h *baseQueryHandler) queryJoinMyself(s *sql.Selector) {
	t := sql.Table(enttopmostgood.Table)
	s.LeftJoin(t).
		On(
			s.C(enttopmostgood.FieldID),
			t.C(enttopmostgood.FieldID),
		).
		AppendSelect(
			t.C(enttopmostgood.FieldEntID),
			t.C(enttopmostgood.FieldAppGoodID),
			t.C(enttopmostgood.FieldTopMostID),
			t.C(enttopmostgood.FieldDisplayIndex),
			t.C(enttopmostgood.FieldUnitPrice),
			t.C(enttopmostgood.FieldCreatedAt),
			t.C(enttopmostgood.FieldUpdatedAt),
		)
}

func (h *baseQueryHandler) queryJoinAppGood(s *sql.Selector) {
	t1 := sql.Table(entappgoodbase.Table)
	t2 := sql.Table(entgoodbase.Table)
	s.Join(t1).
		On(
			s.C(enttopmostgood.FieldAppGoodID),
			t1.C(entappgoodbase.FieldEntID),
		).
		Join(t2).
		On(
			t1.C(entappgoodbase.FieldGoodID),
			t2.C(entgoodbase.FieldEntID),
		).
		AppendSelect(
			t1.C(entappgoodbase.FieldAppID),
			sql.As(t1.C(entappgoodbase.FieldName), "app_good_name"),
			sql.As(t2.C(entappgoodbase.FieldEntID), "good_id"),
			sql.As(t2.C(entappgoodbase.FieldName), "good_name"),
		)
}

func (h *baseQueryHandler) queryJoinTopMost(s *sql.Selector) {
	t := sql.Table(enttopmost.Table)
	s.LeftJoin(t).
		On(
			s.C(enttopmostgood.FieldTopMostID),
			t.C(enttopmost.FieldEntID),
		).
		AppendSelect(
			t.C(enttopmost.FieldTopMostType),
			sql.As(t.C(enttopmost.FieldTitle), "top_most_title"),
			sql.As(t.C(enttopmost.FieldMessage), "top_most_message"),
			sql.As(t.C(enttopmost.FieldTargetURL), "top_most_target_url"),
		)

	if h.Conds != nil && h.Conds.TopMostType != nil {
		_type, ok := h.Conds.TopMostType.Val.(types.GoodTopMostType)
		if !ok {
			return
		}
		s.Where(
			sql.EQ(t.C(enttopmost.FieldTopMostType), _type.String()),
		)
	}
}

func (h *baseQueryHandler) queryJoin() {
	h.stmSelect.Modify(func(s *sql.Selector) {
		h.queryJoinMyself(s)
		h.queryJoinTopMost(s)
		h.queryJoinAppGood(s)
	})
}
