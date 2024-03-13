package appsimulategood

import (
	"fmt"

	"github.com/NpoolPlatform/good-middleware/pkg/db/ent"
	entappsimulategood "github.com/NpoolPlatform/good-middleware/pkg/db/ent/appsimulategood"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	"github.com/shopspring/decimal"

	"github.com/google/uuid"
)

type Req struct {
	ID                 *uint32
	EntID              *uuid.UUID
	AppID              *uuid.UUID
	AppGoodID          *uuid.UUID
	CoinTypeID         *uuid.UUID
	GoodID             *uuid.UUID
	FixedOrderUnits    *decimal.Decimal
	FixedOrderDuration *uint32
	DeletedAt          *uint32
}

func CreateSet(c *ent.AppSimulateGoodCreate, req *Req) *ent.AppSimulateGoodCreate {
	if req.EntID != nil {
		c.SetEntID(*req.EntID)
	}
	if req.AppID != nil {
		c.SetAppID(*req.AppID)
	}
	if req.CoinTypeID != nil {
		c.SetCoinTypeID(*req.CoinTypeID)
	}
	if req.GoodID != nil {
		c.SetGoodID(*req.GoodID)
	}
	if req.AppGoodID != nil {
		c.SetAppGoodID(*req.AppGoodID)
	}
	if req.FixedOrderUnits != nil {
		c.SetFixedOrderUnits(*req.FixedOrderUnits)
	}
	if req.FixedOrderDuration != nil {
		c.SetFixedOrderDuration(*req.FixedOrderDuration)
	}
	return c
}

func UpdateSet(u *ent.AppSimulateGoodUpdateOne, req *Req) *ent.AppSimulateGoodUpdateOne {
	if req.FixedOrderUnits != nil {
		u.SetFixedOrderUnits(*req.FixedOrderUnits)
	}
	if req.FixedOrderDuration != nil {
		u.SetFixedOrderDuration(*req.FixedOrderDuration)
	}
	if req.DeletedAt != nil {
		u.SetDeletedAt(*req.DeletedAt)
	}
	return u
}

type Conds struct {
	ID          *cruder.Cond
	EntID       *cruder.Cond
	AppID       *cruder.Cond
	GoodID      *cruder.Cond
	AppGoodID   *cruder.Cond
	CoinTypeID  *cruder.Cond
	GoodIDs     *cruder.Cond
	CoinTypeIDs *cruder.Cond
}

//nolint:gocyclo,funlen
func SetQueryConds(q *ent.AppSimulateGoodQuery, conds *Conds) (*ent.AppSimulateGoodQuery, error) {
	q.Where(entappsimulategood.DeletedAt(0))
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
			q.Where(entappsimulategood.ID(id))
		default:
			return nil, fmt.Errorf("invalid appsimulategood field")
		}
	}
	if conds.EntID != nil {
		id, ok := conds.EntID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid id")
		}
		switch conds.EntID.Op {
		case cruder.EQ:
			q.Where(entappsimulategood.EntID(id))
		default:
			return nil, fmt.Errorf("invalid appsimulategood field")
		}
	}
	if conds.AppID != nil {
		id, ok := conds.AppID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid appid")
		}
		switch conds.AppID.Op {
		case cruder.EQ:
			q.Where(entappsimulategood.AppID(id))
		default:
			return nil, fmt.Errorf("invalid appsimulategood field")
		}
	}
	if conds.GoodID != nil {
		id, ok := conds.GoodID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid goodid")
		}
		switch conds.GoodID.Op {
		case cruder.EQ:
			q.Where(entappsimulategood.GoodID(id))
		default:
			return nil, fmt.Errorf("invalid appsimulategood field")
		}
	}
	if conds.AppGoodID != nil {
		id, ok := conds.AppGoodID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid appgoodid")
		}
		switch conds.AppGoodID.Op {
		case cruder.EQ:
			q.Where(entappsimulategood.AppGoodID(id))
		default:
			return nil, fmt.Errorf("invalid appsimulategood field")
		}
	}
	if conds.CoinTypeID != nil {
		id, ok := conds.CoinTypeID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid cointypeid")
		}
		switch conds.CoinTypeID.Op {
		case cruder.EQ:
			q.Where(entappsimulategood.CoinTypeID(id))
		default:
			return nil, fmt.Errorf("invalid appsimulategood field")
		}
	}
	if conds.GoodIDs != nil {
		ids, ok := conds.GoodIDs.Val.([]uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid goodids")
		}
		switch conds.GoodIDs.Op {
		case cruder.IN:
			q.Where(entappsimulategood.GoodIDIn(ids...))
		default:
			return nil, fmt.Errorf("invalid appsimulategood field")
		}
	}
	if conds.CoinTypeIDs != nil {
		ids, ok := conds.CoinTypeIDs.Val.([]uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid cointypeids")
		}
		switch conds.CoinTypeIDs.Op {
		case cruder.IN:
			q.Where(entappsimulategood.CoinTypeIDIn(ids...))
		default:
			return nil, fmt.Errorf("invalid appsimulategood field")
		}
	}
	return q, nil
}
