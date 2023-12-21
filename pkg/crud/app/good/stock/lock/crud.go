package appstocklock

import (
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent"
	types "github.com/NpoolPlatform/message/npool/basetypes/good/v1"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Req struct {
	EntID             *uuid.UUID
	AppStockID        *uuid.UUID
	Units             *decimal.Decimal
	AppSpotUnits      *decimal.Decimal
	AppStockLockState *types.AppStockLockState
	DeletedAt         *uint32
}

func CreateSet(c *ent.AppStockLockCreate, req *Req) *ent.AppStockLockCreate {
	if req.EntID != nil {
		c.SetEntID(*req.EntID)
	}
	if req.AppStockID != nil {
		c.SetAppStockID(*req.AppStockID)
	}
	if req.Units != nil {
		c.SetUnits(*req.Units)
	}
	if req.AppSpotUnits != nil {
		c.SetAppSpotUnits(*req.AppSpotUnits)
	}
	if req.AppStockLockState != nil {
		c.SetLockState(req.AppStockLockState.String())
	}
	return c
}

func UpdateSet(u *ent.AppStockLockUpdateOne, req *Req) *ent.AppStockLockUpdateOne {
	if req.AppStockLockState != nil {
		u.SetLockState(req.AppStockLockState.String())
	}
	if req.DeletedAt != nil {
		u.SetDeletedAt(*req.DeletedAt)
	}
	return u
}
