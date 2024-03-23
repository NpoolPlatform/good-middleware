package device

import (
	"context"
	"fmt"
	"time"

	"github.com/NpoolPlatform/good-middleware/pkg/db"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent"
)

type createHandler struct {
	*Handler
	sql string
}

func (h *createHandler) constructSql() {
	comma := ""
	now := uint32(time.Now().Unix())
	_sql := "insert into device_infos "
	_sql += "("
	if h.EntID != nil {
		_sql += "ent_id"
		comma = ", "
	}
	_sql += comma + "type"
	comma = ", "
	_sql += comma + "manufacturer_id"
	if h.PowerConsumption != nil {
		_sql += comma + "power_consumption"
	}
	if h.ShipmentAt != nil {
		_sql += comma + "shipment_at"
	}
	_sql += comma + "created_at"
	_sql += comma + "updated_at"
	_sql += comma + "deleted_at"
	_sql += ")"
	comma = ""
	_sql += " select * from (select "
	if h.EntID != nil {
		_sql += fmt.Sprintf("'%v' as ent_id ", *h.EntID)
		comma = ", "
	}
	_sql += fmt.Sprintf("%v'%v' as type", comma, *h.Type)
	comma = ", "
	_sql += fmt.Sprintf("%v'%v' as manufacturer_id", comma, *h.ManufacturerID)
	if h.PowerConsumption != nil {
		_sql += fmt.Sprintf("%v'%v' as power_consumption", comma, *h.PowerConsumption)
	}
	if h.ShipmentAt != nil {
		_sql += fmt.Sprintf("%v'%v' as shipment_at", comma, *h.ShipmentAt)
	}
	_sql += fmt.Sprintf("%v%v as created_at", comma, now)
	_sql += fmt.Sprintf("%v%v as updated_at", comma, now)
	_sql += fmt.Sprintf("%v0 as deleted_at", comma)
	_sql += ") as tmp "
	_sql += "where not exists ("
	_sql += "select 1 from device_infos "
	_sql += fmt.Sprintf("where type='%v' and manufacturer='%v'", *h.Type, *h.ManufacturerID)
	_sql += " limit 1)"
	h.sql = _sql
}

func (h *createHandler) createDeviceType(ctx context.Context, tx *ent.Tx) error {
	rc, err := tx.ExecContext(ctx, h.sql)
	if err != nil {
		return err
	}
	n, err := rc.RowsAffected()
	if err != nil || n != 1 {
		return fmt.Errorf("fail create devicetype: %v", err)
	}
	return nil
}

func (h *Handler) CreateDeviceType(ctx context.Context) error {
	handler := &createHandler{
		Handler: h,
	}
	handler.constructSql()
	return db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		return handler.createDeviceType(_ctx, tx)
	})
}
