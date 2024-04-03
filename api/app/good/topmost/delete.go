package topmost

import (
	"context"

	topmost1 "github.com/NpoolPlatform/good-middleware/pkg/mw/app/good/topmost"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/app/good/topmost"
)

func (s *Server) DeleteTopMost(ctx context.Context, in *npool.DeleteTopMostRequest) (*npool.DeleteTopMostResponse, error) {
	req := in.GetInfo()
	if req == nil {
		logger.Sugar().Errorw(
			"DeleteTopMost",
			"In", in,
		)
		return &npool.DeleteTopMostResponse{}, status.Error(codes.Aborted, "invalid argument")
	}
	handler, err := topmost1.NewHandler(
		ctx,
		topmost1.WithID(req.ID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteTopMost",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteTopMostResponse{}, status.Error(codes.Aborted, err.Error())
	}

	if err := handler.DeleteTopMost(ctx); err != nil {
		logger.Sugar().Errorw(
			"DeleteTopMost",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteTopMostResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.DeleteTopMostResponse{}, nil
}
