package manufacturer

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	manufacturer1 "github.com/NpoolPlatform/good-middleware/pkg/mw/device/manufacturer"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/device/manufacturer"
)

func (s *Server) GetManufacturer(ctx context.Context, in *npool.GetManufacturerRequest) (*npool.GetManufacturerResponse, error) {
	handler, err := manufacturer1.NewHandler(
		ctx,
		manufacturer1.WithEntID(&in.EntID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetManufacturer",
			"In", in,
			"Error", err,
		)
		return &npool.GetManufacturerResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.GetManufacturer(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetManufacturer",
			"In", in,
			"Error", err,
		)
		return &npool.GetManufacturerResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.GetManufacturerResponse{
		Info: info,
	}, nil
}

func (s *Server) GetManufacturers(ctx context.Context, in *npool.GetManufacturersRequest) (*npool.GetManufacturersResponse, error) {
	handler, err := manufacturer1.NewHandler(
		ctx,
		manufacturer1.WithConds(in.GetConds()),
		manufacturer1.WithOffset(in.GetOffset()),
		manufacturer1.WithLimit(in.GetLimit()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetManufacturers",
			"In", in,
			"Error", err,
		)
		return &npool.GetManufacturersResponse{}, status.Error(codes.Aborted, err.Error())
	}

	infos, total, err := handler.GetManufacturers(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetManufacturers",
			"In", in,
			"Error", err,
		)
		return &npool.GetManufacturersResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.GetManufacturersResponse{
		Infos: infos,
		Total: total,
	}, nil
}
