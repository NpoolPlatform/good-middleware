package device

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	device1 "github.com/NpoolPlatform/good-middleware/pkg/mw/device"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/device"
)

func (s *Server) GetDeviceInfo(ctx context.Context, in *npool.GetDeviceInfoRequest) (*npool.GetDeviceInfoResponse, error) {
	handler, err := device1.NewHandler(
		ctx,
		device1.WithEntID(&in.EntID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetDeviceInfo",
			"In", in,
			"Error", err,
		)
		return &npool.GetDeviceInfoResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.GetDeviceInfo(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetDeviceInfo",
			"In", in,
			"Error", err,
		)
		return &npool.GetDeviceInfoResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.GetDeviceInfoResponse{
		Info: info,
	}, nil
}

func (s *Server) GetDeviceInfos(ctx context.Context, in *npool.GetDeviceInfosRequest) (*npool.GetDeviceInfosResponse, error) {
	handler, err := device1.NewHandler(
		ctx,
		device1.WithConds(in.GetConds()),
		device1.WithOffset(in.GetOffset()),
		device1.WithLimit(in.GetLimit()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetDeviceInfos",
			"In", in,
			"Error", err,
		)
		return &npool.GetDeviceInfosResponse{}, status.Error(codes.Aborted, err.Error())
	}

	infos, total, err := handler.GetDeviceInfos(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetDeviceInfos",
			"In", in,
			"Error", err,
		)
		return &npool.GetDeviceInfosResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.GetDeviceInfosResponse{
		Infos: infos,
		Total: total,
	}, nil
}
