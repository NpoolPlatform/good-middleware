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

func (s *Server) UpdateMalfunction(ctx context.Context, in *npool.UpdateMalfunctionRequest) (*npool.UpdateMalfunctionResponse, error) {
	req := in.GetInfo()
	if req == nil {
		logger.Sugar().Errorw(
			"UpdateMalfunction",
			"In", in,
		)
		return &npool.UpdateMalfunctionResponse{}, status.Error(codes.Aborted, "invalid argument")
	}
	handler, err := malfunction1.NewHandler(
		ctx,
		malfunction1.WithID(req.ID, false),
		malfunction1.WithEntID(req.EntID, false),
		malfunction1.WithTitle(req.Title, false),
		malfunction1.WithMessage(req.Message, false),
		malfunction1.WithStartAt(req.StartAt, false),
		malfunction1.WithDurationSeconds(req.DurationSeconds, false),
		malfunction1.WithCompensateSeconds(req.CompensateSeconds, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateMalfunction",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateMalfunctionResponse{}, status.Error(codes.Aborted, err.Error())
	}

	if err := handler.UpdateMalfunction(ctx); err != nil {
		logger.Sugar().Errorw(
			"UpdateMalfunction",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateMalfunctionResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.UpdateMalfunctionResponse{}, nil
}
