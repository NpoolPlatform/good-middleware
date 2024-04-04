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
		powerrental1.WithGoodID(req.GoodID, true),

		powerrental1.WithDeviceTypeID(req.DeviceTypeID, true),
		powerrental1.WithVendorLocationID(req.VendorLocationID, true),
		powerrental1.WithUnitPrice(req.UnitPrice, true),
		powerrental1.WithQuantityUnit(req.QuantityUnit, true),
		powerrental1.WithQuantityUnitAmount(req.QuantityUnitAmount, true),
		powerrental1.WithDeliveryAt(req.DeliveryAt, true),
		powerrental1.WithUnitLockDeposit(req.UnitLockDeposit, true),
		powerrental1.WithDurationType(req.DurationType, true),

		powerrental1.WithGoodType(req.GoodType, true),
		powerrental1.WithBenefitType(req.BenefitType, true),
		powerrental1.WithName(req.Name, true),
		powerrental1.WithServiceStartAt(req.ServiceStartAt, true),
		powerrental1.WithStartMode(req.StartMode, true),
		powerrental1.WithTestOnly(req.TestOnly, true),
		powerrental1.WithBenefitIntervalHours(req.BenefitIntervalHours, true),
		powerrental1.WithPurchasable(req.Purchasable, true),
		powerrental1.WithOnline(req.Online, true),

		powerrental1.WithStockID(req.StockID, true),
		powerrental1.WithTotal(req.Total, true),
		powerrental1.WithStockMode(req.StockMode, true),
		powerrental1.WithStocks(req.MiningGoodStocks, true),

		powerrental1.WithRewardState(req.RewardState, true),
		powerrental1.WithRewardAt(req.RewardAt, true),
		powerrental1.WithRewardTID(req.RewardTID, true),
		powerrental1.WithNextRewardStartAmount(req.NextRewardStartAmount, true),
		powerrental1.WithRewardAmount(req.RewardAmount, true),
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
