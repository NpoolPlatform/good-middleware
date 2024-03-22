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

func (s *Server) CreateDeviceInfo(ctx context.Context, in *npool.CreateDeviceInfoRequest) (*npool.CreateDeviceInfoResponse, error) {
	req := in.GetInfo()
	if req == nil {
		logger.Sugar().Errorw(
			"CreateDeviceInfo",
			"In", in,
		)
		return &npool.CreateDeviceInfoResponse{}, status.Error(codes.Aborted, "invalid argument")
	}
	handler, err := device1.NewHandler(
		ctx,
		device1.WithEntID(req.EntID, false),
		device1.WithType(req.Type, true),
		device1.WithManufacturer(req.Manufacturer, true),
		device1.WithPowerConsumption(req.PowerConsumption, true),
		device1.WithShipmentAt(req.ShipmentAt, true),
		device1.WithPosters(req.Posters, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateDeviceInfo",
			"In", in,
			"Error", err,
		)
		return &npool.CreateDeviceInfoResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.CreateDeviceInfo(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateDeviceInfo",
			"In", in,
			"Error", err,
		)
		return &npool.CreateDeviceInfoResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.CreateDeviceInfoResponse{
		Info: info,
	}, nil
}
