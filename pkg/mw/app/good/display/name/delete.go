package displayname

import (
	"context"
	"time"

	wlog "github.com/NpoolPlatform/go-service-framework/pkg/wlog"
	appgooddisplaynamecrud "github.com/NpoolPlatform/good-middleware/pkg/crud/app/good/display/name"
	"github.com/NpoolPlatform/good-middleware/pkg/db"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent"
)

type deleteHandler struct {
	*Handler
	now uint32
}

func (h *deleteHandler) deleteDisplayName(ctx context.Context, cli *ent.Client) error {
	if _, err := appgooddisplaynamecrud.UpdateSet(
		cli.AppGoodDisplayName.UpdateOneID(*h.ID),
		&appgooddisplaynamecrud.Req{
			DeletedAt: &h.now,
		},
	).Save(ctx); err != nil {
		return wlog.WrapError(err)
	}
	return nil
}

func (h *Handler) DeleteDisplayName(ctx context.Context) error {
	info, err := h.GetDisplayName(ctx)
	if err != nil {
		return wlog.WrapError(err)
	}
	if info == nil {
		return nil
	}

	handler := &deleteHandler{
		Handler: h,
		now:     uint32(time.Now().Unix()),
	}
	h.ID = &info.ID

	return db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		return handler.deleteDisplayName(_ctx, cli)
	})
}
