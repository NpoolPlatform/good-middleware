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
	ID           *uint32
	EntID        *uuid.UUID
	AppID        *uuid.UUID
	GoodID       *uuid.UUID
	AppGoodID    *uuid.UUID
	CoinTypeID   *uuid.UUID
	TopMostID    *uuid.UUID
	DisplayIndex *uint32
	Posters      []string
	UnitPrice    *decimal.Decimal
	PackagePrice *decimal.Decimal
	DeletedAt    *uint32
}

func CreateSet(c *ent.TopMostGoodCreate, req *Req) *ent.TopMostGoodCreate {
	if req.EntID != nil {
		c.SetEntID(*req.EntID)
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
	if req.UnitPrice != nil {
		c.SetUnitPrice(*req.UnitPrice)
	}
	if req.PackagePrice != nil {
		c.SetPackagePrice(*req.PackagePrice)
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
	if req.UnitPrice != nil {
		u.SetUnitPrice(*req.UnitPrice)
	}
	if req.PackagePrice != nil {
		u.SetPackagePrice(*req.PackagePrice)
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
	ID          *cruder.Cond
	EntID       *cruder.Cond
	AppID       *cruder.Cond
	GoodID      *cruder.Cond
	AppGoodID   *cruder.Cond
	TopMostID   *cruder.Cond
	TopMostType *cruder.Cond
	AppGoodIDs  *cruder.Cond
}

//nolint:gocyclo,funlen
func SetQueryConds(q *ent.TopMostGoodQuery, conds *Conds) (*ent.TopMostGoodQuery, error) {
	q.Where(enttopmostgood.DeletedAt(0))
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
			q.Where(enttopmostgood.ID(id))
		default:
			return nil, fmt.Errorf("invalid topmostgood field")
		}
	}
	if conds.EntID != nil {
		id, ok := conds.EntID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid id")
		}
		switch conds.EntID.Op {
		case cruder.EQ:
			q.Where(enttopmostgood.EntID(id))
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
	if conds.AppGoodIDs != nil {
		ids, ok := conds.AppGoodIDs.Val.([]uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid appgoodids")
		}
		switch conds.AppGoodIDs.Op {
		case cruder.IN:
			q.Where(enttopmostgood.AppGoodIDIn(ids...))
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
