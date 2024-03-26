package mining

import (
	"fmt"

	"github.com/NpoolPlatform/good-middleware/pkg/db/ent"
	entmininggoodstock "github.com/NpoolPlatform/good-middleware/pkg/db/ent/mininggoodstock"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Req struct {
	EntID          *uuid.UUID
	GoodStockID    *uuid.UUID
	MiningPoolID   *uuid.UUID
	PoolGoodUserID *uuid.UUID
	Total          *decimal.Decimal
	SpotQuantity   *decimal.Decimal
	Locked         *decimal.Decimal
	InService      *decimal.Decimal
	WaitStart      *decimal.Decimal
	Sold           *decimal.Decimal
	AppReserved    *decimal.Decimal
	DeletedAt      *uint32
}

func CreateSet(c *ent.MiningGoodStockCreate, req *Req) *ent.MiningGoodStockCreate {
	if req.EntID != nil {
		c.SetEntID(*req.EntID)
	}
	if req.GoodStockID != nil {
		c.SetGoodStockID(*req.GoodStockID)
	}
	if req.MiningPoolID != nil {
		c.SetMiningPoolID(*req.MiningPoolID)
	}
	if req.PoolGoodUserID != nil {
		c.SetPoolGoodUserID(*req.PoolGoodUserID)
	}
	if req.Total != nil {
		c.SetTotal(*req.Total)
	}
	if req.SpotQuantity != nil {
		c.SetSpotQuantity(*req.SpotQuantity)
	}
	c.SetLocked(decimal.NewFromInt(0))
	c.SetInService(decimal.NewFromInt(0))
	c.SetWaitStart(decimal.NewFromInt(0))
	c.SetSold(decimal.NewFromInt(0))
	c.SetAppReserved(decimal.NewFromInt(0))
	return c
}

func UpdateSet(u *ent.MiningGoodStockUpdateOne, req *Req) *ent.MiningGoodStockUpdateOne {
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
	if req.AppReserved != nil {
		u.SetAppReserved(*req.AppReserved)
	}
	if req.DeletedAt != nil {
		u.SetDeletedAt(*req.DeletedAt)
	}
	return u
}

type Conds struct {
	EntID        *cruder.Cond
	GoodStockID  *cruder.Cond
	GoodStockIDs *cruder.Cond
}

func SetQueryConds(q *ent.MiningGoodStockQuery, conds *Conds) (*ent.MiningGoodStockQuery, error) {
	q.Where(entmininggoodstock.DeletedAt(0))
	if conds == nil {
		return q, nil
	}
	if conds.EntID != nil {
		id, ok := conds.EntID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid id")
		}
		switch conds.EntID.Op {
		case cruder.EQ:
			q.Where(entmininggoodstock.EntID(id))
		default:
			return nil, fmt.Errorf("invalid mininggoodstock field")
		}
	}
	if conds.GoodStockID != nil {
		id, ok := conds.GoodStockID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid goodstockid")
		}
		switch conds.GoodStockID.Op {
		case cruder.EQ:
			q.Where(entmininggoodstock.GoodStockID(id))
		default:
			return nil, fmt.Errorf("invalid mininggoodstock field")
		}
	}
	if conds.GoodStockIDs != nil {
		ids, ok := conds.GoodStockIDs.Val.([]uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid goodstockids")
		}
		switch conds.GoodStockIDs.Op {
		case cruder.EQ:
			q.Where(entmininggoodstock.GoodStockIDIn(ids...))
		default:
			return nil, fmt.Errorf("invalid mininggoodstock field")
		}
	}
	return q, nil
}
