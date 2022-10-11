//nolint:nolintlint,dupl
package good

import (
	"context"
	"time"

	"github.com/shopspring/decimal"

	constant "github.com/NpoolPlatform/good-middleware/pkg/message/const"
	commontracer "github.com/NpoolPlatform/good-middleware/pkg/tracer"

	"go.opentelemetry.io/otel"
	scodes "go.opentelemetry.io/otel/codes"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/good"

	goodmw "github.com/NpoolPlatform/good-middleware/pkg/good"

	"github.com/google/uuid"
)

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
	return &npool.GetGoodResponse{}, nil
}

func (s *Server) GetGoods(ctx context.Context, in *npool.GetGoodsRequest) (*npool.GetGoodsResponse, error) {
	return &npool.GetGoodsResponse{}, nil
}

func (s *Server) UpdateGood(ctx context.Context, in *npool.UpdateGoodRequest) (*npool.UpdateGoodResponse, error) {
	return &npool.UpdateGoodResponse{}, nil
}
