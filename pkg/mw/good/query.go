package good

import (
	"context"
	"encoding/json"
	"fmt"

	"entgo.io/ent/dialect/sql"

	goodcrud "github.com/NpoolPlatform/good-middleware/pkg/crud/good"
	"github.com/NpoolPlatform/good-middleware/pkg/db"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent"
	entdevice "github.com/NpoolPlatform/good-middleware/pkg/db/ent/deviceinfo"
	entextrainfo "github.com/NpoolPlatform/good-middleware/pkg/db/ent/extrainfo"
	entgood "github.com/NpoolPlatform/good-middleware/pkg/db/ent/good"
	entgoodreward "github.com/NpoolPlatform/good-middleware/pkg/db/ent/goodreward"
	entstock "github.com/NpoolPlatform/good-middleware/pkg/db/ent/stock"
	entvendorbrand "github.com/NpoolPlatform/good-middleware/pkg/db/ent/vendorbrand"
	entvendorlocation "github.com/NpoolPlatform/good-middleware/pkg/db/ent/vendorlocation"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	types "github.com/NpoolPlatform/message/npool/basetypes/good/v1"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/good"

	"github.com/shopspring/decimal"
)

type queryHandler struct {
	*Handler
	stmSelect *ent.GoodSelect
	stmCount  *ent.GoodSelect
	infos     []*npool.Good
	total     uint32
}

func (h *queryHandler) selectGood(stm *ent.GoodQuery) *ent.GoodSelect {
	return stm.Select(entgood.FieldID)
}

func (h *queryHandler) queryGood(cli *ent.Client) error {
	if h.ID == nil && h.EntID == nil {
		return fmt.Errorf("invalid id")
	}
	stm := cli.Good.Query().Where(entgood.DeletedAt(0))
	if h.ID != nil {
		stm.Where(entgood.ID(*h.ID))
	}
	if h.EntID != nil {
		stm.Where(entgood.EntID(*h.EntID))
	}
	h.stmSelect = h.selectGood(stm)
	return nil
}

func (h *queryHandler) queryGoods(cli *ent.Client) (*ent.GoodSelect, error) {
	stm, err := goodcrud.SetQueryConds(cli.Good.Query(), h.Conds)
	if err != nil {
		return nil, err
	}
	return h.selectGood(stm), nil
}

func (h *queryHandler) queryJoinMyself(s *sql.Selector) {
	t := sql.Table(entgood.Table)
	s.LeftJoin(t).
		On(
			s.C(entgood.FieldID),
			t.C(entgood.FieldID),
		).
		AppendSelect(
			sql.As(t.C(entgood.FieldEntID), "ent_id"),
			sql.As(t.C(entgood.FieldDeviceInfoID), "device_info_id"),
			sql.As(t.C(entgood.FieldCoinTypeID), "coin_type_id"),
			sql.As(t.C(entgood.FieldVendorLocationID), "vendor_location_id"),
			sql.As(t.C(entgood.FieldUnitPrice), "unit_price"),
			sql.As(t.C(entgood.FieldBenefitType), "benefit_type"),
			sql.As(t.C(entgood.FieldGoodType), "good_type"),
			sql.As(t.C(entgood.FieldTitle), "title"),
			sql.As(t.C(entgood.FieldQuantityUnit), "quantity_unit"),
			sql.As(t.C(entgood.FieldQuantityUnitAmount), "quantity_unit_amount"),
			sql.As(t.C(entgood.FieldDeliveryAt), "delivery_at"),
			sql.As(t.C(entgood.FieldStartAt), "start_at"),
			sql.As(t.C(entgood.FieldStartMode), "start_mode"),
			sql.As(t.C(entgood.FieldTestOnly), "test_only"),
			sql.As(t.C(entgood.FieldBenefitIntervalHours), "benefit_interval_hours"),
			sql.As(t.C(entgood.FieldUnitLockDeposit), "unit_lock_deposit"),
			sql.As(t.C(entgood.FieldUnitType), "unit_type"),
			sql.As(t.C(entgood.FieldQuantityCalculateType), "quantity_calculate_type"),
			sql.As(t.C(entgood.FieldDurationType), "duration_type"),
			sql.As(t.C(entgood.FieldDurationCalculateType), "duration_calculate_type"),
			sql.As(t.C(entgood.FieldSettlementType), "settlement_type"),
			sql.As(t.C(entgood.FieldCreatedAt), "created_at"),
			sql.As(t.C(entgood.FieldUpdatedAt), "updated_at"),
		)
}

//nolint:dupl
func (h *queryHandler) queryJoinExtraInfo(s *sql.Selector) {
	t := sql.Table(entextrainfo.Table)
	s.LeftJoin(t).
		On(
			s.C(entgood.FieldEntID),
			t.C(entextrainfo.FieldGoodID),
		).
		OnP(
			sql.EQ(t.C(entextrainfo.FieldDeletedAt), 0),
		).
		AppendSelect(
			sql.As(t.C(entextrainfo.FieldPosters), "posters"),
			sql.As(t.C(entextrainfo.FieldLabels), "labels"),
			sql.As(t.C(entextrainfo.FieldLikes), "likes"),
			sql.As(t.C(entextrainfo.FieldDislikes), "dislikes"),
			sql.As(t.C(entextrainfo.FieldScoreCount), "score_count"),
			sql.As(t.C(entextrainfo.FieldRecommendCount), "recommend_count"),
			sql.As(t.C(entextrainfo.FieldCommentCount), "comment_count"),
			sql.As(t.C(entextrainfo.FieldScore), "score"),
		)
}

func (h *queryHandler) queryJoinReward(s *sql.Selector) {
	t := sql.Table(entgoodreward.Table)
	s.LeftJoin(t).
		On(
			s.C(entgood.FieldEntID),
			t.C(entgoodreward.FieldGoodID),
		).
		OnP(
			sql.EQ(t.C(entgoodreward.FieldDeletedAt), 0),
		)

	if h.Conds != nil && h.Conds.RewardState != nil {
		s.Where(
			sql.EQ(t.C(entgoodreward.FieldRewardState), h.Conds.RewardState.Val.(types.BenefitState).String()),
		)
	}
	if h.Conds != nil && h.Conds.RewardAt != nil {
		switch h.Conds.RewardAt.Op {
		case cruder.EQ:
			s.Where(sql.EQ(t.C(entgoodreward.FieldLastRewardAt), h.Conds.RewardAt.Val))
		case cruder.NEQ:
			s.Where(sql.NEQ(t.C(entgoodreward.FieldLastRewardAt), h.Conds.RewardAt.Val))
		}
	}

	s.AppendSelect(
		sql.As(t.C(entgoodreward.FieldRewardState), "reward_state"),
		sql.As(t.C(entgoodreward.FieldLastRewardAt), "last_reward_at"),
		sql.As(t.C(entgoodreward.FieldRewardTid), "reward_tid"),
		sql.As(t.C(entgoodreward.FieldNextRewardStartAmount), "next_reward_start_amount"),
		sql.As(t.C(entgoodreward.FieldLastRewardAmount), "last_reward_amount"),
		sql.As(t.C(entgoodreward.FieldLastUnitRewardAmount), "last_unit_reward_amount"),
		sql.As(t.C(entgoodreward.FieldTotalRewardAmount), "total_reward_amount"),
	)
}

//nolint:dupl
func (h *queryHandler) queryJoinStock(s *sql.Selector) {
	t := sql.Table(entstock.Table)
	s.LeftJoin(t).
		On(
			s.C(entgood.FieldEntID),
			t.C(entstock.FieldGoodID),
		).
		OnP(
			sql.EQ(t.C(entstock.FieldDeletedAt), 0),
		).
		AppendSelect(
			sql.As(t.C(entstock.FieldEntID), "good_stock_id"),
			sql.As(t.C(entstock.FieldTotal), "good_total"),
			sql.As(t.C(entstock.FieldSpotQuantity), "good_spot_quantity"),
			sql.As(t.C(entstock.FieldLocked), "good_locked"),
			sql.As(t.C(entstock.FieldInService), "good_in_service"),
			sql.As(t.C(entstock.FieldWaitStart), "good_wait_start"),
			sql.As(t.C(entstock.FieldSold), "good_sold"),
			sql.As(t.C(entstock.FieldAppReserved), "good_app_reserved"),
		)
}

func (h *queryHandler) queryJoinDeviceInfo(s *sql.Selector) {
	t := sql.Table(entdevice.Table)
	s.LeftJoin(t).
		On(
			s.C(entgood.FieldDeviceInfoID),
			t.C(entdevice.FieldEntID),
		).
		OnP(
			sql.EQ(t.C(entdevice.FieldDeletedAt), 0),
		).
		AppendSelect(
			sql.As(t.C(entdevice.FieldType), "device_type"),
			sql.As(t.C(entdevice.FieldManufacturer), "device_manufacturer"),
			sql.As(t.C(entdevice.FieldPowerConsumption), "device_power_consumption"),
			sql.As(t.C(entdevice.FieldShipmentAt), "device_shipment_at"),
			sql.As(t.C(entdevice.FieldPosters), "device_posters"),
		)
}

func (h *queryHandler) queryJoinVendorLocation(s *sql.Selector) {
	t1 := sql.Table(entvendorlocation.Table)
	s.LeftJoin(t1).
		On(
			s.C(entgood.FieldVendorLocationID),
			t1.C(entvendorlocation.FieldEntID),
		).
		OnP(
			sql.EQ(t1.C(entvendorlocation.FieldDeletedAt), 0),
		).
		AppendSelect(
			sql.As(t1.C(entvendorlocation.FieldCountry), "vendor_location_country"),
			sql.As(t1.C(entvendorlocation.FieldProvince), "vendor_location_province"),
			sql.As(t1.C(entvendorlocation.FieldCity), "vendor_location_city"),
			sql.As(t1.C(entvendorlocation.FieldAddress), "vendor_location_address"),
		)

	t2 := sql.Table(entvendorbrand.Table)
	s.LeftJoin(t2).
		On(
			t1.C(entvendorlocation.FieldBrandID),
			t2.C(entvendorbrand.FieldEntID),
		).
		OnP(
			sql.EQ(t2.C(entvendorbrand.FieldDeletedAt), 0),
		).
		AppendSelect(
			sql.As(t2.C(entvendorbrand.FieldName), "vendor_brand_name"),
			sql.As(t2.C(entvendorbrand.FieldLogo), "vendor_brand_logo"),
		)
}

func (h *queryHandler) queryJoin() {
	h.stmSelect.Modify(func(s *sql.Selector) {
		h.queryJoinMyself(s)
		h.queryJoinExtraInfo(s)
		h.queryJoinReward(s)
		h.queryJoinStock(s)
		h.queryJoinDeviceInfo(s)
		h.queryJoinVendorLocation(s)
	})
	if h.stmCount == nil {
		return
	}
	h.stmCount.Modify(func(s *sql.Selector) {
		h.queryJoinExtraInfo(s)
		h.queryJoinReward(s)
		h.queryJoinStock(s)
		h.queryJoinDeviceInfo(s)
		h.queryJoinVendorLocation(s)
	})
}

func (h *queryHandler) scan(ctx context.Context) error {
	return h.stmSelect.Scan(ctx, &h.infos)
}

//nolint:funlen,gocyclo
func (h *queryHandler) formalize() {
	for _, info := range h.infos {
		_ = json.Unmarshal([]byte(info.DevicePostersStr), &info.DevicePosters)
		info.GoodType = types.GoodType(types.GoodType_value[info.GoodTypeStr])
		info.StartMode = types.GoodStartMode(types.GoodStartMode_value[info.StartModeStr])
		info.BenefitType = types.BenefitType(types.BenefitType_value[info.BenefitTypeStr])
		info.RewardState = types.BenefitState(types.BenefitState_value[info.RewardStateStr])
		info.UnitType = types.GoodUnitType(types.GoodUnitType_value[info.UnitTypeStr])
		info.QuantityCalculateType = types.GoodUnitCalculateType(types.GoodUnitCalculateType_value[info.QuantityCalculateTypeStr])
		info.DurationType = types.GoodDurationType(types.GoodDurationType_value[info.DurationTypeStr])
		info.SettlementType = types.GoodSettlementType(types.GoodSettlementType_value[info.SettlementTypeStr])
		info.DurationCalculateType = types.GoodUnitCalculateType(types.GoodUnitCalculateType_value[info.DurationCalculateTypeStr])
		_ = json.Unmarshal([]byte(info.PostersStr), &info.Posters)
		amount, err := decimal.NewFromString(info.UnitLockDeposit)
		if err != nil {
			info.UnitLockDeposit = decimal.NewFromInt(0).String()
		} else {
			info.UnitLockDeposit = amount.String()
		}
		amount, err = decimal.NewFromString(info.GoodTotal)
		if err != nil {
			info.GoodTotal = decimal.NewFromInt(0).String()
		} else {
			info.GoodTotal = amount.String()
		}
		amount, err = decimal.NewFromString(info.GoodSpotQuantity)
		if err != nil {
			info.GoodSpotQuantity = decimal.NewFromInt(0).String()
		} else {
			info.GoodSpotQuantity = amount.String()
		}
		amount, err = decimal.NewFromString(info.GoodLocked)
		if err != nil {
			info.GoodLocked = decimal.NewFromInt(0).String()
		} else {
			info.GoodLocked = amount.String()
		}
		amount, err = decimal.NewFromString(info.GoodInService)
		if err != nil {
			info.GoodInService = decimal.NewFromInt(0).String()
		} else {
			info.GoodInService = amount.String()
		}
		amount, err = decimal.NewFromString(info.GoodWaitStart)
		if err != nil {
			info.GoodWaitStart = decimal.NewFromInt(0).String()
		} else {
			info.GoodWaitStart = amount.String()
		}
		amount, err = decimal.NewFromString(info.GoodSold)
		if err != nil {
			info.GoodSold = decimal.NewFromInt(0).String()
		} else {
			info.GoodSold = amount.String()
		}
		amount, err = decimal.NewFromString(info.GoodAppReserved)
		if err != nil {
			info.GoodAppReserved = decimal.NewFromInt(0).String()
		} else {
			info.GoodAppReserved = amount.String()
		}
		amount, err = decimal.NewFromString(info.UnitPrice)
		if err != nil {
			info.UnitPrice = decimal.NewFromInt(0).String()
		} else {
			info.UnitPrice = amount.String()
		}
		amount, err = decimal.NewFromString(info.NextRewardStartAmount)
		if err != nil {
			info.NextRewardStartAmount = decimal.NewFromInt(0).String()
		} else {
			info.NextRewardStartAmount = amount.String()
		}
		amount, err = decimal.NewFromString(info.LastRewardAmount)
		if err != nil {
			info.LastRewardAmount = decimal.NewFromInt(0).String()
		} else {
			info.LastRewardAmount = amount.String()
		}
		amount, err = decimal.NewFromString(info.LastUnitRewardAmount)
		if err != nil {
			info.LastUnitRewardAmount = decimal.NewFromInt(0).String()
		} else {
			info.LastUnitRewardAmount = amount.String()
		}
		amount, err = decimal.NewFromString(info.TotalRewardAmount)
		if err != nil {
			info.TotalRewardAmount = decimal.NewFromInt(0).String()
		} else {
			info.TotalRewardAmount = amount.String()
		}
		amount, err = decimal.NewFromString(info.Score)
		if err != nil {
			info.Score = decimal.NewFromInt(0).String()
		} else {
			info.Score = amount.String()
		}
		amount, err = decimal.NewFromString(info.QuantityUnitAmount)
		if err != nil {
			info.QuantityUnitAmount = decimal.NewFromInt(0).String()
		} else {
			info.QuantityUnitAmount = amount.String()
		}
		labels := []string{}
		_ = json.Unmarshal([]byte(info.LabelsStr), &labels)
		for _, label := range labels {
			info.Labels = append(info.Labels, types.GoodLabel(types.GoodLabel_value[label]))
		}
	}
}

func (h *Handler) GetGood(ctx context.Context) (*npool.Good, error) {
	handler := &queryHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryGood(cli); err != nil {
			return err
		}
		handler.queryJoin()
		return handler.scan(_ctx)
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

func (h *Handler) GetGoods(ctx context.Context) ([]*npool.Good, uint32, error) {
	handler := &queryHandler{
		Handler: h,
	}

	var err error

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		handler.stmSelect, err = handler.queryGoods(cli)
		if err != nil {
			return err
		}
		handler.stmCount, err = handler.queryGoods(cli)
		if err != nil {
			return err
		}

		handler.queryJoin()

		total, err := handler.stmCount.Count(_ctx)
		if err != nil {
			return err
		}
		handler.total = uint32(total)

		handler.stmSelect.
			Offset(int(handler.Offset)).
			Limit(int(handler.Limit))
		return handler.scan(_ctx)
	})
	if err != nil {
		return nil, 0, err
	}

	handler.formalize()

	return handler.infos, handler.total, nil
}
