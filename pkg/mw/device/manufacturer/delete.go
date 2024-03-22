package manufacturer

import (
	"context"
	"time"

	manufacturercrud "github.com/NpoolPlatform/good-middleware/pkg/crud/device/manufacturer"
	"github.com/NpoolPlatform/good-middleware/pkg/db"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent"
)

type deleteHandler struct {
	*Handler
	now uint32
}

func (h *deleteHandler) deleteManufacturer(ctx context.Context, cli *ent.Client) error {
	if _, err := manufacturercrud.UpdateSet(
		cli.DeviceManufacturer.UpdateOneID(*h.ID),
		&manufacturercrud.Req{
			DeletedAt: &h.now,
		},
	).Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *Handler) DeleteManufacturer(ctx context.Context) error {
	handler := &deleteHandler{
		Handler: h,
		now:     uint32(time.Now().Unix()),
	}

	info, err := h.GetManufacturer(ctx)
	if err != nil {
		return err
	}
	if info == nil {
		return nil
	}

	return db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		return handler.deleteManufacturer(_ctx, cli)
	})
}
