package required

import (
	"context"

	required1 "github.com/NpoolPlatform/good-middleware/pkg/mw/good/required"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/good/required"
)

func (s *Server) UpdateRequired(ctx context.Context, in *npool.UpdateRequiredRequest) (*npool.UpdateRequiredResponse, error) {
	req := in.GetInfo()
	if req == nil {
		logger.Sugar().Errorw(
			"UpdateRequired",
			"In", in,
		)
		return &npool.UpdateRequiredResponse{}, status.Error(codes.Aborted, "invalid argument")
	}
	handler, err := required1.NewHandler(
		ctx,
		required1.WithID(req.ID, false),
		required1.WithEntID(req.EntID, false),
		required1.WithMust(req.Must, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateRequired",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateRequiredResponse{}, status.Error(codes.Aborted, err.Error())
	}

	if err := handler.UpdateRequired(ctx); err != nil {
		logger.Sugar().Errorw(
			"UpdateRequired",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateRequiredResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.UpdateRequiredResponse{}, nil
}
