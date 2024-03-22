package device

import (
	"context"
	"fmt"
	"time"

	"github.com/NpoolPlatform/good-middleware/pkg/db"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent"
)

type updateHandler struct {
	*Handler
	sql string
}

func (h *updateHandler) constructSql() {
	comma := ""
	now := uint32(time.Now().Unix())

	_sql := "update device_infos "
	if h.Type != nil {
		_sql += fmt.Sprintf("set type = '%v'", *h.Type)
		comma = ", "
	}
	if h.Manufacturer != nil {
		_sql += fmt.Sprintf("%vmanufacturer = '%v'", comma, *h.Manufacturer)
		comma = ", "
	}
	if h.PowerConsumption != nil {
		_sql += fmt.Sprintf("%vpower_consumption = '%v'", comma, *h.PowerConsumption)
		comma = ", "
	}
	if h.ShipmentAt != nil {
		_sql += fmt.Sprintf("%vshipment_at = '%v'", comma, *h.ShipmentAt)
		comma = ", "
	}
	_sql += fmt.Sprintf("%vupdated_at = %v ", comma, now)
	_sql += "where "
	_sql += fmt.Sprintf("id = %v ", *h.ID)
	_sql += "and not exists ("
	_sql += "select 1 from (select * from device_infos) as di "
	_sql += fmt.Sprintf("where di.type = '%v' and di.manufacturer = '%v' and di.id != %v", *h.Type, h.Manufacturer, *h.ID)
	_sql += ") limit 1"

	h.sql = _sql
}

func (h *updateHandler) updateDeviceType(ctx context.Context, tx *ent.Tx) error {
	rc, err := tx.ExecContext(ctx, h.sql)
	if err != nil {
		return err
	}
	if _, err := rc.RowsAffected(); err != nil {
		return err
	}
	return nil
}

func (h *Handler) UpdateDeviceType(ctx context.Context) error {
	handler := &updateHandler{
		Handler: h,
	}
	handler.constructSql()
	return db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		return handler.updateDeviceType(_ctx, tx)
	})
}
