package device

import (
	"context"

	device1 "github.com/NpoolPlatform/good-middleware/pkg/mw/device"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/device"
)

func (s *Server) DeleteDeviceType(ctx context.Context, in *npool.DeleteDeviceTypeRequest) (*npool.DeleteDeviceTypeResponse, error) {
	req := in.GetInfo()
	if req == nil {
		logger.Sugar().Errorw(
			"DeleteDeviceType",
			"In", in,
		)
		return &npool.DeleteDeviceTypeResponse{}, status.Error(codes.Aborted, "invalid argument")
	}
	handler, err := device1.NewHandler(
		ctx,
		device1.WithID(req.ID, false),
		device1.WithEntID(req.EntID, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteDeviceType",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteDeviceTypeResponse{}, status.Error(codes.Aborted, err.Error())
	}

	if err := handler.DeleteDeviceType(ctx); err != nil {
		logger.Sugar().Errorw(
			"DeleteDeviceType",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteDeviceTypeResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.DeleteDeviceTypeResponse{}, nil
}
