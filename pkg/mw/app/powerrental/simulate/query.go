package appsimulatepowerrental

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"

	appsimulatepowerrentalcrud "github.com/NpoolPlatform/good-middleware/pkg/crud/app/powerrental/simulate"
	"github.com/NpoolPlatform/good-middleware/pkg/db"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent"
	entappgoodbase "github.com/NpoolPlatform/good-middleware/pkg/db/ent/appgoodbase"
	entappsimulatepowerrental "github.com/NpoolPlatform/good-middleware/pkg/db/ent/appsimulatepowerrental"
	entgoodbase "github.com/NpoolPlatform/good-middleware/pkg/db/ent/goodbase"
	entpowerrental "github.com/NpoolPlatform/good-middleware/pkg/db/ent/powerrental"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/app/powerrental/simulate"

	"github.com/shopspring/decimal"
)

type queryHandler struct {
	*Handler
	stmSelect *ent.AppSimulatePowerRentalSelect
	stmCount  *ent.AppSimulatePowerRentalSelect
	infos     []*npool.Simulate
	total     uint32
}

func (h *queryHandler) selectSimulate(stm *ent.AppSimulatePowerRentalQuery) *ent.AppSimulatePowerRentalSelect {
	return stm.Select(entappsimulatepowerrental.FieldID)
}

func (h *queryHandler) querySimulate(cli *ent.Client) error {
	if h.ID == nil && h.EntID == nil {
		return fmt.Errorf("invalid id")
	}
	stm := cli.AppSimulatePowerRental.Query().Where(entappsimulatepowerrental.DeletedAt(0))
	if h.ID != nil {
		stm.Where(entappsimulatepowerrental.ID(*h.ID))
	}
	if h.EntID != nil {
		stm.Where(entappsimulatepowerrental.EntID(*h.EntID))
	}
	h.stmSelect = h.selectSimulate(stm)
	return nil
}

func (h *queryHandler) querySimulates(cli *ent.Client) (*ent.AppSimulatePowerRentalSelect, error) {
	stm, err := appsimulatepowerrentalcrud.SetQueryConds(cli.AppSimulatePowerRental.Query(), h.AppSimulatePowerRentalConds)
	if err != nil {
		return nil, err
	}
	return h.selectSimulate(stm), nil
}

func (h *queryHandler) queryJoinMyself(s *sql.Selector) {
	t1 := sql.Table(entappsimulatepowerrental.Table)
	s.LeftJoin(t1).
		On(
			s.C(entappsimulatepowerrental.FieldID),
			t1.C(entappsimulatepowerrental.FieldID),
		).
		AppendSelect(
			t1.C(entappsimulatepowerrental.FieldID),
			t1.C(entappsimulatepowerrental.FieldEntID),
			t1.C(entappsimulatepowerrental.FieldAppGoodID),
			t1.C(entappsimulatepowerrental.FieldOrderUnits),
			t1.C(entappsimulatepowerrental.FieldOrderDurationSeconds),
			t1.C(entappsimulatepowerrental.FieldCreatedAt),
			t1.C(entappsimulatepowerrental.FieldUpdatedAt),
		)
}

func (h *queryHandler) queryJoinAppGoodBase(s *sql.Selector) {
	t1 := sql.Table(entappgoodbase.Table)
	t2 := sql.Table(entgoodbase.Table)
	t3 := sql.Table(entpowerrental.Table)

	s.LeftJoin(t1).
		On(
			s.C(entappsimulatepowerrental.FieldAppGoodID),
			t1.C(entappgoodbase.FieldEntID),
		).
		LeftJoin(t2).
		On(
			t1.C(entappgoodbase.FieldGoodID),
			t2.C(entgoodbase.FieldEntID),
		).
		LeftJoin(t3).
		On(
			t1.C(entappgoodbase.FieldGoodID),
			t3.C(entgoodbase.FieldEntID),
		).
		AppendSelect(
			t1.C(entappgoodbase.FieldAppID),
			sql.As(t1.C(entappgoodbase.FieldName), "app_good_name"),
			sql.As(t2.C(entgoodbase.FieldEntID), "good_id"),
			sql.As(t2.C(entgoodbase.FieldName), "good_name"),
		)
}

func (h *queryHandler) queryJoin() {
	h.stmSelect.Modify(func(s *sql.Selector) {
		h.queryJoinMyself(s)
		h.queryJoinAppGoodBase(s)
	})
	if h.stmCount == nil {
		return
	}
	h.stmCount.Modify(func(s *sql.Selector) {
		h.queryJoinAppGoodBase(s)
	})
}

func (h *queryHandler) scan(ctx context.Context) error {
	return h.stmSelect.Scan(ctx, &h.infos)
}

func (h *queryHandler) formalize() {
	for _, info := range h.infos {
		info.OrderUnits = func() string { amount, _ := decimal.NewFromString(info.OrderUnits); return amount.String() }()
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
