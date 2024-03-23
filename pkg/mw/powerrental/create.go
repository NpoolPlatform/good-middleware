package powerrental

import (
	"context"
	"fmt"
	"time"

	"github.com/NpoolPlatform/good-middleware/pkg/db"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent"
	goodbase1 "github.com/NpoolPlatform/good-middleware/pkg/mw/good/goodbase"
)

type createHandler struct {
	*Handler
	sqlPowerRental string
	sqlGoodBase    string
}

func (h *createHandler) constructGoodBaseSql(ctx context.Context) error {
	handler, err := goodbase1.NewHandler(
		ctx,
		goodbase1.WithEntID(func() *string { s := h.GoodBaseReq.EntID.String(); return &s }(), false),
		goodbase1.WithGoodType(h.GoodBaseReq.GoodType, true),
		goodbase1.WithBenefitType(h.GoodBaseReq.BenefitType, true),
		goodbase1.WithName(h.GoodBaseReq.Name, true),
		goodbase1.WithServiceStartAt(h.GoodBaseReq.ServiceStartAt, true),
		goodbase1.WithStartMode(h.GoodBaseReq.StartMode, true),
		goodbase1.WithTestOnly(h.GoodBaseReq.TestOnly, false),
		goodbase1.WithBenefitIntervalHours(h.GoodBaseReq.BenefitIntervalHours, true),
		goodbase1.WithPurchasable(h.GoodBaseReq.Purchasable, false),
		goodbase1.WithOnline(h.GoodBaseReq.Online, false),
	)
	if err != nil {
		return err
	}
	h.sqlGoodBase = handler.ConstructCreateSql()
	return nil
}

func (h *createHandler) constructPowerRentalSql() {
	comma := ""
	now := uint32(time.Now().Unix())
	_sql := "insert into power_rentals "
	_sql += "("
	if h.EntID != nil {
		_sql += "ent_id"
		comma = ", "
	}
	_sql += comma + "good_id"
	comma = ", "
	_sql += comma + "device_type_id"
	_sql += comma + "vendor_location_id"
	_sql += comma + "unit_price"
	_sql += comma + "quantity_unit"
	_sql += comma + "quantity_unit_amount"
	_sql += comma + "delivery_at"
	if h.UnitLockDeposit != nil {
		_sql += comma + "unit_lock_deposit"
	}
	if h.DurationType != nil {
		_sql += comma + "duration_type"
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
	_sql += fmt.Sprintf("%v'%v' as good_id", comma, *h.GoodID)
	comma = ", "
	_sql += fmt.Sprintf("%v'%v' as device_type_id", comma, *h.DeviceTypeID)
	_sql += fmt.Sprintf("%v'%v' as vendor_location_id", comma, *h.VendorLocationID)
	_sql += fmt.Sprintf("%v'%v' as unit_price", comma, *h.UnitPrice)
	_sql += fmt.Sprintf("%v'%v' as quantity_unit", comma, *h.QuantityUnit)
	_sql += fmt.Sprintf("%v'%v' as quantity_unit_amount", comma, *h.QuantityUnitAmount)
	_sql += fmt.Sprintf("%v%v as delivery_at", comma, *h.DeliveryAt)
	if h.UnitLockDeposit != nil {
		_sql += fmt.Sprintf("%v%v as unit_lock_deposit", comma, *h.UnitLockDeposit)
	}
	if h.DurationType != nil {
		_sql += fmt.Sprintf("%v'%v' as duration_type", comma, h.DurationType.String())
	}
	_sql += fmt.Sprintf("%v%v as created_at", comma, now)
	_sql += fmt.Sprintf("%v%v as updated_at", comma, now)
	_sql += fmt.Sprintf("%v0 as deleted_at", comma)
	_sql += ") as tmp "
	_sql += "where not exists ("
	_sql += "select 1 from power_rentals "
	_sql += fmt.Sprintf("where good_id = '%v'", *h.GoodID)
	_sql += " limit 1) "
	_sql += "and exists ("
	_sql += "select 1 from good_bases "
	_sql += fmt.Sprintf("where ent_id = '%v'", *h.GoodID)
	_sql += " limit 1) "
	_sql += "and exists ("
	_sql += "select 1 from device_infos "
	_sql += fmt.Sprintf("where ent_id = '%v'", *h.DeviceTypeID)
	_sql += "limit 1) "
	_sql += "and exists ("
	_sql += "select 1 from vendor_locations "
	_sql += fmt.Sprintf("where ent_id = '%v'", *h.VendorLocationID)
	_sql += "limit 1)"
	h.sqlPowerRental = _sql
}

func (h *createHandler) execSql(ctx context.Context, tx *ent.Tx, sql string) error {
	rc, err := tx.ExecContext(ctx, sql)
	if err != nil {
		return err
	}
	n, err := rc.RowsAffected()
	if err != nil || n != 1 {
		return fmt.Errorf("fail create powerrental: %v", err)
	}
	return nil
}

func (h *createHandler) createPowerRental(ctx context.Context, tx *ent.Tx) error {
	return h.execSql(ctx, tx, h.sqlPowerRental)
}

func (h *createHandler) createGoodBase(ctx context.Context, tx *ent.Tx) error {
	return h.execSql(ctx, tx, h.sqlGoodBase)
}

func (h *Handler) CreatePowerRental(ctx context.Context) error {
	handler := &createHandler{
		Handler: h,
	}

	handler.constructPowerRentalSql()
	if err := handler.constructGoodBaseSql(ctx); err != nil {
		return err
	}

	return db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		if err := handler.createGoodBase(_ctx, tx); err != nil {
			return err
		}
		return handler.createPowerRental(_ctx, tx)
	})
}
