package constraint

import (
	"context"

	constraint1 "github.com/NpoolPlatform/good-middleware/pkg/mw/app/good/topmost/constraint"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/app/good/topmost/constraint"
)

func (s *Server) CreateTopMostConstraint(ctx context.Context, in *npool.CreateTopMostConstraintRequest) (*npool.CreateTopMostConstraintResponse, error) {
	req := in.GetInfo()
	if req == nil {
		logger.Sugar().Errorw(
			"CreateTopMostConstraint",
			"In", in,
		)
		return &npool.CreateTopMostConstraintResponse{}, status.Error(codes.Aborted, "invalid argument")
	}
	handler, err := constraint1.NewHandler(
		ctx,
		constraint1.WithEntID(req.EntID, false),
		constraint1.WithTopMostID(req.TopMostID, true),
		constraint1.WithConstraint(req.Constraint, false),
		constraint1.WithTargetValue(req.TargetValue, true),
		constraint1.WithIndex(req.Index, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateTopMostConstraint",
			"In", in,
			"Error", err,
		)
		return &npool.CreateTopMostConstraintResponse{}, status.Error(codes.Aborted, err.Error())
	}

	if err := handler.CreateConstraint(ctx); err != nil {
		logger.Sugar().Errorw(
			"CreateTopMostConstraint",
			"In", in,
			"Error", err,
		)
		return &npool.CreateTopMostConstraintResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.CreateTopMostConstraintResponse{}, nil
}
