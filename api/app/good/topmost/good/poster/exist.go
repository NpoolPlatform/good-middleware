package poster

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	poster1 "github.com/NpoolPlatform/good-middleware/pkg/mw/app/good/topmost/good/poster"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/app/good/topmost/good/poster"
)

func (s *Server) ExistPosterConds(ctx context.Context, in *npool.ExistPosterCondsRequest) (*npool.ExistPosterCondsResponse, error) {
	handler, err := poster1.NewHandler(
		ctx,
		poster1.WithConds(in.GetConds()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistPosterConds",
			"In", in,
			"Error", err,
		)
		return &npool.ExistPosterCondsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	exist, err := handler.ExistPosterConds(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistPosterConds",
			"In", in,
			"Error", err,
		)
		return &npool.ExistPosterCondsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.ExistPosterCondsResponse{
		Info: exist,
	}, nil
}
