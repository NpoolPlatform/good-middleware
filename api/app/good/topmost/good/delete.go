package topmostgood

import (
	"context"

	topmostgood1 "github.com/NpoolPlatform/good-middleware/pkg/mw/app/good/topmost/good"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/app/good/topmost/good"
)

func (s *Server) DeleteTopMostGood(ctx context.Context, in *npool.DeleteTopMostGoodRequest) (*npool.DeleteTopMostGoodResponse, error) {
	req := in.GetInfo()
	if req == nil {
		logger.Sugar().Errorw(
			"DeleteTopMostGood",
			"In", in,
		)
		return &npool.DeleteTopMostGoodResponse{}, status.Error(codes.Aborted, "invalid argument")
	}
	handler, err := topmostgood1.NewHandler(
		ctx,
		topmostgood1.WithID(req.ID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteTopMostGood",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteTopMostGoodResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.DeleteTopMostGood(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteTopMostGood",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteTopMostGoodResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.DeleteTopMostGoodResponse{
		Info: info,
	}, nil
}
