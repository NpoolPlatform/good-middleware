package poster

import (
	"context"

	poster1 "github.com/NpoolPlatform/good-middleware/pkg/mw/app/good/topmost/good/poster"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/app/good/topmost/good/poster"
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
		poster1.WithID(req.ID, false),
		poster1.WithEntID(req.EntID, false),
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
