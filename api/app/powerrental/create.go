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
		powerrental1.WithAppGoodID(req.AppGoodID, false),
		powerrental1.WithPurchasable(req.Purchasable, false),
		powerrental1.WithEnableProductPage(req.EnableProductPage, false),
		powerrental1.WithProductPage(req.ProductPage, false),
		powerrental1.WithOnline(req.Online, false),
		powerrental1.WithVisible(req.Visible, false),
		powerrental1.WithName(req.Name, true),
		powerrental1.WithDisplayIndex(req.DisplayIndex, false),
		powerrental1.WithBanner(req.Banner, false),

		powerrental1.WithServiceStartAt(req.ServiceStartAt, true),
		powerrental1.WithCancelMode(req.CancelMode, false),
		powerrental1.WithCancelableBeforeStartSeconds(req.CancelableBeforeStartSeconds, false),
		powerrental1.WithEnableSetCommission(req.EnableSetCommission, false),
		powerrental1.WithMinOrderAmount(req.MinOrderAmount, false),
		powerrental1.WithMaxOrderAmount(req.MaxOrderAmount, false),
		powerrental1.WithMaxUserAmount(req.MaxUserAmount, false),
		powerrental1.WithMinOrderDuration(req.MinOrderDuration, false),
		powerrental1.WithMaxOrderDuration(req.MaxOrderDuration, false),
		powerrental1.WithUnitPrice(req.UnitPrice, true),
		powerrental1.WithSaleStartAt(req.SaleStartAt, true),
		powerrental1.WithSaleEndAt(req.SaleEndAt, true),
		powerrental1.WithSaleMode(req.SaleMode, true),
		powerrental1.WithFixedDuration(req.FixedDuration, false),
		powerrental1.WithPackageWithRequireds(req.PackageWithRequireds, false),

		powerrental1.WithAppGoodStockID(req.AppGoodStockID, false),
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
