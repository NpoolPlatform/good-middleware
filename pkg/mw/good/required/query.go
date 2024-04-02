package required

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"

	requiredcrud "github.com/NpoolPlatform/good-middleware/pkg/crud/good/required"
	"github.com/NpoolPlatform/good-middleware/pkg/db"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent"
	entgoodbase "github.com/NpoolPlatform/good-middleware/pkg/db/ent/goodbase"
	entrequiredgood "github.com/NpoolPlatform/good-middleware/pkg/db/ent/requiredgood"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/good/required"
)

type queryHandler struct {
	*Handler
	stmSelect *ent.RequiredGoodSelect
	stmCount  *ent.RequiredGoodSelect
	infos     []*npool.Required
	total     uint32
}

func (h *queryHandler) selectRequired(stm *ent.RequiredGoodQuery) *ent.RequiredGoodSelect {
	return stm.Select(entrequiredgood.FieldID)
}

func (h *queryHandler) queryRequired(cli *ent.Client) error {
	if h.ID == nil && h.EntID == nil {
		return fmt.Errorf("invalid id")
	}
	stm := cli.RequiredGood.Query().Where(entrequiredgood.DeletedAt(0))
	if h.ID != nil {
		stm.Where(entrequiredgood.ID(*h.ID))
	}
	if h.EntID != nil {
		stm.Where(entrequiredgood.EntID(*h.EntID))
	}
	h.stmSelect = h.selectRequired(stm)
	return nil
}

func (h *queryHandler) queryRequireds(cli *ent.Client) (*ent.RequiredGoodSelect, error) {
	stm, err := requiredcrud.SetQueryConds(cli.RequiredGood.Query(), h.Conds)
	if err != nil {
		return nil, err
	}
	return h.selectRequired(stm), nil
}

func (h *queryHandler) queryJoinMyself(s *sql.Selector) {
	t := sql.Table(entrequiredgood.Table)
	s.LeftJoin(t).
		On(
			s.C(entrequiredgood.FieldID),
			t.C(entrequiredgood.FieldID),
		).
		AppendSelect(
			sql.As(t.C(entrequiredgood.FieldEntID), "ent_id"),
			sql.As(t.C(entrequiredgood.FieldMainGoodID), "main_good_id"),
			sql.As(t.C(entrequiredgood.FieldRequiredGoodID), "required_good_id"),
			sql.As(t.C(entrequiredgood.FieldMust), "must"),
			sql.As(t.C(entrequiredgood.FieldCreatedAt), "created_at"),
			sql.As(t.C(entrequiredgood.FieldUpdatedAt), "updated_at"),
		)
}

func (h *queryHandler) queryJoinMainGood(s *sql.Selector) {
	t := sql.Table(entgoodbase.Table)
	s.LeftJoin(t).
		On(
			s.C(entrequiredgood.FieldMainGoodID),
			t.C(entgoodbase.FieldEntID),
		).
		AppendSelect(
			sql.As(t.C(entgoodbase.FieldName), "main_good_name"),
		)
}

func (h *queryHandler) queryJoinRequiredGood(s *sql.Selector) {
	t := sql.Table(entgoodbase.Table)
	s.LeftJoin(t).
		On(
			s.C(entrequiredgood.FieldRequiredGoodID),
			t.C(entgoodbase.FieldEntID),
		).
		AppendSelect(
			sql.As(t.C(entgoodbase.FieldName), "required_good_name"),
		)
}

func (h *queryHandler) queryJoin() {
	h.stmSelect.Modify(func(s *sql.Selector) {
		h.queryJoinMyself(s)
		h.queryJoinMainGood(s)
		h.queryJoinRequiredGood(s)
	})
	if h.stmCount == nil {
		return
	}
	h.stmCount.Modify(func(s *sql.Selector) {
		h.queryJoinMainGood(s)
		h.queryJoinRequiredGood(s)
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
