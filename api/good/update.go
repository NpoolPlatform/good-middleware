//nolint:dupl
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
		good1.WithDurationDays(req.DurationDays, false),
		good1.WithCoinTypeID(req.CoinTypeID, false),
		good1.WithVendorLocationID(req.VendorLocationID, false),
		good1.WithPrice(req.Price, false),
		good1.WithTitle(req.Title, false),
		good1.WithUnit(req.Unit, false),
		good1.WithUnitAmount(req.UnitAmount, false),
		good1.WithDeliveryAt(req.DeliveryAt, false),
		good1.WithStartAt(req.StartAt, false),
		good1.WithTestOnly(req.TestOnly, false),
		good1.WithBenefitIntervalHours(req.BenefitIntervalHours, false),
		good1.WithUnitLockDeposit(req.UnitLockDeposit, false),
		good1.WithRewardState(req.RewardState, false),
		good1.WithRewardAt(req.RewardAt, false),
		good1.WithRewardTID(req.RewardTID, false),
		good1.WithNextRewardStartAmount(req.NextRewardStartAmount, false),
		good1.WithRewardAmount(req.RewardAmount, false),
		good1.WithUnitRewardAmount(req.UnitRewardAmount, false),
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
