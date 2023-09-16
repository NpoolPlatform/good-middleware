package appstock

import (
	"context"

	appstock1 "github.com/NpoolPlatform/good-middleware/pkg/mw/app/good/stock"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/app/good/stock"
)

func (s *Server) Expire(ctx context.Context, in *npool.ExpireRequest) (*npool.ExpireResponse, error) {
	handler, err := appstock1.NewHandler(
		ctx,
		appstock1.WithLockID(&in.LockID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"Expire",
			"In", in,
			"Error", err,
		)
		return &npool.ExpireResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.ExpireStock(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"Expire",
			"In", in,
			"Error", err,
		)
		return &npool.ExpireResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.ExpireResponse{
		Info: info,
	}, nil
}
