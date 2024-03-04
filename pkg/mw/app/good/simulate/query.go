package appsimulategood

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"

	appsimulategoodcrud "github.com/NpoolPlatform/good-middleware/pkg/crud/app/good/simulate"
	"github.com/NpoolPlatform/good-middleware/pkg/db"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent"
	entappgood "github.com/NpoolPlatform/good-middleware/pkg/db/ent/appgood"
	entappsimulategood "github.com/NpoolPlatform/good-middleware/pkg/db/ent/appsimulategood"
	entgood "github.com/NpoolPlatform/good-middleware/pkg/db/ent/good"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/app/good/simulate"
	"github.com/shopspring/decimal"
)

type queryHandler struct {
	*Handler
	stmSelect *ent.AppSimulateGoodSelect
	stmCount  *ent.AppSimulateGoodSelect
	infos     []*npool.Simulate
	total     uint32
}

func (h *queryHandler) selectSimulate(stm *ent.AppSimulateGoodQuery) *ent.AppSimulateGoodSelect {
	return stm.Select(entappsimulategood.FieldID)
}

func (h *queryHandler) querySimulate(cli *ent.Client) error {
	if h.ID == nil && h.EntID == nil {
		return fmt.Errorf("invalid id")
	}
	stm := cli.AppSimulateGood.Query().Where(entappsimulategood.DeletedAt(0))
	if h.ID != nil {
		stm.Where(entappsimulategood.ID(*h.ID))
	}
	if h.EntID != nil {
		stm.Where(entappsimulategood.EntID(*h.EntID))
	}
	h.stmSelect = h.selectSimulate(stm)
	return nil
}

func (h *queryHandler) querySimulates(cli *ent.Client) (*ent.AppSimulateGoodSelect, error) {
	stm, err := appsimulategoodcrud.SetQueryConds(cli.AppSimulateGood.Query(), h.Conds)
	if err != nil {
		return nil, err
	}
	return h.selectSimulate(stm), nil
}

func (h *queryHandler) queryJoinMyself(s *sql.Selector) {
	t := sql.Table(entappsimulategood.Table)
	s.LeftJoin(t).
		On(
			s.C(entappsimulategood.FieldID),
			t.C(entappsimulategood.FieldID),
		).
		AppendSelect(
			sql.As(t.C(entappsimulategood.FieldEntID), "ent_id"),
			sql.As(t.C(entappsimulategood.FieldAppID), "app_id"),
			sql.As(t.C(entappsimulategood.FieldGoodID), "good_id"),
			sql.As(t.C(entappsimulategood.FieldAppGoodID), "app_good_id"),
			sql.As(t.C(entappsimulategood.FieldCoinTypeID), "coin_type_id"),
			sql.As(t.C(entappsimulategood.FieldFixedOrderUnits), "fixed_order_units"),
			sql.As(t.C(entappsimulategood.FieldFixedOrderDuration), "fixed_order_duration"),
			sql.As(t.C(entappsimulategood.FieldCreatedAt), "created_at"),
			sql.As(t.C(entappsimulategood.FieldUpdatedAt), "updated_at"),
		)
}

func (h *queryHandler) queryJoinGood(s *sql.Selector) {
	t := sql.Table(entgood.Table)
	s.LeftJoin(t).
		On(
			s.C(entappsimulategood.FieldGoodID),
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
			s.C(entappsimulategood.FieldAppGoodID),
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
		amount, err := decimal.NewFromString(info.FixedOrderUnits)
		if err != nil {
			info.FixedOrderUnits = decimal.NewFromInt(0).String()
		} else {
			info.FixedOrderUnits = amount.String()
		}
	}
}

func (h *Handler) GetSimulate(ctx context.Context) (*npool.Simulate, error) {
	handler := &queryHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.querySimulate(cli); err != nil {
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

func (h *Handler) GetSimulates(ctx context.Context) ([]*npool.Simulate, uint32, error) {
	handler := &queryHandler{
		Handler: h,
	}

	var err error

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		handler.stmSelect, err = handler.querySimulates(cli)
		if err != nil {
			return err
		}
		handler.stmCount, err = handler.querySimulates(cli)
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
