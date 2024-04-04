package poster

import (
	"context"

	poster1 "github.com/NpoolPlatform/good-middleware/pkg/mw/app/good/topmost/poster"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/app/good/topmost/poster"
)

func (s *Server) UpdatePoster(ctx context.Context, in *npool.UpdatePosterRequest) (*npool.UpdatePosterResponse, error) {
	req := in.GetInfo()
	if req == nil {
		logger.Sugar().Errorw(
			"UpdatePoster",
			"In", in,
		)
		return &npool.UpdatePosterResponse{}, status.Error(codes.Aborted, "invalid argument")
	}
	handler, err := poster1.NewHandler(
		ctx,
		poster1.WithID(req.ID, false),
		poster1.WithEntID(req.EntID, false),
		poster1.WithPoster(req.Poster, false),
		poster1.WithIndex(func() *uint8 {
			if req.Index == nil {
				return nil
			}
			u := uint8(*req.Index)
			return &u
		}(), false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdatePoster",
			"In", in,
			"Error", err,
		)
		return &npool.UpdatePosterResponse{}, status.Error(codes.Aborted, err.Error())
	}

	if err := handler.UpdatePoster(ctx); err != nil {
		logger.Sugar().Errorw(
			"UpdatePoster",
			"In", in,
			"Error", err,
		)
		return &npool.UpdatePosterResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.UpdatePosterResponse{}, nil
}
