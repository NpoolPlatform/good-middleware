package good

import (
	"fmt"

	"github.com/NpoolPlatform/good-middleware/pkg/db/ent"
	entgood "github.com/NpoolPlatform/good-middleware/pkg/db/ent/good"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	types "github.com/NpoolPlatform/message/npool/basetypes/good/v1"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Req struct {
	ID                   *uuid.UUID
	DeviceInfoID         *uuid.UUID
	DurationDays         *int32
	CoinTypeID           *uuid.UUID
	VendorLocationID     *uuid.UUID
	Price                *decimal.Decimal
	BenefitType          *types.BenefitType
	GoodType             *types.GoodType
	Title                *string
	Unit                 *string
	UnitAmount           *int32
	SupportCoinTypeIDs   []uuid.UUID
	DeliveryAt           *uint32
	StartAt              *uint32
	TestOnly             *bool
	BenefitIntervalHours *uint32
	UnitLockDeposit      *decimal.Decimal
}

func CreateSet(c *ent.GoodCreate, req *Req) *ent.GoodCreate {
	if req.ID != nil {
		c.SetID(*req.ID)
	}
	if req.DeviceInfoID != nil {
		c.SetDeviceInfoID(*req.DeviceInfoID)
	}
	if req.DurationDays != nil {
		c.SetDurationDays(*req.DurationDays)
	}
	if req.CoinTypeID != nil {
		c.SetCoinTypeID(*req.CoinTypeID)
	}
	if req.VendorLocationID != nil {
		c.SetVendorLocationID(*req.VendorLocationID)
	}
	if req.Price != nil {
		c.SetPrice(*req.Price)
	}
	if req.BenefitType != nil {
		c.SetBenefitType(req.BenefitType.String())
	}
	if req.GoodType != nil {
		c.SetGoodType(req.GoodType.String())
	}
	if req.Title != nil {
		c.SetTitle(*req.Title)
	}
	if req.Unit != nil {
		c.SetUnit(*req.Unit)
	}
	if req.UnitAmount != nil {
		c.SetUnitAmount(*req.UnitAmount)
	}
	if len(req.SupportCoinTypeIDs) > 0 {
		c.SetSupportCoinTypeIds(req.SupportCoinTypeIDs)
	}
	if req.DeliveryAt != nil {
		c.SetDeliveryAt(*req.DeliveryAt)
	}
	if req.StartAt != nil {
		c.SetStartAt(*req.StartAt)
	}
	if req.TestOnly != nil {
		c.SetTestOnly(*req.TestOnly)
	}
	if req.BenefitIntervalHours != nil {
		c.SetBenefitIntervalHours(*req.BenefitIntervalHours)
	}
	if req.UnitLockDeposit != nil {
		c.SetUnitLockDeposit(*req.UnitLockDeposit)
	}
	return c
}

func UpdateSet(u *ent.GoodUpdateOne, req *Req) *ent.GoodUpdateOne {
	if req.DeviceInfoID != nil {
		u.SetDeviceInfoID(*req.DeviceInfoID)
	}
	if req.DurationDays != nil {
		u.SetDurationDays(*req.DurationDays)
	}
	if req.CoinTypeID != nil {
		u.SetCoinTypeID(*req.CoinTypeID)
	}
	if req.VendorLocationID != nil {
		u.SetVendorLocationID(*req.VendorLocationID)
	}
	if req.Price != nil {
		u.SetPrice(*req.Price)
	}
	if req.Title != nil {
		u.SetTitle(*req.Title)
	}
	if req.Unit != nil {
		u.SetUnit(*req.Unit)
	}
	if req.UnitAmount != nil {
		u.SetUnitAmount(*req.UnitAmount)
	}
	if len(req.SupportCoinTypeIDs) > 0 {
		u.SetSupportCoinTypeIds(req.SupportCoinTypeIDs)
	}
	if req.DeliveryAt != nil {
		u.SetDeliveryAt(*req.DeliveryAt)
	}
	if req.StartAt != nil {
		u.SetStartAt(*req.StartAt)
	}
	if req.TestOnly != nil {
		u.SetTestOnly(*req.TestOnly)
	}
	if req.BenefitIntervalHours != nil {
		u.SetBenefitIntervalHours(*req.BenefitIntervalHours)
	}
	if req.UnitLockDeposit != nil {
		u.SetUnitLockDeposit(*req.UnitLockDeposit)
	}
	return u
}

type Conds struct {
	ID               *cruder.Cond
	DeviceInfoID     *cruder.Cond
	CoinTypeID       *cruder.Cond
	VendorLocationID *cruder.Cond
	BenefitType      *cruder.Cond
	GoodType         *cruder.Cond
	IDs              *cruder.Cond
}

//nolint
func SetQueryConds(q *ent.GoodQuery, conds *Conds) (*ent.GoodQuery, error) {
	q.Where(entgood.DeletedAt(0))
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
			q.Where(entgood.ID(id))
		default:
			return nil, fmt.Errorf("invalid good field")
		}
	}
	if conds.DeviceInfoID != nil {
		id, ok := conds.DeviceInfoID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid deviceinfoid")
		}
		switch conds.DeviceInfoID.Op {
		case cruder.EQ:
			q.Where(entgood.DeviceInfoID(id))
		default:
			return nil, fmt.Errorf("invalid good field")
		}
	}
	if conds.CoinTypeID != nil {
		id, ok := conds.CoinTypeID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid cointypeid")
		}
		switch conds.CoinTypeID.Op {
		case cruder.EQ:
			q.Where(entgood.CoinTypeID(id))
		default:
			return nil, fmt.Errorf("invalid good field")
		}
	}
	if conds.VendorLocationID != nil {
		id, ok := conds.VendorLocationID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid vendorlocationid")
		}
		switch conds.VendorLocationID.Op {
		case cruder.EQ:
			q.Where(entgood.VendorLocationID(id))
		default:
			return nil, fmt.Errorf("invalid good field")
		}
	}
	if conds.BenefitType != nil {
		_type, ok := conds.BenefitType.Val.(types.BenefitType)
		if !ok {
			return nil, fmt.Errorf("invalid benefittype")
		}
		switch conds.BenefitType.Op {
		case cruder.EQ:
			q.Where(entgood.BenefitType(_type.String()))
		default:
			return nil, fmt.Errorf("invalid good field")
		}
	}
	if conds.GoodType != nil {
		_type, ok := conds.GoodType.Val.(types.GoodType)
		if !ok {
			return nil, fmt.Errorf("invalid goodtype")
		}
		switch conds.GoodType.Op {
		case cruder.EQ:
			q.Where(entgood.GoodType(_type.String()))
		default:
			return nil, fmt.Errorf("invalid good field")
		}
	}
	if conds.IDs != nil {
		ids, ok := conds.IDs.Val.([]uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid ids")
		}
		switch conds.IDs.Op {
		case cruder.IN:
			q.Where(entgood.IDIn(ids...))
		default:
			return nil, fmt.Errorf("invalid good field")
		}
	}
	return q, nil
}
