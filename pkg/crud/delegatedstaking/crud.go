package good

import (
	"fmt"

	"github.com/NpoolPlatform/good-middleware/pkg/db/ent"
	entdelegatedstaking "github.com/NpoolPlatform/good-middleware/pkg/db/ent/delegatedstaking"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	"github.com/google/uuid"
)

type Req struct {
	EntID                    *uuid.UUID
	GoodID                   *uuid.UUID
	NoStakeRedeemDelayHours  *uint32
	MaxRedeemDelayHours      *uint32
	ContractAddress          *string
	NoStakeBenefitDelayHours *uint32
	DeletedAt                *uint32
}

func CreateSet(c *ent.DelegatedStakingCreate, req *Req) *ent.DelegatedStakingCreate {
	if req.EntID != nil {
		c.SetEntID(*req.EntID)
	}
	if req.GoodID != nil {
		c.SetGoodID(*req.GoodID)
	}
	if req.NoStakeRedeemDelayHours != nil {
		c.SetNoStakeRedeemDelayHours(*req.NoStakeRedeemDelayHours)
	}
	if req.MaxRedeemDelayHours != nil {
		c.SetMaxRedeemDelayHours(*req.MaxRedeemDelayHours)
	}
	if req.ContractAddress != nil {
		c.SetContractAddress(*req.ContractAddress)
	}
	if req.NoStakeBenefitDelayHours != nil {
		c.SetNoStakeBenefitDelayHours(*req.NoStakeBenefitDelayHours)
	}
	return c
}

func UpdateSet(u *ent.DelegatedStakingUpdateOne, req *Req) *ent.DelegatedStakingUpdateOne {
	if req.NoStakeRedeemDelayHours != nil {
		u.SetNoStakeRedeemDelayHours(*req.NoStakeRedeemDelayHours)
	}
	if req.MaxRedeemDelayHours != nil {
		u.SetMaxRedeemDelayHours(*req.MaxRedeemDelayHours)
	}
	if req.ContractAddress != nil {
		u.SetContractAddress(*req.ContractAddress)
	}
	if req.NoStakeBenefitDelayHours != nil {
		u.SetNoStakeBenefitDelayHours(*req.NoStakeBenefitDelayHours)
	}
	if req.DeletedAt != nil {
		u.SetDeletedAt(*req.DeletedAt)
	}
	return u
}

type Conds struct {
	ID      *cruder.Cond
	IDs     *cruder.Cond
	EntID   *cruder.Cond
	EntIDs  *cruder.Cond
	GoodID  *cruder.Cond
	GoodIDs *cruder.Cond
}

//nolint
func SetQueryConds(q *ent.DelegatedStakingQuery, conds *Conds) (*ent.DelegatedStakingQuery, error) {
	q.Where(entdelegatedstaking.DeletedAt(0))
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
			q.Where(entdelegatedstaking.ID(id))
		case cruder.NEQ:
			q.Where(entdelegatedstaking.IDNEQ(id))
		default:
			return nil, fmt.Errorf("invalid delegatedstaking field")
		}
	}
	if conds.IDs != nil {
		ids, ok := conds.IDs.Val.([]uint32)
		if !ok {
			return nil, fmt.Errorf("invalid ids")
		}
		switch conds.IDs.Op {
		case cruder.IN:
			q.Where(entdelegatedstaking.IDIn(ids...))
		default:
			return nil, fmt.Errorf("invalid delegatedstaking field")
		}
	}
	if conds.EntID != nil {
		id, ok := conds.EntID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid entid")
		}
		switch conds.EntID.Op {
		case cruder.EQ:
			q.Where(entdelegatedstaking.EntID(id))
		case cruder.NEQ:
			q.Where(entdelegatedstaking.EntIDNEQ(id))
		default:
			return nil, fmt.Errorf("invalid delegatedstaking field")
		}
	}
	if conds.EntIDs != nil {
		ids, ok := conds.EntIDs.Val.([]uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid entids")
		}
		switch conds.EntIDs.Op {
		case cruder.IN:
			q.Where(entdelegatedstaking.EntIDIn(ids...))
		default:
			return nil, fmt.Errorf("invalid delegatedstaking field")
		}
	}
	if conds.GoodID != nil {
		id, ok := conds.GoodID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid goodid")
		}
		switch conds.GoodID.Op {
		case cruder.EQ:
			q.Where(entdelegatedstaking.GoodID(id))
		case cruder.NEQ:
			q.Where(entdelegatedstaking.GoodIDNEQ(id))
		default:
			return nil, fmt.Errorf("invalid delegatedstaking field")
		}
	}
	if conds.GoodIDs != nil {
		ids, ok := conds.GoodIDs.Val.([]uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid goodids")
		}
		switch conds.GoodIDs.Op {
		case cruder.IN:
			q.Where(entdelegatedstaking.GoodIDIn(ids...))
		default:
			return nil, fmt.Errorf("invalid delegatedstaking field")
		}
	}
	return q, nil
}
