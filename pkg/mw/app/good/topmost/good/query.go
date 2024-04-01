package topmostgood

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"

	topmostgoodcrud "github.com/NpoolPlatform/good-middleware/pkg/crud/app/good/topmost/good"
	"github.com/NpoolPlatform/good-middleware/pkg/db"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent"
	entappgoodbase "github.com/NpoolPlatform/good-middleware/pkg/db/ent/appgoodbase"
	entgoodbase "github.com/NpoolPlatform/good-middleware/pkg/db/ent/goodbase"
	enttopmost "github.com/NpoolPlatform/good-middleware/pkg/db/ent/topmost"
	enttopmostgood "github.com/NpoolPlatform/good-middleware/pkg/db/ent/topmostgood"
	types "github.com/NpoolPlatform/message/npool/basetypes/good/v1"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/app/good/topmost/good"

	"github.com/shopspring/decimal"
)

type queryHandler struct {
	*Handler
	stmSelect *ent.TopMostGoodSelect
	stmCount  *ent.TopMostGoodSelect
	infos     []*npool.TopMostGood
	total     uint32
}

func (h *queryHandler) selectTopMostGood(stm *ent.TopMostGoodQuery) *ent.TopMostGoodSelect {
	return stm.Select(enttopmostgood.FieldID)
}

func (h *queryHandler) queryTopMostGood(cli *ent.Client) error {
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

func (h *queryHandler) queryTopMostGoods(cli *ent.Client) (*ent.TopMostGoodSelect, error) {
	stm, err := topmostgoodcrud.SetQueryConds(cli.TopMostGood.Query(), h.Conds)
	if err != nil {
		return nil, err
	}
	return h.selectTopMostGood(stm), nil
}

func (h *queryHandler) queryJoinMyself(s *sql.Selector) {
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

func (h *queryHandler) queryJoinAppGood(s *sql.Selector) {
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

func (h *queryHandler) queryJoinTopMost(s *sql.Selector) {
	t := sql.Table(enttopmost.Table)
	s.LeftJoin(t).
		On(
			s.C(enttopmostgood.FieldTopMostID),
			t.C(enttopmost.FieldEntID),
		).
		AppendSelect(
			sql.As(t.C(enttopmost.FieldTopMostType), "top_most_type"),
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

func (h *queryHandler) queryJoin() {
	h.stmSelect.Modify(func(s *sql.Selector) {
		h.queryJoinMyself(s)
		h.queryJoinTopMost(s)
		h.queryJoinAppGood(s)
	})
	if h.stmCount == nil {
		return
	}
	h.stmCount.Modify(func(s *sql.Selector) {
		h.queryJoinTopMost(s)
		h.queryJoinAppGood(s)
	})
}

func (h *queryHandler) scan(ctx context.Context) error {
	return h.stmSelect.Scan(ctx, &h.infos)
}

func (h *queryHandler) formalize() {
	for _, info := range h.infos {
		info.UnitPrice = func() string { amount, _ := decimal.NewFromString(info.UnitPrice); return amount.String() }()
		info.TopMostType = types.GoodTopMostType(types.GoodTopMostType_value[info.TopMostTypeStr])
	}
}

func (h *Handler) GetTopMostGood(ctx context.Context) (*npool.TopMostGood, error) {
	handler := &queryHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryTopMostGood(cli); err != nil {
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

func (h *Handler) GetTopMostGoods(ctx context.Context) ([]*npool.TopMostGood, uint32, error) {
	handler := &queryHandler{
		Handler: h,
	}

	var err error

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		handler.stmSelect, err = handler.queryTopMostGoods(cli)
		if err != nil {
			return err
		}
		handler.stmCount, err = handler.queryTopMostGoods(cli)
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
