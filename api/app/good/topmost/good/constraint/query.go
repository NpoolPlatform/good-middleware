package constraint

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	constraint1 "github.com/NpoolPlatform/good-middleware/pkg/mw/app/good/topmost/good/constraint"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/app/good/topmost/good/constraint"
)

func (s *Server) GetTopMostGoodConstraint(ctx context.Context, in *npool.GetTopMostGoodConstraintRequest) (*npool.GetTopMostGoodConstraintResponse, error) {
	handler, err := constraint1.NewHandler(
		ctx,
		constraint1.WithEntID(&in.EntID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetTopMostGoodConstraint",
			"In", in,
			"Error", err,
		)
		return &npool.GetTopMostGoodConstraintResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.GetConstraint(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetTopMostGoodConstraint",
			"In", in,
			"Error", err,
		)
		return &npool.GetTopMostGoodConstraintResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.GetTopMostGoodConstraintResponse{
		Info: info,
	}, nil
}

func (s *Server) GetTopMostGoodConstraints(ctx context.Context, in *npool.GetTopMostGoodConstraintsRequest) (*npool.GetTopMostGoodConstraintsResponse, error) {
	handler, err := constraint1.NewHandler(
		ctx,
		constraint1.WithConds(in.GetConds()),
		constraint1.WithOffset(in.GetOffset()),
		constraint1.WithLimit(in.GetLimit()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetTopMostGoodConstraints",
			"In", in,
			"Error", err,
		)
		return &npool.GetTopMostGoodConstraintsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	infos, total, err := handler.GetConstraints(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetTopMostGoodConstraints",
			"In", in,
			"Error", err,
		)
		return &npool.GetTopMostGoodConstraintsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.GetTopMostGoodConstraintsResponse{
		Infos: infos,
		Total: total,
	}, nil
}
