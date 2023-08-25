package topmost

import (
	"fmt"

	"github.com/NpoolPlatform/good-middleware/pkg/db/ent"
	enttopmost "github.com/NpoolPlatform/good-middleware/pkg/db/ent/topmost"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	types "github.com/NpoolPlatform/message/npool/basetypes/good/v1"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Req struct {
	ID                     *uuid.UUID
	AppID                  *uuid.UUID
	TopMostType            *types.GoodTopMostType
	Title                  *string
	Message                *string
	Posters                []string
	StartAt                *uint32
	EndAt                  *uint32
	ThresholdCredits       *decimal.Decimal
	RegisterElapsedSeconds *uint32
	ThresholdPurchases     *uint32
	ThresholdPaymentAmount *decimal.Decimal
	KycMust                *bool
	DeletedAt              *uint32
}

func CreateSet(c *ent.TopMostCreate, req *Req) *ent.TopMostCreate {
	if req.ID != nil {
		c.SetID(*req.ID)
	}
	if req.AppID != nil {
		c.SetAppID(*req.AppID)
	}
	if req.TopMostType != nil {
		c.SetTopMostType(req.TopMostType.String())
	}
	if req.Title != nil {
		c.SetTitle(*req.Title)
	}
	if req.Message != nil {
		c.SetMessage(*req.Message)
	}
	if len(req.Posters) > 0 {
		c.SetPosters(req.Posters)
	}
	if req.StartAt != nil {
		c.SetStartAt(*req.StartAt)
	}
	if req.EndAt != nil {
		c.SetEndAt(*req.EndAt)
	}
	if req.ThresholdCredits != nil {
		c.SetThresholdCredits(*req.ThresholdCredits)
	}
	if req.RegisterElapsedSeconds != nil {
		c.SetRegisterElapsedSeconds(*req.RegisterElapsedSeconds)
	}
	if req.ThresholdPurchases != nil {
		c.SetThresholdPurchases(*req.ThresholdPurchases)
	}
	if req.ThresholdPaymentAmount != nil {
		c.SetThresholdPaymentAmount(*req.ThresholdPaymentAmount)
	}
	if req.KycMust != nil {
		c.SetKycMust(*req.KycMust)
	}
	return c
}

func UpdateSet(u *ent.TopMostUpdateOne, req *Req) *ent.TopMostUpdateOne {
	if req.Title != nil {
		u.SetTitle(*req.Title)
	}
	if req.Message != nil {
		u.SetMessage(*req.Message)
	}
	if len(req.Posters) > 0 {
		u.SetPosters(req.Posters)
	}
	if req.StartAt != nil {
		u.SetStartAt(*req.StartAt)
	}
	if req.EndAt != nil {
		u.SetEndAt(*req.EndAt)
	}
	if req.ThresholdCredits != nil {
		u.SetThresholdCredits(*req.ThresholdCredits)
	}
	if req.RegisterElapsedSeconds != nil {
		u.SetRegisterElapsedSeconds(*req.RegisterElapsedSeconds)
	}
	if req.ThresholdPurchases != nil {
		u.SetThresholdPurchases(*req.ThresholdPurchases)
	}
	if req.ThresholdPaymentAmount != nil {
		u.SetThresholdPaymentAmount(*req.ThresholdPaymentAmount)
	}
	if req.KycMust != nil {
		u.SetKycMust(*req.KycMust)
	}
	if req.DeletedAt != nil {
		u.SetDeletedAt(*req.DeletedAt)
	}
	return u
}

type Conds struct {
	ID          *cruder.Cond
	AppID       *cruder.Cond
	TopMostType *cruder.Cond
	Title       *cruder.Cond
	StartEnd    *cruder.Cond
}

func SetQueryConds(q *ent.TopMostQuery, conds *Conds) (*ent.TopMostQuery, error) {
	q.Where(enttopmost.DeletedAt(0))
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
			q.Where(enttopmost.ID(id))
		default:
			return nil, fmt.Errorf("invalid topmostgood field")
		}
	}
	if conds.AppID != nil {
		id, ok := conds.AppID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid appid")
		}
		switch conds.AppID.Op {
		case cruder.EQ:
			q.Where(enttopmost.AppID(id))
		default:
			return nil, fmt.Errorf("invalid topmostgood field")
		}
	}
	if conds.TopMostType != nil {
		_type, ok := conds.TopMostType.Val.(types.GoodTopMostType)
		if !ok {
			return nil, fmt.Errorf("invalid topmosttype")
		}
		switch conds.TopMostType.Op {
		case cruder.EQ:
			q.Where(enttopmost.TopMostType(_type.String()))
		default:
			return nil, fmt.Errorf("invalid good field")
		}
	}
	if conds.Title != nil {
		title, ok := conds.Title.Val.(string)
		if !ok {
			return nil, fmt.Errorf("invalid title")
		}
		switch conds.Title.Op {
		case cruder.EQ:
			q.Where(enttopmost.Title(title))
		default:
			return nil, fmt.Errorf("invalid good field")
		}
	}
	if conds.StartEnd != nil {
		ats, ok := conds.StartEnd.Val.([]uint32)
		if !ok || len(ats) != 2 { //nolint
			return nil, fmt.Errorf("invalid startend")
		}
		switch conds.StartEnd.Op {
		case cruder.OVERLAP:
			q.Where(
				enttopmost.Or(
					enttopmost.And(
						enttopmost.StartAtLTE(ats[0]),
						enttopmost.EndAtGTE(ats[0]),
					),
					enttopmost.And(
						enttopmost.StartAtLTE(ats[1]),
						enttopmost.EndAtGTE(ats[1]),
					),
					enttopmost.And(
						enttopmost.StartAtGTE(ats[0]),
						enttopmost.EndAtLTE(ats[1]),
					),
				),
			)
		default:
			return nil, fmt.Errorf("invalid good field")
		}
	}
	return q, nil
}
