//nolint:dupl
package powerrental

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"

	appgoodbasecrud "github.com/NpoolPlatform/good-middleware/pkg/crud/app/good/goodbase"
	mininggoodstockcrud "github.com/NpoolPlatform/good-middleware/pkg/crud/good/stock/mining"
	"github.com/NpoolPlatform/good-middleware/pkg/db"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent"
	entappgoodbase "github.com/NpoolPlatform/good-middleware/pkg/db/ent/appgoodbase"
	entapplegacypowerrental "github.com/NpoolPlatform/good-middleware/pkg/db/ent/applegacypowerrental"
	entapppowerrental "github.com/NpoolPlatform/good-middleware/pkg/db/ent/apppowerrental"
	entdevicetype "github.com/NpoolPlatform/good-middleware/pkg/db/ent/deviceinfo"
	entmanufacturer "github.com/NpoolPlatform/good-middleware/pkg/db/ent/devicemanufacturer"
	entextrainfo "github.com/NpoolPlatform/good-middleware/pkg/db/ent/extrainfo"
	entgoodbase "github.com/NpoolPlatform/good-middleware/pkg/db/ent/goodbase"
	entgoodreward "github.com/NpoolPlatform/good-middleware/pkg/db/ent/goodreward"
	entmininggoodstock "github.com/NpoolPlatform/good-middleware/pkg/db/ent/mininggoodstock"
	entpowerrental "github.com/NpoolPlatform/good-middleware/pkg/db/ent/powerrental"
	entstock "github.com/NpoolPlatform/good-middleware/pkg/db/ent/stock"
	entvendorbrand "github.com/NpoolPlatform/good-middleware/pkg/db/ent/vendorbrand"
	entvendorlocation "github.com/NpoolPlatform/good-middleware/pkg/db/ent/vendorlocation"
	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	types "github.com/NpoolPlatform/message/npool/basetypes/good/v1"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/app/powerrental"
	stockmwpb "github.com/NpoolPlatform/message/npool/good/mw/v1/good/stock"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type queryHandler struct {
	*Handler
	stmSelect        *ent.AppGoodBaseSelect
	stmCount         *ent.AppGoodBaseSelect
	infos            []*npool.PowerRental
	miningGoodStocks []*stockmwpb.MiningGoodStockInfo
	total            uint32
}

func (h *queryHandler) selectAppGoodBase(stm *ent.AppGoodBaseQuery) *ent.AppGoodBaseSelect {
	return stm.Select(entappgoodbase.FieldID)
}

func (h *queryHandler) queryAppGoodBase(cli *ent.Client) error {
	if h.AppGoodID == nil {
		return fmt.Errorf("invalid appgoodid")
	}
	h.stmSelect = h.selectAppGoodBase(
		cli.AppGoodBase.
			Query().
			Where(
				entappgoodbase.DeletedAt(0),
				entappgoodbase.EntID(*h.AppGoodID),
			),
	)
	return nil
}

func (h *queryHandler) queryAppGoodBases(cli *ent.Client) (*ent.AppGoodBaseSelect, error) {
	stm, err := appgoodbasecrud.SetQueryConds(cli.AppGoodBase.Query(), h.AppGoodBaseConds)
	if err != nil {
		return nil, err
	}
	return h.selectAppGoodBase(stm), nil
}

func (h *queryHandler) queryJoinMyself(s *sql.Selector) {
	t1 := sql.Table(entappgoodbase.Table)
	s.Join(t1).
		On(
			s.C(entappgoodbase.FieldID),
			t1.C(entappgoodbase.FieldID),
		).
		AppendSelect(
			t1.C(entappgoodbase.FieldAppID),
			t1.C(entappgoodbase.FieldGoodID),
			sql.As(t1.C(entappgoodbase.FieldPurchasable), "app_good_purchasable"),
			sql.As(t1.C(entappgoodbase.FieldOnline), "app_good_online"),
			t1.C(entappgoodbase.FieldEnableProductPage),
			t1.C(entappgoodbase.FieldProductPage),
			t1.C(entappgoodbase.FieldVisible),
			sql.As(t1.C(entappgoodbase.FieldName), "app_good_name"),
			t1.C(entappgoodbase.FieldDisplayIndex),
			t1.C(entappgoodbase.FieldBanner),
			t1.C(entappgoodbase.FieldCreatedAt),
			t1.C(entappgoodbase.FieldUpdatedAt),
		)
}

func (h *queryHandler) queryJoinGoodBase(s *sql.Selector) {
	t1 := sql.Table(entgoodbase.Table)
	s.Join(t1).
		On(
			s.C(entappgoodbase.FieldGoodID),
			t1.C(entgoodbase.FieldEntID),
		).
		OnP(
			sql.Or(
				sql.EQ(t1.C(entgoodbase.FieldGoodType), types.GoodType_PowerRental.String()),
				sql.EQ(t1.C(entgoodbase.FieldGoodType), types.GoodType_LegacyPowerRental.String()),
			),
		).
		AppendSelect(
			t1.C(entgoodbase.FieldGoodType),
			t1.C(entgoodbase.FieldBenefitType),
			sql.As(t1.C(entgoodbase.FieldName), "good_name"),
			t1.C(entgoodbase.FieldServiceStartAt),
			t1.C(entgoodbase.FieldStartMode),
			t1.C(entgoodbase.FieldTestOnly),
			t1.C(entgoodbase.FieldBenefitIntervalHours),
			sql.As(t1.C(entgoodbase.FieldPurchasable), "good_purchasable"),
			sql.As(t1.C(entgoodbase.FieldOnline), "good_online"),
		)
}

func (h *queryHandler) queryJoinGoodReward(s *sql.Selector) {
	t := sql.Table(entgoodreward.Table)
	s.LeftJoin(t).
		On(
			s.C(entappgoodbase.FieldGoodID),
			t.C(entgoodreward.FieldGoodID),
		).
		OnP(
			sql.EQ(t.C(entextrainfo.FieldDeletedAt), 0),
		).
		AppendSelect(
			t.C(entgoodreward.FieldLastRewardAt),
			t.C(entgoodreward.FieldLastRewardAmount),
			t.C(entgoodreward.FieldTotalRewardAmount),
			t.C(entgoodreward.FieldLastUnitRewardAmount),
		)
}

func (h *queryHandler) queryJoinExtraInfo(s *sql.Selector) {
	t := sql.Table(entextrainfo.Table)
	s.LeftJoin(t).
		On(
			s.C(entappgoodbase.FieldEntID),
			t.C(entextrainfo.FieldAppGoodID),
		).
		OnP(
			sql.EQ(t.C(entextrainfo.FieldDeletedAt), 0),
		).
		AppendSelect(
			t.C(entextrainfo.FieldLikes),
			t.C(entextrainfo.FieldDislikes),
			t.C(entextrainfo.FieldScoreCount),
			t.C(entextrainfo.FieldRecommendCount),
			t.C(entextrainfo.FieldCommentCount),
			t.C(entextrainfo.FieldScore),
		)
}

func (h *queryHandler) queryJoinGoodStock(s *sql.Selector) {
	t1 := sql.Table(entstock.Table)
	s.Join(t1).
		On(
			s.C(entappgoodbase.FieldGoodID),
			t1.C(entstock.FieldGoodID),
		).
		AppendSelect(
			sql.As(t1.C(entstock.FieldEntID), "good_stock_id"),
			sql.As(t1.C(entstock.FieldTotal), "good_total"),
			sql.As(t1.C(entstock.FieldSpotQuantity), "good_spot_quantity"),
		)
}

func (h *queryHandler) queryJoinAppPowerRental(s *sql.Selector) {
	t1 := sql.Table(entapppowerrental.Table)
	s.Join(t1).
		On(
			s.C(entappgoodbase.FieldEntID),
			t1.C(entapppowerrental.FieldAppGoodID),
		).
		AppendSelect(
			t1.C(entapppowerrental.FieldID),
			t1.C(entapppowerrental.FieldEntID),
			t1.C(entapppowerrental.FieldAppGoodID),
			t1.C(entapppowerrental.FieldCancelMode),
			t1.C(entapppowerrental.FieldCancelableBeforeStartSeconds),
			t1.C(entapppowerrental.FieldEnableSetCommission),
			t1.C(entapppowerrental.FieldMinOrderAmount),
			t1.C(entapppowerrental.FieldMaxOrderAmount),
			t1.C(entapppowerrental.FieldMaxUserAmount),
			t1.C(entapppowerrental.FieldMinOrderDuration),
			t1.C(entapppowerrental.FieldMaxOrderDuration),
			t1.C(entapppowerrental.FieldSaleStartAt),
			t1.C(entapppowerrental.FieldSaleEndAt),
			t1.C(entapppowerrental.FieldSaleMode),
			t1.C(entapppowerrental.FieldFixedDuration),
			t1.C(entapppowerrental.FieldPackageWithRequireds),
		)
}

func (h *queryHandler) queryJoinAppLegacyPowerRental(s *sql.Selector) {
	t1 := sql.Table(entapplegacypowerrental.Table)
	s.LeftJoin(t1).
		On(
			s.C(entappgoodbase.FieldEntID),
			t1.C(entapplegacypowerrental.FieldAppGoodID),
		).
		AppendSelect(
			t1.C(entapplegacypowerrental.FieldTechniqueFeeRatio),
		)
}

func (h *queryHandler) queryJoinPowerRental(s *sql.Selector) error {
	t1 := sql.Table(entpowerrental.Table)
	t2 := sql.Table(entdevicetype.Table)
	t3 := sql.Table(entmanufacturer.Table)
	t4 := sql.Table(entvendorlocation.Table)
	t5 := sql.Table(entvendorbrand.Table)

	s.Join(t1).
		On(
			s.C(entappgoodbase.FieldGoodID),
			t1.C(entpowerrental.FieldGoodID),
		).
		OnP(
			sql.EQ(t1.C(entpowerrental.FieldDeletedAt), 0),
		)
	if h.PowerRentalConds != nil && h.PowerRentalConds.ID != nil {
		u, ok := h.PowerRentalConds.ID.Val.(uint32)
		if !ok {
			return fmt.Errorf("invalid id")
		}
		s.OnP(sql.EQ(t1.C(entpowerrental.FieldID), u))
	}
	if h.PowerRentalConds != nil && h.PowerRentalConds.IDs != nil {
		ids, ok := h.PowerRentalConds.IDs.Val.([]uint32)
		if !ok {
			return fmt.Errorf("invalid ids")
		}
		s.OnP(sql.In(t1.C(entpowerrental.FieldID), ids))
	}
	if h.PowerRentalConds != nil && h.PowerRentalConds.EntID != nil {
		uid, ok := h.PowerRentalConds.EntID.Val.(uuid.UUID)
		if !ok {
			return fmt.Errorf("invalid entid")
		}
		s.OnP(sql.EQ(t1.C(entpowerrental.FieldEntID), uid))
	}
	if h.PowerRentalConds != nil && h.PowerRentalConds.EntIDs != nil {
		uids, ok := h.PowerRentalConds.EntIDs.Val.([]uuid.UUID)
		if !ok {
			return fmt.Errorf("invalid entids")
		}
		s.OnP(sql.In(t1.C(entpowerrental.FieldEntID), uids))
	}
	s.Join(t2).
		On(
			t1.C(entpowerrental.FieldDeviceTypeID),
			t2.C(entdevicetype.FieldEntID),
		).
		Join(t3).
		On(
			t2.C(entdevicetype.FieldManufacturerID),
			t3.C(entmanufacturer.FieldEntID),
		).
		Join(t4).
		On(
			t1.C(entpowerrental.FieldVendorLocationID),
			t4.C(entvendorlocation.FieldEntID),
		).
		Join(t5).
		On(
			t4.C(entvendorlocation.FieldBrandID),
			t5.C(entvendorbrand.FieldEntID),
		)
	s.AppendSelect(
		t1.C(entpowerrental.FieldGoodID),
		t1.C(entpowerrental.FieldDeviceTypeID),
		t1.C(entpowerrental.FieldVendorLocationID),
		t1.C(entpowerrental.FieldUnitPrice),
		t1.C(entpowerrental.FieldQuantityUnit),
		t1.C(entpowerrental.FieldQuantityUnitAmount),
		t1.C(entpowerrental.FieldDeliveryAt),
		t1.C(entpowerrental.FieldUnitLockDeposit),
		t1.C(entpowerrental.FieldDurationType),
		t1.C(entpowerrental.FieldUnitLockDeposit),
		t1.C(entpowerrental.FieldStockMode),

		sql.As(t2.C(entdevicetype.FieldType), "device_type"),
		sql.As(t2.C(entdevicetype.FieldPowerConsumption), "device_power_consumption"),
		sql.As(t2.C(entdevicetype.FieldShipmentAt), "device_shipment_at"),
		sql.As(t3.C(entmanufacturer.FieldName), "device_manufacturer_name"),
		sql.As(t3.C(entmanufacturer.FieldLogo), "device_manufacturer_logo"),

		sql.As(t4.C(entvendorlocation.FieldCountry), "vendor_country"),
		sql.As(t4.C(entvendorlocation.FieldProvince), "vendor_province"),
		sql.As(t5.C(entvendorbrand.FieldName), "vendor_brand"),
		sql.As(t5.C(entvendorbrand.FieldLogo), "vendor_logo"),
	)
	return nil
}

func (h *queryHandler) queryJoin() error {
	var err error
	h.stmSelect.Modify(func(s *sql.Selector) {
		h.queryJoinMyself(s)
		h.queryJoinGoodBase(s)
		h.queryJoinGoodStock(s)
		h.queryJoinGoodReward(s)
		h.queryJoinExtraInfo(s)
		h.queryJoinAppPowerRental(s)
		h.queryJoinAppLegacyPowerRental(s)
		err = h.queryJoinPowerRental(s)
	})
	if err != nil {
		return err
	}
	if h.stmCount == nil {
		return nil
	}
	h.stmCount.Modify(func(s *sql.Selector) {
		h.queryJoinGoodBase(s)
		h.queryJoinGoodStock(s)
		h.queryJoinGoodReward(s)
		h.queryJoinExtraInfo(s)
		h.queryJoinAppPowerRental(s)
		h.queryJoinAppLegacyPowerRental(s)
		err = h.queryJoinPowerRental(s)
	})
	if err != nil {
		return err
	}
	return nil
}

func (h *queryHandler) scan(ctx context.Context) error {
	return h.stmSelect.Scan(ctx, &h.infos)
}

func (h *queryHandler) getMiningGoodStocks(ctx context.Context, cli *ent.Client) error {
	goodStockIDs := func() (uids []uuid.UUID) {
		for _, info := range h.infos {
			if info.StockModeStr != types.GoodStockMode_GoodStockByMiningPool.String() {
				continue
			}
			uids = append(uids, uuid.MustParse(info.GoodStockID))
		}
		return
	}()

	stm, err := mininggoodstockcrud.SetQueryConds(
		cli.MiningGoodStock.Query(),
		&mininggoodstockcrud.Conds{
			GoodStockIDs: &cruder.Cond{Op: cruder.IN, Val: goodStockIDs},
		},
	)
	if err != nil {
		return err
	}

	return stm.Select(
		entmininggoodstock.FieldGoodStockID,
		entmininggoodstock.FieldMiningPoolID,
		entmininggoodstock.FieldPoolGoodUserID,
		entmininggoodstock.FieldTotal,
		entmininggoodstock.FieldSpotQuantity,
	).Scan(ctx, &h.miningGoodStocks)
}

//nolint:funlen,gocyclo
func (h *queryHandler) formalize() {
	goodMiningStocks := map[string][]*stockmwpb.MiningGoodStockInfo{}

	for _, stock := range h.miningGoodStocks {
		stock.Total = func() string { amount, _ := decimal.NewFromString(stock.Total); return amount.String() }()
		stock.SpotQuantity = func() string { amount, _ := decimal.NewFromString(stock.SpotQuantity); return amount.String() }()
		goodMiningStocks[stock.GoodStockID] = append(goodMiningStocks[stock.GoodStockID], stock)
	}
	for _, info := range h.infos {
		info.UnitPrice = func() string { amount, _ := decimal.NewFromString(info.UnitPrice); return amount.String() }()
		info.QuantityUnitAmount = func() string { amount, _ := decimal.NewFromString(info.QuantityUnitAmount); return amount.String() }()
		info.UnitLockDeposit = func() string { amount, _ := decimal.NewFromString(info.UnitLockDeposit); return amount.String() }()
		info.MinOrderAmount = func() string { amount, _ := decimal.NewFromString(info.MinOrderAmount); return amount.String() }()
		info.MaxOrderAmount = func() string { amount, _ := decimal.NewFromString(info.MaxOrderAmount); return amount.String() }()
		info.MaxUserAmount = func() string { amount, _ := decimal.NewFromString(info.MaxUserAmount); return amount.String() }()
		info.TechniqueFeeRatio = func() string { amount, _ := decimal.NewFromString(info.TechniqueFeeRatio); return amount.String() }()
		info.GoodTotal = func() string { amount, _ := decimal.NewFromString(info.GoodTotal); return amount.String() }()
		info.GoodSpotQuantity = func() string { amount, _ := decimal.NewFromString(info.GoodSpotQuantity); return amount.String() }()
		info.LastRewardAmount = func() string { amount, _ := decimal.NewFromString(info.LastRewardAmount); return amount.String() }()
		info.LastUnitRewardAmount = func() string { amount, _ := decimal.NewFromString(info.LastUnitRewardAmount); return amount.String() }()
		info.TotalRewardAmount = func() string { amount, _ := decimal.NewFromString(info.TotalRewardAmount); return amount.String() }()
		info.GoodType = types.GoodType(types.GoodType_value[info.GoodTypeStr])
		info.CancelMode = types.CancelMode(types.CancelMode_value[info.CancelModeStr])
		info.SaleMode = types.GoodSaleMode(types.GoodSaleMode_value[info.SaleModeStr])
		info.BenefitType = types.BenefitType(types.BenefitType_value[info.BenefitTypeStr])
		info.DurationType = types.GoodDurationType(types.GoodDurationType_value[info.DurationTypeStr])
		info.StartMode = types.GoodStartMode(types.GoodStartMode_value[info.StartModeStr])
		info.StockMode = types.GoodStockMode(types.GoodStockMode_value[info.StockModeStr])
		info.MiningGoodStocks = goodMiningStocks[info.GoodStockID]
	}
}

func (h *Handler) GetPowerRental(ctx context.Context) (*npool.PowerRental, error) {
	handler := &queryHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryAppGoodBase(cli); err != nil {
			return err
		}
		if err := handler.queryJoin(); err != nil {
			return err
		}
		handler.stmSelect.
			Offset(0).
			Limit(2)
		if err := handler.scan(_ctx); err != nil {
			return err
		}
		return handler.getMiningGoodStocks(_ctx, cli)
	})
	if err != nil {
		return nil, err
	}
	if len(handler.infos) == 0 {
		return nil, nil
	}
	if len(handler.infos) > 1 {
		return nil, fmt.Errorf("too many records")
	}

	handler.formalize()
	return handler.infos[0], nil
}

func (h *Handler) GetPowerRentals(ctx context.Context) ([]*npool.PowerRental, uint32, error) {
	handler := &queryHandler{
		Handler: h,
	}

	var err error

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		handler.stmSelect, err = handler.queryAppGoodBases(cli)
		if err != nil {
			return err
		}
		handler.stmCount, err = handler.queryAppGoodBases(cli)
		if err != nil {
			return err
		}

		if err := handler.queryJoin(); err != nil {
			return err
		}

		total, err := handler.stmCount.Count(_ctx)
		if err != nil {
			return err
		}
		handler.total = uint32(total)

		handler.stmSelect.
			Offset(int(handler.Offset)).
			Limit(int(handler.Limit))
		if err := handler.scan(_ctx); err != nil {
			return err
		}
		return handler.getMiningGoodStocks(_ctx, cli)
	})
	if err != nil {
		return nil, 0, err
	}

	handler.formalize()

	return handler.infos, handler.total, nil
}
