//nolint:dupl
package deviceinfo

import (
	"context"

	deviceinfo1 "github.com/NpoolPlatform/good-middleware/pkg/mw/deviceinfo"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/deviceinfo"
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
	handler, err := deviceinfo1.NewHandler(
		ctx,
		deviceinfo1.WithEntID(req.EntID, false),
		deviceinfo1.WithType(req.Type, true),
		deviceinfo1.WithManufacturer(req.Manufacturer, true),
		deviceinfo1.WithPowerConsumption(req.PowerConsumption, true),
		deviceinfo1.WithShipmentAt(req.ShipmentAt, true),
		deviceinfo1.WithPosters(req.Posters, true),
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
