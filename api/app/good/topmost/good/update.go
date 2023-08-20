//nolint:dupl
package topmostgood

import (
	"context"

	topmostgood1 "github.com/NpoolPlatform/good-middleware/pkg/mw/app/good/topmost/good"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/app/good/topmost/good"
)

func (s *Server) UpdateTopMost(ctx context.Context, in *npool.UpdateTopMostGoodRequest) (*npool.UpdateTopMostGoodResponse, error) {
	req := in.GetInfo()
	if req == nil {
		logger.Sugar().Errorw(
			"UpdateTopMostGood",
			"In", in,
		)
		return &npool.UpdateTopMostGoodResponse{}, status.Error(codes.Aborted, "invalid argument")
	}
	handler, err := topmostgood1.NewHandler(
		ctx,
		topmostgood1.WithID(req.ID, true),
		topmostgood1.WithAppID(req.AppID, true),
		topmostgood1.WithGoodID(req.GoodID, true),
		topmostgood1.WithAppGoodID(req.AppGoodID, true),
		topmostgood1.WithTopMostID(req.TopMostID, true),
		topmostgood1.WithDisplayIndex(req.DisplayIndex, false),
		topmostgood1.WithPosters(req.Posters, false),
		topmostgood1.WithPrice(req.Price, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateTopMostGood",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateTopMostGoodResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.UpdateTopMostGood(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateTopMostGood",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateTopMostGoodResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.UpdateTopMostGoodResponse{
		Info: info,
	}, nil
}