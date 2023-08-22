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

func (s *Server) UpdateDeviceInfo(ctx context.Context, in *npool.UpdateDeviceInfoRequest) (*npool.UpdateDeviceInfoResponse, error) {
	req := in.GetInfo()
	if req == nil {
		logger.Sugar().Errorw(
			"UpdateDeviceInfo",
			"In", in,
		)
		return &npool.UpdateDeviceInfoResponse{}, status.Error(codes.Aborted, "invalid argument")
	}
	handler, err := deviceinfo1.NewHandler(
		ctx,
		deviceinfo1.WithID(req.ID),
		deviceinfo1.WithType(req.Type),
		deviceinfo1.WithManufacturer(req.Manufacturer),
		deviceinfo1.WithPowerConsumption(req.PowerConsumption),
		deviceinfo1.WithShipmentAt(req.ShipmentAt),
		deviceinfo1.WithPosters(req.Posters),
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
