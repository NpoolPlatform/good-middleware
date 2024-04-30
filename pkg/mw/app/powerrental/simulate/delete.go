package appsimulatepowerrental

import (
	"context"
	"time"

	wlog "github.com/NpoolPlatform/go-service-framework/pkg/wlog"
	appsimulatepowerrentalcrud "github.com/NpoolPlatform/good-middleware/pkg/crud/app/powerrental/simulate"
	"github.com/NpoolPlatform/good-middleware/pkg/db"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent"
)

type deleteHandler struct {
	*Handler
	now uint32
}

func (h *deleteHandler) deleteSimulate(ctx context.Context, cli *ent.Client) error {
	if _, err := appsimulatepowerrentalcrud.UpdateSet(
		cli.AppSimulatePowerRental.UpdateOneID(*h.ID),
		&appsimulatepowerrentalcrud.Req{
			DeletedAt: &h.now,
		},
	).Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *Handler) DeleteSimulate(ctx context.Context) error {
	info, err := h.GetSimulate(ctx)
	if err != nil {
		return wlog.WrapError(err)
	}
	if info == nil {
		return wlog.Errorf("invalid simulate")
	}

	h.ID = &info.ID
	handler := &deleteHandler{
		Handler: h,
		now:     uint32(time.Now().Unix()),
	}
	return db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		return handler.deleteSimulate(_ctx, cli)
	})
}
