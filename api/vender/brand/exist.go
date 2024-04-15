package brand

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	brand1 "github.com/NpoolPlatform/good-middleware/pkg/mw/vender/brand"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/vender/brand"
)

func (s *Server) ExistBrandConds(ctx context.Context, in *npool.ExistBrandCondsRequest) (*npool.ExistBrandCondsResponse, error) {
	handler, err := brand1.NewHandler(
		ctx,
		brand1.WithConds(in.GetConds()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistBrandConds",
			"In", in,
			"Error", err,
		)
		return &npool.ExistBrandCondsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	exist, err := handler.ExistBrandConds(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistBrandConds",
			"In", in,
			"Error", err,
		)
		return &npool.ExistBrandCondsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.ExistBrandCondsResponse{
		Info: exist,
	}, nil
}
