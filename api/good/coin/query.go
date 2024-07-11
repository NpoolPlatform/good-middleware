package coin

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	coin1 "github.com/NpoolPlatform/good-middleware/pkg/mw/good/coin"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/good/coin"
)

func (s *Server) GetGoodCoin(ctx context.Context, in *npool.GetGoodCoinRequest) (*npool.GetGoodCoinResponse, error) {
	handler, err := coin1.NewHandler(
		ctx,
		coin1.WithEntID(&in.EntID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetGoodCoin",
			"In", in,
			"Error", err,
		)
		return &npool.GetGoodCoinResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.GetGoodCoin(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetGoodCoin",
			"In", in,
			"Error", err,
		)
		return &npool.GetGoodCoinResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.GetGoodCoinResponse{
		Info: info,
	}, nil
}

func (s *Server) GetGoodCoins(ctx context.Context, in *npool.GetGoodCoinsRequest) (*npool.GetGoodCoinsResponse, error) {
	handler, err := coin1.NewHandler(
		ctx,
		coin1.WithConds(in.GetConds()),
		coin1.WithOffset(in.GetOffset()),
		coin1.WithLimit(in.GetLimit()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetGoodCoins",
			"In", in,
			"Error", err,
		)
		return &npool.GetGoodCoinsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	infos, total, err := handler.GetGoodCoins(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetGoodCoins",
			"In", in,
			"Error", err,
		)
		return &npool.GetGoodCoinsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.GetGoodCoinsResponse{
		Infos: infos,
		Total: total,
	}, nil
}
