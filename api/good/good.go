//nolint:nolintlint,dupl
package good

import (
	"context"
	"time"

	"github.com/shopspring/decimal"

	constant1 "github.com/NpoolPlatform/good-middleware/pkg/const"
	constant "github.com/NpoolPlatform/good-middleware/pkg/message/const"
	commontracer "github.com/NpoolPlatform/good-middleware/pkg/tracer"

	"go.opentelemetry.io/otel"
	scodes "go.opentelemetry.io/otel/codes"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	mgrpb "github.com/NpoolPlatform/message/npool/good/mgr/v1/good"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/good"

	goodmw "github.com/NpoolPlatform/good-middleware/pkg/good"

	"github.com/google/uuid"
)

// nolint
func (s *Server) CreateGood(ctx context.Context, in *npool.CreateGoodRequest) (*npool.CreateGoodResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CreateGood")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	// TODO: Check if device exist
	// TODO: Check inherit from good exist
	// TODO: Check vendor location exist

	if _, err := uuid.Parse(in.GetInfo().GetCoinTypeID()); err != nil {
		logger.Sugar().Errorw("CreateGood", "CoinTypeID", in.GetInfo().GetCoinTypeID(), "error", err)
		return &npool.CreateGoodResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	if in.GetInfo().GetDurationDays() <= 0 {
		logger.Sugar().Errorw("CreateGood", "DurationDays", in.GetInfo().GetDurationDays())
		return &npool.CreateGoodResponse{}, status.Error(codes.InvalidArgument, "DurationDays is invalid")
	}

	if price, err := decimal.NewFromString(in.GetInfo().GetPrice()); err != nil || price.Cmp(decimal.NewFromInt(0)) <= 0 {
		logger.Sugar().Errorw("CreateGood", "Price", in.GetInfo().GetPrice(), "error", err)
		return &npool.CreateGoodResponse{}, status.Error(codes.InvalidArgument, "Price is invalid")
	}

	switch in.GetInfo().GetBenefitType() {
	case mgrpb.BenefitType_BenefitTypePlatform:
	case mgrpb.BenefitType_BenefitTypePool:
	default:
		logger.Sugar().Errorw("CreateGood", "BenefitType", in.GetInfo().GetBenefitType())
		return &npool.CreateGoodResponse{}, status.Error(codes.InvalidArgument, "BenefitType is invalid")
	}

	switch in.GetInfo().GetGoodType() {
	case mgrpb.GoodType_GoodTypeClassicMining:
	case mgrpb.GoodType_GoodTypeUnionMining:
	case mgrpb.GoodType_GoodTypeTechniqueFee:
	case mgrpb.GoodType_GoodTypeElectricityFee:
	default:
		logger.Sugar().Errorw("CreateGood", "GoodType", in.GetInfo().GetGoodType())
		return &npool.CreateGoodResponse{}, status.Error(codes.InvalidArgument, "GoodType is invalid")
	}

	if in.GetInfo().GetTitle() == "" {
		logger.Sugar().Errorw("CreateGood", "Title", in.GetInfo().GetTitle())
		return &npool.CreateGoodResponse{}, status.Error(codes.InvalidArgument, "Title is invalid")
	}

	if in.GetInfo().GetUnitAmount() <= 0 {
		logger.Sugar().Errorw("CreateGood", "UnitAmount", in.GetInfo().GetUnitAmount())
		return &npool.CreateGoodResponse{}, status.Error(codes.InvalidArgument, "UnitAmount is invalid")
	}

	for _, coinTypeID := range in.GetInfo().GetSupportCoinTypeIDs() {
		if _, err := uuid.Parse(coinTypeID); err != nil {
			logger.Sugar().Errorw("CreateGood", "SupportCoinTypeIDs", in.GetInfo().GetSupportCoinTypeIDs(), "error", err)
			return &npool.CreateGoodResponse{}, status.Error(codes.InvalidArgument, err.Error())
		}
	}

	now := uint32(time.Now().Unix())
	if in.GetInfo().GetDeliveryAt() <= now {
		logger.Sugar().Errorw("CreateGood", "DeliveryAt", in.GetInfo().GetDeliveryAt(), "now", now)
		return &npool.CreateGoodResponse{}, status.Error(codes.InvalidArgument, "DeliveryAt is invalid")
	}

	if in.GetInfo().GetStartAt() <= now {
		logger.Sugar().Errorw("CreateGood", "StartAt", in.GetInfo().GetStartAt(), "now", now)
		return &npool.CreateGoodResponse{}, status.Error(codes.InvalidArgument, "StartAt is invalid")
	}

	if in.GetInfo().GetTotal() <= 0 {
		logger.Sugar().Errorw("CreateGood", "Total", in.GetInfo().GetTotal())
		return &npool.CreateGoodResponse{}, status.Error(codes.InvalidArgument, "Total is invalid")
	}

	span = commontracer.TraceInvoker(span, "Good", "mw", "CreateGood")

	info, err := goodmw.CreateGood(ctx, in.GetInfo())
	if err != nil {
		logger.Sugar().Errorw("CreateGood", "Good", in.GetInfo(), "error", err)
		return &npool.CreateGoodResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateGoodResponse{
		Info: info,
	}, nil
}

func (s *Server) GetGood(ctx context.Context, in *npool.GetGoodRequest) (*npool.GetGoodResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetGood")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	if _, err := uuid.Parse(in.GetID()); err != nil {
		logger.Sugar().Errorw("GetGood", "ID", in.GetID(), "error", err)
		return &npool.GetGoodResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "Good", "mw", "GetGood")

	info, err := goodmw.GetGood(ctx, in.GetID())
	if err != nil {
		logger.Sugar().Errorw("GetGood", "ID", in.GetID(), "error", err)
		return &npool.GetGoodResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetGoodResponse{
		Info: info,
	}, nil
}

// nolint
func (s *Server) GetGoods(ctx context.Context, in *npool.GetGoodsRequest) (*npool.GetGoodsResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetGoods")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	if in.GetConds().ID != nil {
		if _, err := uuid.Parse(in.GetConds().GetID().GetValue()); err != nil {
			logger.Sugar().Errorw("GetGoods", "ID", in.GetConds().GetID().GetValue(), "error", err)
			return &npool.GetGoodsResponse{}, status.Error(codes.InvalidArgument, err.Error())
		}
	}

	if in.GetConds().DeviceInfoID != nil {
		if _, err := uuid.Parse(in.GetConds().GetDeviceInfoID().GetValue()); err != nil {
			logger.Sugar().Errorw("GetGoods", "DeviceInfoID", in.GetConds().GetDeviceInfoID().GetValue(), "error", err)
			return &npool.GetGoodsResponse{}, status.Error(codes.InvalidArgument, err.Error())
		}
	}

	if in.GetConds().CoinTypeID != nil {
		if _, err := uuid.Parse(in.GetConds().GetCoinTypeID().GetValue()); err != nil {
			logger.Sugar().Errorw("GetGoods", "CoinTypeID", in.GetConds().GetCoinTypeID().GetValue(), "error", err)
			return &npool.GetGoodsResponse{}, status.Error(codes.InvalidArgument, err.Error())
		}
	}

	if in.GetConds().VendorLocationID != nil {
		if _, err := uuid.Parse(in.GetConds().GetVendorLocationID().GetValue()); err != nil {
			logger.Sugar().Errorw("GetGoods", "VendorLocationID", in.GetConds().GetVendorLocationID().GetValue(), "error", err)
			return &npool.GetGoodsResponse{}, status.Error(codes.InvalidArgument, err.Error())
		}
	}

	if in.GetConds().BenefitType != nil {
		switch mgrpb.BenefitType(in.GetConds().GetBenefitType().GetValue()) {
		case mgrpb.BenefitType_BenefitTypePlatform:
		case mgrpb.BenefitType_BenefitTypePool:
		default:
			logger.Sugar().Errorw("GetGoods", "BenefitType", in.GetConds().GetBenefitType().GetValue())
			return &npool.GetGoodsResponse{}, status.Error(codes.InvalidArgument, "BenefitType is invalid")
		}
	}

	if in.GetConds().GoodType != nil {
		switch mgrpb.GoodType(in.GetConds().GetGoodType().GetValue()) {
		case mgrpb.GoodType_GoodTypeClassicMining:
		case mgrpb.GoodType_GoodTypeUnionMining:
		case mgrpb.GoodType_GoodTypeTechniqueFee:
		case mgrpb.GoodType_GoodTypeElectricityFee:
		default:
			logger.Sugar().Errorw("GetGoods", "GoodType", in.GetConds().GetGoodType().GetValue())
			return &npool.GetGoodsResponse{}, status.Error(codes.InvalidArgument, "GoodType is invalid")
		}
	}

	span = commontracer.TraceInvoker(span, "Good", "mw", "GetGoods")

	limit := in.GetLimit()
	if limit <= 0 {
		limit = constant1.DefaultLimit
	}

	infos, total, err := goodmw.GetGoods(ctx, in.GetConds(), in.GetOffset(), limit)
	if err != nil {
		logger.Sugar().Errorw("GetGoods", "Conds", in.GetConds(), "error", err)
		return &npool.GetGoodsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetGoodsResponse{
		Infos: infos,
		Total: total,
	}, nil
}

func (s *Server) GetManyGoods(ctx context.Context, in *npool.GetManyGoodsRequest) (*npool.GetManyGoodsResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetManyGoods")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	for _, id := range in.GetIDs() {
		if _, err := uuid.Parse(id); err != nil {
			logger.Sugar().Errorw("GetManyGoods", "IDs", in.GetIDs(), "error", err)
			return &npool.GetManyGoodsResponse{}, status.Error(codes.InvalidArgument, err.Error())
		}
	}

	span = commontracer.TraceInvoker(span, "Good", "mw", "GetManyGoods")

	limit := in.GetLimit()
	if limit <= 0 {
		limit = constant1.DefaultLimit
	}

	infos, total, err := goodmw.GetManyGoods(ctx, in.GetIDs(), in.GetOffset(), limit)
	if err != nil {
		logger.Sugar().Errorw("GetManyGoods", "ID", in.GetIDs(), "error", err)
		return &npool.GetManyGoodsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetManyGoodsResponse{
		Infos: infos,
		Total: total,
	}, nil
}

// nolint
func (s *Server) UpdateGood(ctx context.Context, in *npool.UpdateGoodRequest) (*npool.UpdateGoodResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "UpdateGood")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	if _, err := uuid.Parse(in.GetInfo().GetID()); err != nil {
		logger.Sugar().Errorw("UpdateGood", "ID", in.GetInfo().GetID(), "error", err)
		return &npool.UpdateGoodResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	if in.GetInfo().DeviceInfoID != nil {
		if _, err := uuid.Parse(in.GetInfo().GetDeviceInfoID()); err != nil {
			logger.Sugar().Errorw("UpdateGood", "DeviceInfoID", in.GetInfo().GetDeviceInfoID(), "error", err)
			return &npool.UpdateGoodResponse{}, status.Error(codes.InvalidArgument, err.Error())
		}
	}

	if in.GetInfo().DurationDays != nil && in.GetInfo().GetDurationDays() <= 0 {
		logger.Sugar().Errorw("UpdateGood", "DurationDays", in.GetInfo().GetDurationDays())
		return &npool.UpdateGoodResponse{}, status.Error(codes.InvalidArgument, "DurationDays is invalid")
	}

	if in.GetInfo().CoinTypeID != nil {
		if _, err := uuid.Parse(in.GetInfo().GetCoinTypeID()); err != nil {
			logger.Sugar().Errorw("UpdateGood", "CoinTypeID", in.GetInfo().GetCoinTypeID(), "error", err)
			return &npool.UpdateGoodResponse{}, status.Error(codes.InvalidArgument, err.Error())
		}
	}

	if in.GetInfo().InheritFromGoodID != nil {
		if _, err := uuid.Parse(in.GetInfo().GetInheritFromGoodID()); err != nil {
			logger.Sugar().Errorw("UpdateGood", "InheritFromGoodID", in.GetInfo().GetInheritFromGoodID(), "error", err)
			return &npool.UpdateGoodResponse{}, status.Error(codes.InvalidArgument, err.Error())
		}
	}

	if in.GetInfo().VendorLocationID != nil {
		if _, err := uuid.Parse(in.GetInfo().GetVendorLocationID()); err != nil {
			logger.Sugar().Errorw("UpdateGood", "VendorLocationID", in.GetInfo().GetVendorLocationID(), "error", err)
			return &npool.UpdateGoodResponse{}, status.Error(codes.InvalidArgument, err.Error())
		}
	}

	if in.GetInfo().Price != nil {
		if price, err := decimal.NewFromString(in.GetInfo().GetPrice()); err != nil || price.Cmp(decimal.NewFromInt(0)) <= 0 {
			logger.Sugar().Errorw("UpdateGood", "Price", in.GetInfo().GetPrice(), "error", err)
			return &npool.UpdateGoodResponse{}, status.Error(codes.InvalidArgument, "Price is invalid")
		}
	}

	if in.GetInfo().BenefitType != nil {
		switch in.GetInfo().GetBenefitType() {
		case mgrpb.BenefitType_BenefitTypePlatform:
		case mgrpb.BenefitType_BenefitTypePool:
		default:
			logger.Sugar().Errorw("UpdateGood", "BenefitType", in.GetInfo().GetBenefitType())
			return &npool.UpdateGoodResponse{}, status.Error(codes.InvalidArgument, "BenefitType is invalid")
		}
	}

	if in.GetInfo().GoodType != nil {
		switch in.GetInfo().GetGoodType() {
		case mgrpb.GoodType_GoodTypeClassicMining:
		case mgrpb.GoodType_GoodTypeUnionMining:
		case mgrpb.GoodType_GoodTypeTechniqueFee:
		case mgrpb.GoodType_GoodTypeElectricityFee:
		default:
			logger.Sugar().Errorw("UpdateGood", "GoodType", in.GetInfo().GetGoodType())
			return &npool.UpdateGoodResponse{}, status.Error(codes.InvalidArgument, "GoodType is invalid")
		}
	}

	if in.GetInfo().Title != nil {
		if in.GetInfo().GetTitle() == "" {
			logger.Sugar().Errorw("UpdateGood", "Title", in.GetInfo().GetTitle())
			return &npool.UpdateGoodResponse{}, status.Error(codes.InvalidArgument, "Title is invalid")
		}
	}

	if in.GetInfo().UnitAmount != nil {
		if in.GetInfo().GetUnitAmount() <= 0 {
			logger.Sugar().Errorw("UpdateGood", "UnitAmount", in.GetInfo().GetUnitAmount())
			return &npool.UpdateGoodResponse{}, status.Error(codes.InvalidArgument, "UnitAmount is invalid")
		}
	}

	for _, coinTypeID := range in.GetInfo().GetSupportCoinTypeIDs() {
		if _, err := uuid.Parse(coinTypeID); err != nil {
			logger.Sugar().Errorw("UpdateGood", "SupportCoinTypeIDs", in.GetInfo().GetSupportCoinTypeIDs(), "error", err)
			return &npool.UpdateGoodResponse{}, status.Error(codes.InvalidArgument, err.Error())
		}
	}

	now := uint32(time.Now().Unix())

	if in.GetInfo().DeliveryAt != nil {
		if in.GetInfo().GetDeliveryAt() <= now {
			logger.Sugar().Errorw("UpdateGood", "DeliveryAt", in.GetInfo().GetDeliveryAt(), "now", now)
			return &npool.UpdateGoodResponse{}, status.Error(codes.InvalidArgument, "DeliveryAt is invalid")
		}
	}

	if in.GetInfo().StartAt != nil {
		if in.GetInfo().GetStartAt() <= now {
			logger.Sugar().Errorw("UpdateGood", "StartAt", in.GetInfo().GetStartAt(), "now", now)
			return &npool.UpdateGoodResponse{}, status.Error(codes.InvalidArgument, "StartAt is invalid")
		}
	}

	if in.GetInfo().Total != nil {
		if in.GetInfo().GetTotal() <= 0 {
			logger.Sugar().Errorw("UpdateGood", "Total", in.GetInfo().GetTotal())
			return &npool.UpdateGoodResponse{}, status.Error(codes.InvalidArgument, "Total is invalid")
		}
	}

	if in.GetInfo().Sold != nil {
		if in.GetInfo().GetSold() <= 0 {
			logger.Sugar().Errorw("UpdateGood", "Sold", in.GetInfo().GetSold())
			return &npool.UpdateGoodResponse{}, status.Error(codes.InvalidArgument, "Sold is invalid")
		}
	}

	span = commontracer.TraceInvoker(span, "Good", "mw", "UpdateGood")

	info, err := goodmw.UpdateGood(ctx, in.GetInfo())
	if err != nil {
		logger.Sugar().Errorw("UpdateGood", "Good", in.GetInfo())
		return &npool.UpdateGoodResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.UpdateGoodResponse{
		Info: info,
	}, nil
}
