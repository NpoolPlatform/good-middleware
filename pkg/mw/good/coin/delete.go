package coin

import (
	"context"
	"time"

	goodcoincrud "github.com/NpoolPlatform/good-middleware/pkg/crud/good/coin"
	"github.com/NpoolPlatform/good-middleware/pkg/db"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent"
)

type deleteHandler struct {
	*Handler
	now uint32
}

func (h *deleteHandler) deleteGoodCoin(ctx context.Context, cli *ent.Client) error {
	if _, err := goodcoincrud.UpdateSet(
		cli.GoodCoin.UpdateOneID(*h.ID),
		&goodcoincrud.Req{
			DeletedAt: &h.now,
		},
	).Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *Handler) DeleteGoodCoin(ctx context.Context) error {
	handler := &deleteHandler{
		Handler: h,
		now:     uint32(time.Now().Unix()),
	}
	info, err := h.GetGoodCoin(ctx)
	if err != nil {
		return err
	}
	if info == nil {
		return nil
	}
	return db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		return handler.deleteGoodCoin(_ctx, cli)
	})
}
