package appstock

import (
	"context"

	appstock1 "github.com/NpoolPlatform/good-middleware/pkg/mw/app/good/stock"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/app/good/stock"
)

func (s *Server) Reserve(ctx context.Context, in *npool.ReserveRequest) (*npool.ReserveResponse, error) {
	handler, err := appstock1.NewHandler(
		ctx,
		appstock1.WithEntID(&in.EntID, true),
		appstock1.WithAppGoodID(&in.AppGoodID, true),
		appstock1.WithReserved(&in.Units, true),
		appstock1.WithAppMiningGoodStocks(in.AppMiningGoodStocks, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"Reserve",
			"In", in,
			"Error", err,
		)
		return &npool.ReserveResponse{}, status.Error(codes.Aborted, err.Error())
	}

	if err := handler.ReserveStock(ctx); err != nil {
		logger.Sugar().Errorw(
			"Reserve",
			"In", in,
			"Error", err,
		)
		return &npool.ReserveResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.ReserveResponse{}, nil
}
