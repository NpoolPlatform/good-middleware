package manufacturer

import (
	"context"

	manufacturer1 "github.com/NpoolPlatform/good-middleware/pkg/mw/device/manufacturer"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/device/manufacturer"
)

func (s *Server) DeleteManufacturer(ctx context.Context, in *npool.DeleteManufacturerRequest) (*npool.DeleteManufacturerResponse, error) {
	req := in.GetInfo()
	if req == nil {
		logger.Sugar().Errorw(
			"DeleteManufacturer",
			"In", in,
		)
		return &npool.DeleteManufacturerResponse{}, status.Error(codes.Aborted, "invalid argument")
	}
	handler, err := manufacturer1.NewHandler(
		ctx,
		manufacturer1.WithID(req.ID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteManufacturer",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteManufacturerResponse{}, status.Error(codes.Aborted, err.Error())
	}

	if err := handler.DeleteManufacturer(ctx); err != nil {
		logger.Sugar().Errorw(
			"DeleteManufacturer",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteManufacturerResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.DeleteManufacturerResponse{}, nil
}
