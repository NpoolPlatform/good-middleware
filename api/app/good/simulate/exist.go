package appsimulategood

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	appsimulategood1 "github.com/NpoolPlatform/good-middleware/pkg/mw/app/good/simulate"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/app/good/simulate"
)

func (s *Server) ExistSimulateConds(ctx context.Context, in *npool.ExistSimulateCondsRequest) (*npool.ExistSimulateCondsResponse, error) {
	handler, err := appsimulategood1.NewHandler(
		ctx,
		appsimulategood1.WithConds(in.GetConds()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistSimulateConds",
			"In", in,
			"Error", err,
		)
		return &npool.ExistSimulateCondsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	exist, err := handler.ExistSimulateConds(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistSimulateConds",
			"In", in,
			"Error", err,
		)
		return &npool.ExistSimulateCondsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.ExistSimulateCondsResponse{
		Info: exist,
	}, nil
}
