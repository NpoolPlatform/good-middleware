package device

import (
	"context"

	device1 "github.com/NpoolPlatform/good-middleware/pkg/mw/device"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/device"
)

func (s *Server) DeleteDeviceInfo(ctx context.Context, in *npool.DeleteDeviceInfoRequest) (*npool.DeleteDeviceInfoResponse, error) {
	req := in.GetInfo()
	if req == nil {
		logger.Sugar().Errorw(
			"DeleteDeviceInfo",
			"In", in,
		)
		return &npool.DeleteDeviceInfoResponse{}, status.Error(codes.Aborted, "invalid argument")
	}
	handler, err := device1.NewHandler(
		ctx,
		device1.WithID(req.ID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteDeviceInfo",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteDeviceInfoResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.DeleteDeviceInfo(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteDeviceInfo",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteDeviceInfoResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.DeleteDeviceInfoResponse{
		Info: info,
	}, nil
}
