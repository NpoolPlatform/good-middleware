package score

import (
	"fmt"

	"github.com/NpoolPlatform/good-middleware/pkg/db/ent"
	entscore "github.com/NpoolPlatform/good-middleware/pkg/db/ent/score"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Req struct {
	ID     *uuid.UUID
	AppID  *uuid.UUID
	UserID *uuid.UUID
	GoodID *uuid.UUID
	Score  *decimal.Decimal
}

func CreateSet(c *ent.ScoreCreate, req *Req) *ent.ScoreCreate {
	if req.ID != nil {
		c.SetID(*req.ID)
	}
	if req.AppID != nil {
		c.SetAppID(*req.AppID)
	}
	if req.UserID != nil {
		c.SetUserID(*req.UserID)
	}
	if req.GoodID != nil {
		c.SetGoodID(*req.GoodID)
	}
	if req.Score != nil {
		c.SetScore(*req.Score)
	}
	return c
}

func UpdateSet(u *ent.ScoreUpdateOne, req *Req) *ent.ScoreUpdateOne {
	if req.Score != nil {
		u.SetScore(*req.Score)
	}
	return u
}

type Conds struct {
	ID      *cruder.Cond
	AppID   *cruder.Cond
	UserID  *cruder.Cond
	GoodID  *cruder.Cond
	GoodIDs *cruder.Cond
}

func SetQueryConds(q *ent.ScoreQuery, conds *Conds) (*ent.ScoreQuery, error) {
	q.Where(entscore.DeletedAt(0))
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
			q.Where(entscore.ID(id))
		default:
			return nil, fmt.Errorf("invalid score field")
		}
	}
	if conds.AppID != nil {
		id, ok := conds.AppID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid appid")
		}
		switch conds.AppID.Op {
		case cruder.EQ:
			q.Where(entscore.AppID(id))
		default:
			return nil, fmt.Errorf("invalid score field")
		}
	}
	if conds.UserID != nil {
		id, ok := conds.UserID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid userid")
		}
		switch conds.UserID.Op {
		case cruder.EQ:
			q.Where(entscore.UserID(id))
		default:
			return nil, fmt.Errorf("invalid score field")
		}
	}
	if conds.GoodID != nil {
		id, ok := conds.GoodID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid goodid")
		}
		switch conds.GoodID.Op {
		case cruder.EQ:
			q.Where(entscore.GoodID(id))
		default:
			return nil, fmt.Errorf("invalid score field")
		}
	}
	if conds.GoodIDs != nil {
		ids, ok := conds.GoodIDs.Val.([]uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid goodids")
		}
		switch conds.GoodIDs.Op {
		case cruder.EQ:
			q.Where(entscore.GoodIDIn(ids...))
		default:
			return nil, fmt.Errorf("invalid score field")
		}
	}
	return q, nil
}
