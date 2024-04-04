package manufacturer

import (
	"context"

	manufacturer1 "github.com/NpoolPlatform/good-middleware/pkg/mw/device/manufacturer"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/device/manufacturer"
)

func (s *Server) CreateManufacturer(ctx context.Context, in *npool.CreateManufacturerRequest) (*npool.CreateManufacturerResponse, error) {
	req := in.GetInfo()
	if req == nil {
		logger.Sugar().Errorw(
			"CreateManufacturer",
			"In", in,
		)
		return &npool.CreateManufacturerResponse{}, status.Error(codes.Aborted, "invalid argument")
	}
	handler, err := manufacturer1.NewHandler(
		ctx,
		manufacturer1.WithEntID(req.EntID, false),
		manufacturer1.WithName(req.Name, true),
		manufacturer1.WithLogo(req.Logo, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateManufacturer",
			"In", in,
			"Error", err,
		)
		return &npool.CreateManufacturerResponse{}, status.Error(codes.Aborted, err.Error())
	}

	if err := handler.CreateManufacturer(ctx); err != nil {
		logger.Sugar().Errorw(
			"CreateManufacturer",
			"In", in,
			"Error", err,
		)
		return &npool.CreateManufacturerResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.CreateManufacturerResponse{}, nil
}
