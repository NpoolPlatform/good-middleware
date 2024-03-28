package label

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
	now := uint32(time.Now().Unix())
	comma := ""
	_sql := "insert into app_good_labels ("
	if h.EntID != nil {
		_sql += "ent_id"
		comma = ", "
	}
	_sql += comma + "app_good_id"
	comma = ", "
	if h.Icon != nil {
		_sql += comma + "icon"
	}
	if h.IconBgColor != nil {
		_sql += comma + "icon_bg_color"
	}
	_sql += comma + "label"
	if h.LabelBgColor != nil {
		_sql += comma + "label_bg_color"
	}
	_sql += comma + "created_at"
	_sql += comma + "updated_at"
	_sql += comma + "deleted_at"
	_sql += ")"

	comma = ""
	_sql += " select * from ( select "
	if h.EntID != nil {
		_sql += fmt.Sprintf("'%v' as ent_id", *h.EntID)
		comma = ", "
	}
	_sql += fmt.Sprintf("%v'%v' as app_good_id", comma, *h.AppGoodID)
	comma = ", "
	if h.Icon != nil {
		_sql += fmt.Sprintf("%v'%v' as icon", comma, *h.Icon)
	}
	if h.IconBgColor != nil {
		_sql += fmt.Sprintf("%v'%v' as icon_bg_color", comma, *h.IconBgColor)
	}
	_sql += fmt.Sprintf("%v'%v' as label", comma, *h.Label)
	if h.LabelBgColor != nil {
		_sql += fmt.Sprintf("%v'%v' as label_bg_color", comma, *h.LabelBgColor)
	}
	_sql += fmt.Sprintf("%v%v as created_at", comma, now)
	_sql += fmt.Sprintf("%v%v as updated_at", comma, now)
	_sql += fmt.Sprintf("%v0 as deleted_at", comma)
	_sql += ") as tmp "
	_sql += "where exists ("
	_sql += "select 1 from app_good_bases "
	_sql += fmt.Sprintf("where ent_id = '%v'", *h.AppGoodID)
	_sql += " limit 1) "
	_sql += "and not exists ("
	_sql += "select 1 from app_good_labels "
	_sql += fmt.Sprintf("where label = '%v' and app_good_id = '%v'", *h.Label, *h.AppGoodID)
	_sql += " limit 1)"

	h.sql = _sql
}

func (h *createHandler) createLabel(ctx context.Context, tx *ent.Tx) error {
	rc, err := tx.ExecContext(ctx, h.sql)
	if err != nil {
		return err
	}
	n, err := rc.RowsAffected()
	if err != nil || n != 1 {
		return fmt.Errorf("fail create appdefaultgood: %v", err)
	}
	return nil
}

func (h *Handler) CreateLabel(ctx context.Context) error {
	handler := &createHandler{
		Handler: h,
	}
	handler.constructSql()
	return db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		return handler.createLabel(_ctx, tx)
	})
}
