package history

import (
	"context"

	"entgo.io/ent/dialect/sql"

	historycrud "github.com/NpoolPlatform/good-middleware/pkg/crud/good/reward/history"
	"github.com/NpoolPlatform/good-middleware/pkg/db"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent"
	entgoodbase "github.com/NpoolPlatform/good-middleware/pkg/db/ent/goodbase"
	enthistory "github.com/NpoolPlatform/good-middleware/pkg/db/ent/goodrewardhistory"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/good/reward/history"

	"github.com/shopspring/decimal"
)

type queryHandler struct {
	*Handler
	stmSelect *ent.GoodRewardHistorySelect
	stmCount  *ent.GoodRewardHistorySelect
	infos     []*npool.History
	total     uint32
}

func (h *queryHandler) selectHistory(stm *ent.GoodRewardHistoryQuery) *ent.GoodRewardHistorySelect {
	return stm.Select(enthistory.FieldID)
}

func (h *queryHandler) queryHistories(cli *ent.Client) (*ent.GoodRewardHistorySelect, error) {
	stm, err := historycrud.SetQueryConds(cli.GoodRewardHistory.Query(), h.Conds)
	if err != nil {
		return nil, err
	}
	return h.selectHistory(stm), nil
}

func (h *queryHandler) queryJoinMyself(s *sql.Selector) {
	t := sql.Table(enthistory.Table)
	s.LeftJoin(t).
		On(
			s.C(enthistory.FieldID),
			t.C(enthistory.FieldID),
		).
		AppendSelect(
			t.C(enthistory.FieldEntID),
			t.C(enthistory.FieldGoodID),
			t.C(enthistory.FieldRewardDate),
			t.C(enthistory.FieldTid),
			t.C(enthistory.FieldAmount),
			t.C(enthistory.FieldUnitAmount),
			t.C(enthistory.FieldCreatedAt),
			t.C(enthistory.FieldUpdatedAt),
		)
}

func (h *queryHandler) queryJoinGood(s *sql.Selector) {
	t := sql.Table(entgoodbase.Table)
	s.LeftJoin(t).
		On(
			s.C(enthistory.FieldGoodID),
			t.C(entgoodbase.FieldEntID),
		).
		AppendSelect(
			sql.As(t.C(entgoodbase.FieldName), "good_name"),
		)
}

func (h *queryHandler) queryJoin() {
	h.stmSelect.Modify(func(s *sql.Selector) {
		h.queryJoinMyself(s)
		h.queryJoinGood(s)
	})
	if h.stmCount == nil {
		return
	}
	h.stmCount.Modify(func(s *sql.Selector) {
		h.queryJoinGood(s)
	})
}

func (h *queryHandler) scan(ctx context.Context) error {
	return h.stmSelect.Scan(ctx, &h.infos)
}

func (h *queryHandler) formalize() {
	for _, info := range h.infos {
		amount, err := decimal.NewFromString(info.Amount)
		if err != nil {
			info.Amount = decimal.NewFromInt(0).String()
		} else {
			info.Amount = amount.String()
		}
		amount, err = decimal.NewFromString(info.UnitAmount)
		if err != nil {
			info.UnitAmount = decimal.NewFromInt(0).String()
		} else {
			info.UnitAmount = amount.String()
		}
	}
}

func (h *Handler) GetHistories(ctx context.Context) ([]*npool.History, uint32, error) {
	handler := &queryHandler{
		Handler: h,
	}

	var err error
	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		handler.stmSelect, err = handler.queryHistories(cli)
		if err != nil {
			return err
		}
		handler.stmCount, err = handler.queryHistories(cli)
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
