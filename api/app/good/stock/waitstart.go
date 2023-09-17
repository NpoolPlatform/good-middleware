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

func (s *Server) WaitStart(ctx context.Context, in *npool.WaitStartRequest) (*npool.WaitStartResponse, error) {
	handler, err := appstock1.NewHandler(
		ctx,
		appstock1.WithLockID(&in.LockID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"WaitStart",
			"In", in,
			"Error", err,
		)
		return &npool.WaitStartResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.WaitStartStock(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"WaitStart",
			"In", in,
			"Error", err,
		)
		return &npool.WaitStartResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.WaitStartResponse{
		Info: info,
	}, nil
}
