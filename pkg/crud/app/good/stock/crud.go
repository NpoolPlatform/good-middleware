package appstock

import (
	"fmt"

	"github.com/NpoolPlatform/good-middleware/pkg/db/ent"
	entappstock "github.com/NpoolPlatform/good-middleware/pkg/db/ent/appstock"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Req struct {
	ID           *uuid.UUID
	AppID        *uuid.UUID
	GoodID       *uuid.UUID
	Total        *decimal.Decimal
	SpotQuantity *decimal.Decimal
	Locked       *decimal.Decimal
	InService    *decimal.Decimal
	WaitStart    *decimal.Decimal
	Sold         *decimal.Decimal
}

//nolint:funlen,gocyclo
func CreateSet(c *ent.AppStockCreate, req *Req) *ent.AppStockCreate {
	if req.ID != nil {
		c.SetID(*req.ID)
	}
	if req.AppID != nil {
		c.SetAppID(*req.AppID)
	}
	if req.GoodID != nil {
		c.SetGoodID(*req.GoodID)
	}
	if req.Total != nil {
		c.SetTotal(*req.Total)
	}
	if req.SpotQuantity != nil {
		c.SetSpotQuantity(*req.SpotQuantity)
	}
	if req.Locked != nil {
		c.SetLocked(*req.Locked)
	}
	if req.InService != nil {
		c.SetInService(*req.InService)
	}
	if req.WaitStart != nil {
		c.SetWaitStart(*req.WaitStart)
	}
	if req.Sold != nil {
		c.SetSold(*req.Sold)
	}
	return c
}

func UpdateSet(u *ent.AppStockUpdateOne, req *Req) *ent.AppStockUpdateOne {
	if req.Total != nil {
		u.SetTotal(*req.Total)
	}
	if req.SpotQuantity != nil {
		u.SetSpotQuantity(*req.SpotQuantity)
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
	return u
}

type Conds struct {
	ID      *cruder.Cond
	AppID   *cruder.Cond
	GoodID  *cruder.Cond
	GoodIDs *cruder.Cond
	AppIDs  *cruder.Cond
}

func SetQueryConds(q *ent.AppStockQuery, conds *Conds) (*ent.AppStockQuery, error) {
	q.Where(entappstock.DeletedAt(0))
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
			q.Where(entappstock.ID(id))
		default:
			return nil, fmt.Errorf("invalid appstock field")
		}
	}
	if conds.AppID != nil {
		id, ok := conds.AppID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid appid")
		}
		switch conds.AppID.Op {
		case cruder.EQ:
			q.Where(entappstock.AppID(id))
		default:
			return nil, fmt.Errorf("invalid appstock field")
		}
	}
	if conds.GoodID != nil {
		id, ok := conds.GoodID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid goodid")
		}
		switch conds.GoodID.Op {
		case cruder.EQ:
			q.Where(entappstock.GoodID(id))
		default:
			return nil, fmt.Errorf("invalid appstock field")
		}
	}
	if conds.GoodIDs != nil {
		ids, ok := conds.GoodIDs.Val.([]uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid goodids")
		}
		switch conds.GoodIDs.Op {
		case cruder.IN:
			q.Where(entappstock.GoodIDIn(ids...))
		default:
			return nil, fmt.Errorf("invalid appstock field")
		}
	}
	if conds.AppIDs != nil {
		ids, ok := conds.AppIDs.Val.([]uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid goodids")
		}
		switch conds.AppIDs.Op {
		case cruder.IN:
			q.Where(entappstock.AppIDIn(ids...))
		default:
			return nil, fmt.Errorf("invalid appstock field")
		}
	}
	return q, nil
}
