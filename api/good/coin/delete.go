package coin

import (
	"context"

	coin1 "github.com/NpoolPlatform/good-middleware/pkg/mw/good/coin"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/good/coin"
)

func (s *Server) DeleteGoodCoin(ctx context.Context, in *npool.DeleteGoodCoinRequest) (*npool.DeleteGoodCoinResponse, error) {
	req := in.GetInfo()
	if req == nil {
		logger.Sugar().Errorw(
			"DeleteGoodCoin",
			"In", in,
		)
		return &npool.DeleteGoodCoinResponse{}, status.Error(codes.Aborted, "invalid argument")
	}
	handler, err := coin1.NewHandler(
		ctx,
		coin1.WithID(req.ID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteGoodCoin",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteGoodCoinResponse{}, status.Error(codes.Aborted, err.Error())
	}

	if err := handler.DeleteGoodCoin(ctx); err != nil {
		logger.Sugar().Errorw(
			"DeleteGoodCoin",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteGoodCoinResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.DeleteGoodCoinResponse{}, nil
}
