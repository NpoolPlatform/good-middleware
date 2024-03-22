package manufacturer

import (
	"context"
	"fmt"

	manufacturercrud "github.com/NpoolPlatform/good-middleware/pkg/crud/device/manufacturer"
	"github.com/NpoolPlatform/good-middleware/pkg/db"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent"
	entmanufacturer "github.com/NpoolPlatform/good-middleware/pkg/db/ent/devicemanufacturer"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/device/manufacturer"
)

type queryHandler struct {
	*Handler
	stmSelect *ent.DeviceManufacturerSelect
	infos     []*npool.Manufacturer
	total     uint32
}

func (h *queryHandler) selectDeviceManufacturer(stm *ent.DeviceManufacturerQuery) {
	h.stmSelect = stm.Select(
		entmanufacturer.FieldID,
		entmanufacturer.FieldEntID,
		entmanufacturer.FieldName,
		entmanufacturer.FieldLogo,
		entmanufacturer.FieldCreatedAt,
		entmanufacturer.FieldUpdatedAt,
	)
}

func (h *queryHandler) queryDeviceManufacturer(cli *ent.Client) error {
	if h.ID == nil && h.EntID == nil {
		return fmt.Errorf("invalid id")
	}
	stm := cli.DeviceManufacturer.Query().Where(entmanufacturer.DeletedAt(0))
	if h.ID != nil {
		stm.Where(entmanufacturer.ID(*h.ID))
	}
	if h.EntID != nil {
		stm.Where(entmanufacturer.EntID(*h.EntID))
	}
	h.selectDeviceManufacturer(stm)
	return nil
}

func (h *queryHandler) queryDeviceManufacturers(ctx context.Context, cli *ent.Client) error {
	stm, err := manufacturercrud.SetQueryConds(cli.DeviceManufacturer.Query(), h.ManufacturerConds)
	if err != nil {
		return err
	}
	total, err := stm.Count(ctx)
	if err != nil {
		return err
	}
	h.total = uint32(total)
	h.selectDeviceManufacturer(stm)
	return nil
}

func (h *queryHandler) scan(ctx context.Context) error {
	return h.stmSelect.Scan(ctx, &h.infos)
}

func (h *Handler) GetManufacturer(ctx context.Context) (*npool.Manufacturer, error) {
	handler := &queryHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryDeviceManufacturer(cli); err != nil {
			return err
		}
		return handler.scan(_ctx)
	})
	if err != nil {
		return nil, err
	}
	if len(handler.infos) == 0 {
		return nil, nil
	}
	if len(handler.infos) > 1 {
		return nil, fmt.Errorf("too many records")
	}
	return handler.infos[0], nil
}

func (h *Handler) GetManufacturers(ctx context.Context) ([]*npool.Manufacturer, uint32, error) {
	handler := &queryHandler{
		Handler: h,
	}
	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryDeviceManufacturers(_ctx, cli); err != nil {
			return err
		}
		handler.stmSelect.
			Offset(int(h.Offset)).
			Limit(int(h.Limit))
		return handler.scan(_ctx)
	})
	if err != nil {
		return nil, 0, err
	}
	return handler.infos, handler.total, nil
}
