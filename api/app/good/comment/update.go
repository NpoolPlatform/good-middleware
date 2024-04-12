package comment

import (
	"context"

	comment1 "github.com/NpoolPlatform/good-middleware/pkg/mw/app/good/comment"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/app/good/comment"
)

func (s *Server) UpdateComment(ctx context.Context, in *npool.UpdateCommentRequest) (*npool.UpdateCommentResponse, error) {
	req := in.GetInfo()
	if req == nil {
		logger.Sugar().Errorw(
			"UpdateComment",
			"In", in,
		)
		return &npool.UpdateCommentResponse{}, status.Error(codes.Aborted, "invalid argument")
	}
	handler, err := comment1.NewHandler(
		ctx,
		comment1.WithID(req.ID, true),
		comment1.WithContent(req.Content, false),
		comment1.WithAnonymous(req.Anonymous, false),
		comment1.WithHide(req.Hide, false),
		comment1.WithHideReason(req.HideReason, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateComment",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateCommentResponse{}, status.Error(codes.Aborted, err.Error())
	}

	if err := handler.UpdateComment(ctx); err != nil {
		logger.Sugar().Errorw(
			"UpdateComment",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateCommentResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.UpdateCommentResponse{}, nil
}
