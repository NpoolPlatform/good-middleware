package description

import (
	"context"

	description1 "github.com/NpoolPlatform/good-middleware/pkg/mw/app/good/description"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/app/good/description"
)

func (s *Server) UpdateDescription(ctx context.Context, in *npool.UpdateDescriptionRequest) (*npool.UpdateDescriptionResponse, error) {
	req := in.GetInfo()
	if req == nil {
		logger.Sugar().Errorw(
			"UpdateDescription",
			"In", in,
		)
		return &npool.UpdateDescriptionResponse{}, status.Error(codes.Aborted, "invalid argument")
	}
	handler, err := description1.NewHandler(
		ctx,
		description1.WithID(req.ID, false),
		description1.WithEntID(req.EntID, false),
		description1.WithDescription(req.Description, false),
		description1.WithIndex(func() *uint8 {
			if req.Index == nil {
				return nil
			}
			u := uint8(*req.Index)
			return &u
		}(), false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateDescription",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateDescriptionResponse{}, status.Error(codes.Aborted, err.Error())
	}

	if err := handler.UpdateDescription(ctx); err != nil {
		logger.Sugar().Errorw(
			"UpdateDescription",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateDescriptionResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.UpdateDescriptionResponse{}, nil
}
