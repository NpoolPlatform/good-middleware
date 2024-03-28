package powerrental

import (
	"context"
	"fmt"
	"time"

	"github.com/NpoolPlatform/good-middleware/pkg/db"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent"
	appgoodbase1 "github.com/NpoolPlatform/good-middleware/pkg/mw/app/good/goodbase"
	appgoodstock1 "github.com/NpoolPlatform/good-middleware/pkg/mw/app/good/stock"
	appmininggoodstock1 "github.com/NpoolPlatform/good-middleware/pkg/mw/app/good/stock/mining"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type createHandler struct {
	*powerRentalAppGoodQueryHandler
	stockValidator         *validateStockHandler
	sqlAppPowerRental      string
	sqlAppGoodBase         string
	sqlAppGoodStock        string
	sqlAppMiningGoodStocks []string
}

func (h *createHandler) constructAppGoodBaseSql(ctx context.Context) error {
	handler, err := appgoodbase1.NewHandler(
		ctx,
		appgoodbase1.WithEntID(func() *string { s := h.AppGoodBaseReq.EntID.String(); return &s }(), false),
		appgoodbase1.WithAppID(func() *string { s := h.AppGoodBaseReq.AppID.String(); return &s }(), true),
		appgoodbase1.WithGoodID(func() *string { s := h.AppGoodBaseReq.GoodID.String(); return &s }(), true),
		appgoodbase1.WithName(h.AppGoodBaseReq.Name, true),
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
	h.sqlAppGoodBase = handler.ConstructCreateSql()
	return nil
}

func (h *createHandler) constructAppMiningGoodStockSql(ctx context.Context) error {
	for _, poolStock := range h.AppMiningGoodStockReqs {
		handler, err := appmininggoodstock1.NewHandler(
			ctx,
			appmininggoodstock1.WithEntID(func() *string { s := poolStock.EntID.String(); return &s }(), false),
			appmininggoodstock1.WithAppGoodStockID(func() *string { s := poolStock.AppGoodStockID.String(); return &s }(), false),
			appmininggoodstock1.WithMiningGoodStockID(func() *string { s := poolStock.MiningGoodStockID.String(); return &s }(), false),
			appmininggoodstock1.WithReserved(func() *string { s := poolStock.Reserved.String(); return &s }(), true),
		)
		if err != nil {
			return err
		}
		h.sqlAppMiningGoodStocks = append(h.sqlAppMiningGoodStocks, handler.ConstructCreateSql())
	}
	return nil
}

func (h *createHandler) constructAppGoodStockSql(ctx context.Context) error {
	handler, err := appgoodstock1.NewHandler(
		ctx,
		appgoodstock1.WithEntID(func() *string { s := h.AppGoodStockReq.EntID.String(); return &s }(), false),
		appgoodstock1.WithAppGoodID(func() *string { s := h.AppGoodStockReq.AppGoodID.String(); return &s }(), false),
		appgoodstock1.WithReserved(func() *string { s := h.AppGoodStockReq.Reserved.String(); return &s }(), true),
	)
	if err != nil {
		return err
	}
	h.sqlAppGoodStock = handler.ConstructCreateSql()
	return nil
}

func (h *createHandler) constructAppPowerRentalSql() {
	comma := ""
	now := uint32(time.Now().Unix())
	_sql := "insert into app_power_rentals "
	_sql += "("
	if h.EntID != nil {
		_sql += "ent_id"
		comma = ", "
	}
	_sql += comma + "app_good_id"
	comma = ", "
	_sql += comma + "service_start_at"
	if h.CancelMode != nil {
		_sql += comma + "cancel_mode"
	}
	if h.CancelableBeforeStartSeconds != nil {
		_sql += comma + "cancelable_before_start_seconds"
	}
	if h.EnableSetCommission != nil {
		_sql += comma + "enable_set_commission"
	}
	if h.MinOrderAmount != nil {
		_sql += comma + "min_order_amount"
	}
	if h.MaxOrderAmount != nil {
		_sql += comma + "max_order_amount"
	}
	if h.MaxUserAmount != nil {
		_sql += comma + "max_user_amount"
	}
	if h.MinOrderDuration != nil {
		_sql += comma + "min_order_duration"
	}
	if h.MaxOrderDuration != nil {
		_sql += comma + "max_order_duration"
	}
	_sql += comma + "unit_price"
	if h.SaleStartAt != nil {
		_sql += comma + "sale_start_at"
	}
	if h.SaleEndAt != nil {
		_sql += comma + "sale_end_at"
	}
	if h.SaleMode != nil {
		_sql += comma + "sale_mode"
	}
	if h.FixedDuration != nil {
		_sql += comma + "fixed_duration"
	}
	if h.PackageWithRequireds != nil {
		_sql += comma + "package_with_requireds"
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
	_sql += fmt.Sprintf("%v'%v' as app_good_id", comma, *h.AppGoodID)
	comma = ", "
	_sql += fmt.Sprintf("%v%v as service_start_at", comma, *h.ServiceStartAt)
	if h.CancelMode != nil {
		_sql += fmt.Sprintf("%v'%v' as cancel_mode", comma, h.CancelMode.String())
	}
	if h.CancelableBeforeStartSeconds != nil {
		_sql += fmt.Sprintf("%v%v as cancelable_before_start_seconds", comma, *h.CancelableBeforeStartSeconds)
	}
	if h.EnableSetCommission != nil {
		_sql += fmt.Sprintf("%v%v as enable_set_commission", comma, *h.EnableSetCommission)
	}
	if h.MinOrderAmount != nil {
		_sql += fmt.Sprintf("%v%v as min_order_amount", comma, *h.MinOrderAmount)
	}
	if h.MaxOrderAmount != nil {
		_sql += fmt.Sprintf("%v%v as max_order_amount", comma, *h.MaxOrderAmount)
	}
	if h.MaxUserAmount != nil {
		_sql += fmt.Sprintf("%v%v as max_user_amount", comma, *h.MaxUserAmount)
	}
	if h.MinOrderDuration != nil {
		_sql += fmt.Sprintf("%v%v as min_order_duration", comma, *h.MinOrderDuration)
	}
	if h.MaxOrderDuration != nil {
		_sql += fmt.Sprintf("%v%v as max_order_duration", comma, *h.MaxOrderDuration)
	}
	_sql += fmt.Sprintf("%v'%v' as unit_price", comma, *h.UnitPrice)
	if h.SaleStartAt != nil {
		_sql += fmt.Sprintf("%v%v as sale_start_at", comma, *h.SaleStartAt)
	}
	if h.SaleEndAt != nil {
		_sql += fmt.Sprintf("%v%v as sale_end_at", comma, *h.SaleEndAt)
	}
	if h.SaleMode != nil {
		_sql += fmt.Sprintf("%v'%v' as sale_mode", comma, *h.SaleMode)
	}
	if h.FixedDuration != nil {
		_sql += fmt.Sprintf("%v%v as fixed_duration", comma, *h.FixedDuration)
	}
	if h.PackageWithRequireds != nil {
		_sql += fmt.Sprintf("%v%v as package_with_requireds", comma, *h.PackageWithRequireds)
	}
	_sql += fmt.Sprintf("%v%v as created_at", comma, now)
	_sql += fmt.Sprintf("%v%v as updated_at", comma, now)
	_sql += fmt.Sprintf("%v0 as deleted_at", comma)
	_sql += ") as tmp "
	_sql += "where not exists ("
	_sql += "select 1 from ("
	_sql += "select * from app_power_rentals as apr "
	_sql += fmt.Sprintf("where app_good_id = '%v'", *h.AppGoodID)
	_sql += " limit 1) as tmp)"
	_sql += "and exists ("
	_sql += "select 1 from power_rentals "
	_sql += fmt.Sprintf("where good_id = '%v'", *h.AppGoodBaseReq.GoodID)
	_sql += " limit 1)"
	h.sqlAppPowerRental = _sql
}

func (h *createHandler) execSql(ctx context.Context, tx *ent.Tx, sql string) error {
	rc, err := tx.ExecContext(ctx, sql)
	if err != nil {
		return err
	}
	n, err := rc.RowsAffected()
	if err != nil || n != 1 {
		return fmt.Errorf("fail create apppowerrental: %v", err)
	}
	return nil
}

func (h *createHandler) createAppPowerRental(ctx context.Context, tx *ent.Tx) error {
	return h.execSql(ctx, tx, h.sqlAppPowerRental)
}

func (h *createHandler) createAppGoodBase(ctx context.Context, tx *ent.Tx) error {
	return h.execSql(ctx, tx, h.sqlAppGoodBase)
}

func (h *createHandler) validateFixedDurationUnitPrice() error {
	if h.MinOrderDuration != h.MaxOrderDuration {
		return fmt.Errorf("invalid order duration")
	}
	unitPrice := h._ent.powerRental.UnitPrice.Mul(decimal.NewFromInt(int64(*h.MaxOrderDuration)))
	if h.UnitPrice.Cmp(unitPrice) < 0 {
		return fmt.Errorf("invalid unitprice")
	}
	return nil
}

func (h *createHandler) validateUnitPrice(ctx context.Context) error {
	if h.FixedDuration != nil && *h.FixedDuration {
		return h.validateFixedDurationUnitPrice()
	}
	if h.UnitPrice.Cmp(h._ent.powerRental.UnitPrice) < 0 {
		return fmt.Errorf("invalid unitprice")
	}
	return nil
}

func (h *createHandler) _validateStock() error {
	if h.AppGoodStockReq.EntID == nil {
		h.AppGoodStockReq.EntID = func() *uuid.UUID { uid := uuid.New(); return &uid }()
	}
	for _, poolStock := range h.AppMiningGoodStockReqs {
		poolStock.AppGoodStockID = h.AppGoodStockReq.EntID
	}
	if h.AppGoodBaseReq.EntID == nil {
		h.AppGoodBaseReq.EntID = func() *uuid.UUID { uid := uuid.New(); return &uid }()
		h.AppGoodStockReq.AppGoodID = h.AppGoodBaseReq.EntID
	}
	return h.stockValidator.validateStock()
}

func (h *createHandler) createAppStock(ctx context.Context, tx *ent.Tx) error {
	return h.execSql(ctx, tx, h.sqlAppGoodStock)
}

func (h *createHandler) createAppMiningGoodStocks(ctx context.Context, tx *ent.Tx) error {
	for _, _sql := range h.sqlAppMiningGoodStocks {
		if err := h.execSql(ctx, tx, _sql); err != nil {
			return err
		}
	}
	return nil
}

func (h *Handler) CreatePowerRental(ctx context.Context) error {
	handler := &createHandler{
		powerRentalAppGoodQueryHandler: &powerRentalAppGoodQueryHandler{
			Handler: h,
		},
		stockValidator: &validateStockHandler{
			Handler: h,
		},
	}

	if err := handler.requirePowerRentalGood(ctx); err != nil {
		return err
	}
	if err := handler._validateStock(); err != nil {
		return err
	}
	if err := handler.validateUnitPrice(ctx); err != nil {
		return err
	}

	handler.constructAppPowerRentalSql()
	if err := handler.constructAppGoodBaseSql(ctx); err != nil {
		return err
	}
	if err := handler.constructAppGoodStockSql(ctx); err != nil {
		return err
	}
	if err := handler.constructAppMiningGoodStockSql(ctx); err != nil {
		return err
	}

	return db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		if err := handler.createAppGoodBase(_ctx, tx); err != nil {
			return err
		}
		if err := handler.createAppStock(_ctx, tx); err != nil {
			return err
		}
		if err := handler.createAppMiningGoodStocks(_ctx, tx); err != nil {
			return err
		}
		return handler.createAppPowerRental(_ctx, tx)
	})
}