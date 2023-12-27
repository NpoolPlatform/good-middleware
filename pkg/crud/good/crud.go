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
	ID                    *uint32
	EntID                 *uuid.UUID
	DeviceInfoID          *uuid.UUID
	DurationDays          *int32
	CoinTypeID            *uuid.UUID
	VendorLocationID      *uuid.UUID
	UnitPrice             *decimal.Decimal
	BenefitType           *types.BenefitType
	StartMode             *types.GoodStartMode
	GoodType              *types.GoodType
	Title                 *string
	QuantityUnit          *string
	QuantityUnitAmount    *decimal.Decimal
	DeliveryAt            *uint32
	StartAt               *uint32
	TestOnly              *bool
	BenefitIntervalHours  *uint32
	UnitLockDeposit       *decimal.Decimal
	UnitType              *types.GoodUnitType
	QuantityCalculateType *types.GoodUnitCalculateType
	DurationType          *types.GoodDurationType
	DurationCalculateType *types.GoodUnitCalculateType
	SettlementType        *types.GoodSettlementType
	DeletedAt             *uint32
}

//nolint:gocyclo
func CreateSet(c *ent.GoodCreate, req *Req) *ent.GoodCreate {
	if req.EntID != nil {
		c.SetEntID(*req.EntID)
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
	if req.UnitPrice != nil {
		c.SetUnitPrice(*req.UnitPrice)
	}
	if req.BenefitType != nil {
		c.SetBenefitType(req.BenefitType.String())
	}
	if req.StartMode != nil {
		c.SetStartMode(req.StartMode.String())
	}
	if req.GoodType != nil {
		c.SetGoodType(req.GoodType.String())
	}
	if req.Title != nil {
		c.SetTitle(*req.Title)
	}
	if req.QuantityUnit != nil {
		c.SetQuantityUnit(*req.QuantityUnit)
	}
	if req.QuantityUnitAmount != nil {
		c.SetQuantityUnitAmount(*req.QuantityUnitAmount)
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
	if req.UnitType != nil {
		c.SetUnitType(req.UnitType.String())
	}
	if req.QuantityCalculateType != nil {
		c.SetQuantityCalculateType(req.QuantityCalculateType.String())
	}
	if req.DurationType != nil {
		c.SetDurationType(req.DurationType.String())
	}
	if req.SettlementType != nil {
		c.SetSettlementType(req.SettlementType.String())
	}
	if req.DurationCalculateType != nil {
		c.SetDurationCalculateType(req.DurationCalculateType.String())
	}
	return c
}

//nolint:gocyclo
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
	if req.UnitPrice != nil {
		u.SetUnitPrice(*req.UnitPrice)
	}
	if req.BenefitType != nil {
		u.SetBenefitType(req.BenefitType.String())
	}
	if req.StartMode != nil {
		u.SetStartMode(req.StartMode.String())
	}
	if req.Title != nil {
		u.SetTitle(*req.Title)
	}
	if req.QuantityUnit != nil {
		u.SetQuantityUnit(*req.QuantityUnit)
	}
	if req.QuantityUnitAmount != nil {
		u.SetQuantityUnitAmount(*req.QuantityUnitAmount)
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
	if req.UnitType != nil {
		u.SetUnitType(req.UnitType.String())
	}
	if req.QuantityCalculateType != nil {
		u.SetQuantityCalculateType(req.QuantityCalculateType.String())
	}
	if req.DurationType != nil {
		u.SetDurationType(req.DurationType.String())
	}
	if req.SettlementType != nil {
		u.SetSettlementType(req.SettlementType.String())
	}
	if req.DurationCalculateType != nil {
		u.SetDurationCalculateType(req.DurationCalculateType.String())
	}
	if req.DeletedAt != nil {
		u.SetDeletedAt(*req.DeletedAt)
	}
	return u
}

type Conds struct {
	ID               *cruder.Cond
	EntID            *cruder.Cond
	DeviceInfoID     *cruder.Cond
	CoinTypeID       *cruder.Cond
	VendorLocationID *cruder.Cond
	BenefitType      *cruder.Cond
	GoodType         *cruder.Cond
	EntIDs           *cruder.Cond
	IDs              *cruder.Cond
	Title            *cruder.Cond
	RewardState      *cruder.Cond
	RewardAt         *cruder.Cond
}

//nolint
func SetQueryConds(q *ent.GoodQuery, conds *Conds) (*ent.GoodQuery, error) {
	q.Where(entgood.DeletedAt(0))
	if conds == nil {
		return q, nil
	}
	if conds.EntID != nil {
		id, ok := conds.EntID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid id")
		}
		switch conds.EntID.Op {
		case cruder.EQ:
			q.Where(entgood.EntID(id))
		case cruder.NEQ:
			q.Where(entgood.EntIDNEQ(id))
		default:
			return nil, fmt.Errorf("invalid good field")
		}
	}
	if conds.ID != nil {
		id, ok := conds.ID.Val.(uint32)
		if !ok {
			return nil, fmt.Errorf("invalid id")
		}
		switch conds.ID.Op {
		case cruder.EQ:
			q.Where(entgood.ID(id))
		case cruder.NEQ:
			q.Where(entgood.IDNEQ(id))
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
		case cruder.NEQ:
			q.Where(entgood.BenefitTypeNEQ(_type.String()))
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
	if conds.EntIDs != nil {
		ids, ok := conds.EntIDs.Val.([]uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid ids")
		}
		switch conds.EntIDs.Op {
		case cruder.IN:
			q.Where(entgood.EntIDIn(ids...))
		default:
			return nil, fmt.Errorf("invalid good field")
		}
	}
	if conds.IDs != nil {
		ids, ok := conds.IDs.Val.([]uint32)
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
	if conds.Title != nil {
		title, ok := conds.Title.Val.(string)
		if !ok {
			return nil, fmt.Errorf("invalid title")
		}
		switch conds.Title.Op {
		case cruder.EQ:
			q.Where(entgood.Title(title))
		default:
			return nil, fmt.Errorf("invalid good field")
		}
	}
	return q, nil
}
