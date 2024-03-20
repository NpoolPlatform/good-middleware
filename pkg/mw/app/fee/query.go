//nolint:dupl
package fee

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"

	appgoodbasecrud "github.com/NpoolPlatform/good-middleware/pkg/crud/app/good/goodbase"
	"github.com/NpoolPlatform/good-middleware/pkg/db"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent"
	entappfee "github.com/NpoolPlatform/good-middleware/pkg/db/ent/appfee"
	entappgoodbase "github.com/NpoolPlatform/good-middleware/pkg/db/ent/appgoodbase"
	entfee "github.com/NpoolPlatform/good-middleware/pkg/db/ent/fee"
	entgoodbase "github.com/NpoolPlatform/good-middleware/pkg/db/ent/goodbase"
	types "github.com/NpoolPlatform/message/npool/basetypes/good/v1"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/app/fee"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type queryHandler struct {
	*Handler
	stmSelect *ent.AppGoodBaseSelect
	stmCount  *ent.AppGoodBaseSelect
	infos     []*npool.Fee
	total     uint32
}

func (h *queryHandler) selectAppGoodBase(stm *ent.AppGoodBaseQuery) *ent.AppGoodBaseSelect {
	return stm.Select(entappgoodbase.FieldID)
}

func (h *queryHandler) queryAppGoodBase(cli *ent.Client) error {
	if h.AppGoodID == nil {
		return fmt.Errorf("invalid appgoodid")
	}
	h.stmSelect = h.selectAppGoodBase(
		cli.AppGoodBase.
			Query().
			Where(
				entappgoodbase.DeletedAt(0),
				entappgoodbase.EntID(*h.AppGoodID),
			),
	)
	return nil
}

func (h *queryHandler) queryAppGoodBases(cli *ent.Client) (*ent.AppGoodBaseSelect, error) {
	stm, err := appgoodbasecrud.SetQueryConds(cli.AppGoodBase.Query(), h.AppGoodBaseConds)
	if err != nil {
		return nil, err
	}
	return h.selectAppGoodBase(stm), nil
}

func (h *queryHandler) queryJoinMyself(s *sql.Selector) error {
	t1 := sql.Table(entappgoodbase.Table)
	s.LeftJoin(t1).
		On(
			s.C(entappgoodbase.FieldID),
			t1.C(entappgoodbase.FieldID),
		).
		OnP(
			sql.EQ(t1.C(entappgoodbase.FieldDeletedAt), 0),
		).
		AppendSelect(
			sql.As(t1.C(entappgoodbase.FieldEntID), "app_good_id"),
			sql.As(t1.C(entappgoodbase.FieldAppID), "app_id"),
			sql.As(t1.C(entappgoodbase.FieldGoodID), "good_id"),
			sql.As(t1.C(entappgoodbase.FieldName), "name"),
			sql.As(t1.C(entappgoodbase.FieldBanner), "banner"),
			sql.As(t1.C(entappgoodbase.FieldCreatedAt), "created_at"),
			sql.As(t1.C(entappgoodbase.FieldUpdatedAt), "updated_at"),
		)

	t2 := sql.Table(entgoodbase.Table)
	s.LeftJoin(t2).
		On(
			t1.C(entappgoodbase.FieldGoodID),
			t2.C(entgoodbase.FieldEntID),
		).
		OnP(
			sql.Or(
				sql.EQ(t2.C(entgoodbase.FieldGoodType), types.GoodType_TechniqueServiceFee.String()),
				sql.EQ(t2.C(entgoodbase.FieldGoodType), types.GoodType_ElectricityFee.String()),
			),
		).
		OnP(
			sql.EQ(t2.C(entgoodbase.FieldDeletedAt), 0),
		)
	if h.GoodBaseConds != nil && h.GoodBaseConds.EntID != nil {
		uid, ok := h.GoodBaseConds.EntID.Val.(uuid.UUID)
		if !ok {
			return fmt.Errorf("invalid goodid")
		}
		s.OnP(sql.EQ(t2.C(entgoodbase.FieldEntID), uid))
	}
	if h.GoodBaseConds != nil && h.GoodBaseConds.EntIDs != nil {
		uids, ok := h.GoodBaseConds.EntIDs.Val.([]uuid.UUID)
		if !ok {
			return fmt.Errorf("invalid goodids")
		}
		s.OnP(sql.In(t2.C(entgoodbase.FieldEntID), uids))
	}
	s.AppendSelect(
		sql.As(t2.C(entgoodbase.FieldGoodType), "good_type"),
	)

	t3 := sql.Table(entfee.Table)
	s.LeftJoin(t3).
		On(
			t2.C(entgoodbase.FieldEntID),
			t3.C(entfee.FieldGoodID),
		).
		OnP(
			sql.EQ(t3.C(entfee.FieldDeletedAt), 0),
		)
	if h.FeeConds != nil && h.FeeConds.GoodID != nil {
		uid, ok := h.FeeConds.GoodID.Val.(uuid.UUID)
		if !ok {
			return fmt.Errorf("invalid goodid")
		}
		s.OnP(sql.EQ(t3.C(entfee.FieldGoodID), uid))
	}
	if h.FeeConds != nil && h.FeeConds.GoodIDs != nil {
		uids, ok := h.FeeConds.GoodIDs.Val.([]uuid.UUID)
		if !ok {
			return fmt.Errorf("invalid goodids")
		}
		s.OnP(sql.In(t3.C(entfee.FieldGoodID), uids))
	}
	s.AppendSelect(
		t3.C(entfee.FieldSettlementType),
		t3.C(entfee.FieldUnitValue),
		t3.C(entfee.FieldDurationType),
	)
	return nil
}

func (h *queryHandler) queryJoinAppFee(s *sql.Selector) error {
	t1 := sql.Table(entappfee.Table)
	s.LeftJoin(t1).
		On(
			s.C(entappgoodbase.FieldEntID),
			t1.C(entappfee.FieldAppGoodID),
		).
		OnP(
			sql.EQ(t1.C(entappfee.FieldDeletedAt), 0),
		)
	if h.AppFeeConds != nil && h.AppFeeConds.ID != nil {
		u, ok := h.AppFeeConds.ID.Val.(uint32)
		if !ok {
			return fmt.Errorf("invalid id")
		}
		s.OnP(sql.EQ(t1.C(entappfee.FieldID), u))
	}
	if h.AppFeeConds != nil && h.AppFeeConds.IDs != nil {
		ids, ok := h.AppFeeConds.IDs.Val.([]uint32)
		if !ok {
			return fmt.Errorf("invalid ids")
		}
		s.OnP(sql.In(t1.C(entappfee.FieldID), ids))
	}
	if h.AppFeeConds != nil && h.AppFeeConds.EntID != nil {
		uid, ok := h.AppFeeConds.EntID.Val.(uuid.UUID)
		if !ok {
			return fmt.Errorf("invalid entid")
		}
		s.OnP(sql.EQ(t1.C(entappfee.FieldEntID), uid))
	}
	if h.AppFeeConds != nil && h.AppFeeConds.EntIDs != nil {
		uids, ok := h.AppFeeConds.EntIDs.Val.([]uuid.UUID)
		if !ok {
			return fmt.Errorf("invalid entids")
		}
		s.OnP(sql.In(t1.C(entappfee.FieldEntID), uids))
	}
	if h.AppFeeConds != nil && h.AppFeeConds.AppGoodID != nil {
		uid, ok := h.AppFeeConds.AppGoodID.Val.(uuid.UUID)
		if !ok {
			return fmt.Errorf("invalid appgoodid")
		}
		s.OnP(sql.EQ(t1.C(entappfee.FieldAppGoodID), uid))
	}
	if h.AppFeeConds != nil && h.AppFeeConds.AppGoodIDs != nil {
		uids, ok := h.AppFeeConds.AppGoodIDs.Val.([]uuid.UUID)
		if !ok {
			return fmt.Errorf("invalid appgoodids")
		}
		s.OnP(sql.In(t1.C(entappfee.FieldAppGoodID), uids))
	}
	s.AppendSelect(
		sql.As(t1.C(entappfee.FieldID), "id"),
		sql.As(t1.C(entappfee.FieldEntID), "ent_id"),
		sql.As(t1.C(entappfee.FieldUnitValue), "unit_value"),
		sql.As(t1.C(entappfee.FieldMinOrderDuration), "min_order_duration"),
	)
	return nil
}

func (h *queryHandler) queryJoin() error {
	var err error
	h.stmSelect.Modify(func(s *sql.Selector) {
		h.queryJoinMyself(s)
		h.queryJoinAppFee(s)
	})
	if err != nil {
		return err
	}
	if h.stmCount == nil {
		return nil
	}
	h.stmCount.Modify(func(s *sql.Selector) {
	})
	if err != nil {
		return err
	}
	return nil
}

func (h *queryHandler) scan(ctx context.Context) error {
	return h.stmSelect.Scan(ctx, &h.infos)
}

//nolint:funlen,gocyclo
func (h *queryHandler) formalize() {
	for _, info := range h.infos {
		amount, _ := decimal.NewFromString(info.UnitValue)
		info.UnitValue = amount.String()
		info.GoodType = types.GoodType(types.GoodType_value[info.GoodTypeStr])
		info.SettlementType = types.GoodSettlementType(types.GoodSettlementType_value[info.SettlementTypeStr])
		info.DurationType = types.GoodDurationType(types.GoodDurationType_value[info.DurationTypeStr])
	}
}

func (h *Handler) GetFee(ctx context.Context) (*npool.Fee, error) {
	handler := &queryHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryAppGoodBase(cli); err != nil {
			return err
		}
		if err := handler.queryJoin(); err != nil {
			return err
		}
		return handler.scan(_ctx)
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

func (h *Handler) GetFees(ctx context.Context) ([]*npool.Fee, uint32, error) {
	handler := &queryHandler{
		Handler: h,
	}

	var err error

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		handler.stmSelect, err = handler.queryAppGoodBases(cli)
		if err != nil {
			return err
		}
		handler.stmCount, err = handler.queryAppGoodBases(cli)
		if err != nil {
			return err
		}

		if err := handler.queryJoin(); err != nil {
			return err
		}

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
