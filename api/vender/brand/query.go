package brand

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	brand1 "github.com/NpoolPlatform/good-middleware/pkg/mw/vender/brand"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/vender/brand"
)

func (s *Server) GetBrand(ctx context.Context, in *npool.GetBrandRequest) (*npool.GetBrandResponse, error) {
	handler, err := brand1.NewHandler(
		ctx,
		brand1.WithEntID(&in.EntID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetBrand",
			"In", in,
			"Error", err,
		)
		return &npool.GetBrandResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.GetBrand(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetBrand",
			"In", in,
			"Error", err,
		)
		return &npool.GetBrandResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.GetBrandResponse{
		Info: info,
	}, nil
}

func (s *Server) GetBrands(ctx context.Context, in *npool.GetBrandsRequest) (*npool.GetBrandsResponse, error) {
	handler, err := brand1.NewHandler(
		ctx,
		brand1.WithConds(in.GetConds()),
		brand1.WithOffset(in.GetOffset()),
		brand1.WithLimit(in.GetLimit()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetBrands",
			"In", in,
			"Error", err,
		)
		return &npool.GetBrandsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	infos, total, err := handler.GetBrands(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetBrands",
			"In", in,
			"Error", err,
		)
		return &npool.GetBrandsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.GetBrandsResponse{
		Infos: infos,
		Total: total,
	}, nil
}
