package appgood

import (
	"context"

	appgood1 "github.com/NpoolPlatform/good-middleware/pkg/mw/app/good"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/app/good"
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
	handler, err := appgood1.NewHandler(
		ctx,
		appgood1.WithEntID(req.EntID, false),
		appgood1.WithAppID(req.AppID, true),
		appgood1.WithGoodID(req.GoodID, true),
		appgood1.WithOnline(req.Online, true),
		appgood1.WithVisible(req.Visible, true),
		appgood1.WithGoodName(req.GoodName, true),
		appgood1.WithUnitPrice(req.UnitPrice, true),
		appgood1.WithPackagePrice(req.PackagePrice, false),
		appgood1.WithDisplayIndex(req.DisplayIndex, true),
		appgood1.WithSaleStartAt(req.SaleStartAt, true),
		appgood1.WithSaleEndAt(req.SaleEndAt, true),
		appgood1.WithServiceStartAt(req.ServiceStartAt, true),
		appgood1.WithDescriptions(req.Descriptions, true),
		appgood1.WithGoodBanner(req.GoodBanner, true),
		appgood1.WithEnablePurchase(req.EnablePurchase, true),
		appgood1.WithEnableProductPage(req.EnableProductPage, true),
		appgood1.WithCancelMode(req.CancelMode, false),
		appgood1.WithDisplayColors(req.DisplayColors, false),
		appgood1.WithCancellableBeforeStart(req.CancellableBeforeStart, false),
		appgood1.WithProductPage(req.ProductPage, true),
		appgood1.WithEnableSetCommission(req.EnableSetCommission, false),
		appgood1.WithPosters(req.Posters, true),
		appgood1.WithTechniqueFeeRatio(req.TechnicalFeeRatio, false),
		appgood1.WithElectricityFeeRatio(req.ElectricityFeeRatio, false),
		appgood1.WithDisplayNames(req.DisplayNames, false),
		appgood1.WithMinOrderAmount(req.MinOrderAmount, false),
		appgood1.WithMaxOrderAmount(req.MaxOrderAmount, false),
		appgood1.WithMaxUserAmount(req.MaxUserAmount, false),
		appgood1.WithMinOrderDuration(req.MinOrderDuration, false),
		appgood1.WithMaxOrderDuration(req.MaxOrderDuration, false),
		appgood1.WithPackageWithRequireds(req.PackageWithRequireds, false),
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
