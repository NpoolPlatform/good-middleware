package coin

import (
	"context"
	"fmt"

	goodcoincrud "github.com/NpoolPlatform/good-middleware/pkg/crud/good/coin"
	"github.com/NpoolPlatform/good-middleware/pkg/db"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent"
)

type updateHandler struct {
	*Handler
}

func (h *Handler) updateGoodCoin(ctx context.Context, cli *ent.Client) error {
	if _, err := goodcoincrud.UpdateSet(
		cli.GoodCoin.UpdateOneID(*h.ID),
		&h.Req,
	).Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *Handler) UpdateGoodCoin(ctx context.Context) error {
	handler := &updateHandler{
		Handler: h,
	}

	info, err := h.GetGoodCoin(ctx)
	if err != nil {
		return err
	}
	if info == nil {
		return fmt.Errorf("invalid goodcoin")
	}

	return db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		return handler.updateGoodCoin(_ctx, cli)
	})
}
