package appsimulatepowerrental

import (
	"context"
	"fmt"
	"time"

	"github.com/NpoolPlatform/good-middleware/pkg/db"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent"
	apppowerrental1 "github.com/NpoolPlatform/good-middleware/pkg/mw/app/powerrental"
)

type createHandler struct {
	*appPowerRentalHandler
	appPowerRental apppowerrental1.PowerRental
	sql            string
}

func (h *createHandler) constructSql() {
	now := uint32(time.Now().Unix())
	comma := ""
	_sql := "insert into app_simulate_power_rentals ("
	if h.EntID != nil {
		_sql += "ent_id"
		comma = ", "
	}
	_sql += comma + "app_good_id"
	comma = ", "
	_sql += comma + "order_units"
	_sql += comma + "order_duration"
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
	_sql += fmt.Sprintf("%v'%v' as order_units", comma, *h.OrderUnits)
	_sql += fmt.Sprintf("%v%v as order_duration", comma, *h.OrderDuration)
	_sql += fmt.Sprintf("%v%v as created_at", comma, now)
	_sql += fmt.Sprintf("%v%v as updated_at", comma, now)
	_sql += fmt.Sprintf("%v0 as deleted_at", comma)
	_sql += ") as tmp "
	_sql += "where not exists ("
	_sql += "select 1 from app_simulate_power_rentals as aspr "
	_sql += fmt.Sprintf("where aspr.app_good_id = '%v'", *h.AppGoodID)
	_sql += " limit 1)"

	h.sql = _sql
}

func (h *createHandler) createSimulate(ctx context.Context, tx *ent.Tx) error {
	rc, err := tx.ExecContext(ctx, h.sql)
	if err != nil {
		return err
	}
	n, err := rc.RowsAffected()
	if err != nil || n != 1 {
		return fmt.Errorf("fail create appsimulatepowerrental: %v", err)
	}
	return nil
}

func (h *createHandler) queryAppPowerRental(ctx context.Context) error {
	handler, err := apppowerrental1.NewHandler(
		ctx,
		apppowerrental1.WithAppGoodID(func() *string { s := h.AppGoodID.String(); return &s }(), true),
	)
	if err != nil {
		return err
	}
	if h.appPowerRental, err = handler.QueryPowerRentalEnt(ctx); err != nil {
		return err
	}
	return nil
}

func (h *createHandler) validateOrderUnits() error {
	if h.OrderUnits.Cmp(h.appPowerRental.MinOrderAmount()) < 0 ||
		h.OrderUnits.Cmp(h.appPowerRental.MaxOrderAmount()) > 0 {
		return fmt.Errorf("invalid orderunits")
	}
	return nil
}

func (h *createHandler) validateOrderDuration() error {
	if *h.OrderDuration < h.appPowerRental.MinOrderDuration() ||
		*h.OrderDuration > h.appPowerRental.MaxOrderDuration() {
		return fmt.Errorf("invalid orderduration")
	}
	return nil
}

func (h *Handler) CreateSimulate(ctx context.Context) error {
	handler := &createHandler{
		appPowerRentalHandler: &appPowerRentalHandler{
			Handler: h,
		},
	}

	if err := handler.queryAppPowerRental(ctx); err != nil {
		return err
	}
	if err := handler.validateOrderUnits(); err != nil {
		return err
	}
	if err := handler.validateOrderDuration(); err != nil {
		return err
	}

	handler.constructSql()
	return db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		return handler.createSimulate(_ctx, tx)
	})
}
