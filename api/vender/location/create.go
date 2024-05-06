package location

import (
	"context"

	location1 "github.com/NpoolPlatform/good-middleware/pkg/mw/vender/location"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/vender/location"
)

func (s *Server) CreateLocation(ctx context.Context, in *npool.CreateLocationRequest) (*npool.CreateLocationResponse, error) {
	req := in.GetInfo()
	if req == nil {
		logger.Sugar().Errorw(
			"CreateLocation",
			"In", in,
		)
		return &npool.CreateLocationResponse{}, status.Error(codes.Aborted, "invalid argument")
	}
	handler, err := location1.NewHandler(
		ctx,
		location1.WithEntID(req.EntID, false),
		location1.WithCountry(req.Country, true),
		location1.WithProvince(req.Province, true),
		location1.WithCity(req.City, true),
		location1.WithAddress(req.Address, true),
		location1.WithBrandID(req.BrandID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateLocation",
			"In", in,
			"Error", err,
		)
		return &npool.CreateLocationResponse{}, status.Error(codes.Aborted, err.Error())
	}

	if err := handler.CreateLocation(ctx); err != nil {
		logger.Sugar().Errorw(
			"CreateLocation",
			"In", in,
			"Error", err,
		)
		return &npool.CreateLocationResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.CreateLocationResponse{}, nil
}
