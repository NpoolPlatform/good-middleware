//nolint:dupl
package recommend

import (
	"context"

	recommend1 "github.com/NpoolPlatform/good-middleware/pkg/mw/app/good/recommend"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/app/good/recommend"
)

func (s *Server) CreateRecommend(ctx context.Context, in *npool.CreateRecommendRequest) (*npool.CreateRecommendResponse, error) {
	req := in.GetInfo()
	if req == nil {
		logger.Sugar().Errorw(
			"CreateRecommend",
			"In", in,
		)
		return &npool.CreateRecommendResponse{}, status.Error(codes.Aborted, "invalid argument")
	}
	handler, err := recommend1.NewHandler(
		ctx,
		recommend1.WithEntID(req.EntID, false),
		recommend1.WithAppGoodID(req.AppGoodID, true),
		recommend1.WithRecommenderID(req.RecommenderID, true),
		recommend1.WithRecommendIndex(req.RecommendIndex, true),
		recommend1.WithMessage(req.Message, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateRecommend",
			"In", in,
			"Error", err,
		)
		return &npool.CreateRecommendResponse{}, status.Error(codes.Aborted, err.Error())
	}

	if err := handler.CreateRecommend(ctx); err != nil {
		logger.Sugar().Errorw(
			"CreateRecommend",
			"In", in,
			"Error", err,
		)
		return &npool.CreateRecommendResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.CreateRecommendResponse{}, nil
}
