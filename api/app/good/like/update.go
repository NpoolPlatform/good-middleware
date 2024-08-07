package like

import (
	"context"

	like1 "github.com/NpoolPlatform/good-middleware/pkg/mw/app/good/like"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/app/good/like"
)

func (s *Server) UpdateLike(ctx context.Context, in *npool.UpdateLikeRequest) (*npool.UpdateLikeResponse, error) {
	req := in.GetInfo()
	if req == nil {
		logger.Sugar().Errorw(
			"UpdateLike",
			"In", in,
		)
		return &npool.UpdateLikeResponse{}, status.Error(codes.Aborted, "invalid argument")
	}
	handler, err := like1.NewHandler(
		ctx,
		like1.WithID(req.ID, false),
		like1.WithEntID(req.EntID, false),
		like1.WithLike(req.Like, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateLike",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateLikeResponse{}, status.Error(codes.Aborted, err.Error())
	}

	if err := handler.UpdateLike(ctx); err != nil {
		logger.Sugar().Errorw(
			"UpdateLike",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateLikeResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.UpdateLikeResponse{}, nil
}
