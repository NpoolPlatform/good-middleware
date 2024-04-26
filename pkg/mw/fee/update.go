package fee

import (
	"context"

	feecrud "github.com/NpoolPlatform/good-middleware/pkg/crud/fee"
	goodbasecrud "github.com/NpoolPlatform/good-middleware/pkg/crud/good/goodbase"
	"github.com/NpoolPlatform/good-middleware/pkg/db"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent"
)

type updateHandler struct {
	*feeGoodQueryHandler
}

func (h *updateHandler) updateGoodBase(ctx context.Context, tx *ent.Tx) error {
	if _, err := goodbasecrud.UpdateSet(
		tx.GoodBase.UpdateOneID(h.goodBase.ID),
		&goodbasecrud.Req{
			GoodType: h.GoodBaseReq.GoodType,
			Name:     h.GoodBaseReq.Name,
		},
	).Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *updateHandler) updateFee(ctx context.Context, tx *ent.Tx) error {
	if _, err := feecrud.UpdateSet(
		tx.Fee.UpdateOneID(*h.ID),
		&feecrud.Req{
			SettlementType:      h.SettlementType,
			UnitValue:           h.UnitValue,
			DurationDisplayType: h.DurationDisplayType,
		},
	).Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *Handler) UpdateFee(ctx context.Context) error {
	handler := &updateHandler{
		feeGoodQueryHandler: &feeGoodQueryHandler{
			Handler: h,
		},
	}

	if err := handler.requireFeeGood(ctx); err != nil {
		return err
	}

	return db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		if err := handler.updateGoodBase(_ctx, tx); err != nil {
			return err
		}
		return handler.updateFee(_ctx, tx)
	})
}
