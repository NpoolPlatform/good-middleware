package description

import (
	"context"

	wlog "github.com/NpoolPlatform/go-service-framework/pkg/wlog"
	appgooddescriptioncrud "github.com/NpoolPlatform/good-middleware/pkg/crud/app/good/description"
	"github.com/NpoolPlatform/good-middleware/pkg/db"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent"
)

type updateHandler struct {
	*Handler
}

func (h *updateHandler) updateDescription(ctx context.Context, cli *ent.Client) error {
	if _, err := appgooddescriptioncrud.UpdateSet(
		cli.AppGoodDescription.UpdateOneID(*h.ID),
		&h.Req,
	).Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *Handler) UpdateDescription(ctx context.Context) error {
	info, err := h.GetDescription(ctx)
	if err != nil {
		return wlog.WrapError(err)
	}
	if info == nil {
		return wlog.Errorf("invalid description")
	}

	h.ID = &info.ID
	handler := &updateHandler{
		Handler: h,
	}
	return db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		return handler.updateDescription(_ctx, cli)
	})
}
