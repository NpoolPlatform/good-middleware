package like

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	like1 "github.com/NpoolPlatform/good-middleware/pkg/mw/good/like"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/good/like"
)

func (s *Server) GetLike(ctx context.Context, in *npool.GetLikeRequest) (*npool.GetLikeResponse, error) {
	handler, err := like1.NewHandler(
		ctx,
		like1.WithEntID(&in.EntID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetLike",
			"In", in,
			"Error", err,
		)
		return &npool.GetLikeResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.GetLike(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetLike",
			"In", in,
			"Error", err,
		)
		return &npool.GetLikeResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.GetLikeResponse{
		Info: info,
	}, nil
}

func (s *Server) GetLikes(ctx context.Context, in *npool.GetLikesRequest) (*npool.GetLikesResponse, error) {
	handler, err := like1.NewHandler(
		ctx,
		like1.WithConds(in.GetConds()),
		like1.WithOffset(in.GetOffset()),
		like1.WithLimit(in.GetLimit()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetLikes",
			"In", in,
			"Error", err,
		)
		return &npool.GetLikesResponse{}, status.Error(codes.Aborted, err.Error())
	}

	infos, total, err := handler.GetLikes(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetLikes",
			"In", in,
			"Error", err,
		)
		return &npool.GetLikesResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.GetLikesResponse{
		Infos: infos,
		Total: total,
	}, nil
}
