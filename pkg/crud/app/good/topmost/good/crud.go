//nolint:dupl
package topmostgood

import (
	"fmt"

	"github.com/NpoolPlatform/good-middleware/pkg/db/ent"
	enttopmostgood "github.com/NpoolPlatform/good-middleware/pkg/db/ent/topmostgood"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Req struct {
	ID           *uint32
	EntID        *uuid.UUID
	AppGoodID    *uuid.UUID
	TopMostID    *uuid.UUID
	DisplayIndex *uint32
	UnitPrice    *decimal.Decimal
	DeletedAt    *uint32
}

func CreateSet(c *ent.TopMostGoodCreate, req *Req) *ent.TopMostGoodCreate {
	if req.EntID != nil {
		c.SetEntID(*req.EntID)
	}
	if req.AppGoodID != nil {
		c.SetAppGoodID(*req.AppGoodID)
	}
	if req.TopMostID != nil {
		c.SetTopMostID(*req.TopMostID)
	}
	if req.DisplayIndex != nil {
		c.SetDisplayIndex(*req.DisplayIndex)
	}
	if req.UnitPrice != nil {
		c.SetUnitPrice(*req.UnitPrice)
	}
	return c
}

func UpdateSet(u *ent.TopMostGoodUpdateOne, req *Req) *ent.TopMostGoodUpdateOne {
	if req.AppGoodID != nil {
		u.SetAppGoodID(*req.AppGoodID)
	}
	if req.TopMostID != nil {
		u.SetTopMostID(*req.TopMostID)
	}
	if req.DisplayIndex != nil {
		u.SetDisplayIndex(*req.DisplayIndex)
	}
	if req.UnitPrice != nil {
		u.SetUnitPrice(*req.UnitPrice)
	}
	if req.DeletedAt != nil {
		u.SetDeletedAt(*req.DeletedAt)
	}
	return u
}

type Conds struct {
	ID          *cruder.Cond
	EntID       *cruder.Cond
	AppGoodID   *cruder.Cond
	AppGoodIDs  *cruder.Cond
	TopMostID   *cruder.Cond
	TopMostType *cruder.Cond
}

//nolint:gocyclo
func SetQueryConds(q *ent.TopMostGoodQuery, conds *Conds) (*ent.TopMostGoodQuery, error) {
	q.Where(enttopmostgood.DeletedAt(0))
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
			q.Where(enttopmostgood.ID(id))
		default:
			return nil, fmt.Errorf("invalid topmostgood field")
		}
	}
	if conds.EntID != nil {
		id, ok := conds.EntID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid id")
		}
		switch conds.EntID.Op {
		case cruder.EQ:
			q.Where(enttopmostgood.EntID(id))
		default:
			return nil, fmt.Errorf("invalid topmostgood field")
		}
	}
	if conds.AppGoodID != nil {
		id, ok := conds.AppGoodID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid appgoodid")
		}
		switch conds.AppGoodID.Op {
		case cruder.EQ:
			q.Where(enttopmostgood.AppGoodID(id))
		default:
			return nil, fmt.Errorf("invalid topmostgood field")
		}
	}
	if conds.AppGoodIDs != nil {
		ids, ok := conds.AppGoodIDs.Val.([]uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid appgoodids")
		}
		switch conds.AppGoodIDs.Op {
		case cruder.IN:
			q.Where(enttopmostgood.AppGoodIDIn(ids...))
		default:
			return nil, fmt.Errorf("invalid topmostgood field")
		}
	}
	if conds.TopMostID != nil {
		id, ok := conds.TopMostID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid topmostid")
		}
		switch conds.TopMostID.Op {
		case cruder.EQ:
			q.Where(enttopmostgood.TopMostID(id))
		default:
			return nil, fmt.Errorf("invalid topmostgood field")
		}
	}
	return q, nil
}
