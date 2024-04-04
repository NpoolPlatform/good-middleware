package powerrental

import (
	"context"

	powerrental1 "github.com/NpoolPlatform/good-middleware/pkg/mw/powerrental"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/powerrental"
)

func (s *Server) UpdatePowerRental(ctx context.Context, in *npool.UpdatePowerRentalRequest) (*npool.UpdatePowerRentalResponse, error) {
	req := in.GetInfo()
	if req == nil {
		logger.Sugar().Errorw(
			"UpdatePowerRental",
			"In", in,
		)
		return &npool.UpdatePowerRentalResponse{}, status.Error(codes.Aborted, "invalid argument")
	}
	handler, err := powerrental1.NewHandler(
		ctx,
		powerrental1.WithID(req.ID, false),
		powerrental1.WithEntID(req.EntID, false),
		powerrental1.WithGoodID(req.GoodID, false),

		powerrental1.WithDeviceTypeID(req.DeviceTypeID, false),
		powerrental1.WithVendorLocationID(req.VendorLocationID, false),
		powerrental1.WithUnitPrice(req.UnitPrice, false),
		powerrental1.WithQuantityUnit(req.QuantityUnit, false),
		powerrental1.WithQuantityUnitAmount(req.QuantityUnitAmount, false),
		powerrental1.WithDeliveryAt(req.DeliveryAt, false),
		powerrental1.WithUnitLockDeposit(req.UnitLockDeposit, false),
		powerrental1.WithDurationType(req.DurationType, false),

		powerrental1.WithGoodType(req.GoodType, false),
		powerrental1.WithBenefitType(req.BenefitType, false),
		powerrental1.WithName(req.Name, false),
		powerrental1.WithServiceStartAt(req.ServiceStartAt, false),
		powerrental1.WithStartMode(req.StartMode, false),
		powerrental1.WithTestOnly(req.TestOnly, false),
		powerrental1.WithBenefitIntervalHours(req.BenefitIntervalHours, false),
		powerrental1.WithPurchasable(req.Purchasable, false),
		powerrental1.WithOnline(req.Online, false),

		powerrental1.WithTotal(req.Total, false),
		powerrental1.WithStockMode(req.StockMode, false),
		powerrental1.WithStocks(req.MiningGoodStocks, false),

		powerrental1.WithRewardState(req.RewardState, false),
		powerrental1.WithRewardAt(req.RewardAt, false),
		powerrental1.WithRewardTID(req.RewardTID, false),
		powerrental1.WithNextRewardStartAmount(req.NextRewardStartAmount, false),
		powerrental1.WithRewardAmount(req.RewardAmount, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdatePowerRental",
			"In", in,
			"Error", err,
		)
		return &npool.UpdatePowerRentalResponse{}, status.Error(codes.Aborted, err.Error())
	}

	if err := handler.UpdatePowerRental(ctx); err != nil {
		logger.Sugar().Errorw(
			"UpdatePowerRental",
			"In", in,
			"Error", err,
		)
		return &npool.UpdatePowerRentalResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.UpdatePowerRentalResponse{}, nil
}
