package fee

import (
	"context"

	wlog "github.com/NpoolPlatform/go-service-framework/pkg/wlog"
	feecrud "github.com/NpoolPlatform/good-middleware/pkg/crud/fee"
	"github.com/NpoolPlatform/good-middleware/pkg/db"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent"
	goodbase1 "github.com/NpoolPlatform/good-middleware/pkg/mw/good/goodbase"
	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
)

type updateHandler struct {
	*feeGoodQueryHandler
	sqlGoodBase string
}

func (h *updateHandler) constructGoodBaseSQL(ctx context.Context) (err error) {
	handler, _ := goodbase1.NewHandler(
		ctx,
		goodbase1.WithID(&h.goodBase.ID, true),
	)
	handler.Req = *h.GoodBaseReq
	h.sqlGoodBase, err = handler.ConstructUpdateSQL()
	if err != nil {
		if err == cruder.ErrUpdateNothing {
			return nil
		}
		return wlog.WrapError(err)
	}
	return nil
}

func (h *updateHandler) updateGoodBase(ctx context.Context, tx *ent.Tx) error {
	if h.sqlGoodBase == "" {
		return nil
	}
	rc, err := tx.ExecContext(ctx, h.sqlGoodBase)
	if err != nil {
		return wlog.WrapError(err)
	}
	n, err := rc.RowsAffected()
	if err != nil || n != 1 {
		return wlog.Errorf("fail update fee: %v", err)
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
		return wlog.WrapError(err)
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
		return wlog.WrapError(err)
	}

	return db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		if err := handler.updateGoodBase(_ctx, tx); err != nil {
			return wlog.WrapError(err)
		}
		return handler.updateFee(_ctx, tx)
	})
}
