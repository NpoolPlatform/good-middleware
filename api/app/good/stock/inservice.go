//nolint:dupl
package appstock

import (
	"context"

	appstock1 "github.com/NpoolPlatform/good-middleware/pkg/mw/app/good/stock"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/app/good/stock"
)

func (s *Server) InService(ctx context.Context, in *npool.InServiceRequest) (*npool.InServiceResponse, error) {
	handler, err := appstock1.NewHandler(
		ctx,
		appstock1.WithLockID(&in.LockID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"InService",
			"In", in,
			"Error", err,
		)
		return &npool.InServiceResponse{}, status.Error(codes.Aborted, err.Error())
	}

	if err := handler.InServiceStock(ctx); err != nil {
		logger.Sugar().Errorw(
			"InService",
			"In", in,
			"Error", err,
		)
		return &npool.InServiceResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.InServiceResponse{}, nil
}
