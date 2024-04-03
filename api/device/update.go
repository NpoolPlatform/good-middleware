//nolint:dupl
package device

import (
	"context"

	device1 "github.com/NpoolPlatform/good-middleware/pkg/mw/device"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/device"
)

func (s *Server) UpdateDeviceType(ctx context.Context, in *npool.UpdateDeviceTypeRequest) (*npool.UpdateDeviceTypeResponse, error) {
	req := in.GetInfo()
	if req == nil {
		logger.Sugar().Errorw(
			"UpdateDeviceType",
			"In", in,
		)
		return &npool.UpdateDeviceTypeResponse{}, status.Error(codes.Aborted, "invalid argument")
	}
	handler, err := device1.NewHandler(
		ctx,
		device1.WithID(req.ID, true),
		device1.WithType(req.Type, false),
		device1.WithManufacturerID(req.ManufacturerID, false),
		device1.WithPowerConsumption(req.PowerConsumption, false),
		device1.WithShipmentAt(req.ShipmentAt, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateDeviceType",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateDeviceTypeResponse{}, status.Error(codes.Aborted, err.Error())
	}

	if err := handler.UpdateDeviceType(ctx); err != nil {
		logger.Sugar().Errorw(
			"UpdateDeviceType",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateDeviceTypeResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.UpdateDeviceTypeResponse{}, nil
}
