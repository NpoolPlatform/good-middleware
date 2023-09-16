//nolint:dupl
package appstock

import (
	"context"
	"time"

	appstocklockcrud "github.com/NpoolPlatform/good-middleware/pkg/crud/app/good/stock/lock"
	"github.com/NpoolPlatform/good-middleware/pkg/db"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent"
	entappstocklock "github.com/NpoolPlatform/good-middleware/pkg/db/ent/appstocklock"
)

type lockopHandler struct {
	*Handler
	lock *ent.AppStockLock
}

func (h *lockopHandler) getLock(ctx context.Context) error {
	return db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		lock, err := cli.
			AppStockLock.
			Query().
			Where(
				entappstocklock.ID(*h.LockID),
				entappstocklock.DeletedAt(0),
			).
			Only(ctx)
		if err != nil {
			return err
		}
		h.lock = lock
		return nil
	})
}

func (h *lockopHandler) deleteLock(ctx context.Context, tx *ent.Tx) error {
	if _, err := tx.
		AppStockLock.
		UpdateOneID(*h.LockID).
		SetDeletedAt(uint32(time.Now().Unix())).
		Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *lockopHandler) createLock(ctx context.Context, tx *ent.Tx) error {
	if _, err := appstocklockcrud.CreateSet(
		tx.AppStockLock.Create(),
		&appstocklockcrud.Req{
			ID:           h.LockID,
			AppStockID:   h.ID,
			Units:        h.Locked,
			AppSpotUnits: h.AppSpotLocked,
		},
	).Save(ctx); err != nil {
		return err
	}
	return nil
}
