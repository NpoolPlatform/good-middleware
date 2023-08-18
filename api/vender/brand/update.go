package brand

import (
	"context"

	brand1 "github.com/NpoolPlatform/good-middleware/pkg/mw/vender/brand"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/vender/brand"
)

func (s *Server) UpdateBrand(ctx context.Context, in *npool.UpdateBrandRequest) (*npool.UpdateBrandResponse, error) {
	req := in.GetInfo()
	if req == nil {
		logger.Sugar().Errorw(
			"UpdateBrand",
			"In", in,
		)
		return &npool.UpdateBrandResponse{}, status.Error(codes.Aborted, "invalid argument")
	}
	handler, err := brand1.NewHandler(
		ctx,
		brand1.WithID(req.ID),
		brand1.WithName(req.Name),
		brand1.WithLogo(req.Logo),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateBrand",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateBrandResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.UpdateBrand(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateBrand",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateBrandResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.UpdateBrandResponse{
		Info: info,
	}, nil
}
