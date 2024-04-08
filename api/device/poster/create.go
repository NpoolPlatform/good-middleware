//nolint:dupl
package poster

import (
	"context"

	poster1 "github.com/NpoolPlatform/good-middleware/pkg/mw/device/poster"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/device/poster"
)

func (s *Server) CreatePoster(ctx context.Context, in *npool.CreatePosterRequest) (*npool.CreatePosterResponse, error) {
	req := in.GetInfo()
	if req == nil {
		logger.Sugar().Errorw(
			"CreatePoster",
			"In", in,
		)
		return &npool.CreatePosterResponse{}, status.Error(codes.Aborted, "invalid argument")
	}
	handler, err := poster1.NewHandler(
		ctx,
		poster1.WithEntID(req.EntID, false),
		poster1.WithDeviceTypeID(req.DeviceTypeID, true),
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
			"CreatePoster",
			"In", in,
			"Error", err,
		)
		return &npool.CreatePosterResponse{}, status.Error(codes.Aborted, err.Error())
	}

	if err := handler.CreatePoster(ctx); err != nil {
		logger.Sugar().Errorw(
			"CreatePoster",
			"In", in,
			"Error", err,
		)
		return &npool.CreatePosterResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.CreatePosterResponse{}, nil
}
