package powerrental

import (
	"context"

	powerrental1 "github.com/NpoolPlatform/good-middleware/pkg/mw/app/powerrental"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/app/powerrental"
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
		powerrental1.WithAppGoodID(req.AppGoodID, false),

		powerrental1.WithPurchasable(req.Purchasable, false),
		powerrental1.WithEnableProductPage(req.EnableProductPage, false),
		powerrental1.WithProductPage(req.ProductPage, false),
		powerrental1.WithOnline(req.Online, false),
		powerrental1.WithVisible(req.Visible, false),
		powerrental1.WithName(req.Name, false),
		powerrental1.WithDisplayIndex(req.DisplayIndex, false),
		powerrental1.WithBanner(req.Banner, false),

		powerrental1.WithServiceStartAt(req.ServiceStartAt, false),
		powerrental1.WithStartMode(req.StartMode, false),
		powerrental1.WithCancelMode(req.CancelMode, false),
		powerrental1.WithCancelableBeforeStartSeconds(req.CancelableBeforeStartSeconds, false),
		powerrental1.WithEnableSetCommission(req.EnableSetCommission, false),
		powerrental1.WithMinOrderAmount(req.MinOrderAmount, false),
		powerrental1.WithMaxOrderAmount(req.MaxOrderAmount, false),
		powerrental1.WithMaxUserAmount(req.MaxUserAmount, false),
		powerrental1.WithMinOrderDurationSeconds(req.MinOrderDurationSeconds, false),
		powerrental1.WithMaxOrderDurationSeconds(req.MaxOrderDurationSeconds, false),
		powerrental1.WithUnitPrice(req.UnitPrice, false),
		powerrental1.WithSaleStartAt(req.SaleStartAt, false),
		powerrental1.WithSaleEndAt(req.SaleEndAt, false),
		powerrental1.WithSaleMode(req.SaleMode, false),
		powerrental1.WithFixedDuration(req.FixedDuration, false),
		powerrental1.WithPackageWithRequireds(req.PackageWithRequireds, false),
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
