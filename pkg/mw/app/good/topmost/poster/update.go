package poster

import (
	"context"

	wlog "github.com/NpoolPlatform/go-service-framework/pkg/wlog"
	topmostpostercrud "github.com/NpoolPlatform/good-middleware/pkg/crud/app/good/topmost/poster"
	"github.com/NpoolPlatform/good-middleware/pkg/db"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent"
)

type updateHandler struct {
	*Handler
}

func (h *updateHandler) updatePoster(ctx context.Context, cli *ent.Client) error {
	if _, err := topmostpostercrud.UpdateSet(
		cli.TopMostPoster.UpdateOneID(*h.ID),
		&h.Req,
	).Save(ctx); err != nil {
		return wlog.WrapError(err)
	}
	return nil
}

func (h *Handler) UpdatePoster(ctx context.Context) error {
	info, err := h.GetPoster(ctx)
	if err != nil {
		return wlog.WrapError(err)
	}
	if info == nil {
		return wlog.Errorf("invalid poster")
	}

	h.ID = &info.ID
	handler := &updateHandler{
		Handler: h,
	}
	return db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		return handler.updatePoster(_ctx, cli)
	})
}
