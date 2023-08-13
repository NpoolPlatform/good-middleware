package appdefaultgood

import (
	"fmt"

	"github.com/NpoolPlatform/good-middleware/pkg/db/ent"
	entappdefaultgood "github.com/NpoolPlatform/good-middleware/pkg/db/ent/appdefaultgood"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	"github.com/google/uuid"
)

type Req struct {
	ID         *uuid.UUID
	AppID      *uuid.UUID
	CoinTypeID *uuid.UUID
	GoodID     *uuid.UUID
}

func CreateSet(c *ent.AppDefaultGoodCreate, req *Req) *ent.AppDefaultGoodCreate {
	if req.ID != nil {
		c.SetID(*req.ID)
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
	return c
}

func UpdateSet(u *ent.AppDefaultGoodUpdateOne, req *Req) *ent.AppDefaultGoodUpdateOne {
	if req.GoodID != nil {
		u.SetGoodID(*req.GoodID)
	}
	return u
}

type Conds struct {
	ID         *cruder.Cond
	AppID      *cruder.Cond
	CoinTypeID *cruder.Cond
}

func SetQueryConds(q *ent.AppDefaultGoodQuery, conds *Conds) (*ent.AppDefaultGoodQuery, error) {
	q.Where(entappdefaultgood.DeletedAt(0))
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
			q.Where(entappdefaultgood.ID(id))
		default:
			return nil, fmt.Errorf("invalid appdefaultgood field")
		}
	}
	if conds.AppID != nil {
		id, ok := conds.AppID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid appid")
		}
		switch conds.AppID.Op {
		case cruder.EQ:
			q.Where(entappdefaultgood.AppID(id))
		default:
			return nil, fmt.Errorf("invalid appdefaultgood field")
		}
	}
	if conds.CoinTypeID != nil {
		id, ok := conds.CoinTypeID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid cointypeid")
		}
		switch conds.CoinTypeID.Op {
		case cruder.EQ:
			q.Where(entappdefaultgood.CoinTypeID(id))
		default:
			return nil, fmt.Errorf("invalid appdefaultgood field")
		}
	}
	return q, nil
}