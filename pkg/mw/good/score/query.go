package score

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"

	scorecrud "github.com/NpoolPlatform/good-middleware/pkg/crud/good/score"
	"github.com/NpoolPlatform/good-middleware/pkg/db"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent"
	entappgood "github.com/NpoolPlatform/good-middleware/pkg/db/ent/appgood"
	entscore "github.com/NpoolPlatform/good-middleware/pkg/db/ent/score"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/good/score"

	"github.com/shopspring/decimal"
)

type queryHandler struct {
	*Handler
	stmSelect *ent.ScoreSelect
	stmCount  *ent.ScoreSelect
	infos     []*npool.Score
	total     uint32
}

func (h *queryHandler) selectScore(stm *ent.ScoreQuery) *ent.ScoreSelect {
	return stm.Select(entscore.FieldID)
}

func (h *queryHandler) queryScore(cli *ent.Client) error {
	if h.ID == nil && h.EntID == nil {
		return fmt.Errorf("invalid id")
	}
	stm := cli.Score.Query().Where(entscore.DeletedAt(0))
	if h.ID != nil {
		stm.Where(entscore.ID(*h.ID))
	}
	if h.EntID != nil {
		stm.Where(entscore.EntID(*h.EntID))
	}
	h.stmSelect = h.selectScore(stm)
	return nil
}

func (h *queryHandler) queryScores(cli *ent.Client) (*ent.ScoreSelect, error) {
	stm, err := scorecrud.SetQueryConds(cli.Score.Query(), h.Conds)
	if err != nil {
		return nil, err
	}
	return h.selectScore(stm), nil
}

func (h *queryHandler) queryJoinMyself(s *sql.Selector) {
	t := sql.Table(entscore.Table)
	s.LeftJoin(t).
		On(
			s.C(entscore.FieldID),
			t.C(entscore.FieldID),
		).
		AppendSelect(
			sql.As(t.C(entscore.FieldEntID), "ent_id"),
			sql.As(t.C(entscore.FieldAppID), "app_id"),
			sql.As(t.C(entscore.FieldUserID), "user_id"),
			sql.As(t.C(entscore.FieldGoodID), "good_id"),
			sql.As(t.C(entscore.FieldAppGoodID), "app_good_id"),
			sql.As(t.C(entscore.FieldScore), "score"),
			sql.As(t.C(entscore.FieldCreatedAt), "created_at"),
			sql.As(t.C(entscore.FieldUpdatedAt), "updated_at"),
		)
}

func (h *queryHandler) queryJoinAppGood(s *sql.Selector) {
	t := sql.Table(entappgood.Table)
	s.LeftJoin(t).
		On(
			s.C(entscore.FieldAppGoodID),
			t.C(entappgood.FieldEntID),
		).
		AppendSelect(
			sql.As(t.C(entappgood.FieldGoodName), "good_name"),
		)
}

func (h *queryHandler) queryJoin() {
	h.stmSelect.Modify(func(s *sql.Selector) {
		h.queryJoinMyself(s)
		h.queryJoinAppGood(s)
	})
	if h.stmCount == nil {
		return
	}
	h.stmCount.Modify(func(s *sql.Selector) {
		h.queryJoinAppGood(s)
	})
}

func (h *queryHandler) scan(ctx context.Context) error {
	return h.stmSelect.Scan(ctx, &h.infos)
}

func (h *queryHandler) formalize() {
	for _, info := range h.infos {
		amount, err := decimal.NewFromString(info.Score)
		if err != nil {
			info.Score = decimal.NewFromInt(0).String()
		} else {
			info.Score = amount.String()
		}
	}
}

func (h *Handler) GetScore(ctx context.Context) (*npool.Score, error) {
	handler := &queryHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryScore(cli); err != nil {
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

func (h *Handler) GetScores(ctx context.Context) ([]*npool.Score, uint32, error) {
	handler := &queryHandler{
		Handler: h,
	}

	var err error
	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		handler.stmSelect, err = handler.queryScores(cli)
		if err != nil {
			return err
		}
		handler.stmCount, err = handler.queryScores(cli)
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

	handler.formalize()

	return handler.infos, handler.total, nil
}
