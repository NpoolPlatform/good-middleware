package device

import (
	"context"

	device1 "github.com/NpoolPlatform/good-middleware/pkg/mw/device"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/device"
)

func (s *Server) CreateDeviceType(ctx context.Context, in *npool.CreateDeviceTypeRequest) (*npool.CreateDeviceTypeResponse, error) {
	req := in.GetInfo()
	if req == nil {
		logger.Sugar().Errorw(
			"CreateDeviceType",
			"In", in,
		)
		return &npool.CreateDeviceTypeResponse{}, status.Error(codes.Aborted, "invalid argument")
	}
	handler, err := device1.NewHandler(
		ctx,
		device1.WithEntID(req.EntID, false),
		device1.WithType(req.Type, true),
		device1.WithManufacturerID(req.ManufacturerID, true),
		device1.WithPowerConsumption(req.PowerConsumption, true),
		device1.WithShipmentAt(req.ShipmentAt, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateDeviceType",
			"In", in,
			"Error", err,
		)
		return &npool.CreateDeviceTypeResponse{}, status.Error(codes.Aborted, err.Error())
	}

	if err := handler.CreateDeviceType(ctx); err != nil {
		logger.Sugar().Errorw(
			"CreateDeviceType",
			"In", in,
			"Error", err,
		)
		return &npool.CreateDeviceTypeResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.CreateDeviceTypeResponse{}, nil
}
