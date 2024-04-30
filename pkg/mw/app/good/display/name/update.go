package displayname

import (
	"context"

	wlog "github.com/NpoolPlatform/go-service-framework/pkg/wlog"
	appgooddisplaynamecrud "github.com/NpoolPlatform/good-middleware/pkg/crud/app/good/display/name"
	"github.com/NpoolPlatform/good-middleware/pkg/db"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent"
)

type updateHandler struct {
	*Handler
}

func (h *updateHandler) updateDisplayName(ctx context.Context, cli *ent.Client) error {
	if _, err := appgooddisplaynamecrud.UpdateSet(
		cli.AppGoodDisplayName.UpdateOneID(*h.ID),
		&h.Req,
	).Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *Handler) UpdateDisplayName(ctx context.Context) error {
	info, err := h.GetDisplayName(ctx)
	if err != nil {
		return wlog.WrapError(err)
	}
	if info == nil {
		return wlog.Errorf("invalid displayname")
	}

	h.ID = &info.ID
	handler := &updateHandler{
		Handler: h,
	}
	return db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		return handler.updateDisplayName(_ctx, cli)
	})
}
