package manufacturer

import (
	"context"

	manufacturer1 "github.com/NpoolPlatform/good-middleware/pkg/mw/device/manufacturer"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/device/manufacturer"
)

func (s *Server) UpdateManufacturer(ctx context.Context, in *npool.UpdateManufacturerRequest) (*npool.UpdateManufacturerResponse, error) {
	req := in.GetInfo()
	if req == nil {
		logger.Sugar().Errorw(
			"UpdateManufacturer",
			"In", in,
		)
		return &npool.UpdateManufacturerResponse{}, status.Error(codes.Aborted, "invalid argument")
	}
	handler, err := manufacturer1.NewHandler(
		ctx,
		manufacturer1.WithID(req.ID, false),
		manufacturer1.WithEntID(req.EntID, false),
		manufacturer1.WithName(req.Name, false),
		manufacturer1.WithLogo(req.Logo, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateManufacturer",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateManufacturerResponse{}, status.Error(codes.Aborted, err.Error())
	}

	if err := handler.UpdateManufacturer(ctx); err != nil {
		logger.Sugar().Errorw(
			"UpdateManufacturer",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateManufacturerResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.UpdateManufacturerResponse{}, nil
}
