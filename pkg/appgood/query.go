package appgood

import (
	"context"
	"encoding/json"
	"fmt"

	mgrpb "github.com/NpoolPlatform/message/npool/good/mgr/v1/appgood"
	goodmgrpb "github.com/NpoolPlatform/message/npool/good/mgr/v1/good"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/appgood"
	commmgrpb "github.com/NpoolPlatform/message/npool/inspire/mgr/v1/commission"

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

func GetGood(ctx context.Context, id string) (*npool.Good, error) {
	var infos []*npool.Good
	var err error

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		stm := cli.
			AppGood.
			Query().
			Where(
				entappgood.ID(uuid.MustParse(id)),
			).
			Limit(1)

		return join(stm).
			Scan(_ctx, &infos)
	})
	if err != nil {
		logger.Sugar().Errorw("GetGood", "err", err)
		return nil, err
	}
	if len(infos) == 0 {
		return nil, nil
	}

	infos, err = expand(infos)
	if err != nil {
		logger.Sugar().Errorw("GetGood", "err", err)
		return nil, err
	}

	return infos[0], nil
}

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
		if conds.GoodIDs != nil {
			goodIDs := []uuid.UUID{}
			for _, v := range conds.GetGoodIDs().GetValue() {
				uGoodID, err := uuid.Parse(v)
				if err != nil {
					return err
				}
				goodIDs = append(goodIDs, uGoodID)
			}
			stm.Where(
				entappgood.GoodIDIn(goodIDs...),
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

func GetGoodOnly(ctx context.Context, conds *mgrpb.Conds) (*npool.Good, error) {
	var infos []*npool.Good
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
		stm.
			Limit(1)

		return join(stm).
			Scan(_ctx, &infos)
	})
	if err != nil {
		logger.Sugar().Errorw("GetGoods", "err", err)
		return nil, err
	}

	if len(infos) == 0 {
		return nil, nil
	}
	if len(infos) > 1 {
		logger.Sugar().Errorw("err", "too many records")
		return nil, fmt.Errorf("too many records")
	}

	infos, err = expand(infos)
	if err != nil {
		logger.Sugar().Errorw("GetGoods", "err", err)
		return nil, err
	}

	return infos[0], nil
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
			entappgood.FieldUpdatedAt,
			entappgood.FieldSaleStartAt,
			entappgood.FieldSaleEndAt,
			entappgood.FieldServiceStartAt,
			entappgood.FieldTechnicalFeeRatio,
			entappgood.FieldElectricityFeeRatio,
			entappgood.FieldDailyRewardAmount,
			entappgood.FieldCommissionSettleType,
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
					sql.As(t1.C(entgood.FieldBenefitIntervalHours), "benefit_interval_hours"),
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

func expand(infos []*npool.Good) ([]*npool.Good, error) { //nolint
	for _, info := range infos {
		_ = json.Unmarshal([]byte(info.LabelsStr), &info.Labels)
		_ = json.Unmarshal([]byte(info.PostersStr), &info.Posters)
		_ = json.Unmarshal([]byte(info.SupportCoinTypeIDsStr), &info.SupportCoinTypeIDs)

		info.GoodType = goodmgrpb.GoodType(goodmgrpb.GoodType_value[info.GoodTypeStr])
		info.BenefitType = goodmgrpb.BenefitType(goodmgrpb.BenefitType_value[info.BenefitTypeStr])
		info.CommissionSettleType = commmgrpb.SettleType(commmgrpb.SettleType_value[info.CommissionSettleTypeStr])

		info.Visible = info.VisibleInt > 0
		info.Online = info.OnlineInt > 0
	}
	return infos, nil
}
