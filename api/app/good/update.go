package appgood

import (
	"context"

	appgood1 "github.com/NpoolPlatform/good-middleware/pkg/mw/app/good"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/app/good"
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
	handler, err := appgood1.NewHandler(
		ctx,
		appgood1.WithID(req.ID, true),
		appgood1.WithOnline(req.Online, false),
		appgood1.WithVisible(req.Visible, false),
		appgood1.WithGoodName(req.GoodName, false),
		appgood1.WithUnitPrice(req.UnitPrice, false),
		appgood1.WithPackagePrice(req.PackagePrice, false),
		appgood1.WithDisplayIndex(req.DisplayIndex, false),
		appgood1.WithSaleStartAt(req.SaleStartAt, false),
		appgood1.WithSaleEndAt(req.SaleEndAt, false),
		appgood1.WithServiceStartAt(req.ServiceStartAt, false),
		appgood1.WithDescriptions(req.Descriptions, false),
		appgood1.WithPosters(req.Posters, false),
		appgood1.WithGoodBanner(req.GoodBanner, false),
		appgood1.WithEnablePurchase(req.EnablePurchase, false),
		appgood1.WithEnableProductPage(req.EnableProductPage, false),
		appgood1.WithCancelMode(req.CancelMode, false),
		appgood1.WithDisplayColors(req.DisplayColors, false),
		appgood1.WithDisplayNames(req.DisplayNames, false),
		appgood1.WithCancellableBeforeStart(req.CancellableBeforeStart, false),
		appgood1.WithProductPage(req.ProductPage, false),
		appgood1.WithEnableSetCommission(req.EnableSetCommission, false),
		appgood1.WithTechniqueFeeRatio(req.TechnicalFeeRatio, false),
		appgood1.WithElectricityFeeRatio(req.ElectricityFeeRatio, false),
		appgood1.WithMinOrderAmount(req.MinOrderAmount, false),
		appgood1.WithMaxOrderAmount(req.MaxOrderAmount, false),
		appgood1.WithMaxUserAmount(req.MaxUserAmount, false),
		appgood1.WithMinOrderDuration(req.MinOrderDuration, false),
		appgood1.WithMaxOrderDuration(req.MaxOrderDuration, false),
		appgood1.WithPackageWithRequireds(req.PackageWithRequireds, false),
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
