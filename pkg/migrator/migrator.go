package migrator

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"time"

	"github.com/NpoolPlatform/go-service-framework/pkg/config"
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	constant "github.com/NpoolPlatform/go-service-framework/pkg/mysql/const"
	redis2 "github.com/NpoolPlatform/go-service-framework/pkg/redis"
	"github.com/NpoolPlatform/good-middleware/pkg/db"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent"
	entappgoodbase "github.com/NpoolPlatform/good-middleware/pkg/db/ent/appgoodbase"
	entgoodbase "github.com/NpoolPlatform/good-middleware/pkg/db/ent/goodbase"
	servicename "github.com/NpoolPlatform/good-middleware/pkg/servicename"
	types "github.com/NpoolPlatform/message/npool/basetypes/good/v1"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

const (
	keyUsername  = "username"
	keyPassword  = "password"
	keyDBName    = "database_name"
	maxOpen      = 5
	maxIdle      = 2
	MaxLife      = 0
	keyServiceID = "serviceid"
)

func lockKey() string {
	serviceID := config.GetStringValueWithNameSpace(servicename.ServiceDomain, keyServiceID)
	return fmt.Sprintf("%v:%v", basetypes.Prefix_PrefixMigrate, serviceID)
}

func dsn(hostname string) (string, error) {
	username := config.GetStringValueWithNameSpace(constant.MysqlServiceName, keyUsername)
	password := config.GetStringValueWithNameSpace(constant.MysqlServiceName, keyPassword)
	dbname := config.GetStringValueWithNameSpace(hostname, keyDBName)

	svc, err := config.PeekService(constant.MysqlServiceName)
	if err != nil {
		logger.Sugar().Warnw("dsn", "error", err)
		return "", err
	}

	return fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?parseTime=true&interpolateParams=true",
		username, password,
		svc.Address,
		svc.Port,
		dbname,
	), nil
}

func open(hostname string) (conn *sql.DB, err error) {
	hdsn, err := dsn(hostname)
	if err != nil {
		return nil, err
	}

	logger.Sugar().Warnw("open", "hdsn", hdsn)

	conn, err = sql.Open("mysql", hdsn)
	if err != nil {
		return nil, err
	}

	// https://github.com/go-sql-driver/mysql
	// See "Important settings" section.

	conn.SetConnMaxLifetime(time.Minute * MaxLife)
	conn.SetMaxOpenConns(maxOpen)
	conn.SetMaxIdleConns(maxIdle)

	return conn, nil
}

//nolint:funlen,gocyclo
func migrateGood(ctx context.Context, conn *sql.DB) error {
	rows, err := conn.QueryContext(
		ctx,
		`select
		  id,
		  ent_id,
		  device_info_id,
		  coin_type_id,
		  vendor_location_id,
		  unit_price,
		  benefit_type,
		  good_type,
		  title,
		  quantity_unit,
		  quantity_unit_amount,
		  delivery_at,
		  start_at,
		  start_mode,
		  test_only,
		  benefit_interval_hours,
		  unit_lock_deposit,
		  duration_type,
		  settlement_type,
		  created_at,
		  updated_at
		from
		  goods
		where
		  deleted_at = 0`,
	)
	if err != nil {
		return err
	}
	type _good struct {
		ID                   uint32
		EntID                uuid.UUID
		DeviceTypeID         uuid.UUID
		CoinTypeID           uuid.UUID
		VendorLocationID     uuid.UUID
		UnitPrice            decimal.Decimal
		BenefitType          string
		GoodType             string
		Title                string
		QuantityUnit         string
		QuantityUnitAmount   decimal.Decimal
		DeliveryAt           uint32
		StartAt              uint32
		StartMode            string
		TestOnly             bool
		BenefitIntervalHours uint32
		UnitLockDeposit      decimal.Decimal
		DurationDisplayType  string
		SettlementType       string
		CreatedAt            uint32
		UpdatedAt            uint32
	}
	_goods := []*_good{}
	for rows.Next() {
		_g := &_good{}
		if err := rows.Scan(
			&_g.ID,
			&_g.EntID,
			&_g.DeviceTypeID,
			&_g.CoinTypeID,
			&_g.VendorLocationID,
			&_g.UnitPrice,
			&_g.BenefitType,
			&_g.GoodType,
			&_g.Title,
			&_g.QuantityUnit,
			&_g.QuantityUnitAmount,
			&_g.DeliveryAt,
			&_g.StartAt,
			&_g.StartMode,
			&_g.TestOnly,
			&_g.BenefitIntervalHours,
			&_g.UnitLockDeposit,
			&_g.DurationDisplayType,
			&_g.SettlementType,
			&_g.CreatedAt,
			&_g.UpdatedAt,
		); err != nil {
			return err
		}
		_goods = append(_goods, _g)
	}
	for _, good := range _goods {
		if err := db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
			exist, err := tx.
				GoodBase.
				Query().
				Where(
					entgoodbase.EntID(good.EntID),
				).
				Exist(_ctx)
			if err != nil {
				return err
			}
			if exist {
				return nil
			}
			if good.GoodType == types.GoodType_PowerRental.String() {
				good.GoodType = types.GoodType_LegacyPowerRental.String()
			}
			if good.GoodType == "PowerRenting" {
				good.GoodType = types.GoodType_LegacyPowerRental.String()
			}
			if good.BenefitType == types.BenefitType_DefaultBenefitType.String() {
				good.BenefitType = types.BenefitType_BenefitTypeNone.String()
			}
			switch good.GoodType {
			case types.GoodType_MachineRental.String():
				fallthrough //nolint
			case types.GoodType_MachineCustody.String():
				good.BenefitType = types.BenefitType_BenefitTypePool.String()
			case types.GoodType_LegacyPowerRental.String():
				fallthrough //nolint
			case types.GoodType_PowerRental.String():
			case types.GoodType_TechniqueServiceFee.String():
				fallthrough //nolint
			case types.GoodType_ElectricityFee.String():
				good.BenefitType = types.BenefitType_BenefitTypeNone.String()
			}
			if _, err := tx.
				GoodBase.
				Create().
				SetEntID(good.EntID).
				SetGoodType(good.GoodType).
				SetBenefitType(good.BenefitType).
				SetName(good.Title).
				SetServiceStartAt(good.StartAt).
				SetStartMode(good.StartMode).
				SetTestOnly(good.TestOnly).
				SetBenefitIntervalHours(good.BenefitIntervalHours).
				SetPurchasable(true).
				SetOnline(true).
				SetCreatedAt(good.CreatedAt).
				SetUpdatedAt(good.UpdatedAt).
				Save(_ctx); err != nil {
				return err
			}
			switch good.GoodType {
			case types.GoodType_MachineRental.String():
			case types.GoodType_MachineCustody.String():
			case types.GoodType_LegacyPowerRental.String():
				fallthrough //nolint
			case types.GoodType_PowerRental.String():
				if _, err := tx.
					GoodCoin.
					Create().
					SetGoodID(good.EntID).
					SetCoinTypeID(good.CoinTypeID).
					SetMain(true).
					SetCreatedAt(good.CreatedAt).
					Save(_ctx); err != nil {
					return err
				}
				if _, err := tx.
					PowerRental.
					Create().
					SetGoodID(good.EntID).
					SetDeviceTypeID(good.DeviceTypeID).
					SetVendorLocationID(good.VendorLocationID).
					SetUnitPrice(good.UnitPrice).
					SetQuantityUnit(good.QuantityUnit).
					SetQuantityUnitAmount(good.QuantityUnitAmount).
					SetDeliveryAt(good.DeliveryAt).
					SetUnitLockDeposit(good.UnitLockDeposit).
					SetDurationDisplayType(good.DurationDisplayType).
					SetStockMode(types.GoodStockMode_GoodStockByUnique.String()).
					SetCreatedAt(good.CreatedAt).
					Save(_ctx); err != nil {
					return err
				}
			case types.GoodType_TechniqueServiceFee.String():
				fallthrough //nolint
			case types.GoodType_ElectricityFee.String():
				if _, err := tx.
					Fee.
					Create().
					SetGoodID(good.EntID).
					SetSettlementType(good.SettlementType).
					SetUnitValue(good.UnitPrice).
					SetDurationDisplayType(good.DurationDisplayType).
					SetCreatedAt(good.CreatedAt).
					Save(_ctx); err != nil {
					return err
				}
			}
			return nil
		}); err != nil {
			return err
		}
	}
	return nil
}

//nolint:funlen,gocyclo
func migrateAppGood(ctx context.Context, conn *sql.DB) error {
	query := `select
		  ag.id,
		  ag.ent_id,
		  ag.app_id,
		  ag.good_id,
		  ag.online,
		  ag.visible,
		  ag.good_name,
		  ag.unit_price,
		  ag.package_price,
		  ag.display_index,
		  ag.sale_start_at,
		  ag.sale_end_at,
		  ag.service_start_at,
		  ag.technical_fee_ratio,
		  ifnull(ag.descriptions, ''),
		  ag.good_banner,
		  ifnull(ag.display_names, ''),
		  ag.enable_purchase,
		  ag.enable_product_page,
		  ag.cancel_mode,
		  ifnull(ag.display_colors, ''),
		  ag.cancellable_before_start,
		  ag.product_page,
		  ag.enable_set_commission,
		  ifnull(ag.posters, ''),
		  ag.min_order_amount,
		  ag.max_order_amount,
		  ag.max_user_amount,
		  ag.min_order_duration,
		  ag.max_order_duration,
		  ag.package_with_requireds,
		  gb.good_type,
		  ag.created_at,
		  ag.updated_at
		from
		  app_goods as ag
		join
		  good_bases as gb
		on
		  gb.ent_id = ag.good_id
		where
		  ag.deleted_at = 0`
	rows, err := conn.QueryContext(ctx, query)
	if err != nil {
		return err
	}
	type _appGood struct {
		ID                      uint32
		EntID                   uuid.UUID
		AppID                   uuid.UUID
		GoodID                  uuid.UUID
		Online                  bool
		Visible                 bool
		GoodName                string
		UnitPrice               decimal.Decimal
		PackagePrice            decimal.Decimal
		DisplayIndex            int32
		SaleStartAt             uint32
		SaleEndAt               uint32
		ServiceStartAt          uint32
		TechniqueFeeRatio       decimal.Decimal
		DescriptionsStr         string
		Descriptions            []string
		GoodBanner              string
		DisplayNamesStr         string
		DisplayNames            []string
		EnablePurchase          bool
		EnableProductPage       bool
		CancelMode              string
		DisplayColorsStr        string
		DisplayColors           []string
		CancelableBeforeStart   uint32
		ProductPage             string
		EnableSetCommission     bool
		PostersStr              string
		Posters                 []string
		MinOrderAmount          decimal.Decimal
		MaxOrderAmount          decimal.Decimal
		MaxUserAmount           decimal.Decimal
		MinOrderDurationSeconds uint32
		MaxOrderDurationSeconds uint32
		PackageWithRequireds    bool
		GoodType                string
		CreatedAt               uint32
		UpdatedAt               uint32
	}
	_appGoods := []*_appGood{}
	for rows.Next() {
		_ag := &_appGood{}
		if err := rows.Scan(
			&_ag.ID,
			&_ag.EntID,
			&_ag.AppID,
			&_ag.GoodID,
			&_ag.Online,
			&_ag.Visible,
			&_ag.GoodName,
			&_ag.UnitPrice,
			&_ag.PackagePrice,
			&_ag.DisplayIndex,
			&_ag.SaleStartAt,
			&_ag.SaleEndAt,
			&_ag.ServiceStartAt,
			&_ag.TechniqueFeeRatio,
			&_ag.DescriptionsStr,
			&_ag.GoodBanner,
			&_ag.DisplayNamesStr,
			&_ag.EnablePurchase,
			&_ag.EnableProductPage,
			&_ag.CancelMode,
			&_ag.DisplayColorsStr,
			&_ag.CancelableBeforeStart,
			&_ag.ProductPage,
			&_ag.EnableSetCommission,
			&_ag.PostersStr,
			&_ag.MinOrderAmount,
			&_ag.MaxOrderAmount,
			&_ag.MaxUserAmount,
			&_ag.MinOrderDurationSeconds,
			&_ag.MaxOrderDurationSeconds,
			&_ag.PackageWithRequireds,
			&_ag.GoodType,
			&_ag.CreatedAt,
			&_ag.UpdatedAt,
		); err != nil {
			return err
		}
		_appGoods = append(_appGoods, _ag)
	}
	for _, appGood := range _appGoods {
		_ = json.Unmarshal([]byte(appGood.DescriptionsStr), &appGood.Descriptions)
		_ = json.Unmarshal([]byte(appGood.DisplayColorsStr), &appGood.DisplayColors)
		_ = json.Unmarshal([]byte(appGood.DisplayNamesStr), &appGood.DisplayNames)
		_ = json.Unmarshal([]byte(appGood.PostersStr), &appGood.Posters)
		if err := db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
			exist, err := tx.
				AppGoodBase.
				Query().
				Where(
					entappgoodbase.EntID(appGood.EntID),
				).
				Exist(_ctx)
			if err != nil {
				return err
			}
			if exist {
				return nil
			}
			if _, err := tx.
				AppGoodBase.
				Create().
				SetEntID(appGood.EntID).
				SetAppID(appGood.AppID).
				SetGoodID(appGood.GoodID).
				SetPurchasable(appGood.EnablePurchase).
				SetEnableProductPage(appGood.EnableProductPage).
				SetProductPage(appGood.ProductPage).
				SetOnline(appGood.Online).
				SetVisible(appGood.Visible).
				SetName(appGood.GoodName).
				SetDisplayIndex(appGood.DisplayIndex).
				SetBanner(appGood.GoodBanner).
				Save(_ctx); err != nil {
				return err
			}
			switch appGood.GoodType {
			case types.GoodType_MachineRental.String():
			case types.GoodType_MachineCustody.String():
			case types.GoodType_LegacyPowerRental.String():
				if _, err := tx.
					AppLegacyPowerRental.
					Create().
					SetAppGoodID(appGood.EntID).
					SetTechniqueFeeRatio(appGood.TechniqueFeeRatio).
					Save(_ctx); err != nil {
					return err
				}
				fallthrough
			case types.GoodType_PowerRental.String():
				stm := tx.
					AppPowerRental.
					Create().
					SetAppGoodID(appGood.EntID).
					SetServiceStartAt(appGood.ServiceStartAt).
					SetCancelMode(appGood.CancelMode).
					SetCancelableBeforeStartSeconds(appGood.CancelableBeforeStart).
					SetEnableSetCommission(appGood.EnableSetCommission).
					SetMinOrderAmount(appGood.MinOrderAmount).
					SetMaxOrderAmount(appGood.MaxOrderAmount).
					SetMaxUserAmount(appGood.MaxUserAmount).
					SetMinOrderDurationSeconds(appGood.MinOrderDurationSeconds).
					SetMaxOrderDurationSeconds(appGood.MaxOrderDurationSeconds)
				if appGood.PackagePrice.Cmp(decimal.NewFromInt(0)) > 0 {
					stm.
						SetUnitPrice(appGood.PackagePrice).
						SetFixedDuration(true)
				} else {
					stm.SetUnitPrice(appGood.UnitPrice)
				}
				if _, err := stm.SetSaleStartAt(appGood.SaleStartAt).
					SetSaleEndAt(appGood.SaleEndAt).
					SetSaleMode(types.GoodSaleMode_GoodSaleModeMainnetSpot.String()).
					SetPackageWithRequireds(appGood.PackageWithRequireds).
					Save(_ctx); err != nil {
					return err
				}
			case types.GoodType_TechniqueServiceFee.String():
				fallthrough //nolint
			case types.GoodType_ElectricityFee.String():
				if _, err := tx.
					AppFee.
					Create().
					SetAppGoodID(appGood.EntID).
					SetUnitValue(appGood.UnitPrice).
					SetMinOrderDurationSeconds(appGood.MinOrderDurationSeconds).
					Save(_ctx); err != nil {
					return err
				}
			}
			for i, description := range appGood.Descriptions {
				if _, err := tx.
					AppGoodDescription.
					Create().
					SetAppGoodID(appGood.EntID).
					SetDescription(description).
					SetIndex(uint8(i)).
					Save(_ctx); err != nil {
					return err
				}
			}
			for i, name := range appGood.DisplayNames {
				if _, err := tx.
					AppGoodDisplayName.
					Create().
					SetAppGoodID(appGood.EntID).
					SetName(name).
					SetIndex(uint8(i)).
					Save(_ctx); err != nil {
					return err
				}
			}
			for i, color := range appGood.DisplayColors {
				if _, err := tx.
					AppGoodDisplayColor.
					Create().
					SetAppGoodID(appGood.EntID).
					SetColor(color).
					SetIndex(uint8(i)).
					Save(_ctx); err != nil {
					return err
				}
			}
			for i, poster := range appGood.Posters {
				if _, err := tx.
					AppGoodPoster.
					Create().
					SetAppGoodID(appGood.EntID).
					SetPoster(poster).
					SetIndex(uint8(i)).
					Save(_ctx); err != nil {
					return err
				}
			}
			return nil
		}); err != nil {
			return err
		}
	}
	return nil
}

func Migrate(ctx context.Context) error {
	var err error
	var conn *sql.DB

	logger.Sugar().Warnw("Migrate", "Start", "...")
	err = redis2.TryLock(lockKey(), 0)
	if err != nil {
		return err
	}
	defer func(err *error) {
		_ = redis2.Unlock(lockKey())
		logger.Sugar().Warnw("Migrate", "Done", "...", "error", *err)
	}(&err)

	conn, err = open(servicename.ServiceDomain)
	if err != nil {
		return err
	}
	defer func() {
		if err := conn.Close(); err != nil {
			logger.Sugar().Errorw("Close", "Error", err)
		}
	}()

	if err := migrateGood(ctx, conn); err != nil {
		return err
	}
	if err := migrateAppGood(ctx, conn); err != nil {
		return err
	}

	return nil
}
