package coin

import (
	"context"

	coin1 "github.com/NpoolPlatform/good-middleware/pkg/mw/good/coin"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/good/coin"
)

func (s *Server) UpdateGoodCoin(ctx context.Context, in *npool.UpdateGoodCoinRequest) (*npool.UpdateGoodCoinResponse, error) {
	req := in.GetInfo()
	if req == nil {
		logger.Sugar().Errorw(
			"UpdateGoodCoin",
			"In", in,
		)
		return &npool.UpdateGoodCoinResponse{}, status.Error(codes.Aborted, "invalid argument")
	}
	handler, err := coin1.NewHandler(
		ctx,
		coin1.WithID(req.ID, false),
		coin1.WithEntID(req.EntID, false),
		coin1.WithMain(req.Main, false),
		coin1.WithIndex(req.Index, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateGoodCoin",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateGoodCoinResponse{}, status.Error(codes.Aborted, err.Error())
	}

	if err := handler.UpdateGoodCoin(ctx); err != nil {
		logger.Sugar().Errorw(
			"UpdateGoodCoin",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateGoodCoinResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.UpdateGoodCoinResponse{}, nil
}
