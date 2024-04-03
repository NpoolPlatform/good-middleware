package poster

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	poster1 "github.com/NpoolPlatform/good-middleware/pkg/mw/app/good/poster"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/app/good/poster"
)

func (s *Server) GetPoster(ctx context.Context, in *npool.GetPosterRequest) (*npool.GetPosterResponse, error) {
	handler, err := poster1.NewHandler(
		ctx,
		poster1.WithEntID(&in.EntID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetPoster",
			"In", in,
			"Error", err,
		)
		return &npool.GetPosterResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.GetPoster(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetPoster",
			"In", in,
			"Error", err,
		)
		return &npool.GetPosterResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.GetPosterResponse{
		Info: info,
	}, nil
}

func (s *Server) GetPosters(ctx context.Context, in *npool.GetPostersRequest) (*npool.GetPostersResponse, error) {
	handler, err := poster1.NewHandler(
		ctx,
		poster1.WithConds(in.GetConds()),
		poster1.WithOffset(in.GetOffset()),
		poster1.WithLimit(in.GetLimit()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetPosters",
			"In", in,
			"Error", err,
		)
		return &npool.GetPostersResponse{}, status.Error(codes.Aborted, err.Error())
	}

	infos, total, err := handler.GetPosters(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetPosters",
			"In", in,
			"Error", err,
		)
		return &npool.GetPostersResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.GetPostersResponse{
		Infos: infos,
		Total: total,
	}, nil
}
