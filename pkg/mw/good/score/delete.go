package score

import (
	"context"
	"time"

	scorecrud "github.com/NpoolPlatform/good-middleware/pkg/crud/good/score"
	"github.com/NpoolPlatform/good-middleware/pkg/db"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/good/score"
)

type deleteHandler struct {
	*Handler
}

func (h *deleteHandler) deleteScore(ctx context.Context, tx *ent.Tx) error {
	now := uint32(time.Now().Unix())
	if _, err := scorecrud.UpdateSet(
		tx.Score.UpdateOneID(*h.ID),
		&scorecrud.Req{
			DeletedAt: &now,
		},
	).Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *Handler) DeleteScore(ctx context.Context) (*npool.Score, error) {
	info, err := h.GetScore(ctx)
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
		if err := handler.deleteScore(ctx, tx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}
