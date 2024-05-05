package powerrental

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	powerrental1 "github.com/NpoolPlatform/good-middleware/pkg/mw/powerrental"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/powerrental"
)

func (s *Server) ExistPowerRentalConds(ctx context.Context, in *npool.ExistPowerRentalCondsRequest) (*npool.ExistPowerRentalCondsResponse, error) {
	handler, err := powerrental1.NewHandler(
		ctx,
		powerrental1.WithConds(in.GetConds()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistPowerRentalConds",
			"In", in,
			"Error", err,
		)
		return &npool.ExistPowerRentalCondsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	exist, err := handler.ExistPowerRentalConds(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistPowerRentalConds",
			"In", in,
			"Error", err,
		)
		return &npool.ExistPowerRentalCondsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.ExistPowerRentalCondsResponse{
		Info: exist,
	}, nil
}
