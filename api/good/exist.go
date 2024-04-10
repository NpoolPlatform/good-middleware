package good

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	good1 "github.com/NpoolPlatform/good-middleware/pkg/mw/good"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/good"
)

func (s *Server) ExistGoodConds(ctx context.Context, in *npool.ExistGoodCondsRequest) (*npool.ExistGoodCondsResponse, error) {
	handler, err := good1.NewHandler(
		ctx,
		good1.WithConds(in.GetConds()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistGoodConds",
			"In", in,
			"Error", err,
		)
		return &npool.ExistGoodCondsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.ExistGoodConds(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistGoodConds",
			"In", in,
			"Error", err,
		)
		return &npool.ExistGoodCondsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.ExistGoodCondsResponse{
		Info: info,
	}, nil
}
