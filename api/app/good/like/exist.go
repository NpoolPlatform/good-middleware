package like

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	like1 "github.com/NpoolPlatform/good-middleware/pkg/mw/app/good/like"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/app/good/like"
)

func (s *Server) ExistLikeConds(ctx context.Context, in *npool.ExistLikeCondsRequest) (*npool.ExistLikeCondsResponse, error) {
	handler, err := like1.NewHandler(
		ctx,
		like1.WithConds(in.GetConds()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistLikeConds",
			"In", in,
			"Error", err,
		)
		return &npool.ExistLikeCondsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	exist, err := handler.ExistLikeConds(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistLikeConds",
			"In", in,
			"Error", err,
		)
		return &npool.ExistLikeCondsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.ExistLikeCondsResponse{
		Info: exist,
	}, nil
}
