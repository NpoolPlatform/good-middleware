package label

import (
	"context"

	label1 "github.com/NpoolPlatform/good-middleware/pkg/mw/app/good/label"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/app/good/label"
)

func (s *Server) DeleteLabel(ctx context.Context, in *npool.DeleteLabelRequest) (*npool.DeleteLabelResponse, error) {
	req := in.GetInfo()
	if req == nil {
		logger.Sugar().Errorw(
			"DeleteLabel",
			"In", in,
		)
		return &npool.DeleteLabelResponse{}, status.Error(codes.Aborted, "invalid argument")
	}
	handler, err := label1.NewHandler(
		ctx,
		label1.WithID(req.ID, false),
		label1.WithEntID(req.EntID, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteLabel",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteLabelResponse{}, status.Error(codes.Aborted, err.Error())
	}

	if err := handler.DeleteLabel(ctx); err != nil {
		logger.Sugar().Errorw(
			"DeleteLabel",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteLabelResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.DeleteLabelResponse{}, nil
}
