package comment

import (
	"fmt"

	"github.com/NpoolPlatform/good-middleware/pkg/db/ent"
	entcomment "github.com/NpoolPlatform/good-middleware/pkg/db/ent/comment"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Req struct {
	EntID             *uuid.UUID
	AppID             *uuid.UUID
	UserID            *uuid.UUID
	GoodID            *uuid.UUID
	AppGoodID         *uuid.UUID
	OrderID           *uuid.UUID
	Content           *string
	ReplyToID         *uuid.UUID
	Anonymous         *bool
	PurchasedUser     *bool
	TrialUser         *bool
	OrderFirstComment *bool
	Score             *decimal.Decimal
	DeletedAt         *uint32
}

func CreateSet(c *ent.CommentCreate, req *Req) *ent.CommentCreate {
	if req.EntID != nil {
		c.SetEntID(*req.EntID)
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
	if req.AppGoodID != nil {
		c.SetAppGoodID(*req.AppGoodID)
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
	if req.Anonymous != nil {
		c.SetAnonymous(*req.Anonymous)
	}
	if req.PurchasedUser != nil {
		c.SetPurchasedUser(*req.PurchasedUser)
	}
	if req.TrialUser != nil {
		c.SetTrialUser(*req.TrialUser)
	}
	if req.OrderFirstComment != nil {
		c.SetOrderFirstComment(*req.OrderFirstComment)
	}
	if req.Score != nil {
		c.SetScore(*req.Score)
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
	ID         *cruder.Cond
	EntID      *cruder.Cond
	AppID      *cruder.Cond
	UserID     *cruder.Cond
	GoodID     *cruder.Cond
	AppGoodID  *cruder.Cond
	OrderID    *cruder.Cond
	OrderIDs   *cruder.Cond
	AppGoodIDs *cruder.Cond
}

//nolint:funlen,gocyclo
func SetQueryConds(q *ent.CommentQuery, conds *Conds) (*ent.CommentQuery, error) {
	q.Where(entcomment.DeletedAt(0))
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
			q.Where(entcomment.ID(id))
		default:
			return nil, fmt.Errorf("invalid id field")
		}
	}
	if conds.EntID != nil {
		id, ok := conds.EntID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid id")
		}
		switch conds.EntID.Op {
		case cruder.EQ:
			q.Where(entcomment.EntID(id))
		default:
			return nil, fmt.Errorf("invalid id field")
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
			return nil, fmt.Errorf("invalid appid field")
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
			return nil, fmt.Errorf("invalid goodid field")
		}
	}
	if conds.AppGoodID != nil {
		id, ok := conds.AppGoodID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid appgoodid")
		}
		switch conds.AppGoodID.Op {
		case cruder.EQ:
			q.Where(entcomment.AppGoodID(id))
		default:
			return nil, fmt.Errorf("invalid appgoodid field")
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
			return nil, fmt.Errorf("invalid orderid field")
		}
	}
	if conds.OrderIDs != nil {
		ids, ok := conds.OrderIDs.Val.([]uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid orderids")
		}
		switch conds.OrderIDs.Op {
		case cruder.IN:
			q.Where(entcomment.OrderIDIn(ids...))
		default:
			return nil, fmt.Errorf("invalid orderids field")
		}
	}
	if conds.AppGoodIDs != nil {
		ids, ok := conds.AppGoodIDs.Val.([]uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid appgoodids")
		}
		switch conds.AppGoodIDs.Op {
		case cruder.IN:
			q.Where(entcomment.AppGoodIDIn(ids...))
		default:
			return nil, fmt.Errorf("invalid appgoodids field")
		}
	}
	return q, nil
}
