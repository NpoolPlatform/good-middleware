package constraint

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	constraint1 "github.com/NpoolPlatform/good-middleware/pkg/mw/app/good/topmost/good/constraint"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/app/good/topmost/good/constraint"
)

func (s *Server) ExistTopMostGoodConstraintConds(ctx context.Context, in *npool.ExistTopMostGoodConstraintCondsRequest) (*npool.ExistTopMostGoodConstraintCondsResponse, error) {
	handler, err := constraint1.NewHandler(
		ctx,
		constraint1.WithConds(in.GetConds()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistTopMostGoodConstraintConds",
			"In", in,
			"Error", err,
		)
		return &npool.ExistTopMostGoodConstraintCondsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	exist, err := handler.ExistConstraintConds(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistTopMostGoodConstraintConds",
			"In", in,
			"Error", err,
		)
		return &npool.ExistTopMostGoodConstraintCondsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.ExistTopMostGoodConstraintCondsResponse{
		Info: exist,
	}, nil
}
