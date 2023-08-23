package topmost

import (
	"context"
	"time"

	topmostcrud "github.com/NpoolPlatform/good-middleware/pkg/crud/app/good/topmost"
	"github.com/NpoolPlatform/good-middleware/pkg/db"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/app/good/topmost"
)

type deleteHandler struct {
	*Handler
}

func (h *deleteHandler) deleteTopMost(ctx context.Context, tx *ent.Tx) error {
	now := uint32(time.Now().Unix())
	if _, err := topmostcrud.UpdateSet(
		tx.TopMost.UpdateOneID(*h.ID),
		&topmostcrud.Req{
			DeletedAt: &now,
		},
	).Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *Handler) DeleteTopMost(ctx context.Context) (*npool.TopMost, error) {
	info, err := h.GetTopMost(ctx)
	if err != nil {
		return nil, err
	}
	if info == nil {
		return nil, nil
	}

	handler := &deleteHandler{
		Handler: h,
	}

	err = db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		if err := handler.deleteTopMost(_ctx, tx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}
