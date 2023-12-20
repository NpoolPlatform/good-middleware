package appstock

import (
	"context"
	"fmt"

	appstocklockcrud "github.com/NpoolPlatform/good-middleware/pkg/crud/app/good/stock/lock"
	"github.com/NpoolPlatform/good-middleware/pkg/db"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent"
	entappstocklock "github.com/NpoolPlatform/good-middleware/pkg/db/ent/appstocklock"
	types "github.com/NpoolPlatform/message/npool/basetypes/good/v1"
)

type lockopHandler struct {
	*Handler
	lock  *ent.AppStockLock
	state *types.AppStockLockState
}

func (h *lockopHandler) getLock(ctx context.Context) error {
	return db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		lock, err := cli.
			AppStockLock.
			Query().
			Where(
				entappstocklock.EntID(*h.LockID),
				entappstocklock.DeletedAt(0),
			).
			Only(ctx)
		if err != nil {
			return err
		}
		h.lock = lock
		h.EntID = &lock.AppStockID
		h.ID = &lock.ID
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
	for _, state := range stateMap[types.AppStockLockState(types.AppStockLockState_value[h.lock.LockState])] {
		if state == *h.state {
			return nil
		}
	}
	return fmt.Errorf("invalid state")
}

func (h *lockopHandler) updateLock(ctx context.Context, tx *ent.Tx) error {
	if err := h.checkState(); err != nil {
		return err
	}
	stm := tx.
		AppStockLock.
		UpdateOneID(h.lock.ID).
		SetLockState(h.state.String())
	switch *h.state {
	case types.AppStockLockState_AppStockRollback:
		fallthrough //nolint
	case types.AppStockLockState_AppStockChargeBack:
		stm.SetChargeBackState(h.lock.LockState)
	}
	if _, err := stm.Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *lockopHandler) createLock(ctx context.Context, tx *ent.Tx) error {
	if _, err := appstocklockcrud.CreateSet(
		tx.AppStockLock.Create(),
		&appstocklockcrud.Req{
			EntID:        h.LockID,
			AppStockID:   h.EntID,
			Units:        h.Locked,
			AppSpotUnits: h.AppSpotLocked,
		},
	).Save(ctx); err != nil {
		return err
	}
	return nil
}
