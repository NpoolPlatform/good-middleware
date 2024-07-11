package appsimulategood

import (
	"context"

	appsimulategood1 "github.com/NpoolPlatform/good-middleware/pkg/mw/app/powerrental/simulate"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/app/powerrental/simulate"
)

func (s *Server) DeleteSimulate(ctx context.Context, in *npool.DeleteSimulateRequest) (*npool.DeleteSimulateResponse, error) {
	req := in.GetInfo()
	if req == nil {
		logger.Sugar().Errorw(
			"DeleteSimulate",
			"In", in,
		)
		return &npool.DeleteSimulateResponse{}, status.Error(codes.Aborted, "invalid argument")
	}
	handler, err := appsimulategood1.NewHandler(
		ctx,
		appsimulategood1.WithID(req.ID, false),
		appsimulategood1.WithEntID(req.EntID, false),
		appsimulategood1.WithAppGoodID(req.AppGoodID, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteSimulate",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteSimulateResponse{}, status.Error(codes.Aborted, err.Error())
	}

	if err := handler.DeleteSimulate(ctx); err != nil {
		logger.Sugar().Errorw(
			"DeleteSimulate",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteSimulateResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.DeleteSimulateResponse{}, nil
}
