package brand

import (
	"context"

	brand1 "github.com/NpoolPlatform/good-middleware/pkg/mw/vender/brand"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/vender/brand"
)

func (s *Server) DeleteBrand(ctx context.Context, in *npool.DeleteBrandRequest) (*npool.DeleteBrandResponse, error) {
	req := in.GetInfo()
	if req == nil {
		logger.Sugar().Errorw(
			"DeleteBrand",
			"In", in,
		)
		return &npool.DeleteBrandResponse{}, status.Error(codes.Aborted, "invalid argument")
	}
	handler, err := brand1.NewHandler(
		ctx,
		brand1.WithID(req.ID, false),
		brand1.WithEntID(req.EntID, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteBrand",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteBrandResponse{}, status.Error(codes.Aborted, err.Error())
	}

	if err := handler.DeleteBrand(ctx); err != nil {
		logger.Sugar().Errorw(
			"DeleteBrand",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteBrandResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.DeleteBrandResponse{}, nil
}
