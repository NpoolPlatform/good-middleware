package displayname

import (
	"context"

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
	handler := &updateHandler{
		Handler: h,
	}
	return db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		return handler.updateDisplayName(_ctx, cli)
	})
}
