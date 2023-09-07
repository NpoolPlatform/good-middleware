package comment

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	comment1 "github.com/NpoolPlatform/good-middleware/pkg/mw/good/comment"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/good/comment"
)

func (s *Server) GetComment(ctx context.Context, in *npool.GetCommentRequest) (*npool.GetCommentResponse, error) {
	handler, err := comment1.NewHandler(
		ctx,
		comment1.WithID(&in.ID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetComment",
			"In", in,
			"Error", err,
		)
		return &npool.GetCommentResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.GetComment(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetComment",
			"In", in,
			"Error", err,
		)
		return &npool.GetCommentResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.GetCommentResponse{
		Info: info,
	}, nil
}

func (s *Server) GetComments(ctx context.Context, in *npool.GetCommentsRequest) (*npool.GetCommentsResponse, error) {
	handler, err := comment1.NewHandler(
		ctx,
		comment1.WithConds(in.GetConds()),
		comment1.WithOffset(in.GetOffset()),
		comment1.WithLimit(in.GetLimit()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetComments",
			"In", in,
			"Error", err,
		)
		return &npool.GetCommentsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	infos, total, err := handler.GetComments(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetComments",
			"In", in,
			"Error", err,
		)
		return &npool.GetCommentsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.GetCommentsResponse{
		Infos: infos,
		Total: total,
	}, nil
}
