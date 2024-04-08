package appsimulatepowerrental

import (
	"context"
	"fmt"
	"time"

	"github.com/NpoolPlatform/good-middleware/pkg/db"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	"github.com/shopspring/decimal"
)

type updateHandler struct {
	*appPowerRentalHandler
	sql string
}

func (h *updateHandler) constructSQL() error {
	if h.ID == nil && h.EntID == nil && h.AppGoodID == nil {
		return fmt.Errorf("invalid simulateid")
	}

	set := "set "
	now := uint32(time.Now().Unix())

	_sql := "update app_simulate_power_rentals "
	if h.OrderUnits != nil {
		_sql += fmt.Sprintf("%vorder_units= '%v', ", set, *h.OrderUnits)
		set = ""
	}
	if h.OrderDuration != nil {
		_sql += fmt.Sprintf("%vorder_duration= %v, ", set, *h.OrderDuration)
		set = ""
	}
	if set != "" {
		return cruder.ErrUpdateNothing
	}
	_sql += fmt.Sprintf("updated_at = %v ", now)
	_sql += "where "
	and := ""
	if h.ID != nil {
		_sql += fmt.Sprintf("id = %v ", *h.ID)
		and = "and "
	}
	if h.EntID != nil {
		_sql += fmt.Sprintf("%vent_id = '%v' ", and, *h.EntID)
		and = "and "
	}
	if h.AppGoodID != nil {
		_sql += fmt.Sprintf("%vapp_good_id = '%v' ", and, *h.AppGoodID)
	}
	h.sql = _sql
	return nil
}

func (h *updateHandler) updateSimulate(ctx context.Context, tx *ent.Tx) error {
	rc, err := tx.ExecContext(ctx, h.sql)
	if err != nil {
		return err
	}
	if _, err := rc.RowsAffected(); err != nil {
		return err
	}
	return nil
}

func (h *updateHandler) validateOrderUnits() error {
	if h.OrderUnits == nil {
		return nil
	}
	if (h.appPowerRental.MinOrderAmount().Cmp(decimal.NewFromInt(0)) > 0 && h.OrderUnits.Cmp(h.appPowerRental.MinOrderAmount()) < 0) ||
		(h.appPowerRental.MaxOrderAmount().Cmp(decimal.NewFromInt(0)) > 0 && h.OrderUnits.Cmp(h.appPowerRental.MaxOrderAmount()) > 0) {
		return fmt.Errorf("invalid orderunits")
	}
	return nil
}

func (h *updateHandler) validateOrderDuration() error {
	if h.OrderDuration == nil {
		return nil
	}
	if (h.appPowerRental.MinOrderDuration() > 0 && *h.OrderDuration < h.appPowerRental.MinOrderDuration()) ||
		(h.appPowerRental.MaxOrderDuration() > 0 && *h.OrderDuration > h.appPowerRental.MaxOrderDuration()) {
		return fmt.Errorf("invalid orderduration")
	}
	return nil
}

func (h *Handler) UpdateSimulate(ctx context.Context) error {
	handler := &updateHandler{
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

	if err := handler.constructSQL(); err != nil {
		return err
	}

	return db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		return handler.updateSimulate(_ctx, tx)
	})
}
