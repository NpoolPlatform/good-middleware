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
	AppID            *uuid.UUID
	GoodID           *uuid.UUID
	UnitValue        *decimal.Decimal // Per unit per duration, cash or from profit
	MinOrderDuration *uint32
	DeletedAt        *uint32
}

//nolint:gocyclo,funlen
func CreateSet(c *ent.AppFeeCreate, req *Req) *ent.AppFeeCreate {
	if req.EntID != nil {
		c.SetEntID(*req.EntID)
	}
	if req.AppID != nil {
		c.SetAppID(*req.AppID)
	}
	if req.GoodID != nil {
		c.SetGoodID(*req.GoodID)
	}
	if req.UnitValue != nil {
		c.SetUnitValue(*req.UnitValue)
	}
	if req.MinOrderDuration != nil {
		c.SetMinOrderDuration(*req.MinOrderDuration)
	}
	return c
}

//nolint:gocyclo,funlen
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
	ID      *cruder.Cond
	IDs     *cruder.Cond
	EntID   *cruder.Cond
	EntIDs  *cruder.Cond
	AppID   *cruder.Cond
	AppIDs  *cruder.Cond
	GoodID  *cruder.Cond
	GoodIDs *cruder.Cond
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
	if conds.AppID != nil {
		id, ok := conds.AppID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid appid")
		}
		switch conds.AppID.Op {
		case cruder.EQ:
			q.Where(entappfee.AppID(id))
		default:
			return nil, fmt.Errorf("invalid appfee field")
		}
	}
	if conds.AppIDs != nil {
		ids, ok := conds.AppIDs.Val.([]uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid appids")
		}
		switch conds.AppIDs.Op {
		case cruder.IN:
			q.Where(entappfee.AppIDIn(ids...))
		default:
			return nil, fmt.Errorf("invalid appfee field")
		}
	}
	if conds.GoodID != nil {
		id, ok := conds.GoodID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid goodid")
		}
		switch conds.GoodID.Op {
		case cruder.EQ:
			q.Where(entappfee.GoodID(id))
		default:
			return nil, fmt.Errorf("invalid appfee field")
		}
	}
	if conds.GoodIDs != nil {
		ids, ok := conds.GoodIDs.Val.([]uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid goodids")
		}
		switch conds.GoodIDs.Op {
		case cruder.IN:
			q.Where(entappfee.GoodIDIn(ids...))
		default:
			return nil, fmt.Errorf("invalid appfee field")
		}
	}
	return q, nil
}
