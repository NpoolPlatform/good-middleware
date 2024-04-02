package recommend

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	recommend1 "github.com/NpoolPlatform/good-middleware/pkg/mw/app/good/recommend"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/app/good/recommend"
)

func (s *Server) ExistRecommendConds(ctx context.Context, in *npool.ExistRecommendCondsRequest) (*npool.ExistRecommendCondsResponse, error) {
	handler, err := recommend1.NewHandler(
		ctx,
		recommend1.WithConds(in.GetConds()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistRecommendConds",
			"In", in,
			"Error", err,
		)
		return &npool.ExistRecommendCondsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	exist, err := handler.ExistRecommendConds(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistRecommendConds",
			"In", in,
			"Error", err,
		)
		return &npool.ExistRecommendCondsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.ExistRecommendCondsResponse{
		Info: exist,
	}, nil
}
