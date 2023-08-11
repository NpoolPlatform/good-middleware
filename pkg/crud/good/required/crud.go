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
	AppID          *uuid.UUID
	MainGoodID     *uuid.UUID
	RequiredGoodID *uuid.UUID
	Must           *bool
	Commission     *bool
}

func CreateSet(c *ent.RequiredGoodCreate, req *Req) *ent.RequiredGoodCreate {
	if req.ID != nil {
		c.SetID(*req.ID)
	}
	if req.AppID != nil {
		c.SetAppID(*req.AppID)
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

type Conds struct {
	ID          *cruder.Cond
	AppID       *cruder.Cond
	MainGoodID  *cruder.Cond
	MainGoodIDs *cruder.Cond
}

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
	if conds.AppID != nil {
		id, ok := conds.AppID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid appid")
		}
		switch conds.AppID.Op {
		case cruder.EQ:
			q.Where(entrequiredgood.AppID(id))
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
	if conds.MainGoodIDs != nil {
		ids, ok := conds.MainGoodIDs.Val.([]uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid maingoodids")
		}
		switch conds.MainGoodIDs.Op {
		case cruder.IN:
			q.Where(entrequiredgood.MainGoodIDIn(ids...))
		default:
			return nil, fmt.Errorf("invalid requiredgood field")
		}
	}
	return q, nil
}
