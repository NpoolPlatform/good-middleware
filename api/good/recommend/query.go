package recommend

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	recommend1 "github.com/NpoolPlatform/good-middleware/pkg/mw/good/recommend"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/good/recommend"
)

func (s *Server) GetRecommend(ctx context.Context, in *npool.GetRecommendRequest) (*npool.GetRecommendResponse, error) {
	handler, err := recommend1.NewHandler(
		ctx,
		recommend1.WithEntID(&in.EntID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetRecommend",
			"In", in,
			"Error", err,
		)
		return &npool.GetRecommendResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.GetRecommend(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetRecommend",
			"In", in,
			"Error", err,
		)
		return &npool.GetRecommendResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.GetRecommendResponse{
		Info: info,
	}, nil
}

func (s *Server) GetRecommends(ctx context.Context, in *npool.GetRecommendsRequest) (*npool.GetRecommendsResponse, error) {
	handler, err := recommend1.NewHandler(
		ctx,
		recommend1.WithConds(in.GetConds()),
		recommend1.WithOffset(in.GetOffset()),
		recommend1.WithLimit(in.GetLimit()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetRecommends",
			"In", in,
			"Error", err,
		)
		return &npool.GetRecommendsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	infos, total, err := handler.GetRecommends(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetRecommends",
			"In", in,
			"Error", err,
		)
		return &npool.GetRecommendsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.GetRecommendsResponse{
		Infos: infos,
		Total: total,
	}, nil
}
