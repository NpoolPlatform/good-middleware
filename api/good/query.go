package good

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	good1 "github.com/NpoolPlatform/good-middleware/pkg/mw/good"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/good"
)

func (s *Server) GetGood(ctx context.Context, in *npool.GetGoodRequest) (*npool.GetGoodResponse, error) {
	handler, err := good1.NewHandler(
		ctx,
		good1.WithEntID(&in.EntID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetGood",
			"In", in,
			"Error", err,
		)
		return &npool.GetGoodResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.GetGood(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetGood",
			"In", in,
			"Error", err,
		)
		return &npool.GetGoodResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.GetGoodResponse{
		Info: info,
	}, nil
}

func (s *Server) GetGoods(ctx context.Context, in *npool.GetGoodsRequest) (*npool.GetGoodsResponse, error) {
	handler, err := good1.NewHandler(
		ctx,
		good1.WithConds(in.GetConds()),
		good1.WithOffset(in.GetOffset()),
		good1.WithLimit(in.GetLimit()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetGoods",
			"In", in,
			"Error", err,
		)
		return &npool.GetGoodsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	infos, total, err := handler.GetGoods(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetGoods",
			"In", in,
			"Error", err,
		)
		return &npool.GetGoodsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.GetGoodsResponse{
		Infos: infos,
		Total: total,
	}, nil
}
