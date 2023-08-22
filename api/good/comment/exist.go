package comment

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	comment1 "github.com/NpoolPlatform/good-middleware/pkg/mw/good/comment"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/good/comment"
)

func (s *Server) ExistCommentConds(ctx context.Context, in *npool.ExistCommentCondsRequest) (*npool.ExistCommentCondsResponse, error) {
	handler, err := comment1.NewHandler(
		ctx,
		comment1.WithConds(in.GetConds()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistCommentConds",
			"In", in,
			"Error", err,
		)
		return &npool.ExistCommentCondsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	exist, err := handler.ExistCommentConds(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistCommentConds",
			"In", in,
			"Error", err,
		)
		return &npool.ExistCommentCondsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.ExistCommentCondsResponse{
		Info: exist,
	}, nil
}
