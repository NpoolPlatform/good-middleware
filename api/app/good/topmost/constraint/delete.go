package constraint

import (
	"context"

	constraint1 "github.com/NpoolPlatform/good-middleware/pkg/mw/app/good/topmost/constraint"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/app/good/topmost/constraint"
)

func (s *Server) DeleteTopMostConstraint(ctx context.Context, in *npool.DeleteTopMostConstraintRequest) (*npool.DeleteTopMostConstraintResponse, error) {
	req := in.GetInfo()
	if req == nil {
		logger.Sugar().Errorw(
			"DeleteTopMostConstraint",
			"In", in,
		)
		return &npool.DeleteTopMostConstraintResponse{}, status.Error(codes.Aborted, "invalid argument")
	}
	handler, err := constraint1.NewHandler(
		ctx,
		constraint1.WithID(req.ID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteTopMostConstraint",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteTopMostConstraintResponse{}, status.Error(codes.Aborted, err.Error())
	}

	if err := handler.DeleteConstraint(ctx); err != nil {
		logger.Sugar().Errorw(
			"DeleteTopMostConstraint",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteTopMostConstraintResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.DeleteTopMostConstraintResponse{}, nil
}
