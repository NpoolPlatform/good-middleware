package stocklock

import (
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent"
	types "github.com/NpoolPlatform/message/npool/basetypes/good/v1"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Req struct {
	EntID          *uuid.UUID
	StockID        *uuid.UUID
	Units          *decimal.Decimal
	StockLockState *types.StockLockState
	ExLockID       *uuid.UUID
	DeletedAt      *uint32
}

func CreateSet(c *ent.StockLockCreate, req *Req) *ent.StockLockCreate {
	if req.EntID != nil {
		c.SetEntID(*req.EntID)
	}
	if req.StockID != nil {
		c.SetStockID(*req.StockID)
	}
	if req.Units != nil {
		c.SetUnits(*req.Units)
	}
	if req.StockLockState != nil {
		c.SetLockState(req.StockLockState.String())
	}
	if req.ExLockID != nil {
		c.SetExLockID(*req.ExLockID)
	}
	return c
}

func UpdateSet(u *ent.StockLockUpdateOne, req *Req) *ent.StockLockUpdateOne {
	if req.StockLockState != nil {
		u.SetLockState(req.StockLockState.String())
	}
	if req.DeletedAt != nil {
		u.SetDeletedAt(*req.DeletedAt)
	}
	return u
}
