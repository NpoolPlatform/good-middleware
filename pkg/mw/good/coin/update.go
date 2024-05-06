package coin

import (
	"context"

	wlog "github.com/NpoolPlatform/go-service-framework/pkg/wlog"
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
		return wlog.WrapError(err)
	}
	return nil
}

func (h *Handler) UpdateGoodCoin(ctx context.Context) error {
	info, err := h.GetGoodCoin(ctx)
	if err != nil {
		return wlog.WrapError(err)
	}
	if info == nil {
		return wlog.Errorf("invalid goodcoin")
	}

	h.ID = &info.ID
	handler := &updateHandler{
		Handler: h,
	}

	return db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		return handler.updateGoodCoin(_ctx, cli)
	})
}
