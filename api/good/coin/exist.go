package coin

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	coin1 "github.com/NpoolPlatform/good-middleware/pkg/mw/good/coin"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/good/coin"
)

func (s *Server) ExistGoodCoinConds(ctx context.Context, in *npool.ExistGoodCoinCondsRequest) (*npool.ExistGoodCoinCondsResponse, error) {
	handler, err := coin1.NewHandler(
		ctx,
		coin1.WithConds(in.GetConds()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistGoodCoinConds",
			"In", in,
			"Error", err,
		)
		return &npool.ExistGoodCoinCondsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	exist, err := handler.ExistGoodCoinConds(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistGoodCoinConds",
			"In", in,
			"Error", err,
		)
		return &npool.ExistGoodCoinCondsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.ExistGoodCoinCondsResponse{
		Info: exist,
	}, nil
}
