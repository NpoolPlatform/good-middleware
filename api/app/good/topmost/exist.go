package topmost

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	topmost1 "github.com/NpoolPlatform/good-middleware/pkg/mw/app/good/topmost"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/app/good/topmost"
)

func (s *Server) ExistTopMostConds(ctx context.Context, in *npool.ExistTopMostCondsRequest) (*npool.ExistTopMostCondsResponse, error) {
	handler, err := topmost1.NewHandler(
		ctx,
		topmost1.WithConds(in.GetConds()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistTopMostConds",
			"In", in,
			"Error", err,
		)
		return &npool.ExistTopMostCondsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	exist, err := handler.ExistTopMostConds(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistTopMostConds",
			"In", in,
			"Error", err,
		)
		return &npool.ExistTopMostCondsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.ExistTopMostCondsResponse{
		Info: exist,
	}, nil
}
