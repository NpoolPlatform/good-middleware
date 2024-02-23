package appsimulategood

import (
	"context"

	appsimulategoodcrud "github.com/NpoolPlatform/good-middleware/pkg/crud/app/good/simulate"
	"github.com/NpoolPlatform/good-middleware/pkg/db"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/app/good/simulate"
)

type updateHandler struct {
	*Handler
}

func (h *updateHandler) updateSimulate(ctx context.Context, tx *ent.Tx) error {
	if _, err := appsimulategoodcrud.UpdateSet(
		tx.AppSimulateGood.UpdateOneID(*h.ID),
		&appsimulategoodcrud.Req{
			FixedOrderUnits:    h.FixedOrderUnits,
			FixedOrderDuration: h.FixedOrderDuration,
		},
	).Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *Handler) UpdateSimulate(ctx context.Context) (*npool.Simulate, error) {
	info, err := h.GetSimulate(ctx)
	if err != nil {
		return nil, err
	}
	if h.FixedOrderUnits == nil && h.FixedOrderDuration == nil {
		return info, nil
	}

	handler := &updateHandler{
		Handler: h,
	}

	err = db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		if err := handler.updateSimulate(_ctx, tx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return h.GetSimulate(ctx)
}
