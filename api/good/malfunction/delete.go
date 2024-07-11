package malfunction

import (
	"context"

	malfunction1 "github.com/NpoolPlatform/good-middleware/pkg/mw/good/malfunction"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/good/malfunction"
)

func (s *Server) DeleteMalfunction(ctx context.Context, in *npool.DeleteMalfunctionRequest) (*npool.DeleteMalfunctionResponse, error) {
	req := in.GetInfo()
	if req == nil {
		logger.Sugar().Errorw(
			"DeleteMalfunction",
			"In", in,
		)
		return &npool.DeleteMalfunctionResponse{}, status.Error(codes.Aborted, "invalid argument")
	}
	handler, err := malfunction1.NewHandler(
		ctx,
		malfunction1.WithID(req.ID, false),
		malfunction1.WithEntID(req.EntID, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteMalfunction",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteMalfunctionResponse{}, status.Error(codes.Aborted, err.Error())
	}

	if err := handler.DeleteMalfunction(ctx); err != nil {
		logger.Sugar().Errorw(
			"DeleteMalfunction",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteMalfunctionResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.DeleteMalfunctionResponse{}, nil
}
