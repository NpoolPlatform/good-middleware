package deviceinfo

import (
	"context"

	deviceinfo1 "github.com/NpoolPlatform/good-middleware/pkg/mw/deviceinfo"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/deviceinfo"
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
	handler, err := deviceinfo1.NewHandler(
		ctx,
		deviceinfo1.WithID(req.ID),
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
