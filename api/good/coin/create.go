package coin

import (
	"context"

	coin1 "github.com/NpoolPlatform/good-middleware/pkg/mw/good/coin"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/good/coin"
)

func (s *Server) CreateGoodCoin(ctx context.Context, in *npool.CreateGoodCoinRequest) (*npool.CreateGoodCoinResponse, error) {
	req := in.GetInfo()
	if req == nil {
		logger.Sugar().Errorw(
			"CreateGoodCoin",
			"In", in,
		)
		return &npool.CreateGoodCoinResponse{}, status.Error(codes.Aborted, "invalid argument")
	}
	handler, err := coin1.NewHandler(
		ctx,
		coin1.WithEntID(req.EntID, false),
		coin1.WithGoodID(req.GoodID, true),
		coin1.WithCoinTypeID(req.CoinTypeID, false),
		coin1.WithMain(req.Main, true),
		coin1.WithIndex(req.Index, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateGoodCoin",
			"In", in,
			"Error", err,
		)
		return &npool.CreateGoodCoinResponse{}, status.Error(codes.Aborted, err.Error())
	}

	if err := handler.CreateGoodCoin(ctx); err != nil {
		logger.Sugar().Errorw(
			"CreateGoodCoin",
			"In", in,
			"Error", err,
		)
		return &npool.CreateGoodCoinResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.CreateGoodCoinResponse{}, nil
}
