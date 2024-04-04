package poster

import (
	"context"

	poster1 "github.com/NpoolPlatform/good-middleware/pkg/mw/device/poster"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/device/poster"
)

func (s *Server) DeletePoster(ctx context.Context, in *npool.DeletePosterRequest) (*npool.DeletePosterResponse, error) {
	req := in.GetInfo()
	if req == nil {
		logger.Sugar().Errorw(
			"DeletePoster",
			"In", in,
		)
		return &npool.DeletePosterResponse{}, status.Error(codes.Aborted, "invalid argument")
	}
	handler, err := poster1.NewHandler(
		ctx,
		poster1.WithID(req.ID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"DeletePoster",
			"In", in,
			"Error", err,
		)
		return &npool.DeletePosterResponse{}, status.Error(codes.Aborted, err.Error())
	}

	if err := handler.DeletePoster(ctx); err != nil {
		logger.Sugar().Errorw(
			"DeletePoster",
			"In", in,
			"Error", err,
		)
		return &npool.DeletePosterResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.DeletePosterResponse{}, nil
}
