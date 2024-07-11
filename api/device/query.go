package device

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	device1 "github.com/NpoolPlatform/good-middleware/pkg/mw/device"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/device"
)

func (s *Server) GetDeviceType(ctx context.Context, in *npool.GetDeviceTypeRequest) (*npool.GetDeviceTypeResponse, error) {
	handler, err := device1.NewHandler(
		ctx,
		device1.WithEntID(&in.EntID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetDeviceType",
			"In", in,
			"Error", err,
		)
		return &npool.GetDeviceTypeResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.GetDeviceType(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetDeviceType",
			"In", in,
			"Error", err,
		)
		return &npool.GetDeviceTypeResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.GetDeviceTypeResponse{
		Info: info,
	}, nil
}

func (s *Server) GetDeviceTypes(ctx context.Context, in *npool.GetDeviceTypesRequest) (*npool.GetDeviceTypesResponse, error) {
	handler, err := device1.NewHandler(
		ctx,
		device1.WithConds(in.GetConds()),
		device1.WithOffset(in.GetOffset()),
		device1.WithLimit(in.GetLimit()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetDeviceTypes",
			"In", in,
			"Error", err,
		)
		return &npool.GetDeviceTypesResponse{}, status.Error(codes.Aborted, err.Error())
	}

	infos, total, err := handler.GetDeviceTypes(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetDeviceTypes",
			"In", in,
			"Error", err,
		)
		return &npool.GetDeviceTypesResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.GetDeviceTypesResponse{
		Infos: infos,
		Total: total,
	}, nil
}
