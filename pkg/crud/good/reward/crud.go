package goodreward

import (
	"fmt"

	"github.com/NpoolPlatform/good-middleware/pkg/db/ent"
	entgoodreward "github.com/NpoolPlatform/good-middleware/pkg/db/ent/goodreward"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	types "github.com/NpoolPlatform/message/npool/basetypes/good/v1"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Req struct {
	ID                    *uuid.UUID
	GoodID                *uuid.UUID
	RewardState           *types.BenefitState
	LastRewardAt          *uint32
	RewardTID             *uuid.UUID
	NextRewardStartAmount *decimal.Decimal
	LastRewardAmount      *decimal.Decimal
	DeletedAt             *uint32
}

func CreateSet(c *ent.GoodRewardCreate, req *Req) *ent.GoodRewardCreate {
	if req.ID != nil {
		c.SetID(*req.ID)
	}
	if req.GoodID != nil {
		c.SetGoodID(*req.GoodID)
	}
	c.SetRewardState(types.BenefitState_BenefitWait.String())
	return c
}

func UpdateSet(u *ent.GoodRewardUpdateOne, req *Req) *ent.GoodRewardUpdateOne {
	if req.RewardState != nil {
		u.SetRewardState(req.RewardState.String())
	}
	if req.LastRewardAt != nil {
		u.SetLastRewardAt(*req.LastRewardAt)
	}
	if req.RewardTID != nil {
		u.SetRewardTid(*req.RewardTID)
	}
	if req.NextRewardStartAmount != nil {
		u.SetNextRewardStartAmount(*req.NextRewardStartAmount)
	}
	if req.LastRewardAmount != nil {
		u.SetLastRewardAmount(*req.LastRewardAmount)
	}
	if req.DeletedAt != nil {
		u.SetDeletedAt(*req.DeletedAt)
	}
	return u
}

type Conds struct {
	ID          *cruder.Cond
	GoodID      *cruder.Cond
	RewardState *cruder.Cond
}

func SetQueryConds(q *ent.GoodRewardQuery, conds *Conds) (*ent.GoodRewardQuery, error) {
	q.Where(entgoodreward.DeletedAt(0))
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
			q.Where(entgoodreward.ID(id))
		default:
			return nil, fmt.Errorf("invalid goodreward field")
		}
	}
	if conds.GoodID != nil {
		id, ok := conds.GoodID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid goodid")
		}
		switch conds.GoodID.Op {
		case cruder.EQ:
			q.Where(entgoodreward.GoodID(id))
		default:
			return nil, fmt.Errorf("invalid goodreward field")
		}
	}
	if conds.RewardState != nil {
		state, ok := conds.RewardState.Val.(types.BenefitState)
		if !ok {
			return nil, fmt.Errorf("invalid rewardstate")
		}
		switch conds.RewardState.Op {
		case cruder.EQ:
			q.Where(entgoodreward.RewardState(state.String()))
		default:
			return nil, fmt.Errorf("invalid goodreward field")
		}
	}
	return q, nil
}
