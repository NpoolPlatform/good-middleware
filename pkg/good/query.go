package good

import (
	"context"
	"encoding/json"
	"fmt"

	"go.opentelemetry.io/otel"
	scodes "go.opentelemetry.io/otel/codes"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	"entgo.io/ent/dialect/sql"

	constant "github.com/NpoolPlatform/good-middleware/pkg/message/const"
	commontracer "github.com/NpoolPlatform/good-middleware/pkg/tracer"

	mgrpb "github.com/NpoolPlatform/message/npool/good/mgr/v1/good"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/good"

	"github.com/NpoolPlatform/good-manager/pkg/db"
	"github.com/NpoolPlatform/good-manager/pkg/db/ent"

	entdeviceinfo "github.com/NpoolPlatform/good-manager/pkg/db/ent/deviceinfo"
	entextrainfo "github.com/NpoolPlatform/good-manager/pkg/db/ent/extrainfo"
	entgood "github.com/NpoolPlatform/good-manager/pkg/db/ent/good"
	entstock "github.com/NpoolPlatform/good-manager/pkg/db/ent/stock"
	entvendorlocation "github.com/NpoolPlatform/good-manager/pkg/db/ent/vendorlocation"

	goodcrud "github.com/NpoolPlatform/good-manager/pkg/crud/good"

	"github.com/google/uuid"
)

func GetGood(ctx context.Context, id string) (*npool.Good, error) {
	var infos []*npool.Good
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetGood")
	defer span.End()
	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = commontracer.TraceInvoker(span, "good", "middleware", "CRUD")

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		stm := cli.
			Good.
			Query().
			Where(
				entgood.ID(uuid.MustParse(id)),
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
	if len(infos) > 1 {
		logger.Sugar().Errorw("GetGood", "err", "too many records")
		return nil, fmt.Errorf("too many records")
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

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetManyGoods")
	defer span.End()
	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = commontracer.TraceInvoker(span, "good", "middleware", "CRUD")

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		stm, err := goodcrud.SetQueryConds(conds, cli)
		if err != nil {
			return err
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
		logger.Sugar().Errorw("GetManyGoods", "err", err)
		return nil, 0, err
	}
	if len(infos) == 0 {
		return nil, 0, nil
	}

	infos, err = expand(infos)
	if err != nil {
		logger.Sugar().Errorw("GetManyGoods", "err", err)
		return nil, 0, err
	}

	return infos, total, nil
}

func GetManyGoods(ctx context.Context, ids []string, offset, limit int32) ([]*npool.Good, uint32, error) {
	var infos []*npool.Good
	var total uint32
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetManyGoods")
	defer span.End()
	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	uids := []uuid.UUID{}
	for _, id := range ids {
		uids = append(uids, uuid.MustParse(id))
	}

	span = commontracer.TraceInvoker(span, "good", "middleware", "CRUD")

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		stm := cli.
			Good.
			Query().
			Where(
				entgood.IDIn(uids...),
			)

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
		logger.Sugar().Errorw("GetManyGoods", "err", err)
		return nil, 0, err
	}
	if len(infos) == 0 {
		return nil, 0, nil
	}

	infos, err = expand(infos)
	if err != nil {
		logger.Sugar().Errorw("GetManyGoods", "err", err)
		return nil, 0, err
	}

	return infos, total, nil
}

func join(stm *ent.GoodQuery) *ent.GoodSelect {
	return stm.
		Select(
			entgood.FieldID,
			entgood.FieldDeviceInfoID,
			entgood.FieldDurationDays,
			entgood.FieldCoinTypeID,
			entgood.FieldInheritFromGoodID,
			entgood.FieldPrice,
			entgood.FieldBenefitType,
			entgood.FieldGoodType,
			entgood.FieldTitle,
			entgood.FieldUnit,
			entgood.FieldUnitAmount,
			entgood.FieldSupportCoinTypeIds,
			entgood.FieldDeliveryAt,
			entgood.FieldStartAt,
			entgood.FieldTestOnly,
			entgood.FieldCreatedAt,
			entgood.FieldUpdatedAt,
			entgood.FieldVendorLocationID,
			entgood.FieldBenefitIntervalHours,
			entgood.FieldBenefitState,
			entgood.FieldLastBenefitAt,
			entgood.FieldBenefitTids,
			entgood.FieldNextBenefitStartAmount,
			entgood.FieldLastBenefitAmount,
		).
		Modify(func(s *sql.Selector) {
			t1 := sql.Table(entdeviceinfo.Table)
			s.
				LeftJoin(t1).
				On(
					s.C(entgood.FieldDeviceInfoID),
					t1.C(entdeviceinfo.FieldID),
				).
				AppendSelect(
					sql.As(t1.C(entdeviceinfo.FieldType), "device_type"),
					sql.As(t1.C(entdeviceinfo.FieldManufacturer), "device_manufacturer"),
					sql.As(t1.C(entdeviceinfo.FieldPowerComsuption), "device_power_comsuption"),
					sql.As(t1.C(entdeviceinfo.FieldShipmentAt), "device_shipment_at"),
					sql.As(t1.C(entdeviceinfo.FieldPosters), "device_posters"),
				)

			t2 := sql.Table(entgood.Table)
			s.
				LeftJoin(t2).
				On(
					s.C(entgood.FieldInheritFromGoodID),
					t2.C(entgood.FieldID),
				).
				AppendSelect(
					sql.As(t2.C(entgood.FieldTitle), "inherit_from_good_name"),
					sql.As(t2.C(entgood.FieldGoodType), "inherit_from_good_type"),
					sql.As(t2.C(entgood.FieldBenefitType), "inherit_from_good_benefit_type"),
				)

			t3 := sql.Table(entvendorlocation.Table)
			s.
				LeftJoin(t3).
				On(
					s.C(entgood.FieldVendorLocationID),
					t3.C(entvendorlocation.FieldID),
				).
				AppendSelect(
					sql.As(t3.C(entvendorlocation.FieldCountry), "vendor_location_country"),
					sql.As(t3.C(entvendorlocation.FieldProvince), "vendor_location_province"),
					sql.As(t3.C(entvendorlocation.FieldCity), "vendor_location_city"),
					sql.As(t3.C(entvendorlocation.FieldAddress), "vendor_location_address"),
				)

			t4 := sql.Table(entextrainfo.Table)
			s.
				LeftJoin(t4).
				On(
					s.C(entgood.FieldID),
					t4.C(entextrainfo.FieldGoodID),
				).
				AppendSelect(
					sql.As(t4.C(entextrainfo.FieldPosters), "posters"),
					sql.As(t4.C(entextrainfo.FieldLabels), "labels"),
					sql.As(t4.C(entextrainfo.FieldVoteCount), "vote_count"),
					sql.As(t4.C(entextrainfo.FieldRating), "rating"),
				)

			t5 := sql.Table(entstock.Table)
			s.
				LeftJoin(t5).
				On(
					s.C(entgood.FieldID),
					t5.C(entstock.FieldGoodID),
				).
				AppendSelect(
					sql.As(t5.C(entstock.FieldID), "good_stock_id"),
					sql.As(t5.C(entstock.FieldTotal), "good_total"),
					sql.As(t5.C(entstock.FieldLocked), "good_locked"),
					sql.As(t5.C(entstock.FieldInService), "good_in_service"),
					sql.As(t5.C(entstock.FieldWaitStart), "good_wait_start"),
					sql.As(t5.C(entstock.FieldSold), "good_sold"),
				)
		})
}

func expand(infos []*npool.Good) ([]*npool.Good, error) { //nolint
	for _, info := range infos {
		_ = json.Unmarshal([]byte(info.LabelsStr), &info.Labels)
		_ = json.Unmarshal([]byte(info.PostersStr), &info.Posters)
		_ = json.Unmarshal([]byte(info.SupportCoinTypeIDsStr), &info.SupportCoinTypeIDs)
		_ = json.Unmarshal([]byte(info.DevicePostersStr), &info.DevicePosters)
		_ = json.Unmarshal([]byte(info.BenefitTIDsStr), &info.BenefitTIDs)

		info.GoodType = mgrpb.GoodType(mgrpb.GoodType_value[info.GoodTypeStr])
		info.BenefitType = mgrpb.BenefitType(mgrpb.BenefitType_value[info.BenefitTypeStr])
		info.BenefitState = mgrpb.BenefitState(mgrpb.BenefitState_value[info.BenefitStateStr])

		if info.InheritFromGoodTypeStr != nil {
			inheritFromGoodType := mgrpb.GoodType(mgrpb.GoodType_value[*info.InheritFromGoodTypeStr])
			info.InheritFromGoodType = &inheritFromGoodType
		}
		if info.InheritFromGoodBenefitTypeStr != nil {
			inheritFromGoodBenefitType := mgrpb.BenefitType(mgrpb.BenefitType_value[*info.InheritFromGoodBenefitTypeStr])
			info.InheritFromGoodBenefitType = &inheritFromGoodBenefitType
		}
	}
	return infos, nil
}
