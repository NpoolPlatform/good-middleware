package score

import (
	"context"

	score1 "github.com/NpoolPlatform/good-middleware/pkg/mw/good/score"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/good/score"
)

func (s *Server) CreateScore(ctx context.Context, in *npool.CreateScoreRequest) (*npool.CreateScoreResponse, error) {
	req := in.GetInfo()
	if req == nil {
		logger.Sugar().Errorw(
			"CreateScore",
			"In", in,
		)
		return &npool.CreateScoreResponse{}, status.Error(codes.Aborted, "invalid argument")
	}
	handler, err := score1.NewHandler(
		ctx,
		score1.WithID(req.ID, false),
		score1.WithAppID(req.AppID, true),
		score1.WithUserID(req.UserID, true),
		score1.WithAppGoodID(req.AppGoodID, true),
		score1.WithScore(req.Score, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateScore",
			"In", in,
			"Error", err,
		)
		return &npool.CreateScoreResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.CreateScore(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateScore",
			"In", in,
			"Error", err,
		)
		return &npool.CreateScoreResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.CreateScoreResponse{
		Info: info,
	}, nil
}
