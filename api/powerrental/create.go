package powerrental

import (
	"context"

	powerrental1 "github.com/NpoolPlatform/good-middleware/pkg/mw/powerrental"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/powerrental"
)

func (s *Server) CreatePowerRental(ctx context.Context, in *npool.CreatePowerRentalRequest) (*npool.CreatePowerRentalResponse, error) {
	req := in.GetInfo()
	if req == nil {
		logger.Sugar().Errorw(
			"CreatePowerRental",
			"In", in,
		)
		return &npool.CreatePowerRentalResponse{}, status.Error(codes.Aborted, "invalid argument")
	}
	handler, err := powerrental1.NewHandler(
		ctx,
		powerrental1.WithEntID(req.EntID, false),
		powerrental1.WithGoodID(req.GoodID, false),
		powerrental1.WithDeviceTypeID(req.DeviceTypeID, true),
		powerrental1.WithVendorLocationID(req.VendorLocationID, true),
		powerrental1.WithUnitPrice(req.UnitPrice, true),
		powerrental1.WithQuantityUnit(req.QuantityUnit, true),
		powerrental1.WithQuantityUnitAmount(req.QuantityUnitAmount, true),
		powerrental1.WithDeliveryAt(req.DeliveryAt, true),
		powerrental1.WithUnitLockDeposit(req.UnitLockDeposit, false),
		powerrental1.WithDurationType(req.DurationType, false),

		powerrental1.WithGoodType(req.GoodType, true),
		powerrental1.WithBenefitType(req.BenefitType, true),
		powerrental1.WithName(req.Name, true),
		powerrental1.WithServiceStartAt(req.ServiceStartAt, true),
		powerrental1.WithStartMode(req.StartMode, true),
		powerrental1.WithTestOnly(req.TestOnly, false),
		powerrental1.WithBenefitIntervalHours(req.BenefitIntervalHours, true),
		powerrental1.WithPurchasable(req.Purchasable, false),
		powerrental1.WithOnline(req.Online, false),

		powerrental1.WithStockID(req.StockID, false),
		powerrental1.WithTotal(req.Total, true),
		powerrental1.WithStockMode(req.StockMode, true),
		powerrental1.WithStocks(req.MiningGoodStocks, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CreatePowerRental",
			"In", in,
			"Error", err,
		)
		return &npool.CreatePowerRentalResponse{}, status.Error(codes.Aborted, err.Error())
	}

	if err = handler.CreatePowerRental(ctx); err != nil {
		logger.Sugar().Errorw(
			"CreatePowerRental",
			"In", in,
			"Error", err,
		)
		return &npool.CreatePowerRentalResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.CreatePowerRentalResponse{}, nil
}
