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
		topmost1.WithID(req.ID, true),
		topmost1.WithAppID(req.AppID, true),
		topmost1.WithTopMostType(req.TopMostType, true),
		topmost1.WithTitle(req.Title, true),
		topmost1.WithMessage(req.Message, true),
		topmost1.WithPosters(req.Posters, true),
		topmost1.WithStartAt(req.StartAt, true),
		topmost1.WithEndAt(req.EndAt, true),
		topmost1.WithThresholdCredits(req.ThresholdCredits, true),
		topmost1.WithRegisterElapsedSeconds(req.RegisterElapsedSeconds, true),
		topmost1.WithThresholdPurchases(req.ThresholdPurchases, true),
		topmost1.WithThresholdPaymentAmount(req.ThresholdPaymentAmount, true),
		topmost1.WithKycMust(req.KycMust, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateTopMost",
			"In", in,
			"Error", err,
		)
		return &npool.CreateTopMostResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.CreateTopMost(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateTopMost",
			"In", in,
			"Error", err,
		)
		return &npool.CreateTopMostResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.CreateTopMostResponse{
		Info: info,
	}, nil
}
