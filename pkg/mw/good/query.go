package good

import (
	"context"
	"encoding/json"
	"fmt"

	"entgo.io/ent/dialect/sql"

	goodcrud "github.com/NpoolPlatform/good-middleware/pkg/crud/good"
	"github.com/NpoolPlatform/good-middleware/pkg/db"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent"
	entdeviceinfo "github.com/NpoolPlatform/good-middleware/pkg/db/ent/deviceinfo"
	entextrainfo "github.com/NpoolPlatform/good-middleware/pkg/db/ent/extrainfo"
	entgood "github.com/NpoolPlatform/good-middleware/pkg/db/ent/good"
	entgoodreward "github.com/NpoolPlatform/good-middleware/pkg/db/ent/goodreward"
	entstock "github.com/NpoolPlatform/good-middleware/pkg/db/ent/stock"
	entvendorbrand "github.com/NpoolPlatform/good-middleware/pkg/db/ent/vendorbrand"
	entvendorlocation "github.com/NpoolPlatform/good-middleware/pkg/db/ent/vendorlocation"
	types "github.com/NpoolPlatform/message/npool/basetypes/good/v1"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/good"
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

func (h *queryHandler) queryGood(cli *ent.Client) {
	h.stmSelect = h.selectGood(
		cli.Good.
			Query().
			Where(
				entgood.ID(*h.ID),
				entgood.DeletedAt(0),
			),
	)
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
			sql.As(t.C(entgood.FieldDeviceInfoID), "device_info_id"),
			sql.As(t.C(entgood.FieldDurationDays), "duration_days"),
			sql.As(t.C(entgood.FieldCoinTypeID), "coin_type_id"),
			sql.As(t.C(entgood.FieldVendorLocationID), "vendor_location_id"),
			sql.As(t.C(entgood.FieldPrice), "price"),
			sql.As(t.C(entgood.FieldBenefitType), "benefit_type"),
			sql.As(t.C(entgood.FieldGoodType), "good_type"),
			sql.As(t.C(entgood.FieldTitle), "title"),
			sql.As(t.C(entgood.FieldUnit), "unit"),
			sql.As(t.C(entgood.FieldUnitAmount), "unit_amount"),
			sql.As(t.C(entgood.FieldSupportCoinTypeIds), "support_coin_type_ids"),
			sql.As(t.C(entgood.FieldDeliveryAt), "delivery_at"),
			sql.As(t.C(entgood.FieldStartAt), "start_at"),
			sql.As(t.C(entgood.FieldTestOnly), "test_only"),
			sql.As(t.C(entgood.FieldBenefitIntervalHours), "benefit_interval_hours"),
			sql.As(t.C(entgood.FieldUnitLockDeposit), "unit_lock_deposit"),
			sql.As(t.C(entgood.FieldCreatedAt), "created_at"),
			sql.As(t.C(entgood.FieldUpdatedAt), "updated_at"),
		)
}

func (h *queryHandler) queryJoinExtraInfo(s *sql.Selector) {
	t := sql.Table(entextrainfo.Table)
	s.LeftJoin(t).
		On(
			s.C(entgood.FieldID),
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
			sql.As(t.C(entextrainfo.FieldRatingV1), "rating"),
		)
}

func (h *queryHandler) queryJoinReward(s *sql.Selector) {
	t := sql.Table(entgoodreward.Table)
	s.LeftJoin(t).
		On(
			s.C(entgood.FieldID),
			t.C(entgoodreward.FieldGoodID),
		).
		OnP(
			sql.EQ(t.C(entgoodreward.FieldDeletedAt), 0),
		).
		AppendSelect(
			sql.As(t.C(entgoodreward.FieldRewardState), "reward_state"),
			sql.As(t.C(entgoodreward.FieldLastRewardAt), "last_reward_at"),
			sql.As(t.C(entgoodreward.FieldRewardTid), "reward_tid"),
			sql.As(t.C(entgoodreward.FieldNextRewardStartAmount), "next_reward_start_amount"),
			sql.As(t.C(entgoodreward.FieldLastRewardAmount), "last_reward_amount"),
		)
}

func (h *queryHandler) queryJoinStock(s *sql.Selector) {
	t := sql.Table(entstock.Table)
	s.LeftJoin(t).
		On(
			s.C(entgood.FieldID),
			t.C(entstock.FieldGoodID),
		).
		OnP(
			sql.EQ(t.C(entstock.FieldDeletedAt), 0),
		).
		AppendSelect(
			sql.As(t.C(entstock.FieldID), "good_stock_id"),
			sql.As(t.C(entstock.FieldTotal), "good_total"),
			sql.As(t.C(entstock.FieldLocked), "good_locked"),
			sql.As(t.C(entstock.FieldInService), "good_in_service"),
			sql.As(t.C(entstock.FieldWaitStart), "good_wait_start"),
			sql.As(t.C(entstock.FieldSold), "good_sold"),
			sql.As(t.C(entstock.FieldAppLocked), "good_app_locked"),
		)
}

func (h *queryHandler) queryJoinDeviceInfo(s *sql.Selector) {
	t := sql.Table(entdeviceinfo.Table)
	s.LeftJoin(t).
		On(
			s.C(entgood.FieldDeviceInfoID),
			t.C(entdeviceinfo.FieldID),
		).
		OnP(
			sql.EQ(t.C(entdeviceinfo.FieldDeletedAt), 0),
		).
		AppendSelect(
			sql.As(t.C(entdeviceinfo.FieldType), "device_type"),
			sql.As(t.C(entdeviceinfo.FieldManufacturer), "device_manufacturer"),
			sql.As(t.C(entdeviceinfo.FieldPowerComsuption), "device_power_comsuption"),
			sql.As(t.C(entdeviceinfo.FieldShipmentAt), "device_shipment_at"),
			sql.As(t.C(entdeviceinfo.FieldPosters), "device_posters"),
		)
}

func (h *queryHandler) queryJoinVendorLocation(s *sql.Selector) {
	t1 := sql.Table(entvendorlocation.Table)
	s.LeftJoin(t1).
		On(
			s.C(entgood.FieldVendorLocationID),
			t1.C(entvendorlocation.FieldID),
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
			t2.C(entvendorbrand.FieldID),
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

func (h *queryHandler) formalize() {
	for _, info := range h.infos {
		_ = json.Unmarshal([]byte(info.DevicePostersStr), &info.DevicePosters)
		info.GoodType = types.GoodType(types.GoodType_value[info.GoodTypeStr])
		info.BenefitType = types.BenefitType(types.BenefitType_value[info.BenefitTypeStr])
		_ = json.Unmarshal([]byte(info.SupportCoinTypeIDsStr), &info.SupportCoinTypeIDs)
		_ = json.Unmarshal([]byte(info.PostersStr), &info.Posters)
		_ = json.Unmarshal([]byte(info.LabelsStr), &info.Labels)
	}
}

func (h *Handler) GetGood(ctx context.Context) (*npool.Good, error) {
	handler := &queryHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		handler.queryGood(cli)
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
