package fee

import (
	"fmt"

	"entgo.io/ent/dialect/sql"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	goodbasecrud "github.com/NpoolPlatform/good-middleware/pkg/crud/good/goodbase"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent"
	entfee "github.com/NpoolPlatform/good-middleware/pkg/db/ent/fee"
	entgoodbase "github.com/NpoolPlatform/good-middleware/pkg/db/ent/goodbase"
	types "github.com/NpoolPlatform/message/npool/basetypes/good/v1"

	"github.com/google/uuid"
)

type baseQueryHandler struct {
	*Handler
	stmSelect *ent.GoodBaseSelect
}

func (h *baseQueryHandler) selectGoodBase(stm *ent.GoodBaseQuery) *ent.GoodBaseSelect {
	return stm.Select(entgoodbase.FieldID)
}

func (h *baseQueryHandler) queryGoodBase(cli *ent.Client) error {
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

func (h *baseQueryHandler) queryGoodBases(cli *ent.Client) (*ent.GoodBaseSelect, error) {
	stm, err := goodbasecrud.SetQueryConds(cli.GoodBase.Query(), h.GoodBaseConds)
	if err != nil {
		return nil, err
	}
	return h.selectGoodBase(stm), nil
}

func (h *baseQueryHandler) queryJoinMyself(s *sql.Selector) {
	t := sql.Table(entgoodbase.Table)
	s.LeftJoin(t).
		On(
			s.C(entgoodbase.FieldID),
			t.C(entgoodbase.FieldID),
		).
		AppendSelect(
			t.C(entgoodbase.FieldGoodType),
			t.C(entgoodbase.FieldName),
			t.C(entgoodbase.FieldCreatedAt),
			t.C(entgoodbase.FieldUpdatedAt),
		)
}

func (h *baseQueryHandler) queryJoinFee(s *sql.Selector) error {
	t1 := sql.Table(entfee.Table)
	s.LeftJoin(t1).
		On(
			s.C(entgoodbase.FieldEntID),
			t1.C(entfee.FieldGoodID),
		).
		OnP(
			sql.EQ(t1.C(entfee.FieldDeletedAt), 0),
		)
	if h.FeeConds != nil && h.FeeConds.ID != nil {
		u, ok := h.FeeConds.ID.Val.(uint32)
		if !ok {
			return fmt.Errorf("invalid id")
		}
		s.OnP(sql.EQ(t1.C(entfee.FieldID), u))
	}
	if h.FeeConds != nil && h.FeeConds.IDs != nil {
		ids, ok := h.FeeConds.IDs.Val.([]uint32)
		if !ok {
			return fmt.Errorf("invalid ids")
		}
		s.OnP(sql.In(t1.C(entfee.FieldID), ids))
	}
	if h.FeeConds != nil && h.FeeConds.EntID != nil {
		uid, ok := h.FeeConds.EntID.Val.(uuid.UUID)
		if !ok {
			return fmt.Errorf("invalid entid")
		}
		s.OnP(sql.EQ(t1.C(entfee.FieldEntID), uid))
	}
	if h.FeeConds != nil && h.FeeConds.EntIDs != nil {
		uids, ok := h.FeeConds.EntIDs.Val.([]uuid.UUID)
		if !ok {
			return fmt.Errorf("invalid entids")
		}
		s.OnP(sql.In(t1.C(entfee.FieldEntID), uids))
	}
	s.AppendSelect(
		t1.C(entfee.FieldID),
		t1.C(entfee.FieldEntID),
		t1.C(entfee.FieldGoodID),
		t1.C(entfee.FieldSettlementType),
		t1.C(entfee.FieldUnitValue),
		t1.C(entfee.FieldDurationDisplayType),
	)
	return nil
}

func (h *baseQueryHandler) queryJoin() {
	h.stmSelect.Modify(func(s *sql.Selector) {
		h.queryJoinMyself(s)
		if err := h.queryJoinFee(s); err != nil {
			logger.Sugar().Errorw("queryJoinFee", "Error", err)
		}
	})
}
