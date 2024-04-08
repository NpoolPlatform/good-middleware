package constraint

import (
	"context"

	constraint1 "github.com/NpoolPlatform/good-middleware/pkg/mw/app/good/topmost/good/constraint"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/app/good/topmost/good/constraint"
)

func (s *Server) UpdateTopMostGoodConstraint(ctx context.Context, in *npool.UpdateTopMostGoodConstraintRequest) (*npool.UpdateTopMostGoodConstraintResponse, error) {
	req := in.GetInfo()
	if req == nil {
		logger.Sugar().Errorw(
			"UpdateTopMostGoodConstraint",
			"In", in,
		)
		return &npool.UpdateTopMostGoodConstraintResponse{}, status.Error(codes.Aborted, "invalid argument")
	}
	handler, err := constraint1.NewHandler(
		ctx,
		constraint1.WithID(req.ID, false),
		constraint1.WithEntID(req.EntID, false),
		constraint1.WithTargetValue(req.TargetValue, false),
		constraint1.WithIndex(req.Index, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateTopMostGoodConstraint",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateTopMostGoodConstraintResponse{}, status.Error(codes.Aborted, err.Error())
	}

	if err := handler.UpdateConstraint(ctx); err != nil {
		logger.Sugar().Errorw(
			"UpdateTopMostGoodConstraint",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateTopMostGoodConstraintResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.UpdateTopMostGoodConstraintResponse{}, nil
}
