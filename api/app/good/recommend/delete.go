package recommend

import (
	"context"

	recommend1 "github.com/NpoolPlatform/good-middleware/pkg/mw/app/good/recommend"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/app/good/recommend"
)

func (s *Server) DeleteRecommend(ctx context.Context, in *npool.DeleteRecommendRequest) (*npool.DeleteRecommendResponse, error) {
	req := in.GetInfo()
	if req == nil {
		logger.Sugar().Errorw(
			"DeleteRecommend",
			"In", in,
		)
		return &npool.DeleteRecommendResponse{}, status.Error(codes.Aborted, "invalid argument")
	}
	handler, err := recommend1.NewHandler(
		ctx,
		recommend1.WithID(req.ID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteRecommend",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteRecommendResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.DeleteRecommend(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteRecommend",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteRecommendResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.DeleteRecommendResponse{
		Info: info,
	}, nil
}
