package appsimulategood

import (
	"context"
	"fmt"
	"time"

	appsimulategoodcrud "github.com/NpoolPlatform/good-middleware/pkg/crud/app/good/simulate"
	"github.com/NpoolPlatform/good-middleware/pkg/db"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/app/good/simulate"
)

type deleteHandler struct {
	*Handler
}

func (h *deleteHandler) existSimulateCheck(ctx context.Context) error {
	if h.EntID == nil || h.ID == nil {
		return nil
	}
	h.Conds = &appsimulategoodcrud.Conds{
		ID:    &cruder.Cond{Op: cruder.EQ, Val: h.ID},
		EntID: &cruder.Cond{Op: cruder.EQ, Val: h.EntID},
	}
	exist, err := h.ExistSimulateConds(ctx)
	if err != nil {
		return err
	}
	if !exist {
		return fmt.Errorf("invalid id")
	}
	return nil
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

	handler := &deleteHandler{
		Handler: h,
	}

	if err := handler.existSimulateCheck(ctx); err != nil {
		return nil, err
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
