package score

import (
	"context"

	score1 "github.com/NpoolPlatform/good-middleware/pkg/mw/app/good/score"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/app/good/score"
)

func (s *Server) UpdateScore(ctx context.Context, in *npool.UpdateScoreRequest) (*npool.UpdateScoreResponse, error) {
	req := in.GetInfo()
	if req == nil {
		logger.Sugar().Errorw(
			"UpdateScore",
			"In", in,
		)
		return &npool.UpdateScoreResponse{}, status.Error(codes.Aborted, "invalid argument")
	}
	handler, err := score1.NewHandler(
		ctx,
		score1.WithID(req.ID, false),
		score1.WithEntID(req.EntID, false),
		score1.WithScore(req.Score, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateScore",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateScoreResponse{}, status.Error(codes.Aborted, err.Error())
	}

	if err := handler.UpdateScore(ctx); err != nil {
		logger.Sugar().Errorw(
			"UpdateScore",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateScoreResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.UpdateScoreResponse{}, nil
}
