package poster

import (
	"context"

	topmostgoodpostercrud "github.com/NpoolPlatform/good-middleware/pkg/crud/app/good/topmost/good/poster"
	"github.com/NpoolPlatform/good-middleware/pkg/db"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent"
)

type createHandler struct {
	*Handler
}

func (h *createHandler) createPoster(ctx context.Context, cli *ent.Client) error {
	if _, err := topmostgoodpostercrud.CreateSet(
		cli.TopMostGoodPoster.Create(),
		&h.Req,
	).Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *Handler) CreatePoster(ctx context.Context) error {
	handler := &createHandler{
		Handler: h,
	}
	return db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		return handler.createPoster(_ctx, cli)
	})
}
