package appstocklock

import (
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Req struct {
	ID        *uuid.UUID
	Units     *decimal.Decimal
	DeletedAt *uint32
}

func CreateSet(c *ent.AppStockLockCreate, req *Req) *ent.AppStockLockCreate {
	if req.ID != nil {
		c.SetID(*req.ID)
	}
	if req.Units != nil {
		c.SetUnits(*req.Units)
	}
	return c
}

func UpdateSet(u *ent.AppStockLockUpdateOne, req *Req) *ent.AppStockLockUpdateOne {
	if req.DeletedAt != nil {
		u.SetDeletedAt(*req.DeletedAt)
	}
	return u
}
