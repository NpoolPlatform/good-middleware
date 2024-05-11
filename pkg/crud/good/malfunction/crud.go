package goodmalfunction

import (
	wlog "github.com/NpoolPlatform/go-service-framework/pkg/wlog"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent"
	entgoodmalfunction "github.com/NpoolPlatform/good-middleware/pkg/db/ent/goodmalfunction"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	types "github.com/NpoolPlatform/message/npool/basetypes/good/v1"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Req struct {
	EntID  *uuid.UUID
	GoodID *uuid.UUID

	DeletedAt *uint32
}

func CreateSet(c *ent.GoodMalfunctionCreate, req *Req) *ent.GoodMalfunctionCreate {
	if req.EntID != nil {
		c.SetEntID(*req.EntID)
	}
	if req.GoodID != nil {
		c.SetGoodID(*req.GoodID)
	}
	c.SetMalfunctionState(types.BenefitState_BenefitWait.String())
	return c
}

func UpdateSet(u *ent.GoodMalfunctionUpdateOne, req *Req) *ent.GoodMalfunctionUpdateOne {
	if req.MalfunctionState != nil {
		u.SetMalfunctionState(req.MalfunctionState.String())
	}
	if req.LastMalfunctionAt != nil {
		u.SetLastMalfunctionAt(*req.LastMalfunctionAt)
	}
	if req.MalfunctionTID != nil {
		u.SetMalfunctionTid(*req.MalfunctionTID)
	}
	if req.NextMalfunctionStartAmount != nil {
		u.SetNextMalfunctionStartAmount(*req.NextMalfunctionStartAmount)
	}
	if req.LastMalfunctionAmount != nil {
		u.SetLastMalfunctionAmount(*req.LastMalfunctionAmount)
	}
	if req.LastUnitMalfunctionAmount != nil {
		u.SetLastUnitMalfunctionAmount(*req.LastUnitMalfunctionAmount)
	}
	if req.TotalMalfunctionAmount != nil {
		u.SetTotalMalfunctionAmount(*req.TotalMalfunctionAmount)
	}
	if req.DeletedAt != nil {
		u.SetDeletedAt(*req.DeletedAt)
	}
	return u
}

type Conds struct {
	EntID            *cruder.Cond
	GoodID           *cruder.Cond
	MalfunctionState *cruder.Cond
	MalfunctionAt    *cruder.Cond
}

func SetQueryConds(q *ent.GoodMalfunctionQuery, conds *Conds) (*ent.GoodMalfunctionQuery, error) { //nolint:gocyclo
	q.Where(entgoodmalfunction.DeletedAt(0))
	if conds == nil {
		return q, nil
	}
	if conds.EntID != nil {
		id, ok := conds.EntID.Val.(uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid id")
		}
		switch conds.EntID.Op {
		case cruder.EQ:
			q.Where(entgoodmalfunction.EntID(id))
		default:
			return nil, wlog.Errorf("invalid goodmalfunction field")
		}
	}
	if conds.GoodID != nil {
		id, ok := conds.GoodID.Val.(uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid goodid")
		}
		switch conds.GoodID.Op {
		case cruder.EQ:
			q.Where(entgoodmalfunction.GoodID(id))
		default:
			return nil, wlog.Errorf("invalid goodmalfunction field")
		}
	}
	if conds.MalfunctionState != nil {
		state, ok := conds.MalfunctionState.Val.(types.BenefitState)
		if !ok {
			return nil, wlog.Errorf("invalid malfunctionstate")
		}
		switch conds.MalfunctionState.Op {
		case cruder.EQ:
			q.Where(entgoodmalfunction.MalfunctionState(state.String()))
		default:
			return nil, wlog.Errorf("invalid goodmalfunction field")
		}
	}
	if conds.MalfunctionAt != nil {
		at, ok := conds.MalfunctionAt.Val.(uint32)
		if !ok {
			return nil, wlog.Errorf("invalid malfunctionat")
		}
		switch conds.MalfunctionAt.Op {
		case cruder.EQ:
			q.Where(entgoodmalfunction.LastMalfunctionAt(at))
		case cruder.NEQ:
			q.Where(entgoodmalfunction.LastMalfunctionAtNEQ(at))
		case cruder.LT:
			q.Where(entgoodmalfunction.LastMalfunctionAtLT(at))
		case cruder.GT:
			q.Where(entgoodmalfunction.LastMalfunctionAtGT(at))
		default:
			return nil, wlog.Errorf("invalid goodmalfunction field")
		}
	}
	return q, nil
}
