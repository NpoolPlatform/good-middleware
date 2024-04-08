package appfee

import (
	"fmt"

	"github.com/NpoolPlatform/good-middleware/pkg/db/ent"
	entappfee "github.com/NpoolPlatform/good-middleware/pkg/db/ent/appfee"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Req struct {
	EntID            *uuid.UUID
	AppGoodID        *uuid.UUID
	UnitValue        *decimal.Decimal // Per unit per duration, cash or from profit
	MinOrderDuration *uint32
	DeletedAt        *uint32
}

func CreateSet(c *ent.AppFeeCreate, req *Req) *ent.AppFeeCreate {
	if req.EntID != nil {
		c.SetEntID(*req.EntID)
	}
	if req.AppGoodID != nil {
		c.SetAppGoodID(*req.AppGoodID)
	}
	if req.UnitValue != nil {
		c.SetUnitValue(*req.UnitValue)
	}
	if req.MinOrderDuration != nil {
		c.SetMinOrderDuration(*req.MinOrderDuration)
	}
	return c
}

func UpdateSet(u *ent.AppFeeUpdateOne, req *Req) *ent.AppFeeUpdateOne {
	if req.UnitValue != nil {
		u.SetUnitValue(*req.UnitValue)
	}
	if req.MinOrderDuration != nil {
		u.SetMinOrderDuration(*req.MinOrderDuration)
	}
	if req.DeletedAt != nil {
		u.SetDeletedAt(*req.DeletedAt)
	}
	return u
}

type Conds struct {
	ID         *cruder.Cond
	IDs        *cruder.Cond
	EntID      *cruder.Cond
	EntIDs     *cruder.Cond
	AppGoodID  *cruder.Cond
	AppGoodIDs *cruder.Cond
}

//nolint:gocyclo,funlen
func SetQueryConds(q *ent.AppFeeQuery, conds *Conds) (*ent.AppFeeQuery, error) {
	q.Where(entappfee.DeletedAt(0))
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
			q.Where(entappfee.ID(id))
		default:
			return nil, fmt.Errorf("invalid appfee field")
		}
	}
	if conds.IDs != nil {
		ids, ok := conds.IDs.Val.([]uint32)
		if !ok {
			return nil, fmt.Errorf("invalid ids")
		}
		switch conds.IDs.Op {
		case cruder.IN:
			q.Where(entappfee.IDIn(ids...))
		default:
			return nil, fmt.Errorf("invalid appfee field")
		}
	}
	if conds.EntID != nil {
		id, ok := conds.EntID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid entid")
		}
		switch conds.EntID.Op {
		case cruder.EQ:
			q.Where(entappfee.EntID(id))
		default:
			return nil, fmt.Errorf("invalid appfee field")
		}
	}
	if conds.EntIDs != nil {
		ids, ok := conds.EntIDs.Val.([]uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid entids")
		}
		switch conds.EntIDs.Op {
		case cruder.IN:
			q.Where(entappfee.EntIDIn(ids...))
		default:
			return nil, fmt.Errorf("invalid appfee field")
		}
	}
	if conds.AppGoodID != nil {
		id, ok := conds.AppGoodID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid appgoodid")
		}
		switch conds.AppGoodID.Op {
		case cruder.EQ:
			q.Where(entappfee.AppGoodID(id))
		default:
			return nil, fmt.Errorf("invalid appfee field")
		}
	}
	if conds.AppGoodIDs != nil {
		ids, ok := conds.AppGoodIDs.Val.([]uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid appgoodids")
		}
		switch conds.AppGoodIDs.Op {
		case cruder.IN:
			q.Where(entappfee.AppGoodIDIn(ids...))
		default:
			return nil, fmt.Errorf("invalid appfee field")
		}
	}
	return q, nil
}
