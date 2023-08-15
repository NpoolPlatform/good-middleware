package required

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"

	// requiredcrud "github.com/NpoolPlatform/good-middleware/pkg/crud/good/required"
	"github.com/NpoolPlatform/good-middleware/pkg/db"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent"
	entgood "github.com/NpoolPlatform/good-middleware/pkg/db/ent/good"
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

func (h *queryHandler) queryRequired(cli *ent.Client) {
	h.stmSelect = h.selectRequired(
		cli.RequiredGood.
			Query().
			Where(
				entrequiredgood.ID(*h.ID),
				entrequiredgood.DeletedAt(0),
			),
	)
}

func (h *queryHandler) queryJoinMyself(s *sql.Selector) {
	t := sql.Table(entrequiredgood.Table)
	s.LeftJoin(t).
		On(
			s.C(entrequiredgood.FieldID),
			t.C(entrequiredgood.FieldID),
		).
		AppendSelect(
			sql.As(t.C(entrequiredgood.FieldMainGoodID), "main_good_id"),
			sql.As(t.C(entrequiredgood.FieldRequiredGoodID), "required_good_id"),
			sql.As(t.C(entrequiredgood.FieldMust), "must"),
			sql.As(t.C(entrequiredgood.FieldCommission), "commission"),
			sql.As(t.C(entrequiredgood.FieldCreatedAt), "created_at"),
			sql.As(t.C(entrequiredgood.FieldUpdatedAt), "updated_at"),
		)
}

func (h *queryHandler) queryJoinMainGood(s *sql.Selector) {
	t := sql.Table(entgood.Table)
	s.LeftJoin(t).
		On(
			s.C(entrequiredgood.FieldMainGoodID),
			t.C(entgood.FieldID),
		).
		AppendSelect(
			sql.As(t.C(entgood.FieldTitle), "main_good_name"),
		)
}

func (h *queryHandler) queryJoinRequiredGood(s *sql.Selector) {
	t := sql.Table(entgood.Table)
	s.LeftJoin(t).
		On(
			s.C(entrequiredgood.FieldRequiredGoodID),
			t.C(entgood.FieldID),
		).
		AppendSelect(
			sql.As(t.C(entgood.FieldTitle), "required_good_name"),
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
		handler.queryRequired(cli)
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
