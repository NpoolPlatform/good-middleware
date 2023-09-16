package appstocklock

import (
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Req struct {
	ID           *uuid.UUID
	AppStockID   *uuid.UUID
	Units        *decimal.Decimal
	AppSpotUnits *decimal.Decimal
	DeletedAt    *uint32
}

func CreateSet(c *ent.AppStockLockCreate, req *Req) *ent.AppStockLockCreate {
	if req.ID != nil {
		c.SetID(*req.ID)
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
	return c
}

func UpdateSet(u *ent.AppStockLockUpdateOne, req *Req) *ent.AppStockLockUpdateOne {
	if req.DeletedAt != nil {
		u.SetDeletedAt(*req.DeletedAt)
	}
	return u
}
