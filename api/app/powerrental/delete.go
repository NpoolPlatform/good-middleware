package powerrental

import (
	"context"

	powerrental1 "github.com/NpoolPlatform/good-middleware/pkg/mw/app/powerrental"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/app/powerrental"
)

func (s *Server) DeletePowerRental(ctx context.Context, in *npool.DeletePowerRentalRequest) (*npool.DeletePowerRentalResponse, error) {
	req := in.GetInfo()
	if req == nil {
		logger.Sugar().Errorw(
			"DeletePowerRental",
			"In", in,
		)
		return &npool.DeletePowerRentalResponse{}, status.Error(codes.Aborted, "invalid argument")
	}
	handler, err := powerrental1.NewHandler(
		ctx,
		powerrental1.WithID(req.ID, false),
		powerrental1.WithEntID(req.EntID, false),
		powerrental1.WithAppGoodID(req.AppGoodID, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"DeletePowerRental",
			"In", in,
			"Error", err,
		)
		return &npool.DeletePowerRentalResponse{}, status.Error(codes.Aborted, err.Error())
	}

	if err := handler.DeletePowerRental(ctx); err != nil {
		logger.Sugar().Errorw(
			"DeletePowerRental",
			"In", in,
			"Error", err,
		)
		return &npool.DeletePowerRentalResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.DeletePowerRentalResponse{}, nil
}
