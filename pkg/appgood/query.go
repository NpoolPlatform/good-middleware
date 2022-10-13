package appgood

import (
	"context"
	"encoding/json"

	mgrpb "github.com/NpoolPlatform/message/npool/good/mgr/v1/appgood"
	goodmgrpb "github.com/NpoolPlatform/message/npool/good/mgr/v1/good"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/appgood"

	"github.com/NpoolPlatform/good-manager/pkg/db"
	"github.com/NpoolPlatform/good-manager/pkg/db/ent"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	"entgo.io/ent/dialect/sql"

	entappgood "github.com/NpoolPlatform/good-manager/pkg/db/ent/appgood"
	entdeviceinfo "github.com/NpoolPlatform/good-manager/pkg/db/ent/deviceinfo"
	entextra "github.com/NpoolPlatform/good-manager/pkg/db/ent/extrainfo"
	entgood "github.com/NpoolPlatform/good-manager/pkg/db/ent/good"
	entpromotion "github.com/NpoolPlatform/good-manager/pkg/db/ent/promotion"
	entrecommend "github.com/NpoolPlatform/good-manager/pkg/db/ent/recommend"
	entstock "github.com/NpoolPlatform/good-manager/pkg/db/ent/stock"
	entvendorlocation "github.com/NpoolPlatform/good-manager/pkg/db/ent/vendorlocation"

	"github.com/google/uuid"
)

func GetGoods(ctx context.Context, conds *mgrpb.Conds, offset, limit int32) ([]*npool.Good, uint32, error) {
	var infos []*npool.Good
	var total uint32
	var err error

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		stm := cli.
			AppGood.
			Query()

		if conds.ID != nil {
			stm.Where(
				entappgood.ID(uuid.MustParse(conds.GetID().GetValue())),
			)
		}
		if conds.AppID != nil {
			stm.Where(
				entappgood.AppID(uuid.MustParse(conds.GetAppID().GetValue())),
			)
		}
		if conds.GoodID != nil {
			stm.Where(
				entappgood.GoodID(uuid.MustParse(conds.GetGoodID().GetValue())),
			)
		}

		n, err := stm.Count(_ctx)
		if err != nil {
			return err
		}

		total = uint32(n)

		stm.
			Offset(int(offset)).
			Limit(int(limit))

		return join(stm).
			Scan(_ctx, &infos)
	})
	if err != nil {
		logger.Sugar().Errorw("GetGoods", "err", err)
		return nil, 0, err
	}

	infos, err = expand(infos)
	if err != nil {
		logger.Sugar().Errorw("GetGoods", "err", err)
		return nil, 0, err
	}

	return infos, total, nil
}

// nolint
func join(stm *ent.AppGoodQuery) *ent.AppGoodSelect {
	return stm.
		Select(
			entappgood.FieldID,
			entappgood.FieldAppID,
			entappgood.FieldGoodID,
			entappgood.FieldOnline,
			entappgood.FieldVisible,
			entappgood.FieldPrice,
			entappgood.FieldDisplayIndex,
			entappgood.FieldPurchaseLimit,
			entappgood.FieldCommissionPercent,
			entappgood.FieldGoodName,
			entappgood.FieldCreatedAt,
		).
		Modify(func(s *sql.Selector) {
			t1 := sql.Table(entgood.Table)
			t2 := sql.Table(entdeviceinfo.Table)
			t3 := sql.Table(entvendorlocation.Table)

			s.
				LeftJoin(t1).
				On(
					s.C(entappgood.FieldGoodID),
					t1.C(entgood.FieldID),
				).
				LeftJoin(t2).
				On(
					t1.C(entgood.FieldDeviceInfoID),
					t2.C(entdeviceinfo.FieldID),
				).
				LeftJoin(t3).
				On(
					t1.C(entgood.FieldVendorLocationID),
					t3.C(entvendorlocation.FieldID),
				).
				AppendSelect(
					sql.As(t1.C(entgood.FieldDurationDays), "duration_days"),
					sql.As(t1.C(entgood.FieldCoinTypeID), "coin_type_id"),
					sql.As(t1.C(entgood.FieldGoodType), "good_type"),
					sql.As(t1.C(entgood.FieldBenefitType), "benefit_type"),
					sql.As(t1.C(entgood.FieldUnit), "unit"),
					sql.As(t1.C(entgood.FieldUnitAmount), "unit_amount"),
					sql.As(t1.C(entgood.FieldSupportCoinTypeIds), "support_coin_type_ids"),
					sql.As(t1.C(entgood.FieldTestOnly), "test_only"),
					sql.As(t1.C(entgood.FieldStartAt), "start_at"),
					sql.As(t2.C(entdeviceinfo.FieldType), "device_type"),
					sql.As(t2.C(entdeviceinfo.FieldManufacturer), "device_manufacturer"),
					sql.As(t2.C(entdeviceinfo.FieldPowerComsuption), "device_power_comsuption"),
					sql.As(t2.C(entdeviceinfo.FieldShipmentAt), "device_shipment_at"),
					sql.As(t2.C(entdeviceinfo.FieldPosters), "device_posters"),
					sql.As(t3.C(entvendorlocation.FieldCountry), "vendor_location_country"),
				)

			t4 := sql.Table(entextra.Table)
			s.
				LeftJoin(t4).
				On(
					s.C(entappgood.FieldGoodID),
					t4.C(entextra.FieldGoodID),
				).
				AppendSelect(
					sql.As(t4.C(entextra.FieldPosters), "posters"),
					sql.As(t4.C(entextra.FieldLabels), "labels"),
					sql.As(t4.C(entextra.FieldVoteCount), "vote_count"),
					sql.As(t4.C(entextra.FieldRating), "rating"),
				)

			t5 := sql.Table(entstock.Table)
			s.
				LeftJoin(t5).
				On(
					s.C(entappgood.FieldGoodID),
					t5.C(entstock.FieldGoodID),
				).
				AppendSelect(
					sql.As(t5.C(entstock.FieldTotal), "good_total"),
					sql.As(t5.C(entstock.FieldLocked), "good_locked"),
					sql.As(t5.C(entstock.FieldInService), "good_in_service"),
					sql.As(t5.C(entstock.FieldSold), "good_sold"),
				)

			t6 := sql.Table(entpromotion.Table)
			s.
				LeftJoin(t6).
				On(
					s.C(entappgood.FieldGoodID),
					t6.C(entpromotion.FieldGoodID),
				).
				On(
					s.C(entappgood.FieldAppID),
					t6.C(entpromotion.FieldAppID),
				).
				AppendSelect(
					sql.As(t6.C(entpromotion.FieldStartAt), "promotion_start_at"),
					sql.As(t6.C(entpromotion.FieldEndAt), "promotion_end_at"),
					sql.As(t6.C(entpromotion.FieldMessage), "promotion_message"),
					sql.As(t6.C(entpromotion.FieldPrice), "promotion_price"),
					sql.As(t6.C(entpromotion.FieldPosters), "promotion_posters"),
				)

			t7 := sql.Table(entrecommend.Table)
			s.
				LeftJoin(t7).
				On(
					s.C(entappgood.FieldGoodID),
					t7.C(entrecommend.FieldGoodID),
				).
				On(
					s.C(entappgood.FieldAppID),
					t7.C(entrecommend.FieldAppID),
				).
				AppendSelect(
					sql.As(t7.C(entrecommend.FieldRecommenderID), "recommender_id"),
					sql.As(t7.C(entrecommend.FieldMessage), "recommend_message"),
					sql.As(t7.C(entrecommend.FieldRecommendIndex), "recommend_index"),
				)
		})
}

func expand(infos []*npool.Good) ([]*npool.Good, error) {
	for _, info := range infos {
		var labels []string
		if err := json.Unmarshal([]byte(info.LabelsStr), &labels); err != nil {
			return nil, err
		}
		info.Labels = labels

		var posters []string
		if err := json.Unmarshal([]byte(info.PostersStr), &posters); err != nil {
			return nil, err
		}
		info.Posters = posters

		var coinTypeIDs []string
		if err := json.Unmarshal([]byte(info.SupportCoinTypeIDsStr), &coinTypeIDs); err != nil {
			return nil, err
		}
		info.SupportCoinTypeIDs = coinTypeIDs

		info.GoodType = goodmgrpb.GoodType(goodmgrpb.GoodType_value[info.GoodTypeStr])
		info.BenefitType = goodmgrpb.BenefitType(goodmgrpb.BenefitType_value[info.BenefitTypeStr])
	}
	return infos, nil
}
