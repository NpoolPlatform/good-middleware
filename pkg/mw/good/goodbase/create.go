package goodbase

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
	_sql := "insert into good_bases ("
	if h.EntID != nil {
		_sql += "ent_id"
		comma = ", "
	}
	_sql += comma + "good_type"
	comma = ", "
	if h.BenefitType != nil {
		_sql += comma + "benefit_type"
	}
	_sql += comma + "name"
	if h.ServiceStartAt != nil {
		_sql += comma + "service_start_at"
	}
	if h.StartMode != nil {
		_sql += comma + "start_mode"
	}
	if h.TestOnly != nil {
		_sql += comma + "test_only"
	}
	if h.BenefitIntervalHours != nil {
		_sql += comma + "benefit_interval_hours"
	}
	if h.Purchasable != nil {
		_sql += comma + "purchasable"
	}
	if h.Online != nil {
		_sql += comma + "online"
	}
	_sql += comma + "created_at"
	_sql += comma + "updated_at"
	_sql += comma + "deleted_at"
	comma = ""
	_sql += ") select * from (select "
	if h.EntID != nil {
		_sql += fmt.Sprintf("'%v' as ent_id", *h.EntID)
		comma = ", "
	}
	_sql += fmt.Sprintf("%v'%v' as good_type", comma, h.GoodType.String())
	comma = ", "
	if h.BenefitType != nil {
		_sql += fmt.Sprintf("%v'%v' as benefit_type", comma, h.BenefitType.String())
	}
	_sql += fmt.Sprintf("%v'%v' as name", comma, *h.Name)
	if h.ServiceStartAt != nil {
		_sql += fmt.Sprintf("%v'%v' as service_start_at", comma, *h.ServiceStartAt)
	}
	if h.StartMode != nil {
		_sql += fmt.Sprintf("%v'%v' as start_mode", comma, h.StartMode.String())
	}
	if h.TestOnly != nil {
		_sql += fmt.Sprintf("%v'%v' as test_only", comma, *h.TestOnly)
	}
	if h.BenefitIntervalHours != nil {
		_sql += fmt.Sprintf("%v'%v' as benefit_interval_hours", comma, *h.BenefitIntervalHours)
	}
	if h.Purchasable != nil {
		_sql += fmt.Sprintf("%v'%v' as purchasable", comma, *h.Purchasable)
	}
	if h.Online != nil {
		_sql += fmt.Sprintf("%v'%v'  as online", comma, *h.Online)
	}
	_sql += fmt.Sprintf("%v%v as created_at", comma, now)
	_sql += fmt.Sprintf("%v%v as updated_at", comma, now)
	_sql += fmt.Sprintf("%v0 as deleted_at", comma)
	_sql += ") as tmp "
	_sql += "where not exists ("
	_sql += fmt.Sprintf("select 1 from good_bases where name='%v' ", *h.Name)
	_sql += ") limit 1"
	h.sql = _sql
}

func (h *createHandler) createGoodBase(ctx context.Context, tx *ent.Tx) error {
	rc, err := tx.ExecContext(ctx, h.sql)
	if err != nil {
		return err
	}
	if n, err := rc.RowsAffected(); err != nil || n != 1 {
		return fmt.Errorf("fail create goodbase")
	}
	return nil
}

func (h *Handler) CreateGoodBase(ctx context.Context) error {
	handler := &createHandler{
		Handler: h,
	}
	handler.constructSql()
	return db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		return handler.createGoodBase(_ctx, tx)
	})
}
