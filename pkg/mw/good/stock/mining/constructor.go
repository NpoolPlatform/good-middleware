package mining

import (
	"fmt"
	"time"
)

func (h *Handler) ConstructCreateSql() string {
	comma := ""
	now := uint32(time.Now().Unix())
	_sql := "insert into mining_good_stocks "
	_sql += "("
	if h.EntID != nil {
		_sql += "ent_id"
		comma = ", "
	}
	_sql += comma + "good_stock_id"
	comma = ", "
	_sql += comma + "mining_pool_id"
	_sql += comma + "pool_good_user_id"
	_sql += comma + "total"
	_sql += comma + "spot_quantity"
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
	_sql += fmt.Sprintf("%v'%v' as good_stock_id", comma, *h.GoodStockID)
	comma = ", "
	_sql += fmt.Sprintf("%v'%v' as mining_pool_id", comma, *h.MiningPoolID)
	_sql += fmt.Sprintf("%v'%v' as pool_good_user_id", comma, *h.PoolGoodUserID)
	_sql += fmt.Sprintf("%v'%v' as total", comma, *h.Total)
	_sql += fmt.Sprintf("%v'%v' as spot_quantity", comma, *h.Total)
	_sql += fmt.Sprintf("%v%v as created_at", comma, now)
	_sql += fmt.Sprintf("%v%v as updated_at", comma, now)
	_sql += fmt.Sprintf("%v0 as deleted_at", comma)
	_sql += ") as tmp "
	_sql += "where exists ("
	_sql += "select 1 from stocks_v1 "
	_sql += fmt.Sprintf("where ent_id = '%v'", *h.GoodStockID)
	_sql += " limit 1)"

	return _sql
}