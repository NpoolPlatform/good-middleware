package topmostgood

import (
	"context"

	topmostgood1 "github.com/NpoolPlatform/good-middleware/pkg/mw/app/good/topmost/good"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/app/good/topmost/good"
)

func (s *Server) CreateTopMostGood(ctx context.Context, in *npool.CreateTopMostGoodRequest) (*npool.CreateTopMostGoodResponse, error) {
	req := in.GetInfo()
	if req == nil {
		logger.Sugar().Errorw(
			"CreateTopMostGood",
			"In", in,
		)
		return &npool.CreateTopMostGoodResponse{}, status.Error(codes.Aborted, "invalid argument")
	}
	handler, err := topmostgood1.NewHandler(
		ctx,
		topmostgood1.WithID(req.ID, true),
		topmostgood1.WithAppGoodID(req.AppGoodID, true),
		topmostgood1.WithTopMostID(req.TopMostID, true),
		topmostgood1.WithDisplayIndex(req.DisplayIndex, true),
		topmostgood1.WithPosters(req.Posters, true),
		topmostgood1.WithPrice(req.Price, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateTopMostGood",
			"In", in,
			"Error", err,
		)
		return &npool.CreateTopMostGoodResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.CreateTopMostGood(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateTopMostGood",
			"In", in,
			"Error", err,
		)
		return &npool.CreateTopMostGoodResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.CreateTopMostGoodResponse{
		Info: info,
	}, nil
}
