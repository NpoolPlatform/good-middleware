package stock

import (
	"fmt"

	"github.com/NpoolPlatform/good-middleware/pkg/db/ent"
	entstock "github.com/NpoolPlatform/good-middleware/pkg/db/ent/stock"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Req struct {
	ID        *uuid.UUID
	GoodID    *uuid.UUID
	Total     *decimal.Decimal
	Locked    *decimal.Decimal
	InService *decimal.Decimal
	WaitStart *decimal.Decimal
	Sold      *decimal.Decimal
	AppLocked *decimal.Decimal
}

func CreateSet(c *ent.StockCreate, req *Req) *ent.StockCreate {
	if req.ID != nil {
		c.SetID(*req.ID)
	}
	if req.GoodID != nil {
		c.SetGoodID(*req.GoodID)
	}
	if req.Total != nil {
		c.SetTotal(*req.Total)
	}
	c.SetLocked(decimal.NewFromInt(0))
	c.SetInService(decimal.NewFromInt(0))
	c.SetWaitStart(decimal.NewFromInt(0))
	c.SetSold(decimal.NewFromInt(0))
	c.SetAppLocked(decimal.NewFromInt(0))
	return c
}

func UpdateSet(u *ent.StockUpdateOne, req *Req) *ent.StockUpdateOne {
	if req.Total != nil {
		u.SetTotal(*req.Total)
	}
	if req.Locked != nil {
		u.SetLocked(*req.Locked)
	}
	if req.InService != nil {
		u.SetInService(*req.InService)
	}
	if req.WaitStart != nil {
		u.SetWaitStart(*req.WaitStart)
	}
	if req.Sold != nil {
		u.SetSold(*req.Sold)
	}
	if req.AppLocked != nil {
		u.SetAppLocked(*req.AppLocked)
	}
	return u
}

type Conds struct {
	ID     *cruder.Cond
	GoodID *cruder.Cond
}

func SetQueryConds(q *ent.StockQuery, conds *Conds) (*ent.StockQuery, error) {
	q.Where(entstock.DeletedAt(0))
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
			q.Where(entstock.ID(id))
		default:
			return nil, fmt.Errorf("invalid stock field")
		}
	}
	if conds.GoodID != nil {
		id, ok := conds.GoodID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid goodid")
		}
		switch conds.GoodID.Op {
		case cruder.EQ:
			q.Where(entstock.GoodID(id))
		default:
			return nil, fmt.Errorf("invalid stock field")
		}
	}
	return q, nil
}
