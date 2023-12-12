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
		appstock1.WithAppID(&in.AppID, true),
		appstock1.WithGoodID(&in.GoodID, true),
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

	info, err := handler.LockStock(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"Lock",
			"In", in,
			"Error", err,
		)
		return &npool.LockResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.LockResponse{
		Info: info,
	}, nil
}
