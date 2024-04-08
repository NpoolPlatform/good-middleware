package fee

import (
	"fmt"

	"github.com/NpoolPlatform/good-middleware/pkg/db/ent"
	entfee "github.com/NpoolPlatform/good-middleware/pkg/db/ent/fee"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	types "github.com/NpoolPlatform/message/npool/basetypes/good/v1"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Req struct {
	EntID          *uuid.UUID
	GoodID         *uuid.UUID
	SettlementType *types.GoodSettlementType
	UnitValue      *decimal.Decimal // Per unit per duration, cash or from profit
	DurationType   *types.GoodDurationType
	DeletedAt      *uint32
}

func CreateSet(c *ent.FeeCreate, req *Req) *ent.FeeCreate {
	if req.EntID != nil {
		c.SetEntID(*req.EntID)
	}
	if req.GoodID != nil {
		c.SetGoodID(*req.GoodID)
	}
	if req.SettlementType != nil {
		c.SetSettlementType(req.SettlementType.String())
	}
	if req.UnitValue != nil {
		c.SetUnitValue(*req.UnitValue)
	}
	if req.DurationType != nil {
		c.SetDurationType(req.DurationType.String())
	}
	return c
}

func UpdateSet(u *ent.FeeUpdateOne, req *Req) *ent.FeeUpdateOne {
	if req.SettlementType != nil {
		u.SetSettlementType(req.SettlementType.String())
	}
	if req.UnitValue != nil {
		u.SetUnitValue(*req.UnitValue)
	}
	if req.DurationType != nil {
		u.SetDurationType(req.DurationType.String())
	}
	if req.DeletedAt != nil {
		u.SetDeletedAt(*req.DeletedAt)
	}
	return u
}

type Conds struct {
	ID             *cruder.Cond
	IDs            *cruder.Cond
	EntID          *cruder.Cond
	EntIDs         *cruder.Cond
	GoodID         *cruder.Cond
	GoodIDs        *cruder.Cond
	SettlementType *cruder.Cond
}

//nolint:gocyclo,funlen
func SetQueryConds(q *ent.FeeQuery, conds *Conds) (*ent.FeeQuery, error) {
	q.Where(entfee.DeletedAt(0))
	if conds == nil {
		return q, nil
	}
	if conds.ID != nil {
		id, ok := conds.ID.Val.(uint32)
		if !ok {
			return nil, fmt.Errorf("invalid id")
		}
		switch conds.ID.Op {
		case cruder.EQ:
			q.Where(entfee.ID(id))
		default:
			return nil, fmt.Errorf("invalid fee field")
		}
	}
	if conds.IDs != nil {
		ids, ok := conds.IDs.Val.([]uint32)
		if !ok {
			return nil, fmt.Errorf("invalid ids")
		}
		switch conds.IDs.Op {
		case cruder.IN:
			q.Where(entfee.IDIn(ids...))
		default:
			return nil, fmt.Errorf("invalid fee field")
		}
	}
	if conds.EntID != nil {
		id, ok := conds.EntID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid entid")
		}
		switch conds.EntID.Op {
		case cruder.EQ:
			q.Where(entfee.EntID(id))
		default:
			return nil, fmt.Errorf("invalid fee field")
		}
	}
	if conds.EntIDs != nil {
		ids, ok := conds.EntIDs.Val.([]uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid entids")
		}
		switch conds.EntIDs.Op {
		case cruder.IN:
			q.Where(entfee.EntIDIn(ids...))
		default:
			return nil, fmt.Errorf("invalid fee field")
		}
	}
	if conds.GoodID != nil {
		id, ok := conds.GoodID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid goodid")
		}
		switch conds.GoodID.Op {
		case cruder.EQ:
			q.Where(entfee.GoodID(id))
		default:
			return nil, fmt.Errorf("invalid fee field")
		}
	}
	if conds.GoodIDs != nil {
		ids, ok := conds.GoodIDs.Val.([]uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid goodids")
		}
		switch conds.GoodIDs.Op {
		case cruder.IN:
			q.Where(entfee.GoodIDIn(ids...))
		default:
			return nil, fmt.Errorf("invalid fee field")
		}
	}
	if conds.SettlementType != nil {
		e, ok := conds.SettlementType.Val.(types.GoodSettlementType)
		if !ok {
			return nil, fmt.Errorf("invalid settlementtype")
		}
		switch conds.SettlementType.Op {
		case cruder.EQ:
			q.Where(entfee.SettlementType(e.String()))
		case cruder.NEQ:
			q.Where(entfee.SettlementTypeNEQ(e.String()))
		default:
			return nil, fmt.Errorf("invalid settlementtype")
		}
	}
	return q, nil
}
