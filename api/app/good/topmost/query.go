package topmost

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	topmost1 "github.com/NpoolPlatform/good-middleware/pkg/mw/app/good/topmost"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/app/good/topmost"
)

func (s *Server) GetTopMost(ctx context.Context, in *npool.GetTopMostRequest) (*npool.GetTopMostResponse, error) {
	handler, err := topmost1.NewHandler(
		ctx,
		topmost1.WithID(&in.ID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetTopMost",
			"In", in,
			"Error", err,
		)
		return &npool.GetTopMostResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.GetTopMost(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetTopMost",
			"In", in,
			"Error", err,
		)
		return &npool.GetTopMostResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.GetTopMostResponse{
		Info: info,
	}, nil
}

func (s *Server) GetTopMosts(ctx context.Context, in *npool.GetTopMostsRequest) (*npool.GetTopMostsResponse, error) {
	handler, err := topmost1.NewHandler(
		ctx,
		topmost1.WithConds(in.GetConds()),
		topmost1.WithOffset(in.GetOffset()),
		topmost1.WithLimit(in.GetLimit()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetTopMosts",
			"In", in,
			"Error", err,
		)
		return &npool.GetTopMostsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	infos, total, err := handler.GetTopMosts(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetTopMosts",
			"In", in,
			"Error", err,
		)
		return &npool.GetTopMostsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.GetTopMostsResponse{
		Infos: infos,
		Total: total,
	}, nil
}
