package apppledge

import (
	wlog "github.com/NpoolPlatform/go-service-framework/pkg/wlog"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent"
	entapppledge "github.com/NpoolPlatform/good-middleware/pkg/db/ent/apppledge"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	types "github.com/NpoolPlatform/message/npool/basetypes/good/v1"

	"github.com/google/uuid"
)

type Req struct {
	EntID               *uuid.UUID
	AppGoodID           *uuid.UUID
	ServiceStartAt      *uint32
	StartMode           *types.GoodStartMode
	EnableSetCommission *bool
	DeletedAt           *uint32
}

func CreateSet(c *ent.AppPledgeCreate, req *Req) *ent.AppPledgeCreate {
	if req.EntID != nil {
		c.SetEntID(*req.EntID)
	}
	if req.AppGoodID != nil {
		c.SetAppGoodID(*req.AppGoodID)
	}
	if req.ServiceStartAt != nil {
		c.SetServiceStartAt(*req.ServiceStartAt)
	}
	if req.StartMode != nil {
		c.SetStartMode(req.StartMode.String())
	}
	if req.EnableSetCommission != nil {
		c.SetEnableSetCommission(*req.EnableSetCommission)
	}
	return c
}

func UpdateSet(u *ent.AppPledgeUpdateOne, req *Req) *ent.AppPledgeUpdateOne {
	if req.ServiceStartAt != nil {
		u.SetServiceStartAt(*req.ServiceStartAt)
	}
	if req.StartMode != nil {
		u.SetStartMode(req.StartMode.String())
	}
	if req.EnableSetCommission != nil {
		u.SetEnableSetCommission(*req.EnableSetCommission)
	}
	if req.DeletedAt != nil {
		u.SetDeletedAt(*req.DeletedAt)
	}
	return u
}

type Conds struct {
	ID         *cruder.Cond
	IDs        *cruder.Cond
	EntID      *cruder.Cond
	EntIDs     *cruder.Cond
	AppGoodID  *cruder.Cond
	AppGoodIDs *cruder.Cond
}

//nolint:gocyclo,funlen
func SetQueryConds(q *ent.AppPledgeQuery, conds *Conds) (*ent.AppPledgeQuery, error) {
	q.Where(entapppledge.DeletedAt(0))
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
			q.Where(entapppledge.ID(id))
		default:
			return nil, wlog.Errorf("invalid AppPledge field")
		}
	}
	if conds.IDs != nil {
		ids, ok := conds.IDs.Val.([]uint32)
		if !ok {
			return nil, wlog.Errorf("invalid ids")
		}
		switch conds.IDs.Op {
		case cruder.IN:
			q.Where(entapppledge.IDIn(ids...))
		default:
			return nil, wlog.Errorf("invalid AppPledge field")
		}
	}
	if conds.EntID != nil {
		id, ok := conds.EntID.Val.(uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid entid")
		}
		switch conds.EntID.Op {
		case cruder.EQ:
			q.Where(entapppledge.EntID(id))
		default:
			return nil, wlog.Errorf("invalid AppPledge field")
		}
	}
	if conds.EntIDs != nil {
		ids, ok := conds.EntIDs.Val.([]uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid entids")
		}
		switch conds.EntIDs.Op {
		case cruder.IN:
			q.Where(entapppledge.EntIDIn(ids...))
		default:
			return nil, wlog.Errorf("invalid AppPledge field")
		}
	}
	if conds.AppGoodID != nil {
		id, ok := conds.AppGoodID.Val.(uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid appgoodid")
		}
		switch conds.AppGoodID.Op {
		case cruder.EQ:
			q.Where(entapppledge.AppGoodID(id))
		default:
			return nil, wlog.Errorf("invalid AppPledge field")
		}
	}
	if conds.AppGoodIDs != nil {
		ids, ok := conds.AppGoodIDs.Val.([]uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid appgoodids")
		}
		switch conds.AppGoodIDs.Op {
		case cruder.IN:
			q.Where(entapppledge.AppGoodIDIn(ids...))
		default:
			return nil, wlog.Errorf("invalid AppPledge field")
		}
	}
	return q, nil
}
