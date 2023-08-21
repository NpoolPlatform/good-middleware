package topmostgood

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	topmostgood1 "github.com/NpoolPlatform/good-middleware/pkg/mw/app/good/topmost/good"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/app/good/topmost/good"
)

func (s *Server) GetTopMostGood(ctx context.Context, in *npool.GetTopMostGoodRequest) (*npool.GetTopMostGoodResponse, error) {
	handler, err := topmostgood1.NewHandler(
		ctx,
		topmostgood1.WithID(&in.ID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetTopMostGood",
			"In", in,
			"Error", err,
		)
		return &npool.GetTopMostGoodResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.GetTopMostGood(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetTopMostGood",
			"In", in,
			"Error", err,
		)
		return &npool.GetTopMostGoodResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.GetTopMostGoodResponse{
		Info: info,
	}, nil
}

func (s *Server) GetTopMostGoods(ctx context.Context, in *npool.GetTopMostGoodsRequest) (*npool.GetTopMostGoodsResponse, error) {
	handler, err := topmostgood1.NewHandler(
		ctx,
		topmostgood1.WithConds(in.GetConds()),
		topmostgood1.WithOffset(in.GetOffset()),
		topmostgood1.WithLimit(in.GetLimit()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetTopMostGoods",
			"In", in,
			"Error", err,
		)
		return &npool.GetTopMostGoodsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	infos, total, err := handler.GetTopMostGoods(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetTopMostGoods",
			"In", in,
			"Error", err,
		)
		return &npool.GetTopMostGoodsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.GetTopMostGoodsResponse{
		Infos: infos,
		Total: total,
	}, nil
}
