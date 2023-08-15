package comment

import (
	"fmt"

	"github.com/NpoolPlatform/good-middleware/pkg/db/ent"
	entcomment "github.com/NpoolPlatform/good-middleware/pkg/db/ent/comment"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	"github.com/google/uuid"
)

type Req struct {
	ID        *uuid.UUID
	AppID     *uuid.UUID
	UserID    *uuid.UUID
	GoodID    *uuid.UUID
	OrderID   *uuid.UUID
	Content   *string
	ReplyToID *uuid.UUID
	DeletedAt *uint32
}

func CreateSet(c *ent.CommentCreate, req *Req) *ent.CommentCreate {
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
	if req.OrderID != nil {
		c.SetOrderID(*req.OrderID)
	}
	if req.Content != nil {
		c.SetContent(*req.Content)
	}
	if req.ReplyToID != nil {
		c.SetReplyToID(*req.ReplyToID)
	}
	return c
}

func UpdateSet(u *ent.CommentUpdateOne, req *Req) *ent.CommentUpdateOne {
	if req.Content != nil {
		u.SetContent(*req.Content)
	}
	if req.DeletedAt != nil {
		u.SetDeletedAt(*req.DeletedAt)
	}
	return u
}

type Conds struct {
	ID       *cruder.Cond
	AppID    *cruder.Cond
	UserID   *cruder.Cond
	GoodID   *cruder.Cond
	OrderID  *cruder.Cond
	OrderIDs *cruder.Cond
	GoodIDs  *cruder.Cond
}

func SetQueryConds(q *ent.CommentQuery, conds *Conds) (*ent.CommentQuery, error) {
	q.Where(entcomment.DeletedAt(0))
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
			q.Where(entcomment.ID(id))
		default:
			return nil, fmt.Errorf("invalid comment field")
		}
	}
	if conds.AppID != nil {
		id, ok := conds.AppID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid appid")
		}
		switch conds.AppID.Op {
		case cruder.EQ:
			q.Where(entcomment.AppID(id))
		default:
			return nil, fmt.Errorf("invalid comment field")
		}
	}
	if conds.UserID != nil {
		id, ok := conds.UserID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid userid")
		}
		switch conds.UserID.Op {
		case cruder.EQ:
			q.Where(entcomment.UserID(id))
		default:
			return nil, fmt.Errorf("invalid comment field")
		}
	}
	if conds.GoodID != nil {
		id, ok := conds.GoodID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid goodid")
		}
		switch conds.GoodID.Op {
		case cruder.EQ:
			q.Where(entcomment.GoodID(id))
		default:
			return nil, fmt.Errorf("invalid comment field")
		}
	}
	if conds.OrderID != nil {
		id, ok := conds.OrderID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid orderid")
		}
		switch conds.OrderID.Op {
		case cruder.EQ:
			q.Where(entcomment.OrderID(id))
		default:
			return nil, fmt.Errorf("invalid comment field")
		}
	}
	if conds.OrderIDs != nil {
		ids, ok := conds.OrderIDs.Val.([]uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid userids")
		}
		switch conds.OrderIDs.Op {
		case cruder.IN:
			q.Where(entcomment.UserIDIn(ids...))
		default:
			return nil, fmt.Errorf("invalid comment field")
		}
	}
	if conds.GoodIDs != nil {
		ids, ok := conds.GoodIDs.Val.([]uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid goodids")
		}
		switch conds.GoodIDs.Op {
		case cruder.IN:
			q.Where(entcomment.GoodIDIn(ids...))
		default:
			return nil, fmt.Errorf("invalid comment field")
		}
	}
	return q, nil
}
