package appsimulategood

import (
	"context"

	appsimulategood1 "github.com/NpoolPlatform/good-middleware/pkg/mw/app/powerrental/simulate"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/app/powerrental/simulate"
)

func (s *Server) CreateSimulate(ctx context.Context, in *npool.CreateSimulateRequest) (*npool.CreateSimulateResponse, error) {
	req := in.GetInfo()
	if req == nil {
		logger.Sugar().Errorw(
			"CreateSimulate",
			"In", in,
		)
		return &npool.CreateSimulateResponse{}, status.Error(codes.Aborted, "invalid argument")
	}
	handler, err := appsimulategood1.NewHandler(
		ctx,
		appsimulategood1.WithEntID(req.EntID, false),
		appsimulategood1.WithAppGoodID(req.AppGoodID, true),
		appsimulategood1.WithFixedOrderUnits(req.FixedOrderUnits, true),
		appsimulategood1.WithFixedOrderDuration(req.FixedOrderDuration, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateSimulate",
			"In", in,
			"Error", err,
		)
		return &npool.CreateSimulateResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.CreateSimulate(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateSimulate",
			"In", in,
			"Error", err,
		)
		return &npool.CreateSimulateResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.CreateSimulateResponse{
		Info: info,
	}, nil
}
