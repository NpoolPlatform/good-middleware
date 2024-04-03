package comment

import (
	"context"

	comment1 "github.com/NpoolPlatform/good-middleware/pkg/mw/app/good/comment"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/app/good/comment"
)

func (s *Server) DeleteComment(ctx context.Context, in *npool.DeleteCommentRequest) (*npool.DeleteCommentResponse, error) {
	req := in.GetInfo()
	if req == nil {
		logger.Sugar().Errorw(
			"DeleteComment",
			"In", in,
		)
		return &npool.DeleteCommentResponse{}, status.Error(codes.Aborted, "invalid argument")
	}
	handler, err := comment1.NewHandler(
		ctx,
		comment1.WithID(req.ID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteComment",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteCommentResponse{}, status.Error(codes.Aborted, err.Error())
	}

	if err := handler.DeleteComment(ctx); err != nil {
		logger.Sugar().Errorw(
			"DeleteComment",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteCommentResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.DeleteCommentResponse{}, nil
}
