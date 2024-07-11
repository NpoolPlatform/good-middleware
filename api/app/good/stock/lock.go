package appstock

import (
	"context"

	appstock1 "github.com/NpoolPlatform/good-middleware/pkg/mw/app/good/stock"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/app/good/stock"
)

func (s *Server) Lock(ctx context.Context, in *npool.LockRequest) (*npool.LockResponse, error) {
	handler, err := appstock1.NewHandler(
		ctx,
		appstock1.WithEntID(&in.EntID, true),
		appstock1.WithAppGoodID(&in.AppGoodID, true),
		appstock1.WithLocked(&in.Units, true),
		appstock1.WithAppSpotLocked(&in.AppSpotUnits, false),
		appstock1.WithLockID(&in.LockID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"Lock",
			"In", in,
			"Error", err,
		)
		return &npool.LockResponse{}, status.Error(codes.Aborted, err.Error())
	}

	if err := handler.LockStock(ctx); err != nil {
		logger.Sugar().Errorw(
			"Lock",
			"In", in,
			"Error", err,
		)
		return &npool.LockResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.LockResponse{}, nil
}

func (s *Server) Locks(ctx context.Context, in *npool.LocksRequest) (*npool.LocksResponse, error) {
	handler, err := appstock1.NewHandler(
		ctx,
		appstock1.WithStocks(in.Stocks, true),
		appstock1.WithLockID(&in.LockID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"Locks",
			"In", in,
			"Error", err,
		)
		return &npool.LocksResponse{}, status.Error(codes.Aborted, err.Error())
	}

	if err := handler.LockStocks(ctx); err != nil {
		logger.Sugar().Errorw(
			"Locks",
			"In", in,
			"Error", err,
		)
		return &npool.LocksResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.LocksResponse{}, nil
}
