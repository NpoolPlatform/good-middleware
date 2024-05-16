package fee

import (
	"context"

	wlog "github.com/NpoolPlatform/go-service-framework/pkg/wlog"
	appfeecrud "github.com/NpoolPlatform/good-middleware/pkg/crud/app/fee"
	appgoodbasecrud "github.com/NpoolPlatform/good-middleware/pkg/crud/app/good/goodbase"
	"github.com/NpoolPlatform/good-middleware/pkg/db"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent"
)

type createHandler struct {
	*appFeeGoodQueryHandler
}

func (h *createHandler) createAppGoodBase(ctx context.Context, tx *ent.Tx) error {
	if _, err := appgoodbasecrud.CreateSet(
		tx.AppGoodBase.Create(),
		h.AppGoodBaseReq,
	).Save(ctx); err != nil {
		return wlog.WrapError(err)
	}
	return nil
}

func (h *createHandler) createAppFee(ctx context.Context, tx *ent.Tx) error {
	if _, err := appfeecrud.CreateSet(
		tx.AppFee.Create(),
		&h.Req,
	).Save(ctx); err != nil {
		return wlog.WrapError(err)
	}
	return nil
}

func (h *Handler) CreateFee(ctx context.Context) error {
	handler := &createHandler{
		appFeeGoodQueryHandler: &appFeeGoodQueryHandler{
			Handler: h,
		},
	}
	if err := handler.requireFeeGood(ctx); err != nil {
		return wlog.WrapError(err)
	}

	if h.UnitValue == nil {
		h.UnitValue = &handler.fee.UnitValue
	} else if h.UnitValue.LessThan(handler.fee.UnitValue) {
		return wlog.Errorf("invalid unitvalue")
	}

	return db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		if err := handler.createAppGoodBase(_ctx, tx); err != nil {
			return wlog.WrapError(err)
		}
		return handler.createAppFee(_ctx, tx)
	})
}
