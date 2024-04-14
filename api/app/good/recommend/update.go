//nolint:dupl
package recommend

import (
	"context"

	recommend1 "github.com/NpoolPlatform/good-middleware/pkg/mw/app/good/recommend"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/app/good/recommend"
)

func (s *Server) UpdateRecommend(ctx context.Context, in *npool.UpdateRecommendRequest) (*npool.UpdateRecommendResponse, error) {
	req := in.GetInfo()
	if req == nil {
		logger.Sugar().Errorw(
			"UpdateRecommend",
			"In", in,
		)
		return &npool.UpdateRecommendResponse{}, status.Error(codes.Aborted, "invalid argument")
	}
	handler, err := recommend1.NewHandler(
		ctx,
		recommend1.WithID(req.ID, true),
		recommend1.WithRecommendIndex(req.RecommendIndex, false),
		recommend1.WithMessage(req.Message, false),
		recommend1.WithHide(req.Hide, false),
		recommend1.WithHideReason(req.HideReason, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateRecommend",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateRecommendResponse{}, status.Error(codes.Aborted, err.Error())
	}

	if err := handler.UpdateRecommend(ctx); err != nil {
		logger.Sugar().Errorw(
			"UpdateRecommend",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateRecommendResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.UpdateRecommendResponse{}, nil
}
