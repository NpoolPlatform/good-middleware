package powerrental

import (
	"context"

	powerrental1 "github.com/NpoolPlatform/good-middleware/pkg/mw/app/powerrental"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/app/powerrental"
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
		powerrental1.WithAppID(req.AppID, true),
		powerrental1.WithGoodID(req.GoodID, true),
		powerrental1.WithAppGoodID(req.AppGoodID, true),
		powerrental1.WithPurchasable(req.Purchasable, true),
		powerrental1.WithEnableProductPage(req.EnableProductPage, true),
		powerrental1.WithProductPage(req.ProductPage, true),
		powerrental1.WithOnline(req.Online, true),
		powerrental1.WithVisible(req.Visible, true),
		powerrental1.WithName(req.Name, true),
		powerrental1.WithDisplayIndex(req.DisplayIndex, true),
		powerrental1.WithBanner(req.Banner, true),

		powerrental1.WithServiceStartAt(req.ServiceStartAt, true),
		powerrental1.WithCancelMode(req.CancelMode, true),
		powerrental1.WithCancelableBeforeStartSeconds(req.CancelableBeforeStartSeconds, true),
		powerrental1.WithEnableSetCommission(req.EnableSetCommission, true),
		powerrental1.WithMinOrderAmount(req.MinOrderAmount, true),
		powerrental1.WithMaxOrderAmount(req.MaxOrderAmount, true),
		powerrental1.WithMaxUserAmount(req.MaxUserAmount, true),
		powerrental1.WithMinOrderDuration(req.MinOrderDuration, true),
		powerrental1.WithMaxOrderDuration(req.MaxOrderDuration, true),
		powerrental1.WithUnitPrice(req.UnitPrice, true),
		powerrental1.WithSaleStartAt(req.SaleStartAt, true),
		powerrental1.WithSaleEndAt(req.SaleEndAt, true),
		powerrental1.WithSaleMode(req.SaleMode, true),
		powerrental1.WithFixedDuration(req.FixedDuration, true),
		powerrental1.WithPackageWithRequireds(req.PackageWithRequireds, true),

		powerrental1.WithAppGoodStockID(req.AppGoodStockID, true),
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
