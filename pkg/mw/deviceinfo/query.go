package deviceinfo

import (
	"context"
	"encoding/json"
	"fmt"

	deviceinfocrud "github.com/NpoolPlatform/good-middleware/pkg/crud/deviceinfo"
	"github.com/NpoolPlatform/good-middleware/pkg/db"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent"
	entdeviceinfo "github.com/NpoolPlatform/good-middleware/pkg/db/ent/deviceinfo"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/deviceinfo"
)

type queryHandler struct {
	*Handler
	stmSelect *ent.DeviceInfoSelect
	infos     []*npool.DeviceInfo
	total     uint32
}

func (h *queryHandler) selectDeviceInfo(stm *ent.DeviceInfoQuery) {
	h.stmSelect = stm.Select(
		entdeviceinfo.FieldID,
		entdeviceinfo.FieldType,
		entdeviceinfo.FieldManufacturer,
		entdeviceinfo.FieldPowerConsumption,
		entdeviceinfo.FieldShipmentAt,
		entdeviceinfo.FieldPosters,
		entdeviceinfo.FieldCreatedAt,
		entdeviceinfo.FieldUpdatedAt,
	)
}

func (h *queryHandler) queryDeviceInfo(cli *ent.Client) {
	h.selectDeviceInfo(
		cli.DeviceInfo.
			Query().
			Where(
				entdeviceinfo.ID(*h.ID),
				entdeviceinfo.DeletedAt(0),
			),
	)
}

func (h *queryHandler) queryDeviceInfos(ctx context.Context, cli *ent.Client) error {
	stm, err := deviceinfocrud.SetQueryConds(cli.DeviceInfo.Query(), h.Conds)
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
	if h.ID == nil {
		return nil, fmt.Errorf("invalid id")
	}

	handler := &queryHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		handler.queryDeviceInfo(cli)
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
