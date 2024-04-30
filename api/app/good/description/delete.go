package description

import (
	"context"

	description1 "github.com/NpoolPlatform/good-middleware/pkg/mw/app/good/description"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/app/good/description"
)

func (s *Server) DeleteDescription(ctx context.Context, in *npool.DeleteDescriptionRequest) (*npool.DeleteDescriptionResponse, error) {
	req := in.GetInfo()
	if req == nil {
		logger.Sugar().Errorw(
			"DeleteDescription",
			"In", in,
		)
		return &npool.DeleteDescriptionResponse{}, status.Error(codes.Aborted, "invalid argument")
	}
	handler, err := description1.NewHandler(
		ctx,
		description1.WithID(req.ID, false),
		description1.WithEntID(req.EntID, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteDescription",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteDescriptionResponse{}, status.Error(codes.Aborted, err.Error())
	}

	if err := handler.DeleteDescription(ctx); err != nil {
		logger.Sugar().Errorw(
			"DeleteDescription",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteDescriptionResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.DeleteDescriptionResponse{}, nil
}
