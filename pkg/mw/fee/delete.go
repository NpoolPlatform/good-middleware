package fee

import (
	"context"
	"time"

	feecrud "github.com/NpoolPlatform/good-middleware/pkg/crud/fee"
	goodbasecrud "github.com/NpoolPlatform/good-middleware/pkg/crud/good/goodbase"
	"github.com/NpoolPlatform/good-middleware/pkg/db"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent"
)

type deleteHandler struct {
	*feeGoodQueryHandler
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

func (h *deleteHandler) deleteFee(ctx context.Context, tx *ent.Tx) error {
	if _, err := feecrud.UpdateSet(
		tx.Fee.UpdateOneID(*h.ID),
		&feecrud.Req{
			DeletedAt: &h.now,
		},
	).Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *Handler) DeleteFee(ctx context.Context) error {
	handler := &deleteHandler{
		feeGoodQueryHandler: &feeGoodQueryHandler{
			Handler: h,
		},
		now: uint32(time.Now().Unix()),
	}

	if err := handler.getFeeGood(ctx); err != nil {
		return err
	}
	if handler.fee == nil {
		return nil
	}

	return db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		if err := handler.deleteGoodBase(_ctx, tx); err != nil {
			return err
		}
		return handler.deleteFee(_ctx, tx)
	})
}
