package goodbase

import (
	"fmt"
	"time"
)

func (h *Handler) ConstructGoodBaseSql() string {
	comma := ""
	now := uint32(time.Now().Unix())
	_sql := "insert into good_bases "
	_sql += "("
	if h.EntID != nil {
		_sql += "ent_id"
		comma = ", "
	}
	_sql += comma + "good_type"
	comma = ", "
	_sql += comma + "benefit_type"
	_sql += comma + "name"
	_sql += comma + "service_start_at"
	_sql += comma + "start_mode"
	if h.TestOnly != nil {
		_sql += comma + "test_only"
	}
	_sql += comma + "benefit_interval_hours"
	if h.Purchasable != nil {
		_sql += comma + "purchasable"
	}
	if h.Online != nil {
		_sql += comma + "online"
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
	_sql += fmt.Sprintf("%v'%v' as good_type", comma, h.GoodType.String())
	comma = ", "
	_sql += fmt.Sprintf("%v'%v' as benefit_type", comma, h.BenefitType.String())
	_sql += fmt.Sprintf("%v'%v' as name", comma, *h.Name)
	_sql += fmt.Sprintf("%v'%v' as service_start_at", comma, *h.ServiceStartAt)
	_sql += fmt.Sprintf("%v'%v' as start_mode", comma, h.StartMode.String())
	if h.TestOnly != nil {
		_sql += fmt.Sprintf("%v'%v' as test_only", comma, *h.TestOnly)
	}
	_sql += fmt.Sprintf("%v'%v' as benefit_interval_hours", comma, *h.BenefitIntervalHours)
	if h.Purchasable != nil {
		_sql += fmt.Sprintf("%v'%v' as purchasable", comma, *h.Purchasable)
	}
	if h.Online != nil {
		_sql += fmt.Sprintf("%v'%v' as inline", comma, *h.Online)
	}
	_sql += fmt.Sprintf("%v%v as created_at", comma, now)
	_sql += fmt.Sprintf("%v%v as updated_at", comma, now)
	_sql += fmt.Sprintf("%v0 as deleted_at", comma)
	_sql += ") as tmp "
	_sql += "where not exists ("
	_sql += "select 1 from good_coins "
	_sql += fmt.Sprintf("where name = '%v'", *h.Name)
	_sql += ") limit 1"

	return _sql
}
