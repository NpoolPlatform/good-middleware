//nolint:nolintlint,dupl
package appgood

import (
	"context"

	"github.com/shopspring/decimal"

	constant1 "github.com/NpoolPlatform/good-middleware/pkg/const"
	constant "github.com/NpoolPlatform/good-middleware/pkg/message/const"
	commontracer "github.com/NpoolPlatform/good-middleware/pkg/tracer"

	"go.opentelemetry.io/otel"
	scodes "go.opentelemetry.io/otel/codes"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/appgood"

	goodmgrcli "github.com/NpoolPlatform/good-manager/pkg/client/good"
	goodmw "github.com/NpoolPlatform/good-middleware/pkg/appgood"

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

	if in.GetInfo().ID != nil {
		if _, err := uuid.Parse(in.GetInfo().GetID()); err != nil {
			logger.Sugar().Errorw("CreateGood", "ID", in.GetInfo().GetID(), "error", err)
			return &npool.CreateGoodResponse{}, status.Error(codes.InvalidArgument, err.Error())
		}
	}

	if _, err := uuid.Parse(in.GetInfo().GetAppID()); err != nil {
		logger.Sugar().Errorw("CreateGood", "AppID", in.GetInfo().GetAppID(), "error", err)
		return &npool.CreateGoodResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	if price, err := decimal.NewFromString(in.GetInfo().GetPrice()); err != nil || price.Cmp(decimal.NewFromInt(0)) <= 0 {
		logger.Sugar().Errorw("CreateGood", "Price", in.GetInfo().GetPrice(), "error", err)
		return &npool.CreateGoodResponse{}, status.Error(codes.InvalidArgument, "Price is invalid")
	}

	if in.GetInfo().GetGoodName() == "" {
		logger.Sugar().Errorw("CreateGood", "GoodName", in.GetInfo().GetGoodName())
		return &npool.CreateGoodResponse{}, status.Error(codes.InvalidArgument, "GoodName is invalid")
	}

	if in.GetInfo().DisplayIndex != nil {
		if in.GetInfo().GetDisplayIndex() < 0 {
			logger.Sugar().Errorw("CreateGood", "DisplayIndex", in.GetInfo().GetDisplayIndex())
			return &npool.CreateGoodResponse{}, status.Error(codes.InvalidArgument, "DisplayIndex is invalid")
		}
	}

	if in.GetInfo().PurchaseLimit != nil {
		if in.GetInfo().GetPurchaseLimit() <= 0 {
			logger.Sugar().Errorw("CreateGood", "PurchaseLimit", in.GetInfo().GetPurchaseLimit())
			return &npool.CreateGoodResponse{}, status.Error(codes.InvalidArgument, "PurchaseLimit is invalid")
		}
	}

	if in.GetInfo().CommissionPercent != nil {
		if in.GetInfo().GetCommissionPercent() >= 100 || in.GetInfo().GetCommissionPercent() < 0 {
			logger.Sugar().Errorw("CreateGood", "CommissionPercent", in.GetInfo().GetCommissionPercent())
			return &npool.CreateGoodResponse{}, status.Error(codes.InvalidArgument, "CommissionPercent is invalid")
		}
	}

	exist, err := goodmgrcli.ExistGood(ctx, in.GetInfo().GetGoodID())
	if err != nil {
		logger.Sugar().Errorw("CreateGood", "GoodID", in.GetInfo().GetGoodID(), "error", err)
		return &npool.CreateGoodResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}
	if !exist {
		logger.Sugar().Errorw("CreateGood", "GoodID", in.GetInfo().GetGoodID(), "exist", exist)
		return &npool.CreateGoodResponse{}, status.Error(codes.InvalidArgument, "GoodID not exist")
	}

	span = commontracer.TraceInvoker(span, "Good", "mw", "CreateGood")

	info, err := goodmw.CreateGood(ctx, in.GetInfo())
	if err != nil {
		logger.Sugar().Errorw("CreateGood", "Good", in.GetInfo())
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

	if in.GetConds().AppID != nil {
		if _, err := uuid.Parse(in.GetConds().GetAppID().GetValue()); err != nil {
			logger.Sugar().Errorw("GetGoods", "AppID", in.GetConds().GetAppID().GetValue(), "error", err)
			return &npool.GetGoodsResponse{}, status.Error(codes.InvalidArgument, err.Error())
		}
	}

	if in.GetConds().GoodID != nil {
		if _, err := uuid.Parse(in.GetConds().GetGoodID().GetValue()); err != nil {
			logger.Sugar().Errorw("GetGoods", "GoodID", in.GetConds().GetGoodID().GetValue(), "error", err)
			return &npool.GetGoodsResponse{}, status.Error(codes.InvalidArgument, err.Error())
		}

		exist, err := goodmgrcli.ExistGood(ctx, in.GetConds().GetGoodID().GetValue())
		if err != nil {
			logger.Sugar().Errorw("GetGoods", "GoodID", in.GetConds().GetGoodID().GetValue(), "error", err)
			return &npool.GetGoodsResponse{}, status.Error(codes.InvalidArgument, err.Error())
		}
		if !exist {
			logger.Sugar().Errorw("GetGoods", "GoodID", in.GetConds().GetGoodID().GetValue(), "exist", exist)
			return &npool.GetGoodsResponse{}, status.Error(codes.InvalidArgument, "GoodID not exist")
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

	if in.GetInfo().AppID != nil {
		if _, err := uuid.Parse(in.GetInfo().GetAppID()); err != nil {
			logger.Sugar().Errorw("UpdateGood", "AppID", in.GetInfo().GetAppID(), "error", err)
			return &npool.UpdateGoodResponse{}, status.Error(codes.InvalidArgument, err.Error())
		}
	}

	if in.GetInfo().Price != nil {
		if price, err := decimal.NewFromString(in.GetInfo().GetPrice()); err != nil || price.Cmp(decimal.NewFromInt(0)) <= 0 {
			logger.Sugar().Errorw("UpdateGood", "Price", in.GetInfo().GetPrice(), "error", err)
			return &npool.UpdateGoodResponse{}, status.Error(codes.InvalidArgument, "Price is invalid")
		}
	}

	if in.GetInfo().GoodName != nil {
		if in.GetInfo().GetGoodName() == "" {
			logger.Sugar().Errorw("UpdateGood", "GoodName", in.GetInfo().GetGoodName())
			return &npool.UpdateGoodResponse{}, status.Error(codes.InvalidArgument, "GoodName is invalid")
		}
	}

	if in.GetInfo().DisplayIndex != nil {
		if in.GetInfo().GetDisplayIndex() < 0 {
			logger.Sugar().Errorw("UpdateGood", "DisplayIndex", in.GetInfo().GetDisplayIndex())
			return &npool.UpdateGoodResponse{}, status.Error(codes.InvalidArgument, "DisplayIndex is invalid")
		}
	}

	if in.GetInfo().PurchaseLimit != nil {
		if in.GetInfo().GetPurchaseLimit() <= 0 {
			logger.Sugar().Errorw("UpdateGood", "PurchaseLimit", in.GetInfo().GetPurchaseLimit())
			return &npool.UpdateGoodResponse{}, status.Error(codes.InvalidArgument, "PurchaseLimit is invalid")
		}
	}

	if in.GetInfo().CommissionPercent != nil {
		if in.GetInfo().GetCommissionPercent() >= 100 || in.GetInfo().GetCommissionPercent() < 0 {
			logger.Sugar().Errorw("UpdateGood", "CommissionPercent", in.GetInfo().GetCommissionPercent())
			return &npool.UpdateGoodResponse{}, status.Error(codes.InvalidArgument, "CommissionPercent is invalid")
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
