package powerrental

import (
	"entgo.io/ent/dialect/sql"

	logger "github.com/NpoolPlatform/go-service-framework/pkg/logger"
	wlog "github.com/NpoolPlatform/go-service-framework/pkg/wlog"
	goodbasecrud "github.com/NpoolPlatform/good-middleware/pkg/crud/good/goodbase"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent"
	entdevicetype "github.com/NpoolPlatform/good-middleware/pkg/db/ent/deviceinfo"
	entmanufacturer "github.com/NpoolPlatform/good-middleware/pkg/db/ent/devicemanufacturer"
	entgoodbase "github.com/NpoolPlatform/good-middleware/pkg/db/ent/goodbase"
	entgoodcoin "github.com/NpoolPlatform/good-middleware/pkg/db/ent/goodcoin"
	entgoodreward "github.com/NpoolPlatform/good-middleware/pkg/db/ent/goodreward"
	entpowerrental "github.com/NpoolPlatform/good-middleware/pkg/db/ent/powerrental"
	entstock "github.com/NpoolPlatform/good-middleware/pkg/db/ent/stock"
	entvendorbrand "github.com/NpoolPlatform/good-middleware/pkg/db/ent/vendorbrand"
	entvendorlocation "github.com/NpoolPlatform/good-middleware/pkg/db/ent/vendorlocation"
	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	types "github.com/NpoolPlatform/message/npool/basetypes/good/v1"

	"github.com/google/uuid"
)

type baseQueryHandler struct {
	*Handler
	stmSelect *ent.GoodBaseSelect
}

func (h *baseQueryHandler) selectGoodBase(stm *ent.GoodBaseQuery) *ent.GoodBaseSelect {
	return stm.Select(entgoodbase.FieldID)
}

func (h *baseQueryHandler) queryGoodBase(cli *ent.Client) error {
	if h.GoodID == nil {
		return wlog.Errorf("invalid goodid")
	}
	h.stmSelect = h.selectGoodBase(
		cli.GoodBase.
			Query().
			Where(
				entgoodbase.DeletedAt(0),
				entgoodbase.EntID(*h.GoodID),
				entgoodbase.Or(
					entgoodbase.GoodType(types.GoodType_PowerRental.String()),
					entgoodbase.GoodType(types.GoodType_LegacyPowerRental.String()),
				),
			),
	)
	return nil
}

func (h *baseQueryHandler) queryGoodBases(cli *ent.Client) (*ent.GoodBaseSelect, error) {
	stm, err := goodbasecrud.SetQueryConds(cli.GoodBase.Query(), h.GoodBaseConds)
	if err != nil {
		return nil, err
	}
	stm.Where(
		entgoodbase.Or(
			entgoodbase.GoodType(types.GoodType_PowerRental.String()),
			entgoodbase.GoodType(types.GoodType_LegacyPowerRental.String()),
		),
	)
	return h.selectGoodBase(stm), nil
}

func (h *baseQueryHandler) queryJoinMyself(s *sql.Selector) {
	t1 := sql.Table(entgoodbase.Table)
	s.LeftJoin(t1).
		On(
			s.C(entgoodbase.FieldID),
			t1.C(entgoodbase.FieldID),
		).
		AppendSelect(
			t1.C(entgoodbase.FieldGoodType),
			t1.C(entgoodbase.FieldBenefitType),
			t1.C(entgoodbase.FieldName),
			t1.C(entgoodbase.FieldServiceStartAt),
			t1.C(entgoodbase.FieldStartMode),
			t1.C(entgoodbase.FieldTestOnly),
			t1.C(entgoodbase.FieldBenefitIntervalHours),
			t1.C(entgoodbase.FieldPurchasable),
			t1.C(entgoodbase.FieldOnline),
			t1.C(entgoodbase.FieldCreatedAt),
			t1.C(entgoodbase.FieldUpdatedAt),
		)
}

func (h *baseQueryHandler) queryJoinPowerRental(s *sql.Selector) error {
	t1 := sql.Table(entpowerrental.Table)
	t2 := sql.Table(entdevicetype.Table)
	t3 := sql.Table(entmanufacturer.Table)
	t4 := sql.Table(entvendorlocation.Table)
	t5 := sql.Table(entvendorbrand.Table)

	s.Join(t1).
		On(
			s.C(entgoodbase.FieldEntID),
			t1.C(entpowerrental.FieldGoodID),
		).
		OnP(
			sql.EQ(t1.C(entpowerrental.FieldDeletedAt), 0),
		)
	if h.PowerRentalConds.ID != nil {
		u, ok := h.PowerRentalConds.ID.Val.(uint32)
		if !ok {
			return wlog.Errorf("invalid id")
		}
		s.OnP(sql.EQ(t1.C(entpowerrental.FieldID), u))
	}
	if h.PowerRentalConds.IDs != nil {
		ids, ok := h.PowerRentalConds.IDs.Val.([]uint32)
		if !ok {
			return wlog.Errorf("invalid ids")
		}
		s.OnP(sql.In(t1.C(entpowerrental.FieldID), func() (_ids []interface{}) {
			for _, id := range ids {
				_ids = append(_ids, interface{}(id))
			}
			return
		}()...))
	}
	if h.PowerRentalConds.EntID != nil {
		uid, ok := h.PowerRentalConds.EntID.Val.(uuid.UUID)
		if !ok {
			return wlog.Errorf("invalid entid")
		}
		s.OnP(sql.EQ(t1.C(entpowerrental.FieldEntID), uid))
	}
	if h.PowerRentalConds.EntIDs != nil {
		uids, ok := h.PowerRentalConds.EntIDs.Val.([]uuid.UUID)
		if !ok {
			return wlog.Errorf("invalid entids")
		}
		s.OnP(sql.In(t1.C(entpowerrental.FieldEntID), func() (_uids []interface{}) {
			for _, uid := range uids {
				_uids = append(_uids, interface{}(uid))
			}
			return
		}()...))
	}
	s.LeftJoin(t2).
		On(
			t1.C(entpowerrental.FieldDeviceTypeID),
			t2.C(entdevicetype.FieldEntID),
		).
		LeftJoin(t3).
		On(
			t2.C(entdevicetype.FieldManufacturerID),
			t3.C(entmanufacturer.FieldEntID),
		).
		LeftJoin(t4).
		On(
			t1.C(entpowerrental.FieldVendorLocationID),
			t4.C(entvendorlocation.FieldEntID),
		).
		LeftJoin(t5).
		On(
			t4.C(entvendorlocation.FieldBrandID),
			t5.C(entvendorbrand.FieldEntID),
		)
	s.AppendSelect(
		t1.C(entpowerrental.FieldID),
		t1.C(entpowerrental.FieldEntID),
		t1.C(entpowerrental.FieldGoodID),
		t1.C(entpowerrental.FieldDeviceTypeID),
		t1.C(entpowerrental.FieldVendorLocationID),
		t1.C(entpowerrental.FieldUnitPrice),
		t1.C(entpowerrental.FieldQuantityUnit),
		t1.C(entpowerrental.FieldQuantityUnitAmount),
		t1.C(entpowerrental.FieldDeliveryAt),
		t1.C(entpowerrental.FieldUnitLockDeposit),
		t1.C(entpowerrental.FieldDurationDisplayType),
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

func (h *baseQueryHandler) queryJoinReward(s *sql.Selector) {
	t := sql.Table(entgoodreward.Table)
	s.LeftJoin(t).
		On(
			s.C(entgoodbase.FieldEntID),
			t.C(entgoodreward.FieldGoodID),
		).
		OnP(
			sql.EQ(t.C(entgoodreward.FieldDeletedAt), 0),
		)
	if h.RewardConds.RewardState != nil {
		s.OnP(
			sql.EQ(t.C(entgoodreward.FieldRewardState), h.RewardConds.RewardState.Val.(types.BenefitState).String()),
		)
	}
	if h.RewardConds.RewardAt != nil {
		switch h.RewardConds.RewardAt.Op {
		case cruder.EQ:
			s.OnP(sql.EQ(t.C(entgoodreward.FieldLastRewardAt), h.RewardConds.RewardAt.Val))
		case cruder.NEQ:
			s.OnP(sql.NEQ(t.C(entgoodreward.FieldLastRewardAt), h.RewardConds.RewardAt.Val))
		}
	}
	s.AppendSelect(
		t.C(entgoodreward.FieldRewardState),
		t.C(entgoodreward.FieldLastRewardAt),
		t.C(entgoodreward.FieldRewardTid),
		t.C(entgoodreward.FieldNextRewardStartAmount),
		t.C(entgoodreward.FieldLastRewardAmount),
		t.C(entgoodreward.FieldLastUnitRewardAmount),
		t.C(entgoodreward.FieldTotalRewardAmount),
	)
}

func (h *baseQueryHandler) queryJoinGoodCoin(s *sql.Selector) error {
	t := sql.Table(entgoodcoin.Table)
	s.LeftJoin(t).
		On(
			s.C(entgoodbase.FieldEntID),
			t.C(entgoodcoin.FieldGoodID),
		).
		Distinct()
	if h.GoodCoinConds.CoinTypeID != nil {
		id, ok := h.GoodCoinConds.CoinTypeID.Val.(uuid.UUID)
		if !ok {
			return wlog.Errorf("invalid cointypeid")
		}
		s.OnP(
			sql.EQ(t.C(entgoodcoin.FieldCoinTypeID), id),
		)
		s.Where(
			sql.EQ(t.C(entgoodcoin.FieldCoinTypeID), id),
		)
	}
	if h.GoodCoinConds.CoinTypeIDs != nil {
		uids, ok := h.GoodCoinConds.CoinTypeIDs.Val.([]uuid.UUID)
		if !ok {
			return wlog.Errorf("invalid cointypeids")
		}
		_uids := func() (_uids []interface{}) {
			for _, uid := range uids {
				_uids = append(_uids, interface{}(uid))
			}
			return
		}()
		s.OnP(
			sql.In(t.C(entgoodcoin.FieldCoinTypeID), _uids...),
		)
		s.Where(
			sql.In(t.C(entgoodcoin.FieldCoinTypeID), _uids...),
		)
	}
	return nil
}

func (h *baseQueryHandler) queryJoinStock(s *sql.Selector) {
	t1 := sql.Table(entstock.Table)
	s.LeftJoin(t1).
		On(
			s.C(entgoodbase.FieldEntID),
			t1.C(entstock.FieldGoodID),
		).
		AppendSelect(
			sql.As(t1.C(entstock.FieldEntID), "good_stock_id"),
			sql.As(t1.C(entstock.FieldTotal), "good_total"),
			sql.As(t1.C(entstock.FieldSpotQuantity), "good_spot_quantity"),
			sql.As(t1.C(entstock.FieldLocked), "good_locked"),
			sql.As(t1.C(entstock.FieldInService), "good_in_service"),
			sql.As(t1.C(entstock.FieldWaitStart), "good_wait_start"),
			sql.As(t1.C(entstock.FieldSold), "good_sold"),
			sql.As(t1.C(entstock.FieldAppReserved), "good_app_reserved"),
		)
}

func (h *baseQueryHandler) queryJoin() {
	h.stmSelect.Modify(func(s *sql.Selector) {
		h.queryJoinMyself(s)
		h.queryJoinStock(s)
		h.queryJoinReward(s)
		if err := h.queryJoinGoodCoin(s); err != nil {
			logger.Sugar().Errorw("queryJoinGoodCoin", "Error", err)
		}
		if err := h.queryJoinPowerRental(s); err != nil {
			logger.Sugar().Errorw("queryJoinPowerRental", "Error", err)
		}
	})
}
