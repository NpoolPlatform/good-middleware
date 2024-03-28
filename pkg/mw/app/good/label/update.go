package label

import (
	"context"
	"fmt"
	"time"

	"github.com/NpoolPlatform/good-middleware/pkg/db"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent"
	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
)

type updateHandler struct {
	*Handler
	sql       string
	appGoodID string
}

func (h *updateHandler) constructSql() error {
	set := "set "
	now := uint32(time.Now().Unix())

	_sql := "update app_good_labels "
	if h.Icon != nil {
		_sql += fmt.Sprintf("%vicon = '%v', ", set, *h.Icon)
		set = ""
	}
	if h.IconBgColor != nil {
		_sql += fmt.Sprintf("%vicon_bg_color = '%v', ", set, *h.IconBgColor)
		set = ""
	}
	if h.Label != nil {
		_sql += fmt.Sprintf("%vlabel = '%v', ", set, *h.Label)
		set = ""
	}
	if h.LabelBgColor != nil {
		_sql += fmt.Sprintf("%vlabel_bg_color = '%v', ", set, *h.LabelBgColor)
		set = ""
	}
	if set != "" {
		return cruder.ErrUpdateNothing
	}
	_sql += fmt.Sprintf("updated_at = %v ", now)
	_sql += "where "
	_sql += fmt.Sprintf("id = %v ", *h.ID)
	if h.Label != nil {
		_sql += "and not exists ("
		_sql += "select 1 from (select * from app_good_labels "
		_sql += fmt.Sprintf("where app_good_id = '%v' and label = '%v' and id != %v", h.appGoodID, *h.Label, *h.ID)
		_sql += " limit 1) as agl)"
	}

	h.sql = _sql
	return nil
}

func (h *updateHandler) updateLabel(ctx context.Context, tx *ent.Tx) error {
	rc, err := tx.ExecContext(ctx, h.sql)
	if err != nil {
		return err
	}
	if _, err := rc.RowsAffected(); err != nil {
		return err
	}
	return nil
}

func (h *Handler) UpdateLabel(ctx context.Context) error {
	handler := &updateHandler{
		Handler: h,
	}

	info, err := h.GetLabel(ctx)
	if err != nil {
		return err
	}
	if info == nil {
		return fmt.Errorf("invalid label")
	}

	handler.appGoodID = info.AppGoodID
	handler.constructSql()

	return db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		return handler.updateLabel(_ctx, tx)
	})
}
