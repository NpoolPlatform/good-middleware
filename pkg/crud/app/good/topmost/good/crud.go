package topmostgood

import (
	"fmt"

	"github.com/NpoolPlatform/good-middleware/pkg/db/ent"
	enttopmostgood "github.com/NpoolPlatform/good-middleware/pkg/db/ent/topmostgood"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Req struct {
	ID           *uuid.UUID
	AppID        *uuid.UUID
	GoodID       *uuid.UUID
	AppGoodID    *uuid.UUID
	CoinTypeID   *uuid.UUID
	TopMostID    *uuid.UUID
	DisplayIndex *uint32
	Posters      []string
	Price        *decimal.Decimal
	DeletedAt    *uint32
}

func CreateSet(c *ent.TopMostGoodCreate, req *Req) *ent.TopMostGoodCreate {
	if req.ID != nil {
		c.SetID(*req.ID)
	}
	if req.AppID != nil {
		c.SetAppID(*req.AppID)
	}
	if req.GoodID != nil {
		c.SetGoodID(*req.GoodID)
	}
	if req.AppGoodID != nil {
		c.SetAppGoodID(*req.AppGoodID)
	}
	if req.CoinTypeID != nil {
		c.SetCoinTypeID(*req.CoinTypeID)
	}
	if req.TopMostID != nil {
		c.SetTopMostID(*req.TopMostID)
	}
	if req.DisplayIndex != nil {
		c.SetDisplayIndex(*req.DisplayIndex)
	}
	if len(req.Posters) > 0 {
		c.SetPosters(req.Posters)
	}
	if req.Price != nil {
		c.SetPrice(*req.Price)
	}
	return c
}

func UpdateSet(u *ent.TopMostGoodUpdateOne, req *Req) *ent.TopMostGoodUpdateOne {
	if req.GoodID != nil {
		u.SetGoodID(*req.GoodID)
	}
	if req.AppGoodID != nil {
		u.SetAppGoodID(*req.AppGoodID)
	}
	if req.CoinTypeID != nil {
		u.SetCoinTypeID(*req.CoinTypeID)
	}
	if req.TopMostID != nil {
		u.SetTopMostID(*req.TopMostID)
	}
	if req.DisplayIndex != nil {
		u.SetDisplayIndex(*req.DisplayIndex)
	}
	if req.Price != nil {
		u.SetPrice(*req.Price)
	}
	if len(req.Posters) > 0 {
		u.SetPosters(req.Posters)
	}
	if req.DeletedAt != nil {
		u.SetDeletedAt(*req.DeletedAt)
	}
	return u
}

type Conds struct {
	ID        *cruder.Cond
	AppID     *cruder.Cond
	GoodID    *cruder.Cond
	AppGoodID *cruder.Cond
	TopMostID *cruder.Cond
}

//nolint:gocyclo
func SetQueryConds(q *ent.TopMostGoodQuery, conds *Conds) (*ent.TopMostGoodQuery, error) {
	q.Where(enttopmostgood.DeletedAt(0))
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
			q.Where(enttopmostgood.ID(id))
		default:
			return nil, fmt.Errorf("invalid topmostgood field")
		}
	}
	if conds.AppID != nil {
		id, ok := conds.AppID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid appid")
		}
		switch conds.AppID.Op {
		case cruder.EQ:
			q.Where(enttopmostgood.AppID(id))
		default:
			return nil, fmt.Errorf("invalid topmostgood field")
		}
	}
	if conds.GoodID != nil {
		id, ok := conds.GoodID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid goodid")
		}
		switch conds.GoodID.Op {
		case cruder.EQ:
			q.Where(enttopmostgood.GoodID(id))
		default:
			return nil, fmt.Errorf("invalid topmostgood field")
		}
	}
	if conds.AppGoodID != nil {
		id, ok := conds.AppGoodID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid appgoodid")
		}
		switch conds.AppGoodID.Op {
		case cruder.EQ:
			q.Where(enttopmostgood.AppGoodID(id))
		default:
			return nil, fmt.Errorf("invalid topmostgood field")
		}
	}
	if conds.TopMostID != nil {
		id, ok := conds.TopMostID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid topmostid")
		}
		switch conds.TopMostID.Op {
		case cruder.EQ:
			q.Where(enttopmostgood.TopMostID(id))
		default:
			return nil, fmt.Errorf("invalid topmostgood field")
		}
	}
	return q, nil
}
