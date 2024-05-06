package manufacturer

import (
	"context"
	"fmt"
	"time"

	wlog "github.com/NpoolPlatform/go-service-framework/pkg/wlog"
	"github.com/NpoolPlatform/good-middleware/pkg/db"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent"
)

type updateHandler struct {
	*Handler
	sql string
}

func (h *updateHandler) constructSQL() error {
	set := "set "
	now := uint32(time.Now().Unix())

	_sql := "update device_manufacturers "
	if h.Name != nil {
		_sql += fmt.Sprintf("%vname = '%v', ", set, *h.Name)
		set = ""
	}
	if h.Logo != nil {
		_sql += fmt.Sprintf("%vlogo = '%v', ", set, *h.Logo)
		set = ""
	}
	if set != "" {
		return fmt.Errorf("update nothing")
	}

	_sql += fmt.Sprintf("updated_at = %v ", now)
	_sql += "where "
	_sql += fmt.Sprintf("id = %v ", *h.ID)

	if h.Name != nil {
		_sql += "and not exists ("
		_sql += "select 1 from (select * from device_manufacturers) as dm "
		_sql += fmt.Sprintf(
			"where dm.name = '%v' and dm.id != %v and deleted_at = 0",
			*h.Name,
			*h.ID,
		)
		_sql += " limit 1)"
	}

	h.sql = _sql
	return nil
}

func (h *updateHandler) updateManufacturer(ctx context.Context, tx *ent.Tx) error {
	rc, err := tx.ExecContext(ctx, h.sql)
	if err != nil {
		return err
	}
	if _, err := rc.RowsAffected(); err != nil {
		return err
	}
	return nil
}

func (h *Handler) UpdateManufacturer(ctx context.Context) error {
	info, err := h.GetManufacturer(ctx)
	if err != nil {
		return wlog.WrapError(err)
	}
	if info == nil {
		return wlog.Errorf("invalid manufacturer")
	}

	handler := &updateHandler{
		Handler: h,
	}
	h.ID = &info.ID
	if err := handler.constructSQL(); err != nil {
		return err
	}
	return db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		return handler.updateManufacturer(_ctx, tx)
	})
}
