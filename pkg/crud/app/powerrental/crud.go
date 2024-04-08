package apppowerrental

import (
	"fmt"

	"github.com/NpoolPlatform/good-middleware/pkg/db/ent"
	entapppowerrental "github.com/NpoolPlatform/good-middleware/pkg/db/ent/apppowerrental"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	types "github.com/NpoolPlatform/message/npool/basetypes/good/v1"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Req struct {
	EntID                        *uuid.UUID
	AppGoodID                    *uuid.UUID
	ServiceStartAt               *uint32
	CancelMode                   *types.CancelMode
	CancelableBeforeStartSeconds *uint32
	EnableSetCommission          *bool
	MinOrderAmount               *decimal.Decimal
	MaxOrderAmount               *decimal.Decimal
	MaxUserAmount                *decimal.Decimal
	MinOrderDuration             *uint32
	MaxOrderDuration             *uint32
	UnitPrice                    *decimal.Decimal
	SaleStartAt                  *uint32
	SaleEndAt                    *uint32
	SaleMode                     *types.GoodSaleMode
	FixedDuration                *bool
	PackageWithRequireds         *bool
	DeletedAt                    *uint32
}

//nolint:gocyclo
func CreateSet(c *ent.AppPowerRentalCreate, req *Req) *ent.AppPowerRentalCreate {
	if req.EntID != nil {
		c.SetEntID(*req.EntID)
	}
	if req.AppGoodID != nil {
		c.SetAppGoodID(*req.AppGoodID)
	}
	if req.ServiceStartAt != nil {
		c.SetServiceStartAt(*req.ServiceStartAt)
	}
	if req.CancelMode != nil {
		c.SetCancelMode(req.CancelMode.String())
	}
	if req.CancelableBeforeStartSeconds != nil {
		c.SetCancelableBeforeStartSeconds(*req.CancelableBeforeStartSeconds)
	}
	if req.EnableSetCommission != nil {
		c.SetEnableSetCommission(*req.EnableSetCommission)
	}
	if req.MinOrderAmount != nil {
		c.SetMinOrderAmount(*req.MinOrderAmount)
	}
	if req.MaxOrderAmount != nil {
		c.SetMaxOrderAmount(*req.MaxOrderAmount)
	}
	if req.MaxUserAmount != nil {
		c.SetMaxUserAmount(*req.MaxUserAmount)
	}
	if req.MinOrderDuration != nil {
		c.SetMinOrderDuration(*req.MinOrderDuration)
	}
	if req.MaxOrderDuration != nil {
		c.SetMaxOrderDuration(*req.MaxOrderDuration)
	}
	if req.UnitPrice != nil {
		c.SetUnitPrice(*req.UnitPrice)
	}
	if req.SaleStartAt != nil {
		c.SetSaleStartAt(*req.SaleStartAt)
	}
	if req.SaleEndAt != nil {
		c.SetSaleEndAt(*req.SaleEndAt)
	}
	if req.SaleMode != nil {
		c.SetSaleMode(req.SaleMode.String())
	}
	if req.FixedDuration != nil {
		c.SetFixedDuration(*req.FixedDuration)
	}
	if req.PackageWithRequireds != nil {
		c.SetPackageWithRequireds(*req.PackageWithRequireds)
	}
	return c
}

//nolint:gocyclo
func UpdateSet(u *ent.AppPowerRentalUpdateOne, req *Req) *ent.AppPowerRentalUpdateOne {
	if req.ServiceStartAt != nil {
		u.SetServiceStartAt(*req.ServiceStartAt)
	}
	if req.CancelMode != nil {
		u.SetCancelMode(req.CancelMode.String())
	}
	if req.CancelableBeforeStartSeconds != nil {
		u.SetCancelableBeforeStartSeconds(*req.CancelableBeforeStartSeconds)
	}
	if req.EnableSetCommission != nil {
		u.SetEnableSetCommission(*req.EnableSetCommission)
	}
	if req.MinOrderAmount != nil {
		u.SetMinOrderAmount(*req.MinOrderAmount)
	}
	if req.MaxOrderAmount != nil {
		u.SetMaxOrderAmount(*req.MaxOrderAmount)
	}
	if req.MaxUserAmount != nil {
		u.SetMaxUserAmount(*req.MaxUserAmount)
	}
	if req.MinOrderDuration != nil {
		u.SetMinOrderDuration(*req.MinOrderDuration)
	}
	if req.MaxOrderDuration != nil {
		u.SetMaxOrderDuration(*req.MaxOrderDuration)
	}
	if req.UnitPrice != nil {
		u.SetUnitPrice(*req.UnitPrice)
	}
	if req.SaleStartAt != nil {
		u.SetSaleStartAt(*req.SaleStartAt)
	}
	if req.SaleEndAt != nil {
		u.SetSaleEndAt(*req.SaleEndAt)
	}
	if req.SaleMode != nil {
		u.SetSaleMode(req.SaleMode.String())
	}
	if req.FixedDuration != nil {
		u.SetFixedDuration(*req.FixedDuration)
	}
	if req.PackageWithRequireds != nil {
		u.SetPackageWithRequireds(*req.PackageWithRequireds)
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
func SetQueryConds(q *ent.AppPowerRentalQuery, conds *Conds) (*ent.AppPowerRentalQuery, error) {
	q.Where(entapppowerrental.DeletedAt(0))
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
			q.Where(entapppowerrental.ID(id))
		default:
			return nil, fmt.Errorf("invalid apppowerrental field")
		}
	}
	if conds.IDs != nil {
		ids, ok := conds.IDs.Val.([]uint32)
		if !ok {
			return nil, fmt.Errorf("invalid ids")
		}
		switch conds.IDs.Op {
		case cruder.IN:
			q.Where(entapppowerrental.IDIn(ids...))
		default:
			return nil, fmt.Errorf("invalid apppowerrental field")
		}
	}
	if conds.EntID != nil {
		id, ok := conds.EntID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid entid")
		}
		switch conds.EntID.Op {
		case cruder.EQ:
			q.Where(entapppowerrental.EntID(id))
		default:
			return nil, fmt.Errorf("invalid apppowerrental field")
		}
	}
	if conds.EntIDs != nil {
		ids, ok := conds.EntIDs.Val.([]uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid entids")
		}
		switch conds.EntIDs.Op {
		case cruder.IN:
			q.Where(entapppowerrental.EntIDIn(ids...))
		default:
			return nil, fmt.Errorf("invalid apppowerrental field")
		}
	}
	if conds.AppGoodID != nil {
		id, ok := conds.AppGoodID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid appgoodid")
		}
		switch conds.AppGoodID.Op {
		case cruder.EQ:
			q.Where(entapppowerrental.AppGoodID(id))
		default:
			return nil, fmt.Errorf("invalid apppowerrental field")
		}
	}
	if conds.AppGoodIDs != nil {
		ids, ok := conds.AppGoodIDs.Val.([]uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid appgoodids")
		}
		switch conds.AppGoodIDs.Op {
		case cruder.IN:
			q.Where(entapppowerrental.AppGoodIDIn(ids...))
		default:
			return nil, fmt.Errorf("invalid apppowerrental field")
		}
	}
	return q, nil
}
