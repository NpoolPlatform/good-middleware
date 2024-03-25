package appsimulategood

import (
	"context"
	"time"

	appsimulategoodcrud "github.com/NpoolPlatform/good-middleware/pkg/crud/app/powerrental/simulate"
	"github.com/NpoolPlatform/good-middleware/pkg/db"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/app/powerrental/simulate"
)

type deleteHandler struct {
	*Handler
}

func (h *deleteHandler) deleteSimulate(ctx context.Context, tx *ent.Tx) error {
	now := uint32(time.Now().Unix())
	if _, err := appsimulategoodcrud.UpdateSet(
		tx.AppSimulateGood.UpdateOneID(*h.ID),
		&appsimulategoodcrud.Req{
			DeletedAt: &now,
		},
	).Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *Handler) DeleteSimulate(ctx context.Context) (*npool.Simulate, error) {
	info, err := h.GetSimulate(ctx)
	if err != nil {
		return nil, err
	}
	if info == nil {
		return nil, nil
	}

	h.ID = &info.ID

	handler := &deleteHandler{
		Handler: h,
	}

	err = db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		if err := handler.deleteSimulate(_ctx, tx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}
