package device

import (
	"context"
	"encoding/json"
	"fmt"

	devicecrud "github.com/NpoolPlatform/good-middleware/pkg/crud/device"
	"github.com/NpoolPlatform/good-middleware/pkg/db"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent"
	entdevice "github.com/NpoolPlatform/good-middleware/pkg/db/ent/deviceinfo"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/device"
)

type queryHandler struct {
	*Handler
	stmSelect *ent.DeviceInfoSelect
	infos     []*npool.DeviceInfo
	total     uint32
}

func (h *queryHandler) selectDeviceInfo(stm *ent.DeviceInfoQuery) {
	h.stmSelect = stm.Select(
		entdevice.FieldID,
		entdevice.FieldEntID,
		entdevice.FieldType,
		entdevice.FieldManufacturer,
		entdevice.FieldPowerConsumption,
		entdevice.FieldShipmentAt,
		entdevice.FieldPosters,
		entdevice.FieldCreatedAt,
		entdevice.FieldUpdatedAt,
	)
}

func (h *queryHandler) queryDeviceInfo(cli *ent.Client) error {
	if h.ID == nil && h.EntID == nil {
		return fmt.Errorf("invalid id")
	}
	stm := cli.DeviceInfo.Query().Where(entdevice.DeletedAt(0))
	if h.ID != nil {
		stm.Where(entdevice.ID(*h.ID))
	}
	if h.EntID != nil {
		stm.Where(entdevice.EntID(*h.EntID))
	}
	h.selectDeviceInfo(stm)
	return nil
}

func (h *queryHandler) queryDeviceInfos(ctx context.Context, cli *ent.Client) error {
	stm, err := devicecrud.SetQueryConds(cli.DeviceInfo.Query(), h.Conds)
	if err != nil {
		return err
	}
	total, err := stm.Count(ctx)
	if err != nil {
		return err
	}
	h.total = uint32(total)
	h.selectDeviceInfo(stm)
	return nil
}

func (h *queryHandler) scan(ctx context.Context) error {
	return h.stmSelect.Scan(ctx, &h.infos)
}

func (h *queryHandler) formalize() {
	for _, info := range h.infos {
		_ = json.Unmarshal([]byte(info.PostersStr), &info.Posters)
	}
}

func (h *Handler) GetDeviceInfo(ctx context.Context) (*npool.DeviceInfo, error) {
	handler := &queryHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryDeviceInfo(cli); err != nil {
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
	handler.formalize()
	return handler.infos[0], nil
}

func (h *Handler) GetDeviceInfos(ctx context.Context) ([]*npool.DeviceInfo, uint32, error) {
	handler := &queryHandler{
		Handler: h,
	}
	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryDeviceInfos(_ctx, cli); err != nil {
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
	handler.formalize()
	return handler.infos, handler.total, nil
}
