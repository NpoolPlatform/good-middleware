package device

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"

	devicecrud "github.com/NpoolPlatform/good-middleware/pkg/crud/device"
	"github.com/NpoolPlatform/good-middleware/pkg/db"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent"
	entdevicetype "github.com/NpoolPlatform/good-middleware/pkg/db/ent/deviceinfo"
	entmanufacturer "github.com/NpoolPlatform/good-middleware/pkg/db/ent/devicemanufacturer"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/device"
)

type queryHandler struct {
	*Handler
	stmSelect *ent.DeviceInfoSelect
	stmCount  *ent.DeviceInfoSelect
	infos     []*npool.DeviceType
	total     uint32
}

func (h *queryHandler) selectDeviceType(stm *ent.DeviceInfoQuery) {
	h.stmSelect = stm.Select(entdevicetype.FieldID)
}

func (h *queryHandler) queryDeviceType(cli *ent.Client) error {
	if h.ID == nil && h.EntID == nil {
		return fmt.Errorf("invalid id")
	}
	stm := cli.DeviceInfo.Query().Where(entdevicetype.DeletedAt(0))
	if h.ID != nil {
		stm.Where(entdevicetype.ID(*h.ID))
	}
	if h.EntID != nil {
		stm.Where(entdevicetype.EntID(*h.EntID))
	}
	h.selectDeviceType(stm)
	return nil
}

func (h *queryHandler) queryDeviceTypes(ctx context.Context, cli *ent.Client) error {
	stm, err := devicecrud.SetQueryConds(cli.DeviceInfo.Query(), h.DeviceConds)
	if err != nil {
		return err
	}
	total, err := stm.Count(ctx)
	if err != nil {
		return err
	}
	h.total = uint32(total)
	h.selectDeviceType(stm)
	return nil
}

func (h *queryHandler) queryJoinMyself(s *sql.Selector) {
	t1 := sql.Table(entdevicetype.Table)
	s.LeftJoin(t1).
		On(
			s.C(entdevicetype.FieldEntID),
			t1.C(entdevicetype.FieldEntID),
		).
		AppendSelect(
			t1.C(entdevicetype.FieldEntID),
			t1.C(entdevicetype.FieldType),
			t1.C(entdevicetype.FieldManufacturerID),
			t1.C(entdevicetype.FieldPowerConsumption),
			t1.C(entdevicetype.FieldShipmentAt),
			t1.C(entdevicetype.FieldCreatedAt),
			t1.C(entdevicetype.FieldUpdatedAt),
		)
}

func (h *queryHandler) queryJoinManufacturer(s *sql.Selector) {
	t1 := sql.Table(entmanufacturer.Table)
	s.LeftJoin(t1).
		On(
			s.C(entdevicetype.FieldManufacturerID),
			t1.C(entmanufacturer.FieldEntID),
		).
		AppendSelect(
			sql.As(t1.C(entmanufacturer.FieldName), "manufacturer_name"),
			sql.As(t1.C(entmanufacturer.FieldLogo), "manufacturer_logo"),
		)
}

func (h *queryHandler) queryJoin() {
	h.stmSelect.Modify(func(s *sql.Selector) {
		h.queryJoinMyself(s)
		h.queryJoinManufacturer(s)
	})
	if h.stmCount == nil {
		return
	}
	h.stmSelect.Modify(func(s *sql.Selector) {
		h.queryJoinManufacturer(s)
	})
}

func (h *queryHandler) scan(ctx context.Context) error {
	return h.stmSelect.Scan(ctx, &h.infos)
}

func (h *queryHandler) formalize() {
	// TODO: nothing todo
}

func (h *Handler) GetDeviceType(ctx context.Context) (*npool.DeviceType, error) {
	handler := &queryHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryDeviceType(cli); err != nil {
			return err
		}
		handler.queryJoin()
		handler.stmSelect.
			Offset(0).
			Limit(2)
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

func (h *Handler) GetDeviceTypes(ctx context.Context) ([]*npool.DeviceType, uint32, error) {
	handler := &queryHandler{
		Handler: h,
	}
	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryDeviceTypes(_ctx, cli); err != nil {
			return err
		}
		handler.queryJoin()
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
