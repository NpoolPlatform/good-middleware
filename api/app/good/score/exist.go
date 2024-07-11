package score

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	score1 "github.com/NpoolPlatform/good-middleware/pkg/mw/app/good/score"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/app/good/score"
)

func (s *Server) ExistScoreConds(ctx context.Context, in *npool.ExistScoreCondsRequest) (*npool.ExistScoreCondsResponse, error) {
	handler, err := score1.NewHandler(
		ctx,
		score1.WithConds(in.GetConds()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistScoreConds",
			"In", in,
			"Error", err,
		)
		return &npool.ExistScoreCondsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	exist, err := handler.ExistScoreConds(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistScoreConds",
			"In", in,
			"Error", err,
		)
		return &npool.ExistScoreCondsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.ExistScoreCondsResponse{
		Info: exist,
	}, nil
}
