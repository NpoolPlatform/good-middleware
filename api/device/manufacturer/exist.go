package manufacturer

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	manufacturer1 "github.com/NpoolPlatform/good-middleware/pkg/mw/device/manufacturer"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/device/manufacturer"
)

func (s *Server) ExistManufacturerConds(ctx context.Context, in *npool.ExistManufacturerCondsRequest) (*npool.ExistManufacturerCondsResponse, error) {
	handler, err := manufacturer1.NewHandler(
		ctx,
		manufacturer1.WithConds(in.GetConds()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistManufacturerConds",
			"In", in,
			"Error", err,
		)
		return &npool.ExistManufacturerCondsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	exist, err := handler.ExistManufacturerConds(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistManufacturerConds",
			"In", in,
			"Error", err,
		)
		return &npool.ExistManufacturerCondsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.ExistManufacturerCondsResponse{
		Info: exist,
	}, nil
}
