package constraint

import (
	"context"

	constraint1 "github.com/NpoolPlatform/good-middleware/pkg/mw/app/good/topmost/good/constraint"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/app/good/topmost/good/constraint"
)

func (s *Server) DeleteTopMostGoodConstraint(ctx context.Context, in *npool.DeleteTopMostGoodConstraintRequest) (*npool.DeleteTopMostGoodConstraintResponse, error) {
	req := in.GetInfo()
	if req == nil {
		logger.Sugar().Errorw(
			"DeleteTopMostGoodConstraint",
			"In", in,
		)
		return &npool.DeleteTopMostGoodConstraintResponse{}, status.Error(codes.Aborted, "invalid argument")
	}
	handler, err := constraint1.NewHandler(
		ctx,
		constraint1.WithID(req.ID, false),
		constraint1.WithEntID(req.EntID, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteTopMostGoodConstraint",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteTopMostGoodConstraintResponse{}, status.Error(codes.Aborted, err.Error())
	}

	if err := handler.DeleteConstraint(ctx); err != nil {
		logger.Sugar().Errorw(
			"DeleteTopMostGoodConstraint",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteTopMostGoodConstraintResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.DeleteTopMostGoodConstraintResponse{}, nil
}
