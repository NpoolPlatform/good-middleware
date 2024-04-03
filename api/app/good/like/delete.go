package like

import (
	"context"

	like1 "github.com/NpoolPlatform/good-middleware/pkg/mw/app/good/like"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/app/good/like"
)

func (s *Server) DeleteLike(ctx context.Context, in *npool.DeleteLikeRequest) (*npool.DeleteLikeResponse, error) {
	req := in.GetInfo()
	if req == nil {
		logger.Sugar().Errorw(
			"DeleteLike",
			"In", in,
		)
		return &npool.DeleteLikeResponse{}, status.Error(codes.Aborted, "invalid argument")
	}
	handler, err := like1.NewHandler(
		ctx,
		like1.WithID(req.ID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteLike",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteLikeResponse{}, status.Error(codes.Aborted, err.Error())
	}

	if err := handler.DeleteLike(ctx); err != nil {
		logger.Sugar().Errorw(
			"DeleteLike",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteLikeResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.DeleteLikeResponse{}, nil
}
