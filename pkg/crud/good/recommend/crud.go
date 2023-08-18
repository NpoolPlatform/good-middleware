package recommend

import (
	"fmt"

	"github.com/NpoolPlatform/good-middleware/pkg/db/ent"
	entrecommend "github.com/NpoolPlatform/good-middleware/pkg/db/ent/recommend"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Req struct {
	ID             *uuid.UUID
	AppID          *uuid.UUID
	GoodID         *uuid.UUID
	RecommenderID  *uuid.UUID
	Message        *string
	RecommendIndex *decimal.Decimal
	DeletedAt      *uint32
}

func CreateSet(c *ent.RecommendCreate, req *Req) *ent.RecommendCreate {
	if req.ID != nil {
		c.SetID(*req.ID)
	}
	if req.AppID != nil {
		c.SetAppID(*req.AppID)
	}
	if req.GoodID != nil {
		c.SetGoodID(*req.GoodID)
	}
	if req.RecommenderID != nil {
		c.SetRecommenderID(*req.RecommenderID)
	}
	if req.Message != nil {
		c.SetMessage(*req.Message)
	}
	if req.RecommendIndex != nil {
		c.SetRecommendIndex(*req.RecommendIndex)
	}
	return c
}

func UpdateSet(u *ent.RecommendUpdateOne, req *Req) *ent.RecommendUpdateOne {
	if req.Message != nil {
		u.SetMessage(*req.Message)
	}
	if req.RecommendIndex != nil {
		u.SetRecommendIndex(*req.RecommendIndex)
	}
	if req.DeletedAt != nil {
		u.SetDeletedAt(*req.DeletedAt)
	}
	return u
}

type Conds struct {
	ID            *cruder.Cond
	AppID         *cruder.Cond
	RecommenderID *cruder.Cond
	GoodID        *cruder.Cond
	GoodIDs       *cruder.Cond
}

//nolint:gocyclo
func SetQueryConds(q *ent.RecommendQuery, conds *Conds) (*ent.RecommendQuery, error) {
	q.Where(entrecommend.DeletedAt(0))
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
			q.Where(entrecommend.ID(id))
		default:
			return nil, fmt.Errorf("invalid recommend field")
		}
	}
	if conds.AppID != nil {
		id, ok := conds.AppID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid appid")
		}
		switch conds.AppID.Op {
		case cruder.EQ:
			q.Where(entrecommend.AppID(id))
		default:
			return nil, fmt.Errorf("invalid recommend field")
		}
	}
	if conds.RecommenderID != nil {
		id, ok := conds.RecommenderID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid recommenderid")
		}
		switch conds.RecommenderID.Op {
		case cruder.EQ:
			q.Where(entrecommend.RecommenderID(id))
		default:
			return nil, fmt.Errorf("invalid recommend field")
		}
	}
	if conds.GoodID != nil {
		id, ok := conds.GoodID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid goodid")
		}
		switch conds.GoodID.Op {
		case cruder.EQ:
			q.Where(entrecommend.GoodID(id))
		default:
			return nil, fmt.Errorf("invalid recommend field")
		}
	}
	if conds.GoodIDs != nil {
		ids, ok := conds.GoodIDs.Val.([]uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid goodids")
		}
		switch conds.GoodIDs.Op {
		case cruder.IN:
			q.Where(entrecommend.GoodIDIn(ids...))
		default:
			return nil, fmt.Errorf("invalid recommend field")
		}
	}
	return q, nil
}
