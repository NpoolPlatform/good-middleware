package constraint

import (
	"context"

	constraint1 "github.com/NpoolPlatform/good-middleware/pkg/mw/app/good/topmost/constraint"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/app/good/topmost/constraint"
)

func (s *Server) UpdateTopMostConstraint(ctx context.Context, in *npool.UpdateTopMostConstraintRequest) (*npool.UpdateTopMostConstraintResponse, error) {
	req := in.GetInfo()
	if req == nil {
		logger.Sugar().Errorw(
			"UpdateTopMostConstraint",
			"In", in,
		)
		return &npool.UpdateTopMostConstraintResponse{}, status.Error(codes.Aborted, "invalid argument")
	}
	handler, err := constraint1.NewHandler(
		ctx,
		constraint1.WithID(req.ID, false),
		constraint1.WithEntID(req.EntID, false),
		constraint1.WithTargetValue(req.TargetValue, true),
		constraint1.WithIndex(req.Index, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateTopMostConstraint",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateTopMostConstraintResponse{}, status.Error(codes.Aborted, err.Error())
	}

	if err := handler.UpdateConstraint(ctx); err != nil {
		logger.Sugar().Errorw(
			"UpdateTopMostConstraint",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateTopMostConstraintResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.UpdateTopMostConstraintResponse{}, nil
}
