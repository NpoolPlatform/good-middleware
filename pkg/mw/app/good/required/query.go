package required

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"

	requiredcrud "github.com/NpoolPlatform/good-middleware/pkg/crud/app/good/required"
	"github.com/NpoolPlatform/good-middleware/pkg/db"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent"
	entappgoodbase "github.com/NpoolPlatform/good-middleware/pkg/db/ent/appgoodbase"
	entgoodbase "github.com/NpoolPlatform/good-middleware/pkg/db/ent/goodbase"
	entrequiredappgood "github.com/NpoolPlatform/good-middleware/pkg/db/ent/requiredappgood"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/app/good/required"
)

type queryHandler struct {
	*Handler
	stmSelect *ent.RequiredAppGoodSelect
	stmCount  *ent.RequiredAppGoodSelect
	infos     []*npool.Required
	total     uint32
}

func (h *queryHandler) selectRequired(stm *ent.RequiredAppGoodQuery) *ent.RequiredAppGoodSelect {
	return stm.Select(entrequiredappgood.FieldID)
}

func (h *queryHandler) queryRequired(cli *ent.Client) error {
	if h.ID == nil && h.EntID == nil {
		return fmt.Errorf("invalid id")
	}
	stm := cli.RequiredAppGood.Query().Where(entrequiredappgood.DeletedAt(0))
	if h.ID != nil {
		stm.Where(entrequiredappgood.ID(*h.ID))
	}
	if h.EntID != nil {
		stm.Where(entrequiredappgood.EntID(*h.EntID))
	}
	h.stmSelect = h.selectRequired(stm)
	return nil
}

func (h *queryHandler) queryRequireds(cli *ent.Client) (*ent.RequiredAppGoodSelect, error) {
	stm, err := requiredcrud.SetQueryConds(cli.RequiredAppGood.Query(), h.RequiredConds)
	if err != nil {
		return nil, err
	}
	return h.selectRequired(stm), nil
}

func (h *queryHandler) queryJoinMyself(s *sql.Selector) {
	t := sql.Table(entrequiredappgood.Table)
	s.LeftJoin(t).
		On(
			s.C(entrequiredappgood.FieldID),
			t.C(entrequiredappgood.FieldID),
		).
		AppendSelect(
			t.C(entrequiredappgood.FieldEntID),
			t.C(entrequiredappgood.FieldMainAppGoodID),
			t.C(entrequiredappgood.FieldRequiredAppGoodID),
			t.C(entrequiredappgood.FieldMust),
			t.C(entrequiredappgood.FieldCreatedAt),
			t.C(entrequiredappgood.FieldUpdatedAt),
		)
}

func (h *queryHandler) queryJoinMainAppGood(s *sql.Selector) {
	t1 := sql.Table(entappgoodbase.Table)
	t2 := sql.Table(entgoodbase.Table)
	s.LeftJoin(t1).
		On(
			s.C(entrequiredappgood.FieldMainAppGoodID),
			t1.C(entappgoodbase.FieldEntID),
		).
		LeftJoin(t2).
		On(
			t1.C(entappgoodbase.FieldGoodID),
			t2.C(entgoodbase.FieldEntID),
		).
		AppendSelect(
			t1.C(entappgoodbase.FieldAppID),
			sql.As(t1.C(entappgoodbase.FieldName), "main_app_good_name"),
			sql.As(t1.C(entappgoodbase.FieldGoodID), "main_good_id"),
			sql.As(t2.C(entgoodbase.FieldName), "main_good_name"),
		)
}

func (h *queryHandler) queryJoinRequiredAppGood(s *sql.Selector) {
	t1 := sql.Table(entappgoodbase.Table)
	t2 := sql.Table(entgoodbase.Table)
	s.LeftJoin(t1).
		On(
			s.C(entrequiredappgood.FieldRequiredAppGoodID),
			t1.C(entappgoodbase.FieldEntID),
		).
		LeftJoin(t2).
		On(
			t1.C(entappgoodbase.FieldGoodID),
			t2.C(entgoodbase.FieldEntID),
		).
		AppendSelect(
			sql.As(t1.C(entappgoodbase.FieldName), "required_app_good_name"),
			sql.As(t1.C(entappgoodbase.FieldGoodID), "required_good_id"),
			sql.As(t2.C(entgoodbase.FieldName), "required_good_name"),
		)
}

func (h *queryHandler) queryJoin() {
	h.stmSelect.Modify(func(s *sql.Selector) {
		h.queryJoinMyself(s)
		h.queryJoinMainAppGood(s)
		h.queryJoinRequiredAppGood(s)
	})
	if h.stmCount == nil {
		return
	}
	h.stmCount.Modify(func(s *sql.Selector) {
		h.queryJoinMainAppGood(s)
		h.queryJoinRequiredAppGood(s)
	})
}

func (h *queryHandler) scan(ctx context.Context) error {
	return h.stmSelect.Scan(ctx, &h.infos)
}

func (h *Handler) GetRequired(ctx context.Context) (*npool.Required, error) {
	handler := &queryHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryRequired(cli); err != nil {
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

	return handler.infos[0], nil
}

func (h *Handler) GetRequireds(ctx context.Context) ([]*npool.Required, uint32, error) {
	handler := &queryHandler{
		Handler: h,
	}

	var err error
	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		handler.stmSelect, err = handler.queryRequireds(cli)
		if err != nil {
			return err
		}
		handler.stmCount, err = handler.queryRequireds(cli)
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

		return handler.scan(ctx)
	})
	if err != nil {
		return nil, 0, err
	}

	return handler.infos, handler.total, nil
}
