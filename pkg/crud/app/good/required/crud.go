package requiredappgood

import (
	"fmt"

	"github.com/NpoolPlatform/good-middleware/pkg/db/ent"
	entrequiredappgood "github.com/NpoolPlatform/good-middleware/pkg/db/ent/requiredappgood"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	"github.com/google/uuid"
)

type Req struct {
	ID                *uint32
	EntID             *uuid.UUID
	MainAppGoodID     *uuid.UUID
	RequiredAppGoodID *uuid.UUID
	Must              *bool
	DeletedAt         *uint32
}

func CreateSet(c *ent.RequiredAppGoodCreate, req *Req) *ent.RequiredAppGoodCreate {
	if req.EntID != nil {
		c.SetEntID(*req.EntID)
	}
	if req.MainAppGoodID != nil {
		c.SetMainAppGoodID(*req.MainAppGoodID)
	}
	if req.RequiredAppGoodID != nil {
		c.SetRequiredAppGoodID(*req.RequiredAppGoodID)
	}
	if req.Must != nil {
		c.SetMust(*req.Must)
	}
	return c
}

func UpdateSet(u *ent.RequiredAppGoodUpdateOne, req *Req) *ent.RequiredAppGoodUpdateOne {
	if req.Must != nil {
		u.SetMust(*req.Must)
	}
	if req.DeletedAt != nil {
		u.SetDeletedAt(*req.DeletedAt)
	}
	return u
}

type Conds struct {
	ID                *cruder.Cond
	EntID             *cruder.Cond
	MainAppGoodID     *cruder.Cond
	RequiredAppGoodID *cruder.Cond
	GoodID            *cruder.Cond
	GoodIDs           *cruder.Cond
	Must              *cruder.Cond
}

//nolint:funlen,gocyclo
func SetQueryConds(q *ent.RequiredAppGoodQuery, conds *Conds) (*ent.RequiredAppGoodQuery, error) {
	q.Where(entrequiredappgood.DeletedAt(0))
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
			q.Where(entrequiredappgood.ID(id))
		default:
			return nil, fmt.Errorf("invalid requiredappgood field")
		}
	}
	if conds.EntID != nil {
		id, ok := conds.EntID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid id")
		}
		switch conds.EntID.Op {
		case cruder.EQ:
			q.Where(entrequiredappgood.EntID(id))
		default:
			return nil, fmt.Errorf("invalid requiredappgood field")
		}
	}
	if conds.MainAppGoodID != nil {
		id, ok := conds.MainAppGoodID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid mainappgoodid")
		}
		switch conds.MainAppGoodID.Op {
		case cruder.EQ:
			q.Where(entrequiredappgood.MainAppGoodID(id))
		default:
			return nil, fmt.Errorf("invalid requiredappgood field")
		}
	}
	if conds.RequiredAppGoodID != nil {
		id, ok := conds.RequiredAppGoodID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid requiredappgoodid")
		}
		switch conds.RequiredAppGoodID.Op {
		case cruder.EQ:
			q.Where(entrequiredappgood.RequiredAppGoodID(id))
		default:
			return nil, fmt.Errorf("invalid requiredappgood field")
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
				entrequiredappgood.Or(
					entrequiredappgood.MainAppGoodID(id),
					entrequiredappgood.RequiredAppGoodID(id),
				),
			)
		default:
			return nil, fmt.Errorf("invalid requiredappgood field")
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
				entrequiredappgood.Or(
					entrequiredappgood.MainAppGoodIDIn(ids...),
					entrequiredappgood.RequiredAppGoodIDIn(ids...),
				),
			)
		default:
			return nil, fmt.Errorf("invalid requiredappgood field")
		}
	}
	if conds.Must != nil {
		must, ok := conds.Must.Val.(bool)
		if !ok {
			return nil, fmt.Errorf("invalid must")
		}
		switch conds.Must.Op {
		case cruder.EQ:
			q.Where(entrequiredappgood.MustEQ(must))
		default:
			return nil, fmt.Errorf("invalid requiredappgood field")
		}
	}
	return q, nil
}
