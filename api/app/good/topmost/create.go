//nolint:dupl
package topmost

import (
	"context"

	topmost1 "github.com/NpoolPlatform/good-middleware/pkg/mw/app/good/topmost"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/app/good/topmost"
)

func (s *Server) CreateTopMost(ctx context.Context, in *npool.CreateTopMostRequest) (*npool.CreateTopMostResponse, error) {
	req := in.GetInfo()
	if req == nil {
		logger.Sugar().Errorw(
			"CreateTopMost",
			"In", in,
		)
		return &npool.CreateTopMostResponse{}, status.Error(codes.Aborted, "invalid argument")
	}
	handler, err := topmost1.NewHandler(
		ctx,
		topmost1.WithEntID(req.EntID, false),
		topmost1.WithAppID(req.AppID, true),
		topmost1.WithTopMostType(req.TopMostType, true),
		topmost1.WithTitle(req.Title, true),
		topmost1.WithMessage(req.Message, true),
		topmost1.WithTargetURL(req.TargetUrl, false),
		topmost1.WithStartAt(req.StartAt, true),
		topmost1.WithEndAt(req.EndAt, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateTopMost",
			"In", in,
			"Error", err,
		)
		return &npool.CreateTopMostResponse{}, status.Error(codes.Aborted, err.Error())
	}

	if err := handler.CreateTopMost(ctx); err != nil {
		logger.Sugar().Errorw(
			"CreateTopMost",
			"In", in,
			"Error", err,
		)
		return &npool.CreateTopMostResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.CreateTopMostResponse{}, nil
}
