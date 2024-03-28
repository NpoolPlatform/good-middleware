package appstock

import (
	"fmt"
	"time"
)

func (h *Handler) ConstructCreateSql() string {
	comma := ""
	now := uint32(time.Now().Unix())
	_sql := "insert into app_stocks "
	_sql += "("
	if h.EntID != nil {
		_sql += "ent_id"
		comma = ", "
	}
	_sql += comma + "app_good_id"
	comma = ", "
	_sql += comma + "reserved"
	_sql += comma + "spot_quantity"
	_sql += comma + "locked"
	_sql += comma + "in_service"
	_sql += comma + "wait_start"
	_sql += comma + "sold"
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
	_sql += fmt.Sprintf("%v'%v' as app_good_id", comma, *h.AppGoodID)
	comma = ", "
	_sql += fmt.Sprintf("%v'%v' as reserved", comma, *h.Reserved)
	_sql += fmt.Sprintf("%v'%v' as spot_quantity", comma, *h.Reserved)
	_sql += fmt.Sprintf("%v'0' as locked", comma)
	_sql += fmt.Sprintf("%v'0' as in_service", comma)
	_sql += fmt.Sprintf("%v'0' as wait_start", comma)
	_sql += fmt.Sprintf("%v'0' as sold", comma)
	_sql += fmt.Sprintf("%v%v as created_at", comma, now)
	_sql += fmt.Sprintf("%v%v as updated_at", comma, now)
	_sql += fmt.Sprintf("%v0 as deleted_at", comma)
	_sql += ") as tmp "
	_sql += "where exists ("
	_sql += "select 1 from app_good_bases "
	_sql += fmt.Sprintf("where ent_id = '%v'", *h.AppGoodID)
	_sql += " limit 1)"

	return _sql
}
