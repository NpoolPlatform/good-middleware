package powerrental

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	apppowerrental1 "github.com/NpoolPlatform/good-middleware/pkg/mw/app/powerrental"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/app/powerrental"
)

func (s *Server) ExistPowerRentalConds(ctx context.Context, in *npool.ExistPowerRentalCondsRequest) (*npool.ExistPowerRentalCondsResponse, error) {
	handler, err := apppowerrental1.NewHandler(
		ctx,
		apppowerrental1.WithConds(in.GetConds()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistPowerRentalConds",
			"In", in,
			"Error", err,
		)
		return &npool.ExistPowerRentalCondsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.ExistPowerRentalConds(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistPowerRentalConds",
			"In", in,
			"Error", err,
		)
		return &npool.ExistPowerRentalCondsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.ExistPowerRentalCondsResponse{
		Info: info,
	}, nil
}
