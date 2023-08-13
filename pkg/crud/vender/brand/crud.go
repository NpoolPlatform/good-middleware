package brand

import (
	"fmt"

	"github.com/NpoolPlatform/good-middleware/pkg/db/ent"
	entvendorbrand "github.com/NpoolPlatform/good-middleware/pkg/db/ent/vendorbrand"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	"github.com/google/uuid"
)

type Req struct {
	ID        *uuid.UUID
	Name      *string
	Logo      *string
	DeletedAt *uint32
}

func CreateSet(c *ent.VendorBrandCreate, req *Req) *ent.VendorBrandCreate {
	if req.ID != nil {
		c.SetID(*req.ID)
	}
	if req.Name != nil {
		c.SetName(*req.Name)
	}
	if req.Logo != nil {
		c.SetLogo(*req.Logo)
	}
	return c
}

func UpdateSet(u *ent.VendorBrandUpdateOne, req *Req) *ent.VendorBrandUpdateOne {
	if req.Name != nil {
		u.SetName(*req.Name)
	}
	if req.Logo != nil {
		u.SetLogo(*req.Logo)
	}
	if req.DeletedAt != nil {
		u.SetDeletedAt(*req.DeletedAt)
	}
	return u
}

type Conds struct {
	ID   *cruder.Cond
	Name *cruder.Cond
}

func SetQueryConds(q *ent.VendorBrandQuery, conds *Conds) (*ent.VendorBrandQuery, error) {
	q.Where(entvendorbrand.DeletedAt(0))
	if conds == nil {
		return q, nil
	}
	if conds.ID != nil {
		id, ok := conds.ID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid id")
		}
		switch conds.ID.Op {
		case cruder.EQ:
			q.Where(entvendorbrand.ID(id))
		case cruder.NEQ:
			q.Where(entvendorbrand.IDNEQ(id))
		default:
			return nil, fmt.Errorf("invalid vendorbrand field")
		}
	}
	if conds.Name != nil {
		name, ok := conds.Name.Val.(string)
		if !ok {
			return nil, fmt.Errorf("invalid name")
		}
		switch conds.Name.Op {
		case cruder.EQ:
			q.Where(entvendorbrand.Name(name))
		default:
			return nil, fmt.Errorf("invalid vendorbrand field")
		}
	}
	return q, nil
}
