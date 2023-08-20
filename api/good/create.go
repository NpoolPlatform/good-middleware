package good

import (
	"context"

	good1 "github.com/NpoolPlatform/good-middleware/pkg/mw/good"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/good"
)

func (s *Server) CreateGood(ctx context.Context, in *npool.CreateGoodRequest) (*npool.CreateGoodResponse, error) {
	req := in.GetInfo()
	if req == nil {
		logger.Sugar().Errorw(
			"CreateGood",
			"In", in,
		)
		return &npool.CreateGoodResponse{}, status.Error(codes.Aborted, "invalid argument")
	}
	handler, err := good1.NewHandler(
		ctx,
		good1.WithID(req.ID, true),
		good1.WithDeviceInfoID(req.DeviceInfoID, true),
		good1.WithDurationDays(req.DurationDays, true),
		good1.WithCoinTypeID(req.CoinTypeID, true),
		good1.WithVendorLocationID(req.VendorLocationID, true),
		good1.WithPrice(req.Price, true),
		good1.WithBenefitType(req.BenefitType, true),
		good1.WithGoodType(req.GoodType, true),
		good1.WithTitle(req.Title, true),
		good1.WithUnit(req.Unit, true),
		good1.WithUnitAmount(req.UnitAmount, true),
		good1.WithSupportCoinTypeIDs(req.SupportCoinTypeIDs, true),
		good1.WithDeliveryAt(req.DeliveryAt, true),
		good1.WithStartAt(req.StartAt, true),
		good1.WithTestOnly(req.TestOnly, true),
		good1.WithTotal(req.Total, true),
		good1.WithPosters(req.Posters, true),
		good1.WithLabels(req.Labels, true),
		good1.WithBenefitIntervalHours(req.BenefitIntervalHours, true),
		good1.WithUnitLockDeposit(req.UnitLockDeposit, true),
		good1.WithAppReserved(req.AppReserved, true),
		good1.WithRewardState(req.RewardState, true),
		good1.WithRewardAt(req.RewardAt, true),
		good1.WithRewardTID(req.RewardTID, true),
		good1.WithNextRewardStartAmount(req.NextRewardStartAmount, true),
		good1.WithRewardAmount(req.RewardAmount, true),
		good1.WithUnitRewardAmount(req.UnitRewardAmount, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateGood",
			"In", in,
			"Error", err,
		)
		return &npool.CreateGoodResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.CreateGood(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateGood",
			"In", in,
			"Error", err,
		)
		return &npool.CreateGoodResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.CreateGoodResponse{
		Info: info,
	}, nil
}