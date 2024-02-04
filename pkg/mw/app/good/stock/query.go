package appstock

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"

	appstockcrud "github.com/NpoolPlatform/good-middleware/pkg/crud/app/good/stock"
	"github.com/NpoolPlatform/good-middleware/pkg/db"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent"
	entappgood "github.com/NpoolPlatform/good-middleware/pkg/db/ent/appgood"
	entappstock "github.com/NpoolPlatform/good-middleware/pkg/db/ent/appstock"
	entgood "github.com/NpoolPlatform/good-middleware/pkg/db/ent/good"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/app/good/stock"

	"github.com/shopspring/decimal"
)

type queryHandler struct {
	*Handler
	stmSelect *ent.AppStockSelect
	stmCount  *ent.AppStockSelect
	infos     []*npool.Stock
	total     uint32
}

func (h *queryHandler) selectStock(stm *ent.AppStockQuery) *ent.AppStockSelect {
	return stm.Select(entappstock.FieldID)
}

func (h *queryHandler) queryStock(cli *ent.Client) error {
	if h.ID == nil && h.EntID == nil {
		return fmt.Errorf("invalid id")
	}
	stm := cli.AppStock.Query().Where(entappstock.DeletedAt(0))
	if h.ID != nil {
		stm.Where(entappstock.ID(*h.ID))
	}
	if h.EntID != nil {
		stm.Where(entappstock.EntID(*h.EntID))
	}
	h.stmSelect = h.selectStock(stm)
	return nil
}

func (h *queryHandler) queryStocks(cli *ent.Client) (*ent.AppStockSelect, error) {
	stm, err := appstockcrud.SetQueryConds(cli.AppStock.Query(), &appstockcrud.Conds{
		EntIDs: &cruder.Cond{Op: cruder.IN, Val: h.EntIDs},
	})
	if err != nil {
		return nil, err
	}
	return h.selectStock(stm), nil
}

func (h *queryHandler) queryJoinMyself(s *sql.Selector) {
	t := sql.Table(entappstock.Table)
	s.LeftJoin(t).
		On(
			s.C(entappstock.FieldID),
			t.C(entappstock.FieldID),
		).
		AppendSelect(
			sql.As(t.C(entappstock.FieldEntID), "ent_id"),
			sql.As(t.C(entappstock.FieldAppID), "app_id"),
			sql.As(t.C(entappstock.FieldGoodID), "good_id"),
			sql.As(t.C(entappstock.FieldAppGoodID), "app_good_id"),
			sql.As(t.C(entappstock.FieldReserved), "reserved"),
			sql.As(t.C(entappstock.FieldSpotQuantity), "spot_quantity"),
			sql.As(t.C(entappstock.FieldLocked), "locked"),
			sql.As(t.C(entappstock.FieldWaitStart), "wait_start"),
			sql.As(t.C(entappstock.FieldInService), "in_service"),
			sql.As(t.C(entappstock.FieldSold), "sold"),
			sql.As(t.C(entappstock.FieldCreatedAt), "created_at"),
			sql.As(t.C(entappstock.FieldUpdatedAt), "updated_at"),
		)
}

func (h *queryHandler) queryJoinGood(s *sql.Selector) {
	t := sql.Table(entgood.Table)
	s.LeftJoin(t).
		On(
			s.C(entappstock.FieldGoodID),
			t.C(entgood.FieldEntID),
		).
		AppendSelect(
			sql.As(t.C(entgood.FieldTitle), "good_name"),
		)
}

func (h *queryHandler) queryJoinAppGood(s *sql.Selector) {
	t := sql.Table(entappgood.Table)
	s.LeftJoin(t).
		On(
			s.C(entappstock.FieldAppGoodID),
			t.C(entappgood.FieldEntID),
		).
		AppendSelect(
			sql.As(t.C(entappgood.FieldGoodName), "app_good_name"),
		)
}

func (h *queryHandler) queryJoin() {
	h.stmSelect.Modify(func(s *sql.Selector) {
		h.queryJoinMyself(s)
		h.queryJoinGood(s)
		h.queryJoinAppGood(s)
	})
	if h.stmCount == nil {
		return
	}
	h.stmCount.Modify(func(s *sql.Selector) {
		h.queryJoinGood(s)
		h.queryJoinAppGood(s)
	})
}

func (h *queryHandler) scan(ctx context.Context) error {
	return h.stmSelect.Scan(ctx, &h.infos)
}

func (h *queryHandler) formalize() {
	for _, info := range h.infos {
		amount, err := decimal.NewFromString(info.Reserved)
		if err != nil {
			info.Reserved = decimal.NewFromInt(0).String()
		} else {
			info.Reserved = amount.String()
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
		amount, err = decimal.NewFromString(info.SpotQuantity)
		if err != nil {
			info.SpotQuantity = decimal.NewFromInt(0).String()
		} else {
			info.SpotQuantity = amount.String()
		}
	}
}

func (h *Handler) GetStock(ctx context.Context) (*npool.Stock, error) {
	handler := &queryHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryStock(cli); err != nil {
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

func (h *Handler) GetStocks(ctx context.Context) ([]*npool.Stock, uint32, error) {
	handler := &queryHandler{
		Handler: h,
	}

	var err error

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		handler.stmSelect, err = handler.queryStocks(cli)
		if err != nil {
			return err
		}
		handler.stmCount, err = handler.queryStocks(cli)
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
			Limit(len(h.EntIDs))
		return handler.scan(_ctx)
	})
	if err != nil {
		return nil, 0, err
	}

	handler.formalize()

	return handler.infos, handler.total, nil
}
