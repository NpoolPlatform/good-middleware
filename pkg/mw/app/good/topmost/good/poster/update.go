package poster

import (
	"context"

	topmostgoodpostercrud "github.com/NpoolPlatform/good-middleware/pkg/crud/app/good/topmost/good/poster"
	"github.com/NpoolPlatform/good-middleware/pkg/db"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent"
)

type updateHandler struct {
	*Handler
}

func (h *updateHandler) updatePoster(ctx context.Context, cli *ent.Client) error {
	if _, err := topmostgoodpostercrud.UpdateSet(
		cli.TopMostGoodPoster.UpdateOneID(*h.ID),
		&h.Req,
	).Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *Handler) UpdatePoster(ctx context.Context) error {
	handler := &updateHandler{
		Handler: h,
	}
	return db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		return handler.updatePoster(_ctx, cli)
	})
}
