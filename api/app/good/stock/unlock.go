package appstock

import (
	"context"

	appstock1 "github.com/NpoolPlatform/good-middleware/pkg/mw/app/good/stock"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/app/good/stock"
)

func (s *Server) Unlock(ctx context.Context, in *npool.UnlockRequest) (*npool.UnlockResponse, error) {
	handler, err := appstock1.NewHandler(
		ctx,
		appstock1.WithLockID(&in.LockID, true),
		appstock1.WithRollback(&in.Rollback, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"Unlock",
			"In", in,
			"Error", err,
		)
		return &npool.UnlockResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.UnlockStock(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"Unlock",
			"In", in,
			"Error", err,
		)
		return &npool.UnlockResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.UnlockResponse{
		Info: info,
	}, nil
}
