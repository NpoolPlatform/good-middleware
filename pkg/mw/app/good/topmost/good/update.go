package topmostgood

import (
	"context"

	topmostgoodcrud "github.com/NpoolPlatform/good-middleware/pkg/crud/app/good/topmost/good"
	"github.com/NpoolPlatform/good-middleware/pkg/db"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/app/good/topmost/good"
)

type updateHandler struct {
	*Handler
}

func (h *updateHandler) updateTopMostGood(ctx context.Context, tx *ent.Tx) error {
	if _, err := topmostgoodcrud.UpdateSet(
		tx.TopMostGood.UpdateOneID(*h.ID),
		&topmostgoodcrud.Req{
			DisplayIndex: h.DisplayIndex,
			UnitPrice:    h.UnitPrice,
		},
	).Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *Handler) UpdateTopMostGood(ctx context.Context) (*npool.TopMostGood, error) {
	handler := &updateHandler{
		Handler: h,
	}

	err := db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		if err := handler.updateTopMostGood(_ctx, tx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return h.GetTopMostGood(ctx)
}
