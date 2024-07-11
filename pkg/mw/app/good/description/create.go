package description

import (
	"context"

	appgooddescriptioncrud "github.com/NpoolPlatform/good-middleware/pkg/crud/app/good/description"
	"github.com/NpoolPlatform/good-middleware/pkg/db"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent"
)

type createHandler struct {
	*Handler
}

func (h *createHandler) createDescription(ctx context.Context, cli *ent.Client) error {
	if _, err := appgooddescriptioncrud.CreateSet(
		cli.AppGoodDescription.Create(),
		&h.Req,
	).Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *Handler) CreateDescription(ctx context.Context) error {
	handler := &createHandler{
		Handler: h,
	}
	return db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		return handler.createDescription(_ctx, cli)
	})
}
