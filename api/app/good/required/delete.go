package required

import (
	"context"

	required1 "github.com/NpoolPlatform/good-middleware/pkg/mw/app/good/required"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/app/good/required"
)

func (s *Server) DeleteRequired(ctx context.Context, in *npool.DeleteRequiredRequest) (*npool.DeleteRequiredResponse, error) {
	req := in.GetInfo()
	if req == nil {
		logger.Sugar().Errorw(
			"DeleteRequired",
			"In", in,
		)
		return &npool.DeleteRequiredResponse{}, status.Error(codes.Aborted, "invalid argument")
	}
	handler, err := required1.NewHandler(
		ctx,
		required1.WithID(req.ID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteRequired",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteRequiredResponse{}, status.Error(codes.Aborted, err.Error())
	}

	if err := handler.DeleteRequired(ctx); err != nil {
		logger.Sugar().Errorw(
			"DeleteRequired",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteRequiredResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.DeleteRequiredResponse{}, nil
}
