//nolint:dupl
package brand

import (
	"context"

	brand1 "github.com/NpoolPlatform/good-middleware/pkg/mw/vender/brand"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/vender/brand"
)

func (s *Server) CreateBrand(ctx context.Context, in *npool.CreateBrandRequest) (*npool.CreateBrandResponse, error) {
	req := in.GetInfo()
	if req == nil {
		logger.Sugar().Errorw(
			"CreateBrand",
			"In", in,
		)
		return &npool.CreateBrandResponse{}, status.Error(codes.Aborted, "invalid argument")
	}
	handler, err := brand1.NewHandler(
		ctx,
		brand1.WithID(req.ID, false),
		brand1.WithName(req.Name, true),
		brand1.WithLogo(req.Logo, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateBrand",
			"In", in,
			"Error", err,
		)
		return &npool.CreateBrandResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.CreateBrand(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateBrand",
			"In", in,
			"Error", err,
		)
		return &npool.CreateBrandResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.CreateBrandResponse{
		Info: info,
	}, nil
}
