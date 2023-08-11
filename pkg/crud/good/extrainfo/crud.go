package extrainfo

import (
	"fmt"

	"github.com/NpoolPlatform/good-middleware/pkg/db/ent"
	entextrainfo "github.com/NpoolPlatform/good-middleware/pkg/db/ent/extrainfo"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Req struct {
	ID        *uuid.UUID
	GoodID    *uuid.UUID
	Posters   []string
	Labels    []string
	VoteCount *uint32
	Rating    *decimal.Decimal
}

func CreateSet(c *ent.ExtraInfoCreate, req *Req) *ent.ExtraInfoCreate {
	if req.ID != nil {
		c.SetID(*req.ID)
	}
	if req.GoodID != nil {
		c.SetGoodID(*req.GoodID)
	}
	if len(req.Posters) > 0 {
		c.SetPosters(req.Posters)
	}
	if len(req.Labels) > 0 {
		c.SetLabels(req.Labels)
	}
	if req.VoteCount != nil {
		c.SetVoteCount(*req.VoteCount)
	}
	if req.Rating != nil {
		c.SetRatingV1(*req.Rating)
	}
	return c
}

func UpdateSet(u *ent.ExtraInfoUpdateOne, req *Req) *ent.ExtraInfoUpdateOne {
	if len(req.Posters) > 0 {
		u.SetPosters(req.Posters)
	}
	if len(req.Labels) > 0 {
		u.SetLabels(req.Labels)
	}
	if req.VoteCount != nil {
		u.SetVoteCount(*req.VoteCount)
	}
	if req.Rating != nil {
		u.SetRatingV1(*req.Rating)
	}
	return u
}

type Conds struct {
	ID     *cruder.Cond
	GoodID *cruder.Cond
}

func SetQueryConds(q *ent.ExtraInfoQuery, conds *Conds) (*ent.ExtraInfoQuery, error) {
	q.Where(entextrainfo.DeletedAt(0))
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
			q.Where(entextrainfo.ID(id))
		default:
			return nil, fmt.Errorf("invalid extrainfo field")
		}
	}
	if conds.GoodID != nil {
		id, ok := conds.GoodID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid goodid")
		}
		switch conds.GoodID.Op {
		case cruder.EQ:
			q.Where(entextrainfo.GoodID(id))
		default:
			return nil, fmt.Errorf("invalid extrainfo field")
		}
	}
	return q, nil
}
