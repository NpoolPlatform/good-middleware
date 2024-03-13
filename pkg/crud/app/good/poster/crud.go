package appgoodposter

import (
	"fmt"

	"github.com/NpoolPlatform/good-middleware/pkg/db/ent"
	entappgoodposter "github.com/NpoolPlatform/good-middleware/pkg/db/ent/appgoodposter"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	"github.com/google/uuid"
)

type Req struct {
	EntID     *uuid.UUID
	AppGoodID *uuid.UUID
	Poster    *string
	Index     *uint8
	DeletedAt *uint32
}

//nolint:gocyclo,funlen
func CreateSet(c *ent.AppGoodPosterCreate, req *Req) *ent.AppGoodPosterCreate {
	if req.EntID != nil {
		c.SetEntID(*req.EntID)
	}
	if req.AppGoodID != nil {
		c.SetAppGoodID(*req.AppGoodID)
	}
	if req.Poster != nil {
		c.SetPoster(*req.Poster)
	}
	if req.Index != nil {
		c.SetIndex(*req.Index)
	}
	return c
}

//nolint:gocyclo,funlen
func UpdateSet(u *ent.AppGoodPosterUpdateOne, req *Req) *ent.AppGoodPosterUpdateOne {
	if req.Poster != nil {
		u.SetPoster(*req.Poster)
	}
	if req.Index != nil {
		u.SetIndex(*req.Index)
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
func SetQueryConds(q *ent.AppGoodPosterQuery, conds *Conds) (*ent.AppGoodPosterQuery, error) {
	q.Where(entappgoodposter.DeletedAt(0))
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
			q.Where(entappgoodposter.ID(id))
		default:
			return nil, fmt.Errorf("invalid appgoodposter field")
		}
	}
	if conds.IDs != nil {
		ids, ok := conds.IDs.Val.([]uint32)
		if !ok {
			return nil, fmt.Errorf("invalid ids")
		}
		switch conds.IDs.Op {
		case cruder.IN:
			q.Where(entappgoodposter.IDIn(ids...))
		default:
			return nil, fmt.Errorf("invalid appgoodposter field")
		}
	}
	if conds.EntID != nil {
		id, ok := conds.EntID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid entid")
		}
		switch conds.EntID.Op {
		case cruder.EQ:
			q.Where(entappgoodposter.EntID(id))
		default:
			return nil, fmt.Errorf("invalid appgoodposter field")
		}
	}
	if conds.EntIDs != nil {
		ids, ok := conds.EntIDs.Val.([]uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid entids")
		}
		switch conds.EntIDs.Op {
		case cruder.IN:
			q.Where(entappgoodposter.EntIDIn(ids...))
		default:
			return nil, fmt.Errorf("invalid appgoodposter field")
		}
	}
	if conds.AppGoodID != nil {
		id, ok := conds.AppGoodID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid appgoodid")
		}
		switch conds.AppGoodID.Op {
		case cruder.EQ:
			q.Where(entappgoodposter.AppGoodID(id))
		default:
			return nil, fmt.Errorf("invalid appgoodposter field")
		}
	}
	if conds.AppGoodIDs != nil {
		ids, ok := conds.AppGoodIDs.Val.([]uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid appgoodids")
		}
		switch conds.AppGoodIDs.Op {
		case cruder.IN:
			q.Where(entappgoodposter.AppGoodIDIn(ids...))
		default:
			return nil, fmt.Errorf("invalid appgoodposter field")
		}
	}
	return q, nil
}
