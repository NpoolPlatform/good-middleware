package topmostgood

import (
	"context"
	"encoding/json"
	"fmt"

	"entgo.io/ent/dialect/sql"

	topmostgoodcrud "github.com/NpoolPlatform/good-middleware/pkg/crud/app/good/topmost/good"
	"github.com/NpoolPlatform/good-middleware/pkg/db"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent"
	entappgood "github.com/NpoolPlatform/good-middleware/pkg/db/ent/appgood"
	entgood "github.com/NpoolPlatform/good-middleware/pkg/db/ent/good"
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
			sql.As(t.C(enttopmostgood.FieldEntID), "ent_id"),
			sql.As(t.C(enttopmostgood.FieldAppID), "app_id"),
			sql.As(t.C(enttopmostgood.FieldGoodID), "good_id"),
			sql.As(t.C(enttopmostgood.FieldAppGoodID), "app_good_id"),
			sql.As(t.C(enttopmostgood.FieldCoinTypeID), "coin_type_id"),
			sql.As(t.C(enttopmostgood.FieldTopMostID), "top_most_id"),
			sql.As(t.C(enttopmostgood.FieldDisplayIndex), "display_index"),
			sql.As(t.C(enttopmostgood.FieldPosters), "posters"),
			sql.As(t.C(enttopmostgood.FieldUnitPrice), "unit_price"),
			sql.As(t.C(enttopmostgood.FieldPackagePrice), "package_price"),
			sql.As(t.C(enttopmostgood.FieldCreatedAt), "created_at"),
			sql.As(t.C(enttopmostgood.FieldUpdatedAt), "updated_at"),
		)
}

func (h *queryHandler) queryJoinGood(s *sql.Selector) {
	t := sql.Table(entgood.Table)
	s.LeftJoin(t).
		On(
			s.C(enttopmostgood.FieldGoodID),
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
			s.C(enttopmostgood.FieldAppGoodID),
			t.C(entappgood.FieldEntID),
		).
		AppendSelect(
			sql.As(t.C(entappgood.FieldGoodName), "app_good_name"),
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
		h.queryJoinGood(s)
		h.queryJoinTopMost(s)
		h.queryJoinAppGood(s)
	})
	if h.stmCount == nil {
		return
	}
	h.stmCount.Modify(func(s *sql.Selector) {
		h.queryJoinGood(s)
		h.queryJoinTopMost(s)
		h.queryJoinAppGood(s)
	})
}

func (h *queryHandler) scan(ctx context.Context) error {
	return h.stmSelect.Scan(ctx, &h.infos)
}

func (h *queryHandler) formalize() {
	for _, info := range h.infos {
		amount, err := decimal.NewFromString(info.UnitPrice)
		if err != nil {
			info.UnitPrice = decimal.NewFromInt(0).String()
		} else {
			info.UnitPrice = amount.String()
		}
		amount, err = decimal.NewFromString(info.PackagePrice)
		if err != nil {
			info.PackagePrice = decimal.NewFromInt(0).String()
		} else {
			info.PackagePrice = amount.String()
		}
		_ = json.Unmarshal([]byte(info.PostersStr), &info.Posters)
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
