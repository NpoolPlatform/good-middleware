package location

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	location1 "github.com/NpoolPlatform/good-middleware/pkg/mw/vender/location"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/vender/location"
)

func (s *Server) GetLocation(ctx context.Context, in *npool.GetLocationRequest) (*npool.GetLocationResponse, error) {
	handler, err := location1.NewHandler(
		ctx,
		location1.WithEntID(&in.EntID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetLocation",
			"In", in,
			"Error", err,
		)
		return &npool.GetLocationResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.GetLocation(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetLocation",
			"In", in,
			"Error", err,
		)
		return &npool.GetLocationResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.GetLocationResponse{
		Info: info,
	}, nil
}

func (s *Server) GetLocations(ctx context.Context, in *npool.GetLocationsRequest) (*npool.GetLocationsResponse, error) {
	handler, err := location1.NewHandler(
		ctx,
		location1.WithConds(in.GetConds()),
		location1.WithOffset(in.GetOffset()),
		location1.WithLimit(in.GetLimit()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetLocations",
			"In", in,
			"Error", err,
		)
		return &npool.GetLocationsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	infos, total, err := handler.GetLocations(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetLocations",
			"In", in,
			"Error", err,
		)
		return &npool.GetLocationsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.GetLocationsResponse{
		Infos: infos,
		Total: total,
	}, nil
}
