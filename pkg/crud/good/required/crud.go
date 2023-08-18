package requiredgood

import (
	"fmt"

	"github.com/NpoolPlatform/good-middleware/pkg/db/ent"
	entrequiredgood "github.com/NpoolPlatform/good-middleware/pkg/db/ent/requiredgood"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	"github.com/google/uuid"
)

type Req struct {
	ID             *uuid.UUID
	MainGoodID     *uuid.UUID
	RequiredGoodID *uuid.UUID
	Must           *bool
	Commission     *bool
	DeletedAt      *uint32
}

func CreateSet(c *ent.RequiredGoodCreate, req *Req) *ent.RequiredGoodCreate {
	if req.ID != nil {
		c.SetID(*req.ID)
	}
	if req.MainGoodID != nil {
		c.SetMainGoodID(*req.MainGoodID)
	}
	if req.RequiredGoodID != nil {
		c.SetRequiredGoodID(*req.RequiredGoodID)
	}
	if req.Must != nil {
		c.SetMust(*req.Must)
	}
	if req.Commission != nil {
		c.SetCommission(*req.Commission)
	}
	return c
}

func UpdateSet(u *ent.RequiredGoodUpdateOne, req *Req) *ent.RequiredGoodUpdateOne {
	if req.Must != nil {
		u.SetMust(*req.Must)
	}
	if req.Commission != nil {
		u.SetCommission(*req.Commission)
	}
	if req.DeletedAt != nil {
		u.SetDeletedAt(*req.DeletedAt)
	}
	return u
}

type Conds struct {
	ID             *cruder.Cond
	MainGoodID     *cruder.Cond
	RequiredGoodID *cruder.Cond
	GoodID         *cruder.Cond
	GoodIDs        *cruder.Cond
}

//nolint:gocyclo
func SetQueryConds(q *ent.RequiredGoodQuery, conds *Conds) (*ent.RequiredGoodQuery, error) {
	q.Where(entrequiredgood.DeletedAt(0))
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
			q.Where(entrequiredgood.ID(id))
		default:
			return nil, fmt.Errorf("invalid requiredgood field")
		}
	}
	if conds.MainGoodID != nil {
		id, ok := conds.MainGoodID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid maingoodid")
		}
		switch conds.MainGoodID.Op {
		case cruder.EQ:
			q.Where(entrequiredgood.MainGoodID(id))
		default:
			return nil, fmt.Errorf("invalid requiredgood field")
		}
	}
	if conds.RequiredGoodID != nil {
		id, ok := conds.RequiredGoodID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid requiredgoodid")
		}
		switch conds.RequiredGoodID.Op {
		case cruder.EQ:
			q.Where(entrequiredgood.RequiredGoodID(id))
		default:
			return nil, fmt.Errorf("invalid requiredgood field")
		}
	}
	if conds.GoodID != nil {
		id, ok := conds.GoodID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid goodid")
		}
		switch conds.GoodID.Op {
		case cruder.EQ:
			q.Where(
				entrequiredgood.Or(
					entrequiredgood.MainGoodID(id),
					entrequiredgood.RequiredGoodID(id),
				),
			)
		default:
			return nil, fmt.Errorf("invalid requiredgood field")
		}
	}
	if conds.GoodIDs != nil {
		ids, ok := conds.GoodIDs.Val.([]uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid goodids")
		}
		switch conds.GoodIDs.Op {
		case cruder.IN:
			q.Where(
				entrequiredgood.Or(
					entrequiredgood.MainGoodIDIn(ids...),
					entrequiredgood.RequiredGoodIDIn(ids...),
				),
			)
		default:
			return nil, fmt.Errorf("invalid requiredgood field")
		}
	}
	return q, nil
}
