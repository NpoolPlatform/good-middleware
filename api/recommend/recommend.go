//nolint:nolintlint,dupl
package recommend

import (
	"context"
	"fmt"

	constant "github.com/NpoolPlatform/good-middleware/pkg/message/const"
	commontracer "github.com/NpoolPlatform/good-middleware/pkg/tracer"

	"go.opentelemetry.io/otel"
	scodes "go.opentelemetry.io/otel/codes"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/recommend"

	recommendmw "github.com/NpoolPlatform/good-middleware/pkg/recommend"

	"github.com/google/uuid"
)

// nolint
func (s *Server) CreateRecommend(ctx context.Context, in *npool.CreateRecommendRequest) (*npool.CreateRecommendResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CreateRecommend")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	if _, err = uuid.Parse(in.GetInfo().GetAppID()); err != nil {
		logger.Sugar().Errorw("validate", "TargetAppID", in.GetInfo().GetAppID(), "error", err)
		return &npool.CreateRecommendResponse{}, status.Error(codes.InvalidArgument, fmt.Sprintf("AppID is invalid: %v", err))
	}

	if _, err := uuid.Parse(in.GetInfo().GetGoodID()); err != nil {
		logger.Sugar().Errorw("validate", "GoodID", in.GetInfo().GetGoodID(), "error", err)
		return &npool.CreateRecommendResponse{}, status.Error(codes.InvalidArgument, fmt.Sprintf("GoodID is invalid: %v", err))
	}

	if _, err := uuid.Parse(in.GetInfo().GetRecommenderID()); err != nil {
		logger.Sugar().Errorw("validate", "RecommenderID", in.GetInfo().GetRecommenderID(), "error", err)
		return &npool.CreateRecommendResponse{}, status.Error(codes.InvalidArgument, fmt.Sprintf("RecommenderID is invalid: %v", err))
	}

	if in.GetInfo().GetMessage() == "" {
		logger.Sugar().Errorw("validate", "Message", in.GetInfo().GetMessage(), "error", err)
		return &npool.CreateRecommendResponse{}, status.Error(codes.InvalidArgument, fmt.Sprintf("Message is empty"))
	}

	span = commontracer.TraceInvoker(span, "Recommend", "mw", "CreateRecommend")

	info, err := recommendmw.CreateRecommend(ctx, in.GetInfo())
	if err != nil {
		logger.Sugar().Errorw("CreateRecommend", "Recommend", in.GetInfo(), "error", err)
		return &npool.CreateRecommendResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateRecommendResponse{
		Info: info,
	}, nil
}
