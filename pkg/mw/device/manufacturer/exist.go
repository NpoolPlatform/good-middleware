package manufacturer

import (
	"context"

	manufacturercrud "github.com/NpoolPlatform/good-middleware/pkg/crud/device/manufacturer"
	"github.com/NpoolPlatform/good-middleware/pkg/db"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent"
	entdevicemanufacturer "github.com/NpoolPlatform/good-middleware/pkg/db/ent/devicemanufacturer"
)

func (h *Handler) ExistManufacturer(ctx context.Context) (exist bool, err error) {
	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		exist, err = cli.
			DeviceManufacturer.
			Query().
			Where(
				entdevicemanufacturer.EntID(*h.EntID),
				entdevicemanufacturer.DeletedAt(0),
			).
			Exist(_ctx)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return false, err
	}
	return exist, nil
}

func (h *Handler) ExistManufacturerConds(ctx context.Context) (exist bool, err error) {
	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		stm, err := manufacturercrud.SetQueryConds(cli.DeviceManufacturer.Query(), h.ManufacturerConds)
		if err != nil {
			return err
		}
		exist, err = stm.Exist(_ctx)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return false, err
	}
	return exist, nil
}
