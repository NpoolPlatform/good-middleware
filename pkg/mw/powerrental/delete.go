package powerrental

import (
	"context"
	"time"

	goodbasecrud "github.com/NpoolPlatform/good-middleware/pkg/crud/good/goodbase"
	powerrentalcrud "github.com/NpoolPlatform/good-middleware/pkg/crud/powerrental"
	"github.com/NpoolPlatform/good-middleware/pkg/db"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent"
)

type deleteHandler struct {
	*powerRentalGoodQueryHandler
	now uint32
}

func (h *deleteHandler) deleteGoodBase(ctx context.Context, tx *ent.Tx) error {
	if _, err := goodbasecrud.UpdateSet(
		tx.GoodBase.UpdateOneID(h.goodBase.ID),
		&goodbasecrud.Req{
			DeletedAt: &h.now,
		},
	).Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *deleteHandler) deletePowerRental(ctx context.Context, tx *ent.Tx) error {
	if _, err := powerrentalcrud.UpdateSet(
		tx.PowerRental.UpdateOneID(h.powerRental.ID),
		&powerrentalcrud.Req{
			DeletedAt: &h.now,
		},
	).Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *Handler) DeletePowerRental(ctx context.Context) error {
	handler := &deleteHandler{
		powerRentalGoodQueryHandler: &powerRentalGoodQueryHandler{
			Handler: h,
		},
		now: uint32(time.Now().Unix()),
	}

	if err := handler.getPowerRentalGood(ctx); err != nil {
		return err
	}
	if handler.powerRental == nil {
		return nil
	}

	return db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		if err := handler.deleteGoodBase(_ctx, tx); err != nil {
			return err
		}
		return handler.deletePowerRental(_ctx, tx)
	})
}
