package comment

import (
	"context"

	comment1 "github.com/NpoolPlatform/good-middleware/pkg/mw/app/good/comment"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/app/good/comment"
)

func (s *Server) CreateComment(ctx context.Context, in *npool.CreateCommentRequest) (*npool.CreateCommentResponse, error) {
	req := in.GetInfo()
	if req == nil {
		logger.Sugar().Errorw(
			"CreateComment",
			"In", in,
		)
		return &npool.CreateCommentResponse{}, status.Error(codes.Aborted, "invalid argument")
	}
	handler, err := comment1.NewHandler(
		ctx,
		comment1.WithEntID(req.EntID, false),
		comment1.WithUserID(req.UserID, true),
		comment1.WithAppGoodID(req.AppGoodID, true),
		comment1.WithOrderID(req.OrderID, false),
		comment1.WithContent(req.Content, true),
		comment1.WithReplyToID(req.ReplyToID, false),
		comment1.WithAnonymous(req.Anonymous, false),
		comment1.WithPurchasedUser(req.PurchasedUser, false),
		comment1.WithTrialUser(req.TrialUser, false),
		comment1.WithScore(req.Score, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateComment",
			"In", in,
			"Error", err,
		)
		return &npool.CreateCommentResponse{}, status.Error(codes.Aborted, err.Error())
	}

	if err := handler.CreateComment(ctx); err != nil {
		logger.Sugar().Errorw(
			"CreateComment",
			"In", in,
			"Error", err,
		)
		return &npool.CreateCommentResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.CreateCommentResponse{}, nil
}
