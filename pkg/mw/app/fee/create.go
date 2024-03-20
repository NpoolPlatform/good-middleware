package fee

import (
	"context"

	appfeecrud "github.com/NpoolPlatform/good-middleware/pkg/crud/app/fee"
	appgoodbasecrud "github.com/NpoolPlatform/good-middleware/pkg/crud/app/good/goodbase"
	"github.com/NpoolPlatform/good-middleware/pkg/db"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent"
)

type createHandler struct {
	*Handler
}

func (h *createHandler) createAppGoodBase(ctx context.Context, tx *ent.Tx) error {
	if _, err := appgoodbasecrud.CreateSet(
		tx.AppGoodBase.Create(),
		h.AppGoodBaseReq,
	).Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *createHandler) createAppFee(ctx context.Context, tx *ent.Tx) error {
	if _, err := appfeecrud.CreateSet(
		tx.AppFee.Create(),
		&h.Req,
	).Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *Handler) CreateFee(ctx context.Context) error {
	handler := &createHandler{
		Handler: h,
	}
	return db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		if err := handler.createAppGoodBase(_ctx, tx); err != nil {
			return err
		}
		return handler.createAppFee(_ctx, tx)
	})
}
