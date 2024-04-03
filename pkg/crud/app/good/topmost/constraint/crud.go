package topmostconstraint

import (
	"fmt"

	"github.com/NpoolPlatform/good-middleware/pkg/db/ent"
	enttopmostconstraint "github.com/NpoolPlatform/good-middleware/pkg/db/ent/topmostconstraint"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	types "github.com/NpoolPlatform/message/npool/basetypes/good/v1"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Req struct {
	EntID       *uuid.UUID
	TopMostID   *uuid.UUID
	Constraint  *types.GoodTopMostConstraint
	TargetValue *decimal.Decimal
	Index       *uint8
	DeletedAt   *uint32
}

//nolint:gocyclo,funlen
func CreateSet(c *ent.TopMostConstraintCreate, req *Req) *ent.TopMostConstraintCreate {
	if req.EntID != nil {
		c.SetEntID(*req.EntID)
	}
	if req.TopMostID != nil {
		c.SetTopMostID(*req.TopMostID)
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
func UpdateSet(u *ent.TopMostConstraintUpdateOne, req *Req) *ent.TopMostConstraintUpdateOne {
	if req.Constraint != nil {
		u.SetConstraint(req.Constraint.String())
	}
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
	ID         *cruder.Cond
	IDs        *cruder.Cond
	EntID      *cruder.Cond
	EntIDs     *cruder.Cond
	TopMostID  *cruder.Cond
	TopMostIDs *cruder.Cond
}

//nolint:gocyclo,funlen
func SetQueryConds(q *ent.TopMostConstraintQuery, conds *Conds) (*ent.TopMostConstraintQuery, error) {
	q.Where(enttopmostconstraint.DeletedAt(0))
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
			q.Where(enttopmostconstraint.ID(id))
		default:
			return nil, fmt.Errorf("invalid topmostconstraint field")
		}
	}
	if conds.IDs != nil {
		ids, ok := conds.IDs.Val.([]uint32)
		if !ok {
			return nil, fmt.Errorf("invalid ids")
		}
		switch conds.IDs.Op {
		case cruder.IN:
			q.Where(enttopmostconstraint.IDIn(ids...))
		default:
			return nil, fmt.Errorf("invalid topmostconstraint field")
		}
	}
	if conds.EntID != nil {
		id, ok := conds.EntID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid entid")
		}
		switch conds.EntID.Op {
		case cruder.EQ:
			q.Where(enttopmostconstraint.EntID(id))
		default:
			return nil, fmt.Errorf("invalid topmostconstraint field")
		}
	}
	if conds.EntIDs != nil {
		ids, ok := conds.EntIDs.Val.([]uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid entids")
		}
		switch conds.EntIDs.Op {
		case cruder.IN:
			q.Where(enttopmostconstraint.EntIDIn(ids...))
		default:
			return nil, fmt.Errorf("invalid topmostconstraint field")
		}
	}
	if conds.TopMostID != nil {
		id, ok := conds.TopMostID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid topmostid")
		}
		switch conds.TopMostID.Op {
		case cruder.EQ:
			q.Where(enttopmostconstraint.TopMostID(id))
		default:
			return nil, fmt.Errorf("invalid topmostconstraint field")
		}
	}
	if conds.TopMostIDs != nil {
		ids, ok := conds.TopMostIDs.Val.([]uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid topmostids")
		}
		switch conds.TopMostIDs.Op {
		case cruder.IN:
			q.Where(enttopmostconstraint.TopMostIDIn(ids...))
		default:
			return nil, fmt.Errorf("invalid topmostconstraint field")
		}
	}
	return q, nil
}
