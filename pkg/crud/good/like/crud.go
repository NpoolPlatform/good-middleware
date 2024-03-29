package like

import (
	"fmt"

	"github.com/NpoolPlatform/good-middleware/pkg/db/ent"
	entlike "github.com/NpoolPlatform/good-middleware/pkg/db/ent/like"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	"github.com/google/uuid"
)

type Req struct {
	EntID     *uuid.UUID
	AppID     *uuid.UUID
	UserID    *uuid.UUID
	GoodID    *uuid.UUID
	AppGoodID *uuid.UUID
	Like      *bool
	DeletedAt *uint32
}

func CreateSet(c *ent.LikeCreate, req *Req) *ent.LikeCreate {
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
	if req.Like != nil {
		c.SetLike(*req.Like)
	}
	return c
}

func UpdateSet(u *ent.LikeUpdateOne, req *Req) *ent.LikeUpdateOne {
	if req.Like != nil {
		u.SetLike(*req.Like)
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
	AppGoodIDs *cruder.Cond
}

//nolint
func SetQueryConds(q *ent.LikeQuery, conds *Conds) (*ent.LikeQuery, error) {
	q.Where(entlike.DeletedAt(0))
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
			q.Where(entlike.ID(id))
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
			q.Where(entlike.EntID(id))
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
			q.Where(entlike.AppID(id))
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
			q.Where(entlike.UserID(id))
		default:
			return nil, fmt.Errorf("invalid userid field")
		}
	}
	if conds.GoodID != nil {
		id, ok := conds.GoodID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid goodid")
		}
		switch conds.GoodID.Op {
		case cruder.EQ:
			q.Where(entlike.GoodID(id))
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
			q.Where(entlike.AppGoodID(id))
		default:
			return nil, fmt.Errorf("invalid appgoodid field")
		}
	}
	if conds.AppGoodIDs != nil {
		ids, ok := conds.AppGoodIDs.Val.([]uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid appgoodids")
		}
		switch conds.AppGoodIDs.Op {
		case cruder.IN:
			q.Where(entlike.AppGoodIDIn(ids...))
		default:
			return nil, fmt.Errorf("invalid appgoodids field")
		}
	}
	return q, nil
}
