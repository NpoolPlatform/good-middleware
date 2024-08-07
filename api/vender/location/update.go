package location

import (
	"context"

	location1 "github.com/NpoolPlatform/good-middleware/pkg/mw/vender/location"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/vender/location"
)

func (s *Server) UpdateLocation(ctx context.Context, in *npool.UpdateLocationRequest) (*npool.UpdateLocationResponse, error) {
	req := in.GetInfo()
	if req == nil {
		logger.Sugar().Errorw(
			"UpdateLocation",
			"In", in,
		)
		return &npool.UpdateLocationResponse{}, status.Error(codes.Aborted, "invalid argument")
	}
	handler, err := location1.NewHandler(
		ctx,
		location1.WithID(req.ID, false),
		location1.WithEntID(req.EntID, false),
		location1.WithCountry(req.Country, false),
		location1.WithProvince(req.Province, false),
		location1.WithCity(req.City, false),
		location1.WithAddress(req.Address, false),
		location1.WithBrandID(req.BrandID, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateLocation",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateLocationResponse{}, status.Error(codes.Aborted, err.Error())
	}

	if err := handler.UpdateLocation(ctx); err != nil {
		logger.Sugar().Errorw(
			"UpdateLocation",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateLocationResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.UpdateLocationResponse{}, nil
}
