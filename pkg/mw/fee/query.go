//nolint:dupl
package fee

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"

	goodbasecrud "github.com/NpoolPlatform/good-middleware/pkg/crud/good/goodbase"
	"github.com/NpoolPlatform/good-middleware/pkg/db"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent"
	entfee "github.com/NpoolPlatform/good-middleware/pkg/db/ent/fee"
	entgoodbase "github.com/NpoolPlatform/good-middleware/pkg/db/ent/goodbase"
	types "github.com/NpoolPlatform/message/npool/basetypes/good/v1"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/fee"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type queryHandler struct {
	*Handler
	stmSelect *ent.GoodBaseSelect
	stmCount  *ent.GoodBaseSelect
	infos     []*npool.Fee
	total     uint32
}

func (h *queryHandler) selectGoodBase(stm *ent.GoodBaseQuery) *ent.GoodBaseSelect {
	return stm.Select(entgoodbase.FieldID)
}

func (h *queryHandler) queryGoodBase(cli *ent.Client) error {
	if h.GoodID == nil {
		return fmt.Errorf("invalid goodid")
	}
	h.stmSelect = h.selectGoodBase(
		cli.GoodBase.
			Query().
			Where(
				entgoodbase.DeletedAt(0),
				entgoodbase.EntID(*h.GoodID),
				entgoodbase.Or(
					entgoodbase.GoodType(types.GoodType_TechniqueServiceFee.String()),
					entgoodbase.GoodType(types.GoodType_ElectricityFee.String()),
				),
			),
	)
	return nil
}

func (h *queryHandler) queryGoodBases(cli *ent.Client) (*ent.GoodBaseSelect, error) {
	stm, err := goodbasecrud.SetQueryConds(cli.GoodBase.Query(), h.GoodBaseConds)
	if err != nil {
		return nil, err
	}
	return h.selectGoodBase(stm), nil
}

func (h *queryHandler) queryJoinMyself(s *sql.Selector) {
	t := sql.Table(entgoodbase.Table)
	s.LeftJoin(t).
		On(
			s.C(entgoodbase.FieldID),
			t.C(entgoodbase.FieldID),
		).
		AppendSelect(
			sql.As(t.C(entgoodbase.FieldGoodType), "good_type"),
			sql.As(t.C(entgoodbase.FieldName), "name"),
			sql.As(t.C(entgoodbase.FieldCreatedAt), "created_at"),
			sql.As(t.C(entgoodbase.FieldUpdatedAt), "updated_at"),
		)
}

func (h *queryHandler) queryJoinFee(s *sql.Selector) error {
	t1 := sql.Table(entfee.Table)
	s.LeftJoin(t1).
		On(
			s.C(entgoodbase.FieldEntID),
			t1.C(entfee.FieldGoodID),
		).
		OnP(
			sql.EQ(t1.C(entfee.FieldDeletedAt), 0),
		)
	if h.FeeConds.ID != nil {
		u, ok := h.FeeConds.ID.Val.(uint32)
		if !ok {
			return fmt.Errorf("invalid id")
		}
		s.OnP(sql.EQ(t1.C(entfee.FieldID), u))
	}
	if h.FeeConds.IDs != nil {
		ids, ok := h.FeeConds.IDs.Val.([]uint32)
		if !ok {
			return fmt.Errorf("invalid ids")
		}
		s.OnP(sql.In(t1.C(entfee.FieldID), ids))
	}
	if h.FeeConds.EntID != nil {
		uid, ok := h.FeeConds.EntID.Val.(uuid.UUID)
		if !ok {
			return fmt.Errorf("invalid entid")
		}
		s.OnP(sql.EQ(t1.C(entfee.FieldEntID), uid))
	}
	if h.FeeConds.EntIDs != nil {
		uids, ok := h.FeeConds.EntIDs.Val.([]uuid.UUID)
		if !ok {
			return fmt.Errorf("invalid entids")
		}
		s.OnP(sql.In(t1.C(entfee.FieldEntID), uids))
	}
	s.AppendSelect(
		sql.As(t1.C(entfee.FieldID), "id"),
		sql.As(t1.C(entfee.FieldEntID), "ent_id"),
		sql.As(t1.C(entfee.FieldGoodID), "good_id"),
		t1.C(entfee.FieldSettlementType),
		t1.C(entfee.FieldUnitValue),
		t1.C(entfee.FieldDurationType),
	)
	return nil
}

func (h *queryHandler) queryJoin() error {
	var err error
	h.stmSelect.Modify(func(s *sql.Selector) {
		h.queryJoinMyself(s)
		err = h.queryJoinFee(s)
	})
	if err != nil {
		return err
	}
	if h.stmCount == nil {
		return nil
	}
	h.stmCount.Modify(func(s *sql.Selector) {
		err = h.queryJoinFee(s)
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
	}
}

func (h *Handler) GetFee(ctx context.Context) (*npool.Fee, error) {
	handler := &queryHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryGoodBase(cli); err != nil {
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
		handler.stmSelect, err = handler.queryGoodBases(cli)
		if err != nil {
			return err
		}
		handler.stmCount, err = handler.queryGoodBases(cli)
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
