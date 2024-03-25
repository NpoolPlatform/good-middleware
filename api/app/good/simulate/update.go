package appsimulategood

import (
	"context"

	appsimulategood1 "github.com/NpoolPlatform/good-middleware/pkg/mw/app/powerrental/simulate"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/app/powerrental/simulate"
)

func (s *Server) UpdateSimulate(ctx context.Context, in *npool.UpdateSimulateRequest) (*npool.UpdateSimulateResponse, error) {
	req := in.GetInfo()
	if req == nil {
		logger.Sugar().Errorw(
			"UpdateSimulate",
			"In", in,
		)
		return &npool.UpdateSimulateResponse{}, status.Error(codes.Aborted, "invalid argument")
	}
	handler, err := appsimulategood1.NewHandler(
		ctx,
		appsimulategood1.WithID(req.ID, true),
		appsimulategood1.WithFixedOrderUnits(req.FixedOrderUnits, false),
		appsimulategood1.WithFixedOrderDuration(req.FixedOrderDuration, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateSimulate",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateSimulateResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.UpdateSimulate(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateSimulate",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateSimulateResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.UpdateSimulateResponse{
		Info: info,
	}, nil
}
