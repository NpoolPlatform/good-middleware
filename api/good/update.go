package good

import (
	"context"

	good1 "github.com/NpoolPlatform/good-middleware/pkg/mw/good"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/good"
)

func (s *Server) UpdateGood(ctx context.Context, in *npool.UpdateGoodRequest) (*npool.UpdateGoodResponse, error) {
	req := in.GetInfo()
	if req == nil {
		logger.Sugar().Errorw(
			"UpdateGood",
			"In", in,
		)
		return &npool.UpdateGoodResponse{}, status.Error(codes.Aborted, "invalid argument")
	}
	handler, err := good1.NewHandler(
		ctx,
		good1.WithID(req.ID, true),
		good1.WithDeviceInfoID(req.DeviceInfoID, false),
		good1.WithCoinTypeID(req.CoinTypeID, false),
		good1.WithVendorLocationID(req.VendorLocationID, false),
		good1.WithUnitPrice(req.UnitPrice, false),
		good1.WithTitle(req.Title, false),
		good1.WithQuantityUnit(req.QuantityUnit, false),
		good1.WithQuantityUnitAmount(req.QuantityUnitAmount, false),
		good1.WithDeliveryAt(req.DeliveryAt, false),
		good1.WithStartAt(req.StartAt, false),
		good1.WithStartMode(req.StartMode, false),
		good1.WithTestOnly(req.TestOnly, false),
		good1.WithBenefitIntervalHours(req.BenefitIntervalHours, false),
		good1.WithUnitLockDeposit(req.UnitLockDeposit, false),
		good1.WithRewardState(req.RewardState, false),
		good1.WithRewardAt(req.RewardAt, false),
		good1.WithRewardTID(req.RewardTID, false),
		good1.WithNextRewardStartAmount(req.NextRewardStartAmount, false),
		good1.WithRewardAmount(req.RewardAmount, false),
		good1.WithUnitRewardAmount(req.UnitRewardAmount, false),
		good1.WithPosters(req.Posters, false),
		good1.WithLabels(req.Labels, false),
		good1.WithBenefitType(req.BenefitType, false),
		good1.WithTotal(req.Total, false),
		good1.WithUnitType(req.UnitType, false),
		good1.WithQuantityCalculateType(req.QuantityCalculateType, false),
		good1.WithDurationType(req.DurationType, false),
		good1.WithDurationCalculateType(req.DurationCalculateType, false),
		good1.WithSettlementType(req.SettlementType, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateGood",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateGoodResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.UpdateGood(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateGood",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateGoodResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.UpdateGoodResponse{
		Info: info,
	}, nil
}
