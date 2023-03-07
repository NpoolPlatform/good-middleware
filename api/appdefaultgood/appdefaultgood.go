//nolint:nolintlint,dupl
package appdefaultgood

import (
	"context"

	constant1 "github.com/NpoolPlatform/good-middleware/pkg/const"
	constant "github.com/NpoolPlatform/good-middleware/pkg/message/const"
	commontracer "github.com/NpoolPlatform/good-middleware/pkg/tracer"

	"go.opentelemetry.io/otel"
	scodes "go.opentelemetry.io/otel/codes"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/appdefaultgood"

	"github.com/google/uuid"

	mgrcli "github.com/NpoolPlatform/good-manager/pkg/client/appdefaultgood"
)

func (s *Server) GetAppDefaultGood(
	ctx context.Context,
	in *npool.GetAppDefaultGoodRequest,
) (
	*npool.GetAppDefaultGoodResponse,
	error,
) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetAppDefaultGood")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	if _, err := uuid.Parse(in.GetID()); err != nil {
		logger.Sugar().Errorw("GetAppDefaultGood", "ID", in.GetID(), "error", err)
		return &npool.GetAppDefaultGoodResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "AppDefaultGood", "mw", "GetAppDefaultGood")

	info, err := mgrcli.GetAppDefaultGood(ctx, in.GetID())
	if err != nil {
		logger.Sugar().Errorw("GetAppDefaultGood", "ID", in.GetID(), "error", err)
		return &npool.GetAppDefaultGoodResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetAppDefaultGoodResponse{
		Info: info,
	}, nil
}

func (s *Server) GetAppDefaultGoodOnly(
	ctx context.Context,
	in *npool.GetAppDefaultGoodOnlyRequest,
) (
	*npool.GetAppDefaultGoodOnlyResponse,
	error,
) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetAppDefaultGood")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	if in.Conds == nil {
		logger.Sugar().Errorw("GetAppDefaultGoods", "error", err)
		return &npool.GetAppDefaultGoodOnlyResponse{}, status.Error(codes.InvalidArgument, "conds is empty")
	}

	if in.GetConds().ID != nil {
		if _, err := uuid.Parse(in.GetConds().GetID().GetValue()); err != nil {
			logger.Sugar().Errorw("GetAppDefaultGoods", "ID", in.GetConds().GetID().GetValue(), "error", err)
			return &npool.GetAppDefaultGoodOnlyResponse{}, status.Error(codes.InvalidArgument, err.Error())
		}
	}

	if in.GetConds().AppID != nil {
		if _, err := uuid.Parse(in.GetConds().GetAppID().GetValue()); err != nil {
			logger.Sugar().Errorw("GetAppDefaultGoods", "AppID", in.GetConds().GetAppID().GetValue(), "error", err)
			return &npool.GetAppDefaultGoodOnlyResponse{}, status.Error(codes.InvalidArgument, err.Error())
		}
	}

	if in.GetConds().GoodID != nil {
		if _, err := uuid.Parse(in.GetConds().GetGoodID().GetValue()); err != nil {
			logger.Sugar().Errorw("GetAppDefaultGoods", "GoodID", in.GetConds().GetGoodID().GetValue(), "error", err)
			return &npool.GetAppDefaultGoodOnlyResponse{}, status.Error(codes.InvalidArgument, err.Error())
		}
	}

	if in.GetConds().CoinTypeID != nil {
		if _, err := uuid.Parse(in.GetConds().GetCoinTypeID().GetValue()); err != nil {
			logger.Sugar().Errorw("GetAppDefaultGoods", "CoinTypeID", in.GetConds().GetCoinTypeID().GetValue(), "error", err)
			return &npool.GetAppDefaultGoodOnlyResponse{}, status.Error(codes.InvalidArgument, err.Error())
		}
	}

	span = commontracer.TraceInvoker(span, "AppDefaultGood", "mw", "GetAppDefaultGood")

	info, err := mgrcli.GetAppDefaultGoodOnly(ctx, in.Conds)
	if err != nil {
		logger.Sugar().Errorw("GetAppDefaultGood", "error", err)
		return &npool.GetAppDefaultGoodOnlyResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetAppDefaultGoodOnlyResponse{
		Info: info,
	}, nil
}

func (s *Server) GetAppDefaultGoods(
	ctx context.Context,
	in *npool.GetAppDefaultGoodsRequest,
) (
	*npool.GetAppDefaultGoodsResponse,
	error,
) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetAppDefaultGoods")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	if in.Conds == nil {
		logger.Sugar().Errorw("GetAppDefaultGoods", "error", err)
		return &npool.GetAppDefaultGoodsResponse{}, status.Error(codes.InvalidArgument, "conds is empty")
	}

	if in.GetConds().ID != nil {
		if _, err := uuid.Parse(in.GetConds().GetID().GetValue()); err != nil {
			logger.Sugar().Errorw("GetAppDefaultGoods", "ID", in.GetConds().GetID().GetValue(), "error", err)
			return &npool.GetAppDefaultGoodsResponse{}, status.Error(codes.InvalidArgument, err.Error())
		}
	}

	if in.GetConds().AppID != nil {
		if _, err := uuid.Parse(in.GetConds().GetAppID().GetValue()); err != nil {
			logger.Sugar().Errorw("GetAppDefaultGoods", "AppID", in.GetConds().GetAppID().GetValue(), "error", err)
			return &npool.GetAppDefaultGoodsResponse{}, status.Error(codes.InvalidArgument, err.Error())
		}
	}

	if in.GetConds().GoodID != nil {
		if _, err := uuid.Parse(in.GetConds().GetGoodID().GetValue()); err != nil {
			logger.Sugar().Errorw("GetAppDefaultGoods", "GoodID", in.GetConds().GetGoodID().GetValue(), "error", err)
			return &npool.GetAppDefaultGoodsResponse{}, status.Error(codes.InvalidArgument, err.Error())
		}
	}

	if in.GetConds().CoinTypeID != nil {
		if _, err := uuid.Parse(in.GetConds().GetCoinTypeID().GetValue()); err != nil {
			logger.Sugar().Errorw("GetAppDefaultGoods", "CoinTypeID", in.GetConds().GetCoinTypeID().GetValue(), "error", err)
			return &npool.GetAppDefaultGoodsResponse{}, status.Error(codes.InvalidArgument, err.Error())
		}
	}

	span = commontracer.TraceInvoker(span, "AppDefaultGood", "mw", "GetAppDefaultGoods")

	limit := in.GetLimit()
	if limit <= 0 {
		limit = constant1.DefaultLimit
	}

	infos, total, err := mgrcli.GetAppDefaultGoods(ctx, in.GetConds(), in.GetOffset(), limit)
	if err != nil {
		logger.Sugar().Errorw("GetAppDefaultGoods", "Conds", in.GetConds(), "error", err)
		return &npool.GetAppDefaultGoodsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetAppDefaultGoodsResponse{
		Infos: infos,
		Total: total,
	}, nil
}
