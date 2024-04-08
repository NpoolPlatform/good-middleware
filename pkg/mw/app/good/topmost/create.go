package topmost

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

//nolint:goconst
func (h *createHandler) constructSQL() {
	now := uint32(time.Now().Unix())
	comma := ""
	_sql := "insert into top_mosts ("
	if h.EntID != nil {
		_sql += "ent_id"
		comma = ", "
	}
	_sql += comma + "app_id"
	comma = ", "
	_sql += comma + "top_most_type"
	_sql += comma + "title"
	_sql += comma + "message"
	if h.TargetURL != nil {
		_sql += comma + "target_url"
	}
	_sql += comma + "start_at"
	_sql += comma + "end_at"
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
	_sql += fmt.Sprintf("%v'%v' as app_id", comma, *h.AppID)
	comma = ", "
	_sql += fmt.Sprintf("%v'%v' as top_most_type", comma, h.TopMostType.String())
	_sql += fmt.Sprintf("%v'%v' as title", comma, *h.Title)
	_sql += fmt.Sprintf("%v'%v' as message", comma, *h.Message)
	if h.TargetURL != nil {
		_sql += fmt.Sprintf("%v'%v' as target_url", comma, *h.TargetURL)
	}
	_sql += fmt.Sprintf("%v'%v' as start_at", comma, *h.StartAt)
	_sql += fmt.Sprintf("%v'%v' as end_at", comma, *h.EndAt)
	_sql += fmt.Sprintf("%v%v as created_at", comma, now)
	_sql += fmt.Sprintf("%v%v as updated_at", comma, now)
	_sql += fmt.Sprintf("%v0 as deleted_at", comma)
	_sql += ") as tmp "
	_sql += "where not exists ("
	_sql += "select 1 from top_mosts as tm "
	_sql += fmt.Sprintf("where tm.app_id = '%v' and tm.top_most_type ='%v'", *h.AppID, h.TopMostType.String())
	_sql += " limit 1)"

	h.sql = _sql
}

func (h *createHandler) createTopMost(ctx context.Context, tx *ent.Tx) error {
	rc, err := tx.ExecContext(ctx, h.sql)
	if err != nil {
		return err
	}
	n, err := rc.RowsAffected()
	if err != nil || n != 1 {
		return fmt.Errorf("fail create topmost: %v", err)
	}
	return nil
}

func (h *Handler) CreateTopMost(ctx context.Context) error {
	if err := h.formalizeStartEnd(ctx); err != nil {
		return err
	}
	if err := h.checkPromotion(ctx); err != nil {
		return err
	}

	handler := &createHandler{
		Handler: h,
	}
	handler.constructSQL()
	return db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		return handler.createTopMost(_ctx, tx)
	})
}
