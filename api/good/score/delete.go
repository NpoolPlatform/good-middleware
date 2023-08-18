package score

import (
	"context"

	score1 "github.com/NpoolPlatform/good-middleware/pkg/mw/good/score"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/good/score"
)

func (s *Server) DeleteScore(ctx context.Context, in *npool.DeleteScoreRequest) (*npool.DeleteScoreResponse, error) {
	req := in.GetInfo()
	if req == nil {
		logger.Sugar().Errorw(
			"DeleteScore",
			"In", in,
		)
		return &npool.DeleteScoreResponse{}, status.Error(codes.Aborted, "invalid argument")
	}
	handler, err := score1.NewHandler(
		ctx,
		score1.WithID(req.ID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteScore",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteScoreResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.DeleteScore(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteScore",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteScoreResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.DeleteScoreResponse{
		Info: info,
	}, nil
}
