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
	sql            string
	manufacturerID string
}

func (h *updateHandler) constructSql() {
	comma := ""
	now := uint32(time.Now().Unix())

	_sql := "update device_infos "
	_sql += fmt.Sprintf("set type = '%v'", *h.Type)
	comma = ", "
	_sql += fmt.Sprintf("%vmanufacturer_id = '%v'", comma, h.manufacturerID)
	if h.PowerConsumption != nil {
		_sql += fmt.Sprintf("%vpower_consumption = '%v'", comma, *h.PowerConsumption)
	}
	if h.ShipmentAt != nil {
		_sql += fmt.Sprintf("%vshipment_at = '%v'", comma, *h.ShipmentAt)
	}
	_sql += fmt.Sprintf("%vupdated_at = %v ", comma, now)
	_sql += "where "
	_sql += fmt.Sprintf("id = %v ", *h.ID)
	_sql += "and not exists ("
	_sql += "select 1 from (select * from device_infos) as di "
	_sql += fmt.Sprintf("where di.type = '%v' and di.manufacturer_id = '%v' and di.id != %v", *h.Type, h.manufacturerID, *h.ID)
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

	info, err := h.GetDeviceType(ctx)
	if err != nil {
		return err
	}
	if info == nil {
		return fmt.Errorf("invalid devicetype")
	}

	if h.Type == nil {
		h.Type = &info.Type
	}
	if h.ManufacturerID == nil {
		handler.manufacturerID = info.ManufacturerID
	} else {
		handler.manufacturerID = h.ManufacturerID.String()
	}

	handler.constructSql()
	return db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		return handler.updateDeviceType(_ctx, tx)
	})
}
