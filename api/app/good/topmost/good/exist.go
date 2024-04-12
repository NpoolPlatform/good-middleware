package topmostgood

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	topmostgood1 "github.com/NpoolPlatform/good-middleware/pkg/mw/app/good/topmost/good"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/app/good/topmost/good"
)

func (s *Server) ExistTopMostGoodConds(ctx context.Context, in *npool.ExistTopMostGoodCondsRequest) (*npool.ExistTopMostGoodCondsResponse, error) {
	handler, err := topmostgood1.NewHandler(
		ctx,
		topmostgood1.WithConds(in.GetConds()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistTopMostGoodConds",
			"In", in,
			"Error", err,
		)
		return &npool.ExistTopMostGoodCondsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	exist, err := handler.ExistTopMostGoodConds(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistTopMostGoodConds",
			"In", in,
			"Error", err,
		)
		return &npool.ExistTopMostGoodCondsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.ExistTopMostGoodCondsResponse{
		Info: exist,
	}, nil
}
