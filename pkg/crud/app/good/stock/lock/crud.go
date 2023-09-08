package appstocklock

import (
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent"

	"github.com/google/uuid"
)

type Req struct {
	ID        *uuid.UUID
	DeletedAt *uint32
}

func CreateSet(c *ent.AppStockLockCreate, req *Req) *ent.AppStockLockCreate {
	if req.ID != nil {
		c.SetID(*req.ID)
	}
	return c
}

func UpdateSet(u *ent.AppStockLockUpdateOne, req *Req) *ent.AppStockLockUpdateOne {
	if req.DeletedAt != nil {
		u.SetDeletedAt(*req.DeletedAt)
	}
	return u
}
