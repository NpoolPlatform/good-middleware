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
		good1.WithEntID(req.EntID, false),
		good1.WithDeviceInfoID(req.DeviceInfoID, true),
		good1.WithCoinTypeID(req.CoinTypeID, true),
		good1.WithVendorLocationID(req.VendorLocationID, true),
		good1.WithUnitPrice(req.UnitPrice, true),
		good1.WithBenefitType(req.BenefitType, true),
		good1.WithGoodType(req.GoodType, true),
		good1.WithTitle(req.Title, true),
		good1.WithQuantityUnit(req.QuantityUnit, true),
		good1.WithQuantityUnitAmount(req.QuantityUnitAmount, true),
		good1.WithDeliveryAt(req.DeliveryAt, true),
		good1.WithStartAt(req.StartAt, true),
		good1.WithStartMode(req.StartMode, false),
		good1.WithTestOnly(req.TestOnly, true),
		good1.WithTotal(req.Total, true),
		good1.WithPosters(req.Posters, true),
		good1.WithLabels(req.Labels, true),
		good1.WithBenefitIntervalHours(req.BenefitIntervalHours, true),
		good1.WithUnitLockDeposit(req.UnitLockDeposit, true),
		good1.WithUnitType(req.UnitType, false),
		good1.WithQuantityCalculateType(req.QuantityCalculateType, false),
		good1.WithDurationType(req.DurationType, false),
		good1.WithDurationCalculateType(req.DurationCalculateType, false),
		good1.WithSettlementType(req.SettlementType, false),
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
