package powerrental

import (
	"context"
	"fmt"
	"time"

	"github.com/NpoolPlatform/good-middleware/pkg/db"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent"
	appgoodbase1 "github.com/NpoolPlatform/good-middleware/pkg/mw/app/good/goodbase"
	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	"github.com/shopspring/decimal"
)

type updateHandler struct {
	*powerRentalAppGoodQueryHandler
	sqlAppPowerRental string
	sqlAppGoodBase    string
}

func (h *updateHandler) constructAppGoodBaseSql(ctx context.Context) error {
	handler, err := appgoodbase1.NewHandler(
		ctx,
		appgoodbase1.WithEntID(func() *string { s := h.AppGoodBaseReq.EntID.String(); return &s }(), true),
		appgoodbase1.WithAppID(func() *string { s := h.AppGoodBaseReq.AppID.String(); return &s }(), false),
		appgoodbase1.WithGoodID(func() *string { s := h.AppGoodBaseReq.GoodID.String(); return &s }(), false),
		appgoodbase1.WithName(h.AppGoodBaseReq.Name, false),
		appgoodbase1.WithPurchasable(h.AppGoodBaseReq.Purchasable, false),
		appgoodbase1.WithEnableProductPage(h.AppGoodBaseReq.EnableProductPage, false),
		appgoodbase1.WithProductPage(h.AppGoodBaseReq.ProductPage, false),
		appgoodbase1.WithOnline(h.AppGoodBaseReq.Online, false),
		appgoodbase1.WithVisible(h.AppGoodBaseReq.Visible, false),
		appgoodbase1.WithDisplayIndex(h.AppGoodBaseReq.DisplayIndex, false),
		appgoodbase1.WithBanner(h.AppGoodBaseReq.Banner, false),
	)
	if err != nil {
		return err
	}
	h.sqlAppGoodBase, err = handler.ConstructUpdateSql()
	if err != cruder.ErrUpdateNothing {
		return err
	}
	return nil
}

func (h *updateHandler) constructAppPowerRentalSql() error {
	set := "set "
	now := uint32(time.Now().Unix())
	_sql := "update app_power_rentals "
	if h.ServiceStartAt != nil {
		_sql += fmt.Sprintf("%vservice_start_at = %v,", set, *h.ServiceStartAt)
		set = ""
	}
	if h.CancelMode != nil {
		_sql += fmt.Sprintf("%vcancel_mode = '%v',", set, h.CancelMode.String())
		set = ""
	}
	if h.CancelableBeforeStartSeconds != nil {
		_sql += fmt.Sprintf("%vcancelable_before_start_seconds = %v,", set, *h.CancelableBeforeStartSeconds)
		set = ""
	}
	if h.EnableSetCommission != nil {
		_sql += fmt.Sprintf("%venable_set_commission = %v,", set, *h.EnableSetCommission)
		set = ""
	}
	if h.MinOrderAmount != nil {
		_sql += fmt.Sprintf("%vmin_order_amount = '%v',", set, *h.MinOrderAmount)
		set = ""
	}
	if h.MaxOrderAmount != nil {
		_sql += fmt.Sprintf("%vmax_order_amount = '%v',", set, *h.MaxOrderAmount)
		set = ""
	}
	if h.MaxUserAmount != nil {
		_sql += fmt.Sprintf("%vmax_user_amount = '%v',", set, *h.MaxUserAmount)
		set = ""
	}
	if h.MinOrderDuration != nil {
		_sql += fmt.Sprintf("%vmin_order_duration = %v,", set, *h.MinOrderDuration)
		set = ""
	}
	if h.MaxOrderDuration != nil {
		_sql += fmt.Sprintf("%vmax_order_duration = %v,", set, *h.MaxOrderDuration)
		set = ""
	}
	if h.UnitPrice != nil {
		_sql += fmt.Sprintf("%vunit_price = '%v',", set, *h.UnitPrice)
		set = ""
	}
	if h.SaleStartAt != nil {
		_sql += fmt.Sprintf("%vsale_start_at = %v,", set, *h.SaleStartAt)
		set = ""
	}
	if h.SaleEndAt != nil {
		_sql += fmt.Sprintf("%vsale_end_at = %v,", set, *h.SaleEndAt)
		set = ""
	}
	if h.SaleMode != nil {
		_sql += fmt.Sprintf("%vsale_mode = '%v',", set, h.SaleMode.String())
		set = ""
	}
	if h.FixedDuration != nil {
		_sql += fmt.Sprintf("%vfixed_duration = %v,", set, *h.FixedDuration)
		set = ""
	}
	if h.PackageWithRequireds != nil {
		_sql += fmt.Sprintf("%vpackage_with_requireds = %v,", set, *h.PackageWithRequireds)
		set = ""
	}
	if set != "" {
		return cruder.ErrUpdateNothing
	}
	_sql += fmt.Sprintf("updated_at = %v", now)
	_sql += fmt.Sprintf(" where id = %v ", *h.ID)
	_sql += fmt.Sprintf(" and ent_id = '%v' ", *h.EntID)
	_sql += fmt.Sprintf(" and app_good_id = '%v'", *h.AppGoodID)
	h.sqlAppPowerRental = _sql
	return nil
}

func (h *updateHandler) execSql(ctx context.Context, tx *ent.Tx, sql string) error {
	if sql == "" {
		return nil
	}
	rc, err := tx.ExecContext(ctx, sql)
	if err != nil {
		return err
	}
	if _, err := rc.RowsAffected(); err != nil {
		return fmt.Errorf("fail update apppowerrental: %v", err)
	}
	return nil
}

func (h *updateHandler) updateAppPowerRental(ctx context.Context, tx *ent.Tx) error {
	return h.execSql(ctx, tx, h.sqlAppPowerRental)
}

func (h *updateHandler) updateAppGoodBase(ctx context.Context, tx *ent.Tx) error {
	return h.execSql(ctx, tx, h.sqlAppGoodBase)
}

func (h *updateHandler) validateFixedDurationUnitPrice() error {
	if h.MinOrderDuration != h.MaxOrderDuration {
		return fmt.Errorf("invalid order duration")
	}
	unitPrice := h._ent.powerRental.UnitPrice.Mul(decimal.NewFromInt(int64(*h.MaxOrderDuration)))
	if h.UnitPrice.Cmp(unitPrice) < 0 {
		return fmt.Errorf("invalid unitprice")
	}
	return nil
}

func (h *updateHandler) validateUnitPrice(ctx context.Context) error {
	if h.UnitPrice == nil && h.FixedDuration == nil && h.MinOrderDuration == nil && h.MaxOrderDuration == nil {
		return nil
	}

	if h.FixedDuration == nil {
		h.FixedDuration = &h._ent.appPowerRental.FixedDuration
	}
	if h.MinOrderDuration == nil {
		h.MinOrderDuration = &h._ent.appPowerRental.MinOrderDuration
	}
	if h.MaxOrderDuration == nil {
		h.MaxOrderDuration = &h._ent.appPowerRental.MaxOrderDuration
	}
	if h.UnitPrice == nil {
		h.UnitPrice = &h._ent.appPowerRental.UnitPrice
	}

	if *h.FixedDuration {
		return h.validateFixedDurationUnitPrice()
	}
	if h.UnitPrice.Cmp(h._ent.powerRental.UnitPrice) < 0 {
		return fmt.Errorf("invalid unitprice")
	}
	return nil
}

func (h *Handler) UpdatePowerRental(ctx context.Context) error {
	handler := &updateHandler{
		powerRentalAppGoodQueryHandler: &powerRentalAppGoodQueryHandler{
			Handler: h,
		},
	}

	if err := handler.requireAppPowerRentalAppGood(ctx); err != nil {
		return err
	}

	if h.ID == nil {
		h.ID = &handler._ent.appPowerRental.ID
	}
	if h.EntID == nil {
		h.EntID = &handler._ent.appPowerRental.EntID
	}
	if h.AppGoodBaseReq.GoodID == nil {
		h.AppGoodBaseReq.GoodID = &handler._ent.appGoodBase.GoodID
	}
	if h.AppGoodID == nil {
		h.AppGoodID = &handler._ent.appPowerRental.AppGoodID
		h.AppGoodBaseReq.EntID = h.AppGoodID
	}
	if h.AppGoodBaseReq.AppID == nil {
		h.AppGoodBaseReq.AppID = &handler._ent.appGoodBase.AppID
	}

	if err := handler.validateUnitPrice(ctx); err != nil {
		return err
	}

	handler.constructAppPowerRentalSql()
	if err := handler.constructAppGoodBaseSql(ctx); err != nil {
		return err
	}

	return db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		if err := handler.updateAppGoodBase(_ctx, tx); err != nil {
			return err
		}
		return handler.updateAppPowerRental(_ctx, tx)
	})
}
