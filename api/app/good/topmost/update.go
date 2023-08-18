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
		topmost1.WithID(req.ID, true),
		topmost1.WithAppID(req.AppID, true),
		topmost1.WithTopMostType(req.TopMostType, true),
		topmost1.WithTitle(req.Title, false),
		topmost1.WithMessage(req.Message, false),
		topmost1.WithPosters(req.Posters, false),
		topmost1.WithStartAt(req.StartAt, false),
		topmost1.WithEndAt(req.EndAt, false),
		topmost1.WithThresholdCredits(req.ThresholdCredits, false),
		topmost1.WithRegisterElapsedSeconds(req.RegisterElapsedSeconds, false),
		topmost1.WithThresholdPurchases(req.ThresholdPurchases, false),
		topmost1.WithThresholdPaymentAmount(req.ThresholdPaymentAmount, false),
		topmost1.WithKycMust(req.KycMust, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateTopMost",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateTopMostResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.UpdateTopMost(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateTopMost",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateTopMostResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.UpdateTopMostResponse{
		Info: info,
	}, nil
}
