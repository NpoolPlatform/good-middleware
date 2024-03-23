package powerrental

import (
	"context"
	"fmt"
	"time"

	"github.com/NpoolPlatform/good-middleware/pkg/db"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent"
	goodbase1 "github.com/NpoolPlatform/good-middleware/pkg/mw/good/goodbase"
	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
)

type updateHandler struct {
	*Handler
	sqlPowerRental string
	sqlGoodBase    string
}

func (h *updateHandler) constructGoodBaseSql(ctx context.Context) error {
	handler, err := goodbase1.NewHandler(
		ctx,
		goodbase1.WithEntID(func() *string { s := h.GoodBaseReq.EntID.String(); return &s }(), false),
		goodbase1.WithGoodType(h.GoodBaseReq.GoodType, false),
		goodbase1.WithBenefitType(h.GoodBaseReq.BenefitType, false),
		goodbase1.WithName(h.GoodBaseReq.Name, false),
		goodbase1.WithServiceStartAt(h.GoodBaseReq.ServiceStartAt, false),
		goodbase1.WithStartMode(h.GoodBaseReq.StartMode, false),
		goodbase1.WithTestOnly(h.GoodBaseReq.TestOnly, false),
		goodbase1.WithBenefitIntervalHours(h.GoodBaseReq.BenefitIntervalHours, false),
		goodbase1.WithPurchasable(h.GoodBaseReq.Purchasable, false),
		goodbase1.WithOnline(h.GoodBaseReq.Online, false),
	)
	if err != nil {
		return err
	}
	h.sqlGoodBase, err = handler.ConstructUpdateSql()
	if err == cruder.ErrUpdateNothing {
		return nil
	}
	return err
}

func (h *updateHandler) constructPowerRentalSql() error {
	set := "set "
	now := uint32(time.Now().Unix())

	_sql := "update power_rentals "
	if h.DeviceTypeID != nil {
		_sql += fmt.Sprintf("%vdevice_type_id = '%v', ", set, *h.DeviceTypeID)
		set = ""
	}
	if h.VendorLocationID != nil {
		_sql += fmt.Sprintf("%vvendor_location_id = '%v', ", set, *h.VendorLocationID)
		set = ""
	}
	if h.UnitPrice != nil {
		_sql += fmt.Sprintf("%vunit_price = '%v', ", set, *h.UnitPrice)
		set = ""
	}
	if h.QuantityUnit != nil {
		_sql += fmt.Sprintf("%vquantity_unit = '%v', ", set, *h.QuantityUnit)
		set = ""
	}
	if h.QuantityUnitAmount != nil {
		_sql += fmt.Sprintf("%vquantity_unit_amount = '%v', ", set, *h.QuantityUnitAmount)
		set = ""
	}
	if h.DeliveryAt != nil {
		_sql += fmt.Sprintf("%vdelivery_at = %v, ", set, *h.DeliveryAt)
		set = ""
	}
	if h.DurationType != nil {
		_sql += fmt.Sprintf("%vduration_type = '%v', ", set, h.DurationType.String())
		set = ""
	}
	if set != "" {
		return fmt.Errorf("update nothing")
	}
	_sql += fmt.Sprintf("updated_at = %v ", now)
	_sql += fmt.Sprintf("where id = %v ", *h.ID)
	_sql += "and not exists ("
	_sql += "select 1 from ("
	_sql += "select * from power_rentals as pr "
	_sql += fmt.Sprintf("where pr.good_id = '%v'", *h.GoodID)
	_sql += " limit 1) as tmp)"
	if h.DeviceTypeID != nil {
		_sql += " and exists ("
		_sql += "select 1 from device_infos "
		_sql += fmt.Sprintf("where ent_id = '%v'", *h.DeviceTypeID)
		_sql += " limit 1) "
	}
	if h.VendorLocationID != nil {
		_sql += " and exists ("
		_sql += "select 1 from vendor_locations "
		_sql += fmt.Sprintf("where ent_id = '%v'", *h.VendorLocationID)
		_sql += " limit 1) "
	}
	h.sqlPowerRental = _sql
	return nil
}

func (h *updateHandler) execSql(ctx context.Context, tx *ent.Tx, sql string) error {
	rc, err := tx.ExecContext(ctx, sql)
	if err != nil {
		return err
	}
	if _, err := rc.RowsAffected(); err != nil {
		return err
	}
	return nil
}

func (h *updateHandler) updatePowerRental(ctx context.Context, tx *ent.Tx) error {
	return h.execSql(ctx, tx, h.sqlPowerRental)
}

func (h *updateHandler) updateGoodBase(ctx context.Context, tx *ent.Tx) error {
	if h.sqlGoodBase == "" {
		return nil
	}
	return h.execSql(ctx, tx, h.sqlGoodBase)
}

func (h *Handler) UpdatePowerRental(ctx context.Context) error {
	handler := &updateHandler{
		Handler: h,
	}

	handler.constructPowerRentalSql()
	if err := handler.constructGoodBaseSql(ctx); err != nil {
		return err
	}

	return db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		if err := handler.updateGoodBase(_ctx, tx); err != nil {
			return err
		}
		return handler.updatePowerRental(_ctx, tx)
	})
}
