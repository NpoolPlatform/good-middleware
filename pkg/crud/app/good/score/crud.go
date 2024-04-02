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
	EntID     *uuid.UUID
	UserID    *uuid.UUID
	AppGoodID *uuid.UUID
	Score     *decimal.Decimal
	CommentID *uuid.UUID
	DeletedAt *uint32
}

func CreateSet(c *ent.ScoreCreate, req *Req) *ent.ScoreCreate {
	if req.EntID != nil {
		c.SetEntID(*req.EntID)
	}
	if req.UserID != nil {
		c.SetUserID(*req.UserID)
	}
	if req.AppGoodID != nil {
		c.SetAppGoodID(*req.AppGoodID)
	}
	if req.Score != nil {
		c.SetScore(*req.Score)
	}
	if req.CommentID != nil {
		c.SetCommentID(*req.CommentID)
	}
	return c
}

func UpdateSet(u *ent.ScoreUpdateOne, req *Req) *ent.ScoreUpdateOne {
	if req.DeletedAt != nil {
		u.SetDeletedAt(*req.DeletedAt)
	}
	if req.Score != nil {
		u.SetScore(*req.Score)
	}
	if req.CommentID != nil {
		u.SetCommentID(*req.CommentID)
	}
	return u
}

type Conds struct {
	ID         *cruder.Cond
	EntID      *cruder.Cond
	UserID     *cruder.Cond
	AppGoodID  *cruder.Cond
	AppGoodIDs *cruder.Cond
}

//nolint
func SetQueryConds(q *ent.ScoreQuery, conds *Conds) (*ent.ScoreQuery, error) {
	q.Where(entscore.DeletedAt(0))
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
			q.Where(entscore.ID(id))
		default:
			return nil, fmt.Errorf("invalid score field")
		}
	}
	if conds.EntID != nil {
		id, ok := conds.EntID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid id")
		}
		switch conds.EntID.Op {
		case cruder.EQ:
			q.Where(entscore.EntID(id))
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
	if conds.AppGoodID != nil {
		id, ok := conds.AppGoodID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid appgoodid")
		}
		switch conds.AppGoodID.Op {
		case cruder.EQ:
			q.Where(entscore.AppGoodID(id))
		default:
			return nil, fmt.Errorf("invalid appgoodid field")
		}
	}
	if conds.AppGoodIDs != nil {
		ids, ok := conds.AppGoodIDs.Val.([]uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid goodids")
		}
		switch conds.AppGoodIDs.Op {
		case cruder.IN:
			q.Where(entscore.AppGoodIDIn(ids...))
		default:
			return nil, fmt.Errorf("invalid appgoodids field")
		}
	}
	return q, nil
}
