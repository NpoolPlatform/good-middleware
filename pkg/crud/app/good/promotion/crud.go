package promotion

import (
	"fmt"

	"github.com/NpoolPlatform/good-manager/pkg/db/ent"
	entpromotion "github.com/NpoolPlatform/good-manager/pkg/db/ent/promotion"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Req struct {
	ID      *uuid.UUID
	AppID   *uuid.UUID
	GoodID  *uuid.UUID
	Message *string
	StartAt *uint32
	EndAt   *uint32
	Price   *decimal.Decimal
	Posters []string
}

func CreateSet(c *ent.PromotionCreate, req *Req) *ent.PromotionCreate {
	if req.ID != nil {
		c.SetID(*req.ID)
	}
	if req.AppID != nil {
		c.SetAppID(*req.AppID)
	}
	if req.GoodID != nil {
		c.SetGoodID(*req.GoodID)
	}
	if req.Message != nil {
		c.SetMessage(*req.Message)
	}
	if req.StartAt != nil {
		c.SetStartAt(*req.StartAt)
	}
	if req.EndAt != nil {
		c.SetEndAt(*req.EndAt)
	}
	if req.Price != nil {
		c.SetPrice(*req.Price)
	}
	if len(req.Posters) > 0 {
		c.SetPosters(req.Posters)
	}
	return c
}

func UpdateSet(u *ent.PromotionUpdateOne, req *Req) *ent.PromotionUpdateOne {
	if req.Message != nil {
		u.SetMessage(*req.Message)
	}
	if req.StartAt != nil {
		u.SetStartAt(*req.StartAt)
	}
	if req.EndAt != nil {
		u.SetEndAt(*req.EndAt)
	}
	if req.Price != nil {
		u.SetPrice(*req.Price)
	}
	if len(req.Posters) > 0 {
		u.SetPosters(req.Posters)
	}
	return u
}

type Conds struct {
	ID      *cruder.Cond
	AppID   *cruder.Cond
	GoodID  *cruder.Cond
	StartAt *cruder.Cond
	EndAt   *cruder.Cond
}

func SetQueryConds(q *ent.PromotionQuery, conds *Conds) (*ent.PromotionQuery, error) {
	q.Where(entpromotion.DeletedAt(0))
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
			q.Where(entpromotion.ID(id))
		case cruder.NEQ:
			q.Where(entpromotion.IDNEQ(id))
		default:
			return nil, fmt.Errorf("invalid promotion field")
		}
	}
	if conds.AppID != nil {
		id, ok := conds.AppID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid appid")
		}
		switch conds.AppID.Op {
		case cruder.EQ:
			q.Where(entpromotion.AppID(id))
		default:
			return nil, fmt.Errorf("invalid promotion field")
		}
	}
	if conds.GoodID != nil {
		id, ok := conds.GoodID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid goodid")
		}
		switch conds.GoodID.Op {
		case cruder.EQ:
			q.Where(entpromotion.GoodID(id))
		default:
			return nil, fmt.Errorf("invalid promotion field")
		}
	}
	if conds.StartAt != nil {
		at, ok := conds.StartAt.Val.(uint32)
		if !ok {
			return nil, fmt.Errorf("ivnalid startat")
		}
		switch conds.StartAt.Op {
		case cruder.LT:
			q.Where(entpromotion.StartAtLT(at))
		case cruder.LTE:
			q.Where(entpromotion.StartAtLTE(at))
		case cruder.GT:
			q.Where(entpromotion.StartAtGT(at))
		case cruder.GTE:
			q.Where(entpromotion.StartAtGTE(at))
		default:
			return nil, fmt.Errorf("invalid promotion field")
		}
	}
	if conds.EndAt != nil {
		at, ok := conds.EndAt.Val.(uint32)
		if !ok {
			return nil, fmt.Errorf("ivnalid endat")
		}
		switch conds.EndAt.Op {
		case cruder.LT:
			q.Where(entpromotion.EndAtLT(at))
		case cruder.LTE:
			q.Where(entpromotion.EndAtLTE(at))
		case cruder.GT:
			q.Where(entpromotion.EndAtGT(at))
		case cruder.GTE:
			q.Where(entpromotion.EndAtGTE(at))
		default:
			return nil, fmt.Errorf("invalid promotion field")
		}
	}
	return q, nil
}
