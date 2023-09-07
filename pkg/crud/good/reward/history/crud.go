package goodrewardhistory

import (
	"fmt"

	"github.com/NpoolPlatform/good-middleware/pkg/db/ent"
	entgoodrewardhistory "github.com/NpoolPlatform/good-middleware/pkg/db/ent/goodrewardhistory"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Req struct {
	ID         *uuid.UUID
	GoodID     *uuid.UUID
	RewardDate *uint32
	TID        *uuid.UUID
	Amount     *decimal.Decimal
	UnitAmount *decimal.Decimal
}

func CreateSet(c *ent.GoodRewardHistoryCreate, req *Req) *ent.GoodRewardHistoryCreate {
	if req.ID != nil {
		c.SetID(*req.ID)
	}
	if req.GoodID != nil {
		c.SetGoodID(*req.GoodID)
	}
	if req.RewardDate != nil {
		c.SetRewardDate(*req.RewardDate)
	}
	if req.TID != nil {
		c.SetTid(*req.TID)
	}
	if req.Amount != nil {
		c.SetAmount(*req.Amount)
	}
	if req.UnitAmount != nil {
		c.SetUnitAmount(*req.UnitAmount)
	}
	return c
}

func UpdateSet(u *ent.GoodRewardHistoryUpdateOne, req *Req) *ent.GoodRewardHistoryUpdateOne {
	return u
}

type Conds struct {
	ID         *cruder.Cond
	GoodID     *cruder.Cond
	GoodIDs    *cruder.Cond
	RewardDate *cruder.Cond
	StartAt    *cruder.Cond
	EndAt      *cruder.Cond
}

//nolint:gocyclo,funlen
func SetQueryConds(q *ent.GoodRewardHistoryQuery, conds *Conds) (*ent.GoodRewardHistoryQuery, error) {
	q.Where(entgoodrewardhistory.DeletedAt(0))
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
			q.Where(entgoodrewardhistory.ID(id))
		default:
			return nil, fmt.Errorf("invalid goodrewardhistory field")
		}
	}
	if conds.GoodID != nil {
		id, ok := conds.GoodID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid goodid")
		}
		switch conds.GoodID.Op {
		case cruder.EQ:
			q.Where(entgoodrewardhistory.GoodID(id))
		default:
			return nil, fmt.Errorf("invalid goodrewardhistory field")
		}
	}
	if conds.GoodIDs != nil {
		ids, ok := conds.GoodIDs.Val.([]uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid goodids")
		}
		switch conds.GoodIDs.Op {
		case cruder.IN:
			q.Where(entgoodrewardhistory.GoodIDIn(ids...))
		default:
			return nil, fmt.Errorf("invalid goodrewardhistory field")
		}
	}
	if conds.RewardDate != nil {
		date, ok := conds.RewardDate.Val.(uint32)
		if !ok {
			return nil, fmt.Errorf("invalid rewarddate")
		}
		switch conds.RewardDate.Op {
		case cruder.EQ:
			q.Where(entgoodrewardhistory.RewardDate(date))
		case cruder.LTE:
			q.Where(entgoodrewardhistory.RewardDateLTE(date))
		case cruder.GTE:
			q.Where(entgoodrewardhistory.RewardDateGTE(date))
		default:
			return nil, fmt.Errorf("invalid goodrewardhistory field")
		}
	}
	if conds.StartAt != nil {
		date, ok := conds.StartAt.Val.(uint32)
		if !ok {
			return nil, fmt.Errorf("invalid startat")
		}
		q.Where(entgoodrewardhistory.RewardDateGTE(date))
	}
	if conds.EndAt != nil {
		date, ok := conds.EndAt.Val.(uint32)
		if !ok {
			return nil, fmt.Errorf("invalid endat")
		}
		q.Where(entgoodrewardhistory.RewardDateLTE(date))
	}
	return q, nil
}
