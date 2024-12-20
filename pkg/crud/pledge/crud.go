package pledge

import (
	wlog "github.com/NpoolPlatform/go-service-framework/pkg/wlog"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent"
	entpledge "github.com/NpoolPlatform/good-middleware/pkg/db/ent/pledge"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	types "github.com/NpoolPlatform/message/npool/basetypes/good/v1"

	"github.com/google/uuid"
)

type Req struct {
	EntID              *uuid.UUID
	GoodID             *uuid.UUID
	ContractCodeURL    *string
	ContractCodeBranch *string
	ContractState      *types.ContractState
	DeletedAt          *uint32
}

func CreateSet(c *ent.PledgeCreate, req *Req) *ent.PledgeCreate {
	if req.EntID != nil {
		c.SetEntID(*req.EntID)
	}
	if req.GoodID != nil {
		c.SetGoodID(*req.GoodID)
	}
	if req.ContractCodeURL != nil {
		c.SetContractCodeURL(*req.ContractCodeURL)
	}
	if req.ContractCodeBranch != nil {
		c.SetContractCodeBranch(*req.ContractCodeBranch)
	}
	if req.ContractState != nil {
		c.SetContractState(req.ContractState.String())
	}
	return c
}

func UpdateSet(u *ent.PledgeUpdateOne, req *Req) *ent.PledgeUpdateOne {
	if req.ContractCodeURL != nil {
		u.SetContractCodeURL(*req.ContractCodeURL)
	}
	if req.ContractCodeBranch != nil {
		u.SetContractCodeBranch(*req.ContractCodeBranch)
	}
	if req.ContractState != nil {
		u.SetContractState(req.ContractState.String())
	}
	if req.DeletedAt != nil {
		u.SetDeletedAt(*req.DeletedAt)
	}
	return u
}

type Conds struct {
	ID             *cruder.Cond
	IDs            *cruder.Cond
	EntID          *cruder.Cond
	EntIDs         *cruder.Cond
	GoodID         *cruder.Cond
	GoodIDs        *cruder.Cond
	ContractState  *cruder.Cond
	ContractStates *cruder.Cond
}

// nolint
func SetQueryConds(q *ent.PledgeQuery, conds *Conds) (*ent.PledgeQuery, error) {
	q.Where(entpledge.DeletedAt(0))
	if conds == nil {
		return q, nil
	}
	if conds.ID != nil {
		id, ok := conds.ID.Val.(uint32)
		if !ok {
			return nil, wlog.Errorf("invalid id")
		}
		switch conds.ID.Op {
		case cruder.EQ:
			q.Where(entpledge.ID(id))
		case cruder.NEQ:
			q.Where(entpledge.IDNEQ(id))
		default:
			return nil, wlog.Errorf("invalid pledge field")
		}
	}
	if conds.IDs != nil {
		ids, ok := conds.IDs.Val.([]uint32)
		if !ok {
			return nil, wlog.Errorf("invalid ids")
		}
		switch conds.IDs.Op {
		case cruder.IN:
			q.Where(entpledge.IDIn(ids...))
		default:
			return nil, wlog.Errorf("invalid pledge field")
		}
	}
	if conds.EntID != nil {
		id, ok := conds.EntID.Val.(uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid entid")
		}
		switch conds.EntID.Op {
		case cruder.EQ:
			q.Where(entpledge.EntID(id))
		case cruder.NEQ:
			q.Where(entpledge.EntIDNEQ(id))
		default:
			return nil, wlog.Errorf("invalid pledge field")
		}
	}
	if conds.EntIDs != nil {
		ids, ok := conds.EntIDs.Val.([]uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid entids")
		}
		switch conds.EntIDs.Op {
		case cruder.IN:
			q.Where(entpledge.EntIDIn(ids...))
		default:
			return nil, wlog.Errorf("invalid pledge field")
		}
	}
	if conds.GoodID != nil {
		id, ok := conds.GoodID.Val.(uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid goodid")
		}
		switch conds.GoodID.Op {
		case cruder.EQ:
			q.Where(entpledge.GoodID(id))
		case cruder.NEQ:
			q.Where(entpledge.GoodIDNEQ(id))
		default:
			return nil, wlog.Errorf("invalid pledge field")
		}
	}
	if conds.GoodIDs != nil {
		ids, ok := conds.GoodIDs.Val.([]uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid goodids")
		}
		switch conds.GoodIDs.Op {
		case cruder.IN:
			q.Where(entpledge.GoodIDIn(ids...))
		default:
			return nil, wlog.Errorf("invalid pledge field")
		}
	}
	if conds.ContractState != nil {
		_mode, ok := conds.ContractState.Val.(types.ContractState)
		if !ok {
			return nil, wlog.Errorf("invalid contractstate")
		}
		switch conds.ContractState.Op {
		case cruder.EQ:
			q.Where(entpledge.ContractState(_mode.String()))
		case cruder.NEQ:
			q.Where(entpledge.ContractStateNEQ(_mode.String()))
		default:
			return nil, wlog.Errorf("invalid pledge field")
		}
	}
	if conds.ContractStates != nil {
		_types, ok := conds.ContractStates.Val.([]types.ContractState)
		if !ok {
			return nil, wlog.Errorf("invalid contractstates")
		}
		es := []string{}
		for _, _type := range _types {
			es = append(es, _type.String())
		}
		switch conds.ContractStates.Op {
		case cruder.IN:
			q.Where(entpledge.ContractStateIn(es...))
		default:
			return nil, wlog.Errorf("invalid pledge field")
		}
	}
	return q, nil
}
