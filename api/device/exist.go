package device

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	device1 "github.com/NpoolPlatform/good-middleware/pkg/mw/device"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/device"
)

func (s *Server) ExistDeviceTypeConds(ctx context.Context, in *npool.ExistDeviceTypeCondsRequest) (*npool.ExistDeviceTypeCondsResponse, error) {
	handler, err := device1.NewHandler(
		ctx,
		device1.WithConds(in.GetConds()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistDeviceTypeConds",
			"In", in,
			"Error", err,
		)
		return &npool.ExistDeviceTypeCondsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	exist, err := handler.ExistDeviceTypeConds(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistDeviceTypeConds",
			"In", in,
			"Error", err,
		)
		return &npool.ExistDeviceTypeCondsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.ExistDeviceTypeCondsResponse{
		Info: exist,
	}, nil
}
