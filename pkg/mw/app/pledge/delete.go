package pledge

import (
	"context"
	"time"

	wlog "github.com/NpoolPlatform/go-service-framework/pkg/wlog"
	appgoodbasecrud "github.com/NpoolPlatform/good-middleware/pkg/crud/app/good/goodbase"
	apppledgecrud "github.com/NpoolPlatform/good-middleware/pkg/crud/app/pledge"
	"github.com/NpoolPlatform/good-middleware/pkg/db"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent"
)

type deleteHandler struct {
	*pledgeAppGoodQueryHandler
	now uint32
}

func (h *deleteHandler) deleteAppGoodBase(ctx context.Context, tx *ent.Tx) error {
	if _, err := appgoodbasecrud.UpdateSet(
		tx.AppGoodBase.UpdateOneID(h._ent.appGoodBase.ID),
		&appgoodbasecrud.Req{
			DeletedAt: &h.now,
		},
	).Save(ctx); err != nil {
		return wlog.WrapError(err)
	}
	return nil
}

func (h *deleteHandler) deletePledge(ctx context.Context, tx *ent.Tx) error {
	if _, err := apppledgecrud.UpdateSet(
		tx.AppPledge.UpdateOneID(h._ent.appPledge.ID),
		&apppledgecrud.Req{
			DeletedAt: &h.now,
		},
	).Save(ctx); err != nil {
		return wlog.WrapError(err)
	}
	return nil
}

func (h *Handler) DeletePledge(ctx context.Context) error {
	handler := &deleteHandler{
		pledgeAppGoodQueryHandler: &pledgeAppGoodQueryHandler{
			Handler: h,
		},
		now: uint32(time.Now().Unix()),
	}

	if err := handler.getAppPledgeAppGood(ctx); err != nil {
		return wlog.WrapError(err)
	}
	if handler._ent.appPledge == nil {
		return nil
	}

	return db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		if err := handler.deleteAppGoodBase(_ctx, tx); err != nil {
			return wlog.WrapError(err)
		}
		return handler.deletePledge(_ctx, tx)
	})
}
