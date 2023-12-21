package score

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	score1 "github.com/NpoolPlatform/good-middleware/pkg/mw/good/score"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/good/score"
)

func (s *Server) GetScore(ctx context.Context, in *npool.GetScoreRequest) (*npool.GetScoreResponse, error) {
	handler, err := score1.NewHandler(
		ctx,
		score1.WithEntID(&in.EntID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetScore",
			"In", in,
			"Error", err,
		)
		return &npool.GetScoreResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.GetScore(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetScore",
			"In", in,
			"Error", err,
		)
		return &npool.GetScoreResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.GetScoreResponse{
		Info: info,
	}, nil
}

func (s *Server) GetScores(ctx context.Context, in *npool.GetScoresRequest) (*npool.GetScoresResponse, error) {
	handler, err := score1.NewHandler(
		ctx,
		score1.WithConds(in.GetConds()),
		score1.WithOffset(in.GetOffset()),
		score1.WithLimit(in.GetLimit()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetScores",
			"In", in,
			"Error", err,
		)
		return &npool.GetScoresResponse{}, status.Error(codes.Aborted, err.Error())
	}

	infos, total, err := handler.GetScores(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetScores",
			"In", in,
			"Error", err,
		)
		return &npool.GetScoresResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.GetScoresResponse{
		Infos: infos,
		Total: total,
	}, nil
}
