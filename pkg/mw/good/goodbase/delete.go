package goodbase

import (
	"context"
	"time"

	goodbasecrud "github.com/NpoolPlatform/good-middleware/pkg/crud/good/goodbase"
	"github.com/NpoolPlatform/good-middleware/pkg/db"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent"
)

type deleteHandler struct {
	*Handler
	now uint32
}

func (h *deleteHandler) deleteGoodBase(ctx context.Context, cli *ent.Client) error {
	if _, err := goodbasecrud.UpdateSet(
		cli.GoodBase.UpdateOneID(*h.ID),
		&goodbasecrud.Req{
			DeletedAt: &h.now,
		},
	).Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *Handler) DeleteGoodBase(ctx context.Context) error {
	handler := &deleteHandler{
		Handler: h,
		now:     uint32(time.Now().Unix()),
	}
	info, err := h.GetGoodBase(ctx)
	if err != nil {
		return err
	}
	if info == nil {
		return nil
	}

	h.ID = func() *uint32 { id := info.ID(); return &id }()
	return db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		return handler.deleteGoodBase(_ctx, cli)
	})
}
