package topmostgoodconstraint

import (
	"fmt"

	"github.com/NpoolPlatform/good-middleware/pkg/db/ent"
	enttopmostgoodconstraint "github.com/NpoolPlatform/good-middleware/pkg/db/ent/topmostgoodconstraint"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	types "github.com/NpoolPlatform/message/npool/basetypes/good/v1"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Req struct {
	EntID         *uuid.UUID
	TopMostGoodID *uuid.UUID
	Constraint    *types.GoodTopMostConstraint
	TargetValue   *decimal.Decimal
	Index         *uint8
	DeletedAt     *uint32
}

//nolint:gocyclo,funlen
func CreateSet(c *ent.TopMostGoodConstraintCreate, req *Req) *ent.TopMostGoodConstraintCreate {
	if req.EntID != nil {
		c.SetEntID(*req.EntID)
	}
	if req.TopMostGoodID != nil {
		c.SetTopMostGoodID(*req.TopMostGoodID)
	}
	if req.Constraint != nil {
		c.SetConstraint(req.Constraint.String())
	}
	if req.TargetValue != nil {
		c.SetTargetValue(*req.TargetValue)
	}
	if req.Index != nil {
		c.SetIndex(*req.Index)
	}
	return c
}

//nolint:gocyclo,funlen
func UpdateSet(u *ent.TopMostGoodConstraintUpdateOne, req *Req) *ent.TopMostGoodConstraintUpdateOne {
	if req.TargetValue != nil {
		u.SetTargetValue(*req.TargetValue)
	}
	if req.Index != nil {
		u.SetIndex(*req.Index)
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
	TopMostGoodID  *cruder.Cond
	TopMostGoodIDs *cruder.Cond
}

//nolint:gocyclo,funlen
func SetQueryConds(q *ent.TopMostGoodConstraintQuery, conds *Conds) (*ent.TopMostGoodConstraintQuery, error) {
	q.Where(enttopmostgoodconstraint.DeletedAt(0))
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
			q.Where(enttopmostgoodconstraint.ID(id))
		default:
			return nil, fmt.Errorf("invalid topmostgoodconstraint field")
		}
	}
	if conds.IDs != nil {
		ids, ok := conds.IDs.Val.([]uint32)
		if !ok {
			return nil, fmt.Errorf("invalid ids")
		}
		switch conds.IDs.Op {
		case cruder.IN:
			q.Where(enttopmostgoodconstraint.IDIn(ids...))
		default:
			return nil, fmt.Errorf("invalid topmostgoodconstraint field")
		}
	}
	if conds.EntID != nil {
		id, ok := conds.EntID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid entid")
		}
		switch conds.EntID.Op {
		case cruder.EQ:
			q.Where(enttopmostgoodconstraint.EntID(id))
		default:
			return nil, fmt.Errorf("invalid topmostgoodconstraint field")
		}
	}
	if conds.EntIDs != nil {
		ids, ok := conds.EntIDs.Val.([]uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid entids")
		}
		switch conds.EntIDs.Op {
		case cruder.IN:
			q.Where(enttopmostgoodconstraint.EntIDIn(ids...))
		default:
			return nil, fmt.Errorf("invalid topmostgoodconstraint field")
		}
	}
	if conds.TopMostGoodID != nil {
		id, ok := conds.TopMostGoodID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid topmostgoodid")
		}
		switch conds.TopMostGoodID.Op {
		case cruder.EQ:
			q.Where(enttopmostgoodconstraint.TopMostGoodID(id))
		default:
			return nil, fmt.Errorf("invalid topmostgoodconstraint field")
		}
	}
	if conds.TopMostGoodIDs != nil {
		ids, ok := conds.TopMostGoodIDs.Val.([]uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid topmostgoodids")
		}
		switch conds.TopMostGoodIDs.Op {
		case cruder.IN:
			q.Where(enttopmostgoodconstraint.TopMostGoodIDIn(ids...))
		default:
			return nil, fmt.Errorf("invalid topmostgoodconstraint field")
		}
	}
	return q, nil
}
