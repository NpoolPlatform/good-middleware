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

func (s *Server) UpdateDeviceInfo(ctx context.Context, in *npool.UpdateDeviceInfoRequest) (*npool.UpdateDeviceInfoResponse, error) {
	req := in.GetInfo()
	if req == nil {
		logger.Sugar().Errorw(
			"UpdateDeviceInfo",
			"In", in,
		)
		return &npool.UpdateDeviceInfoResponse{}, status.Error(codes.Aborted, "invalid argument")
	}
	handler, err := device1.NewHandler(
		ctx,
		device1.WithID(req.ID, true),
		device1.WithType(req.Type, false),
		device1.WithManufacturer(req.Manufacturer, false),
		device1.WithPowerConsumption(req.PowerConsumption, false),
		device1.WithShipmentAt(req.ShipmentAt, false),
		device1.WithPosters(req.Posters, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateDeviceInfo",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateDeviceInfoResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.UpdateDeviceInfo(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateDeviceInfo",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateDeviceInfoResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.UpdateDeviceInfoResponse{
		Info: info,
	}, nil
}
