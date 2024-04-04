package constraint

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	constraint1 "github.com/NpoolPlatform/good-middleware/pkg/mw/app/good/topmost/constraint"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/app/good/topmost/constraint"
)

func (s *Server) GetTopMostConstraint(ctx context.Context, in *npool.GetTopMostConstraintRequest) (*npool.GetTopMostConstraintResponse, error) {
	handler, err := constraint1.NewHandler(
		ctx,
		constraint1.WithEntID(&in.EntID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetTopMostConstraint",
			"In", in,
			"Error", err,
		)
		return &npool.GetTopMostConstraintResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.GetConstraint(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetTopMostConstraint",
			"In", in,
			"Error", err,
		)
		return &npool.GetTopMostConstraintResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.GetTopMostConstraintResponse{
		Info: info,
	}, nil
}

func (s *Server) GetTopMostConstraints(ctx context.Context, in *npool.GetTopMostConstraintsRequest) (*npool.GetTopMostConstraintsResponse, error) {
	handler, err := constraint1.NewHandler(
		ctx,
		constraint1.WithConds(in.GetConds()),
		constraint1.WithOffset(in.GetOffset()),
		constraint1.WithLimit(in.GetLimit()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetTopMostConstraints",
			"In", in,
			"Error", err,
		)
		return &npool.GetTopMostConstraintsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	infos, total, err := handler.GetConstraints(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetTopMostConstraints",
			"In", in,
			"Error", err,
		)
		return &npool.GetTopMostConstraintsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.GetTopMostConstraintsResponse{
		Infos: infos,
		Total: total,
	}, nil
}
