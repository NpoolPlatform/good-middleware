package topmostposter

import (
	wlog "github.com/NpoolPlatform/go-service-framework/pkg/wlog"

	"github.com/NpoolPlatform/good-middleware/pkg/db/ent"
	enttopmostposter "github.com/NpoolPlatform/good-middleware/pkg/db/ent/topmostposter"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	"github.com/google/uuid"
)

type Req struct {
	EntID     *uuid.UUID
	TopMostID *uuid.UUID
	Poster    *string
	Index     *uint8
	DeletedAt *uint32
}

func CreateSet(c *ent.TopMostPosterCreate, req *Req) *ent.TopMostPosterCreate {
	if req.EntID != nil {
		c.SetEntID(*req.EntID)
	}
	if req.TopMostID != nil {
		c.SetTopMostID(*req.TopMostID)
	}
	if req.Poster != nil {
		c.SetPoster(*req.Poster)
	}
	if req.Index != nil {
		c.SetIndex(*req.Index)
	}
	return c
}

func UpdateSet(u *ent.TopMostPosterUpdateOne, req *Req) *ent.TopMostPosterUpdateOne {
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
	TopMostID  *cruder.Cond
	TopMostIDs *cruder.Cond
}

//nolint:gocyclo,funlen
func SetQueryConds(q *ent.TopMostPosterQuery, conds *Conds) (*ent.TopMostPosterQuery, error) {
	q.Where(enttopmostposter.DeletedAt(0))
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
			q.Where(enttopmostposter.ID(id))
		default:
			return nil, wlog.Errorf("invalid topmostposter field")
		}
	}
	if conds.IDs != nil {
		ids, ok := conds.IDs.Val.([]uint32)
		if !ok {
			return nil, wlog.Errorf("invalid ids")
		}
		switch conds.IDs.Op {
		case cruder.IN:
			q.Where(enttopmostposter.IDIn(ids...))
		default:
			return nil, wlog.Errorf("invalid topmostposter field")
		}
	}
	if conds.EntID != nil {
		id, ok := conds.EntID.Val.(uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid entid")
		}
		switch conds.EntID.Op {
		case cruder.EQ:
			q.Where(enttopmostposter.EntID(id))
		default:
			return nil, wlog.Errorf("invalid topmostposter field")
		}
	}
	if conds.EntIDs != nil {
		ids, ok := conds.EntIDs.Val.([]uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid entids")
		}
		switch conds.EntIDs.Op {
		case cruder.IN:
			q.Where(enttopmostposter.EntIDIn(ids...))
		default:
			return nil, wlog.Errorf("invalid topmostposter field")
		}
	}
	if conds.TopMostID != nil {
		id, ok := conds.TopMostID.Val.(uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid topmostid")
		}
		switch conds.TopMostID.Op {
		case cruder.EQ:
			q.Where(enttopmostposter.TopMostID(id))
		default:
			return nil, wlog.Errorf("invalid topmostposter field")
		}
	}
	if conds.TopMostIDs != nil {
		ids, ok := conds.TopMostIDs.Val.([]uuid.UUID)
		if !ok {
			return nil, wlog.Errorf("invalid topmostids")
		}
		switch conds.TopMostIDs.Op {
		case cruder.IN:
			q.Where(enttopmostposter.TopMostIDIn(ids...))
		default:
			return nil, wlog.Errorf("invalid topmostposter field")
		}
	}
	return q, nil
}
