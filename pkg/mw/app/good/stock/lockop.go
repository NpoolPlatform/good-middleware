package appstock

import (
	"context"
	"fmt"

	appstocklockcrud "github.com/NpoolPlatform/good-middleware/pkg/crud/app/good/stock/lock"
	"github.com/NpoolPlatform/good-middleware/pkg/db"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent"
	entappstocklock "github.com/NpoolPlatform/good-middleware/pkg/db/ent/appstocklock"
	types "github.com/NpoolPlatform/message/npool/basetypes/good/v1"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type LockStock struct {
	EntID         *uuid.UUID
	GoodID        *uuid.UUID
	AppGoodID     *uuid.UUID
	Locked        *decimal.Decimal
	AppSpotLocked *decimal.Decimal
}

type lockopHandler struct {
	*Handler
	locks []*ent.AppStockLock
	state *types.AppStockLockState
}

func (h *lockopHandler) getLocks(ctx context.Context) error {
	return db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		locks, err := cli.
			AppStockLock.
			Query().
			Where(
				entappstocklock.ExLockID(*h.LockID),
				entappstocklock.DeletedAt(0),
			).
			All(ctx)
		if err != nil {
			return err
		}
		h.locks = locks
		if len(locks) > 0 {
			h.EntID = &locks[0].AppStockID
		}
		return nil
	})
}

func (h *lockopHandler) checkState() error {
	stateMap := map[types.AppStockLockState][]types.AppStockLockState{
		types.AppStockLockState_AppStockLocked: {
			types.AppStockLockState_AppStockWaitStart,
			types.AppStockLockState_AppStockRollback,
			types.AppStockLockState_AppStockCanceled,
		},
		types.AppStockLockState_AppStockWaitStart: {types.AppStockLockState_AppStockInService, types.AppStockLockState_AppStockChargeBack},
		types.AppStockLockState_AppStockInService: {types.AppStockLockState_AppStockExpired, types.AppStockLockState_AppStockChargeBack},
	}
	for _, lock := range h.locks {
		validState := false
		for _, state := range stateMap[types.AppStockLockState(types.AppStockLockState_value[lock.LockState])] {
			if state == *h.state {
				validState = true
				break
			}
		}
		if !validState {
			return fmt.Errorf("invalid state")
		}
	}
	return nil
}

func (h *lockopHandler) updateLocks(ctx context.Context, tx *ent.Tx) error {
	if len(h.locks) == 0 {
		return fmt.Errorf("invalid locks")
	}
	if err := h.checkState(); err != nil {
		return err
	}
	lockIDs := []uint32{}
	for _, lock := range h.locks {
		lockIDs = append(lockIDs, lock.ID)
	}
	stm := tx.
		AppStockLock.
		Update().
		Where(
			entappstocklock.IDIn(lockIDs...),
			entappstocklock.DeletedAt(0),
		).
		SetLockState(h.state.String())
	switch *h.state {
	case types.AppStockLockState_AppStockRollback:
		fallthrough //nolint
	case types.AppStockLockState_AppStockChargeBack:
		stm.SetChargeBackState(h.locks[0].LockState)
	}
	if _, err := stm.Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *lockopHandler) createLocks(ctx context.Context, tx *ent.Tx) error {
	if len(h.Stocks) == 0 {
		if _, err := appstocklockcrud.CreateSet(
			tx.AppStockLock.Create(),
			&appstocklockcrud.Req{
				EntID:        h.LockID,
				AppStockID:   h.EntID,
				Units:        h.Locked,
				AppSpotUnits: h.AppSpotLocked,
				ExLockID:     h.LockID,
			},
		).Save(ctx); err != nil {
			return err
		}
		return nil
	}

	for _, stock := range h.Stocks {
		if _, err := appstocklockcrud.CreateSet(
			tx.AppStockLock.Create(),
			&appstocklockcrud.Req{
				AppStockID:   stock.EntID,
				Units:        stock.Locked,
				AppSpotUnits: stock.AppSpotLocked,
				ExLockID:     h.LockID,
			},
		).Save(ctx); err != nil {
			return err
		}
	}

	return nil
}
