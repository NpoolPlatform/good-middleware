package constraint

import (
	"context"
	"fmt"
	"time"

	"github.com/NpoolPlatform/good-middleware/pkg/db"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent"

	"github.com/google/uuid"
)

type createHandler struct {
	*Handler
	sql string
}

//nolint:goconst
func (h *createHandler) constructSQL() {
	now := uint32(time.Now().Unix())
	comma := ""
	_sql := "insert into top_most_good_constraints ("
	if h.EntID != nil {
		_sql += "ent_id"
		comma = ", "
	}
	_sql += comma + "top_most_good_id"
	comma = ", "
	_sql += comma + "`constraint`"
	_sql += comma + "target_value"
	if h.Index != nil {
		_sql += comma + "`index`"
	}
	_sql += comma + "created_at"
	_sql += comma + "updated_at"
	_sql += comma + "deleted_at"
	_sql += ")"

	comma = ""
	_sql += " select * from (select "
	if h.EntID != nil {
		_sql += fmt.Sprintf("'%v' as ent_id", *h.EntID)
		comma = ", "
	}
	_sql += fmt.Sprintf("%v'%v' as top_most_good_id", comma, *h.TopMostGoodID)
	comma = ", "
	_sql += fmt.Sprintf("%v'%v' as `constraint`", comma, h.Constraint.String())
	_sql += fmt.Sprintf("%v'%v' as target_value", comma, *h.TargetValue)
	if h.Index != nil {
		_sql += fmt.Sprintf("%v%v as `index`", comma, *h.Index)
	}
	_sql += fmt.Sprintf("%v%v as created_at", comma, now)
	_sql += fmt.Sprintf("%v%v as updated_at", comma, now)
	_sql += fmt.Sprintf("%v0 as deleted_at", comma)
	_sql += ") as tmp "
	_sql += "where not exists ("
	_sql += "select 1 from top_most_good_constraints as tmgc "
	_sql += fmt.Sprintf("where tmgc.top_most_good_id = '%v' and tmgc.constraint = '%v' and deleted_at = 0", *h.TopMostGoodID, h.Constraint.String())
	_sql += " limit 1) and exists ("
	_sql += "select 1 from top_most_goods "
	_sql += fmt.Sprintf("where ent_id = '%v'", *h.TopMostGoodID)
	_sql += " limit 1)"

	h.sql = _sql
}

func (h *createHandler) createConstraint(ctx context.Context, tx *ent.Tx) error {
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

func (h *Handler) CreateConstraint(ctx context.Context) error {
	handler := &createHandler{
		Handler: h,
	}
	if h.EntID == nil {
		h.EntID = func() *uuid.UUID { s := uuid.New(); return &s }()
	}
	handler.constructSQL()
	return db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		return handler.createConstraint(_ctx, tx)
	})
}
