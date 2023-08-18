package appgood

import (
	"context"
	"encoding/json"
	"fmt"

	"entgo.io/ent/dialect/sql"

	appgoodcrud "github.com/NpoolPlatform/good-middleware/pkg/crud/app/good"
	"github.com/NpoolPlatform/good-middleware/pkg/db"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent"
	entappgood "github.com/NpoolPlatform/good-middleware/pkg/db/ent/appgood"
	entappstock "github.com/NpoolPlatform/good-middleware/pkg/db/ent/appstock"
	entdeviceinfo "github.com/NpoolPlatform/good-middleware/pkg/db/ent/deviceinfo"
	entextrainfo "github.com/NpoolPlatform/good-middleware/pkg/db/ent/extrainfo"
	entgood "github.com/NpoolPlatform/good-middleware/pkg/db/ent/good"
	entstock "github.com/NpoolPlatform/good-middleware/pkg/db/ent/stock"
	entvendorbrand "github.com/NpoolPlatform/good-middleware/pkg/db/ent/vendorbrand"
	entvendorlocation "github.com/NpoolPlatform/good-middleware/pkg/db/ent/vendorlocation"
	types "github.com/NpoolPlatform/message/npool/basetypes/good/v1"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/app/good"

	"github.com/shopspring/decimal"
)

type queryHandler struct {
	*Handler
	stmSelect *ent.AppGoodSelect
	stmCount  *ent.AppGoodSelect
	infos     []*npool.Good
	total     uint32
}

func (h *queryHandler) selectGood(stm *ent.AppGoodQuery) *ent.AppGoodSelect {
	return stm.Select(entgood.FieldID)
}

func (h *queryHandler) queryGood(cli *ent.Client) {
	h.stmSelect = h.selectGood(
		cli.AppGood.
			Query().
			Where(
				entappgood.ID(*h.ID),
				entappgood.DeletedAt(0),
			),
	)
}

func (h *queryHandler) queryGoods(cli *ent.Client) (*ent.AppGoodSelect, error) {
	stm, err := appgoodcrud.SetQueryConds(cli.AppGood.Query(), h.Conds)
	if err != nil {
		return nil, err
	}
	return h.selectGood(stm), nil
}

func (h *queryHandler) queryJoinMyself(s *sql.Selector) {
	t := sql.Table(entappgood.Table)
	s.LeftJoin(t).
		On(
			s.C(entappgood.FieldID),
			t.C(entappgood.FieldID),
		).
		AppendSelect(
			sql.As(t.C(entappgood.FieldAppID), "app_id"),
			sql.As(t.C(entappgood.FieldGoodID), "good_id"),
			sql.As(t.C(entappgood.FieldOnline), "online"),
			sql.As(t.C(entappgood.FieldVisible), "visible"),
			sql.As(t.C(entappgood.FieldGoodName), "good_name"),
			sql.As(t.C(entappgood.FieldPrice), "price"),
			sql.As(t.C(entappgood.FieldDisplayIndex), "display_index"),
			sql.As(t.C(entappgood.FieldPurchaseLimit), "purchase_limit"),
			sql.As(t.C(entappgood.FieldSaleStartAt), "sale_start_at"),
			sql.As(t.C(entappgood.FieldSaleEndAt), "sale_end_at"),
			sql.As(t.C(entappgood.FieldServiceStartAt), "service_start_at"),
			sql.As(t.C(entappgood.FieldDescriptions), "descriptions"),
			sql.As(t.C(entappgood.FieldGoodBanner), "good_banner"),
			sql.As(t.C(entappgood.FieldDisplayNames), "display_names"),
			sql.As(t.C(entappgood.FieldEnablePurchase), "enable_purchase"),
			sql.As(t.C(entappgood.FieldEnableProductPage), "enable_product_page"),
			sql.As(t.C(entappgood.FieldCancelMode), "cancel_mode"),
			sql.As(t.C(entappgood.FieldUserPurchaseLimit), "user_purchase_limit"),
			sql.As(t.C(entappgood.FieldDisplayColors), "display_colors"),
			sql.As(t.C(entappgood.FieldCancellableBeforeStart), "cancellable_before_start"),
			sql.As(t.C(entappgood.FieldProductPage), "product_page"),
			sql.As(t.C(entappgood.FieldEnableSetCommission), "enable_set_commission"),
			sql.As(t.C(entappgood.FieldPosters), "posters"),
			sql.As(t.C(entappgood.FieldCreatedAt), "created_at"),
			sql.As(t.C(entappgood.FieldUpdatedAt), "updated_at"),
		)
}

func (h *queryHandler) queryJoinGood(s *sql.Selector) {
	t1 := sql.Table(entgood.Table)
	t2 := sql.Table(entdeviceinfo.Table)
	t3 := sql.Table(entvendorlocation.Table)
	t4 := sql.Table(entvendorbrand.Table)

	s.LeftJoin(t1).
		On(
			s.C(entappgood.FieldGoodID),
			t1.C(entgood.FieldID),
		).
		LeftJoin(t2).
		On(
			t1.C(entgood.FieldDeviceInfoID),
			t2.C(entdeviceinfo.FieldID),
		).
		OnP(
			sql.EQ(t2.C(entdeviceinfo.FieldDeletedAt), 0),
		).
		LeftJoin(t3).
		On(
			t1.C(entgood.FieldVendorLocationID),
			t3.C(entvendorlocation.FieldID),
		).
		OnP(
			sql.EQ(t3.C(entvendorlocation.FieldDeletedAt), 0),
		).
		LeftJoin(t4).
		On(
			t3.C(entvendorlocation.FieldBrandID),
			t4.C(entvendorbrand.FieldID),
		).
		OnP(
			sql.EQ(t4.C(entvendorbrand.FieldDeletedAt), 0),
		).
		AppendSelect(
			sql.As(t1.C(entgood.FieldDeviceInfoID), "device_info_id"),
			sql.As(t1.C(entgood.FieldVendorLocationID), "vendor_location_id"),
			sql.As(t1.C(entgood.FieldDurationDays), "duration_days"),
			sql.As(t1.C(entgood.FieldCoinTypeID), "coin_type_id"),
			sql.As(t1.C(entgood.FieldBenefitType), "benefit_type"),
			sql.As(t1.C(entgood.FieldGoodType), "good_type"),
			sql.As(t1.C(entgood.FieldUnit), "unit"),
			sql.As(t1.C(entgood.FieldUnitAmount), "unit_amount"),
			sql.As(t1.C(entgood.FieldSupportCoinTypeIds), "support_coin_type_ids"),
			sql.As(t1.C(entgood.FieldStartAt), "start_at"),
			sql.As(t1.C(entgood.FieldTestOnly), "test_only"),
			sql.As(t1.C(entgood.FieldBenefitIntervalHours), "benefit_interval_hours"),
			sql.As(t2.C(entdeviceinfo.FieldType), "device_type"),
			sql.As(t2.C(entdeviceinfo.FieldManufacturer), "device_manufacturer"),
			sql.As(t2.C(entdeviceinfo.FieldPowerComsuption), "device_power_comsuption"),
			sql.As(t2.C(entdeviceinfo.FieldShipmentAt), "device_shipment_at"),
			sql.As(t2.C(entdeviceinfo.FieldPosters), "device_posters"),
			sql.As(t3.C(entvendorlocation.FieldCountry), "vendor_location_country"),
			sql.As(t4.C(entvendorbrand.FieldName), "vendor_brand_name"),
			sql.As(t4.C(entvendorbrand.FieldLogo), "vendor_brand_logo"),
		)
}

func (h *queryHandler) queryJoinExtraInfo(s *sql.Selector) {
	t := sql.Table(entextrainfo.Table)
	s.LeftJoin(t).
		On(
			s.C(entappgood.FieldGoodID),
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

func (h *queryHandler) queryJoinStock(s *sql.Selector) {
	t := sql.Table(entstock.Table)
	s.LeftJoin(t).
		On(
			s.C(entappgood.FieldGoodID),
			t.C(entstock.FieldGoodID),
		).
		OnP(
			sql.EQ(t.C(entstock.FieldDeletedAt), 0),
		).
		AppendSelect(
			sql.As(t.C(entstock.FieldTotal), "good_total"),
			sql.As(t.C(entstock.FieldLocked), "good_locked"),
			sql.As(t.C(entstock.FieldInService), "good_in_service"),
			sql.As(t.C(entstock.FieldWaitStart), "good_wait_start"),
			sql.As(t.C(entstock.FieldSold), "good_sold"),
		)
}

func (h *queryHandler) queryJoinAppStock(s *sql.Selector) {
	t := sql.Table(entappstock.Table)
	s.LeftJoin(t).
		On(
			s.C(entappgood.FieldGoodID),
			t.C(entappstock.FieldAppGoodID),
		).
		OnP(
			sql.EQ(t.C(entappstock.FieldDeletedAt), 0),
		).
		AppendSelect(
			sql.As(t.C(entappstock.FieldID), "app_good_stock_id"),
			sql.As(t.C(entappstock.FieldReserved), "app_good_reserved"),
			sql.As(t.C(entappstock.FieldSpotQuantity), "app_spot_quantity"),
			sql.As(t.C(entappstock.FieldLocked), "app_good_locked"),
			sql.As(t.C(entappstock.FieldInService), "app_good_in_service"),
			sql.As(t.C(entappstock.FieldWaitStart), "app_good_wait_start"),
			sql.As(t.C(entappstock.FieldSold), "app_good_sold"),
		)
}

func (h *queryHandler) queryJoin() {
	h.stmSelect.Modify(func(s *sql.Selector) {
		h.queryJoinMyself(s)
		h.queryJoinGood(s)
		h.queryJoinExtraInfo(s)
		h.queryJoinStock(s)
		h.queryJoinAppStock(s)
	})
	if h.stmCount == nil {
		return
	}
	h.stmCount.Modify(func(s *sql.Selector) {
		h.queryJoinGood(s)
		h.queryJoinExtraInfo(s)
		h.queryJoinStock(s)
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
		amount, err := decimal.NewFromString(info.GoodTotal)
		if err != nil {
			info.GoodTotal = decimal.NewFromInt(0).String()
		} else {
			info.GoodTotal = amount.String()
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
		amount, err = decimal.NewFromString(info.Price)
		if err != nil {
			info.Price = decimal.NewFromInt(0).String()
		} else {
			info.Price = amount.String()
		}
		amount, err = decimal.NewFromString(info.Score)
		if err != nil {
			info.Score = decimal.NewFromInt(0).String()
		} else {
			info.Score = amount.String()
		}
		amount, err = decimal.NewFromString(info.UserPurchaseLimit)
		if err != nil {
			info.UserPurchaseLimit = decimal.NewFromInt(0).String()
		} else {
			info.UserPurchaseLimit = amount.String()
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
		handler.queryGood(cli)
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