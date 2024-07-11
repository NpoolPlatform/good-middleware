//nolint:dupl
package malfunction

import (
	"context"

	malfunction1 "github.com/NpoolPlatform/good-middleware/pkg/mw/good/malfunction"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/good/malfunction"
)

func (s *Server) CreateMalfunction(ctx context.Context, in *npool.CreateMalfunctionRequest) (*npool.CreateMalfunctionResponse, error) {
	req := in.GetInfo()
	if req == nil {
		logger.Sugar().Errorw(
			"CreateMalfunction",
			"In", in,
		)
		return &npool.CreateMalfunctionResponse{}, status.Error(codes.Aborted, "invalid argument")
	}
	handler, err := malfunction1.NewHandler(
		ctx,
		malfunction1.WithEntID(req.EntID, false),
		malfunction1.WithGoodID(req.GoodID, true),
		malfunction1.WithTitle(req.Title, true),
		malfunction1.WithMessage(req.Message, true),
		malfunction1.WithStartAt(req.StartAt, true),
		malfunction1.WithDurationSeconds(req.DurationSeconds, false),
		malfunction1.WithCompensateSeconds(req.CompensateSeconds, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateMalfunction",
			"In", in,
			"Error", err,
		)
		return &npool.CreateMalfunctionResponse{}, status.Error(codes.Aborted, err.Error())
	}

	if err := handler.CreateMalfunction(ctx); err != nil {
		logger.Sugar().Errorw(
			"CreateMalfunction",
			"In", in,
			"Error", err,
		)
		return &npool.CreateMalfunctionResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.CreateMalfunctionResponse{}, nil
}
