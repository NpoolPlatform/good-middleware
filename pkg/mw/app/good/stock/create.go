package appstock

import (
	"context"

	appgoodstockcrud "github.com/NpoolPlatform/good-middleware/pkg/crud/app/good/stock"
	"github.com/NpoolPlatform/good-middleware/pkg/db"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent"
)

type createHandler struct {
	*Handler
}

func (h *createHandler) _createStock(ctx context.Context, cli *ent.Client) error {
	if _, err := appgoodstockcrud.CreateSet(
		cli.AppStock.Create(),
		&h.Req,
	).Save(ctx); err != nil {
		return err
	}
	return nil
}

// Only for test. Stock should always be created with good
func (h *Handler) createStock(ctx context.Context) error {
	handler := &createHandler{
		Handler: h,
	}
	return db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		return handler._createStock(ctx, cli)
	})
}
