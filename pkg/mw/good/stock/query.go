package stock

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"

	stockcrud "github.com/NpoolPlatform/good-middleware/pkg/crud/good/stock"
	"github.com/NpoolPlatform/good-middleware/pkg/db"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent"
	entgood "github.com/NpoolPlatform/good-middleware/pkg/db/ent/good"
	entstock "github.com/NpoolPlatform/good-middleware/pkg/db/ent/stock"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/good/stock"

	"github.com/shopspring/decimal"
)

type queryHandler struct {
	*Handler
	stmSelect *ent.StockSelect
	stmCount  *ent.StockSelect
	infos     []*npool.Stock
	total     uint32
}

func (h *queryHandler) selectStock(stm *ent.StockQuery) *ent.StockSelect {
	return stm.Select(entstock.FieldID)
}

func (h *queryHandler) queryStock(cli *ent.Client) {
	h.stmSelect = h.selectStock(
		cli.Stock.
			Query().
			Where(
				entstock.ID(*h.ID),
				entstock.DeletedAt(0),
			),
	)
}

func (h *queryHandler) queryStocks(cli *ent.Client) (*ent.StockSelect, error) {
	stm, err := stockcrud.SetQueryConds(cli.Stock.Query(), h.Conds)
	if err != nil {
		return nil, err
	}
	return h.selectStock(stm), nil
}

func (h *queryHandler) queryJoinMyself(s *sql.Selector) {
	t := sql.Table(entstock.Table)
	s.LeftJoin(t).
		On(
			s.C(entstock.FieldID),
			t.C(entstock.FieldID),
		).
		AppendSelect(
			sql.As(t.C(entstock.FieldGoodID), "good_id"),
			sql.As(t.C(entstock.FieldTotal), "total"),
			sql.As(t.C(entstock.FieldLocked), "locked"),
			sql.As(t.C(entstock.FieldInService), "in_service"),
			sql.As(t.C(entstock.FieldWaitStart), "wait_start"),
			sql.As(t.C(entstock.FieldSold), "sold"),
			sql.As(t.C(entstock.FieldAppReserved), "app_reserved"),
			sql.As(t.C(entstock.FieldCreatedAt), "created_at"),
			sql.As(t.C(entstock.FieldUpdatedAt), "updated_at"),
		)
}

func (h *queryHandler) queryJoinGood(s *sql.Selector) {
	t := sql.Table(entgood.Table)
	s.LeftJoin(t).
		On(
			s.C(entstock.FieldGoodID),
			t.C(entgood.FieldID),
		).
		AppendSelect(
			sql.As(t.C(entgood.FieldTitle), "good_name"),
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
		amount, err := decimal.NewFromString(info.Total)
		if err != nil {
			info.Total = decimal.NewFromInt(0).String()
		} else {
			info.Total = amount.String()
		}
		amount, err = decimal.NewFromString(info.Locked)
		if err != nil {
			info.Locked = decimal.NewFromInt(0).String()
		} else {
			info.Locked = amount.String()
		}
		amount, err = decimal.NewFromString(info.InService)
		if err != nil {
			info.InService = decimal.NewFromInt(0).String()
		} else {
			info.InService = amount.String()
		}
		amount, err = decimal.NewFromString(info.WaitStart)
		if err != nil {
			info.WaitStart = decimal.NewFromInt(0).String()
		} else {
			info.WaitStart = amount.String()
		}
		amount, err = decimal.NewFromString(info.Sold)
		if err != nil {
			info.Sold = decimal.NewFromInt(0).String()
		} else {
			info.Sold = amount.String()
		}
		amount, err = decimal.NewFromString(info.AppReserved)
		if err != nil {
			info.AppReserved = decimal.NewFromInt(0).String()
		} else {
			info.AppReserved = amount.String()
		}
	}
}

func (h *Handler) GetStock(ctx context.Context) (*npool.Stock, error) {
	handler := &queryHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		handler.queryStock(cli)
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
