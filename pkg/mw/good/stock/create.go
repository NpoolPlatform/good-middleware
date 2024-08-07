package stock

import (
	"context"

	stockcrud "github.com/NpoolPlatform/good-middleware/pkg/crud/good/stock"
	"github.com/NpoolPlatform/good-middleware/pkg/db"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent"
)

type createHandler struct {
	*Handler
}

func (h *createHandler) _createStock(ctx context.Context, cli *ent.Client) error {
	h.Req.SpotQuantity = h.Req.Total
	if _, err := stockcrud.CreateSet(
		cli.Stock.Create(),
		&h.Req,
	).Save(ctx); err != nil {
		return err
	}
	return nil
}

// Only for test. Stock should always be created with good
func (h *Handler) CreateStock(ctx context.Context) error {
	handler := &createHandler{
		Handler: h,
	}
	return db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		return handler._createStock(ctx, cli)
	})
}
