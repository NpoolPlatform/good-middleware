package comment

import (
	"context"

	comment1 "github.com/NpoolPlatform/good-middleware/pkg/mw/good/comment"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/good/comment"
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
	)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateComment",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateCommentResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.UpdateComment(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateComment",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateCommentResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.UpdateCommentResponse{
		Info: info,
	}, nil
}
