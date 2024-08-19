package topmost

import (
	"context"

	topmost1 "github.com/NpoolPlatform/good-middleware/pkg/mw/app/good/topmost"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/app/good/topmost"
)

func (s *Server) UpdateTopMost(ctx context.Context, in *npool.UpdateTopMostRequest) (*npool.UpdateTopMostResponse, error) {
	req := in.GetInfo()
	if req == nil {
		logger.Sugar().Errorw(
			"UpdateTopMost",
			"In", in,
		)
		return &npool.UpdateTopMostResponse{}, status.Error(codes.Aborted, "invalid argument")
	}
	handler, err := topmost1.NewHandler(
		ctx,
		topmost1.WithID(req.ID, false),
		topmost1.WithEntID(req.EntID, false),
		topmost1.WithTopMostType(req.TopMostType, false),
		topmost1.WithTitle(req.Title, false),
		topmost1.WithMessage(req.Message, false),
		topmost1.WithTargetURL(req.TargetUrl, false),
		topmost1.WithStartAt(req.StartAt, false),
		topmost1.WithEndAt(req.EndAt, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateTopMost",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateTopMostResponse{}, status.Error(codes.Aborted, err.Error())
	}

	if err := handler.UpdateTopMost(ctx); err != nil {
		logger.Sugar().Errorw(
			"UpdateTopMost",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateTopMostResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.UpdateTopMostResponse{}, nil
}
