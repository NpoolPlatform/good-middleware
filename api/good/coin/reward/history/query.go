package history

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	history1 "github.com/NpoolPlatform/good-middleware/pkg/mw/good/coin/reward/history"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/good/coin/reward/history"
)

func (s *Server) GetHistories(ctx context.Context, in *npool.GetHistoriesRequest) (*npool.GetHistoriesResponse, error) {
	handler, err := history1.NewHandler(
		ctx,
		history1.WithConds(in.GetConds()),
		history1.WithOffset(in.GetOffset()),
		history1.WithLimit(in.GetLimit()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetHistories",
			"In", in,
			"Error", err,
		)
		return &npool.GetHistoriesResponse{}, status.Error(codes.Aborted, err.Error())
	}

	infos, total, err := handler.GetHistories(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetHistories",
			"In", in,
			"Error", err,
		)
		return &npool.GetHistoriesResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.GetHistoriesResponse{
		Infos: infos,
		Total: total,
	}, nil
}
