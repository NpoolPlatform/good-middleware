package recommend

import (
	"context"

	recommend1 "github.com/NpoolPlatform/good-middleware/pkg/mw/good/recommend"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/good/recommend"
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
		recommend1.WithID(req.ID, false),
		recommend1.WithAppID(req.AppID, true),
		recommend1.WithGoodID(req.GoodID, true),
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

	info, err := handler.CreateRecommend(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateRecommend",
			"In", in,
			"Error", err,
		)
		return &npool.CreateRecommendResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.CreateRecommendResponse{
		Info: info,
	}, nil
}
