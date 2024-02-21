package appsimulategood

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	appsimulategood1 "github.com/NpoolPlatform/good-middleware/pkg/mw/app/good/simulate"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/app/good/simulate"
)

func (s *Server) GetSimulate(ctx context.Context, in *npool.GetSimulateRequest) (*npool.GetSimulateResponse, error) {
	handler, err := appsimulategood1.NewHandler(
		ctx,
		appsimulategood1.WithEntID(&in.EntID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetSimulate",
			"In", in,
			"Error", err,
		)
		return &npool.GetSimulateResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.GetSimulate(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetSimulate",
			"In", in,
			"Error", err,
		)
		return &npool.GetSimulateResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.GetSimulateResponse{
		Info: info,
	}, nil
}

func (s *Server) GetSimulates(ctx context.Context, in *npool.GetSimulatesRequest) (*npool.GetSimulatesResponse, error) {
	handler, err := appsimulategood1.NewHandler(
		ctx,
		appsimulategood1.WithConds(in.GetConds()),
		appsimulategood1.WithOffset(in.GetOffset()),
		appsimulategood1.WithLimit(in.GetLimit()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetSimulates",
			"In", in,
			"Error", err,
		)
		return &npool.GetSimulatesResponse{}, status.Error(codes.Aborted, err.Error())
	}

	infos, total, err := handler.GetSimulates(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetSimulates",
			"In", in,
			"Error", err,
		)
		return &npool.GetSimulatesResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.GetSimulatesResponse{
		Infos: infos,
		Total: total,
	}, nil
}
