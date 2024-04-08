package good

import (
	"fmt"

	"github.com/NpoolPlatform/good-middleware/pkg/db/ent"
	entgoodcoin "github.com/NpoolPlatform/good-middleware/pkg/db/ent/goodcoin"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	"github.com/google/uuid"
)

type Req struct {
	EntID      *uuid.UUID
	GoodID     *uuid.UUID
	CoinTypeID *uuid.UUID
	Main       *bool
	DeletedAt  *uint32
}

func CreateSet(c *ent.GoodCoinCreate, req *Req) *ent.GoodCoinCreate {
	if req.EntID != nil {
		c.SetEntID(*req.EntID)
	}
	if req.GoodID != nil {
		c.SetGoodID(*req.GoodID)
	}
	if req.CoinTypeID != nil {
		c.SetCoinTypeID(*req.CoinTypeID)
	}
	if req.Main != nil {
		c.SetMain(*req.Main)
	}
	return c
}

func UpdateSet(u *ent.GoodCoinUpdateOne, req *Req) *ent.GoodCoinUpdateOne {
	if req.Main != nil {
		u.SetMain(*req.Main)
	}
	if req.DeletedAt != nil {
		u.SetDeletedAt(*req.DeletedAt)
	}
	return u
}

type Conds struct {
	ID          *cruder.Cond
	IDs         *cruder.Cond
	EntID       *cruder.Cond
	EntIDs      *cruder.Cond
	GoodID      *cruder.Cond
	GoodIDs     *cruder.Cond
	CoinTypeID  *cruder.Cond
	CoinTypeIDs *cruder.Cond
	Main        *cruder.Cond
}

//nolint
func SetQueryConds(q *ent.GoodCoinQuery, conds *Conds) (*ent.GoodCoinQuery, error) {
	q.Where(entgoodcoin.DeletedAt(0))
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
			q.Where(entgoodcoin.ID(id))
		case cruder.NEQ:
			q.Where(entgoodcoin.IDNEQ(id))
		default:
			return nil, fmt.Errorf("invalid goodcoin field")
		}
	}
	if conds.IDs != nil {
		ids, ok := conds.IDs.Val.([]uint32)
		if !ok {
			return nil, fmt.Errorf("invalid ids")
		}
		switch conds.IDs.Op {
		case cruder.IN:
			q.Where(entgoodcoin.IDIn(ids...))
		default:
			return nil, fmt.Errorf("invalid goodcoin field")
		}
	}
	if conds.EntID != nil {
		id, ok := conds.EntID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid entid")
		}
		switch conds.EntID.Op {
		case cruder.EQ:
			q.Where(entgoodcoin.EntID(id))
		case cruder.NEQ:
			q.Where(entgoodcoin.EntIDNEQ(id))
		default:
			return nil, fmt.Errorf("invalid goodcoin field")
		}
	}
	if conds.EntIDs != nil {
		ids, ok := conds.EntIDs.Val.([]uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid entids")
		}
		switch conds.EntIDs.Op {
		case cruder.IN:
			q.Where(entgoodcoin.EntIDIn(ids...))
		default:
			return nil, fmt.Errorf("invalid goodcoin field")
		}
	}
	if conds.GoodID != nil {
		id, ok := conds.GoodID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid goodid")
		}
		switch conds.GoodID.Op {
		case cruder.EQ:
			q.Where(entgoodcoin.GoodID(id))
		case cruder.NEQ:
			q.Where(entgoodcoin.GoodIDNEQ(id))
		default:
			return nil, fmt.Errorf("invalid goodcoin field")
		}
	}
	if conds.GoodIDs != nil {
		ids, ok := conds.GoodIDs.Val.([]uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid goodids")
		}
		switch conds.GoodIDs.Op {
		case cruder.IN:
			q.Where(entgoodcoin.GoodIDIn(ids...))
		default:
			return nil, fmt.Errorf("invalid goodcoin field")
		}
	}
	if conds.CoinTypeID != nil {
		id, ok := conds.CoinTypeID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid cointypeid")
		}
		switch conds.CoinTypeID.Op {
		case cruder.EQ:
			q.Where(entgoodcoin.CoinTypeID(id))
		case cruder.NEQ:
			q.Where(entgoodcoin.CoinTypeIDNEQ(id))
		default:
			return nil, fmt.Errorf("invalid goodcoin field")
		}
	}
	if conds.CoinTypeIDs != nil {
		ids, ok := conds.CoinTypeIDs.Val.([]uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid cointypeids")
		}
		switch conds.CoinTypeIDs.Op {
		case cruder.IN:
			q.Where(entgoodcoin.CoinTypeIDIn(ids...))
		default:
			return nil, fmt.Errorf("invalid goodcoin field")
		}
	}
	if conds.Main != nil {
		b, ok := conds.Main.Val.(bool)
		if !ok {
			return nil, fmt.Errorf("invalid main")
		}
		switch conds.Main.Op {
		case cruder.EQ:
			q.Where(entgoodcoin.Main(b))
		case cruder.NEQ:
			q.Where(entgoodcoin.MainNEQ(b))
		}
	}
	return q, nil
}
