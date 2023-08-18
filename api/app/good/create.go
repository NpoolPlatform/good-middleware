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
		appgood1.WithID(req.ID, true),
		appgood1.WithAppID(req.AppID, true),
		appgood1.WithGoodID(req.GoodID, true),
		appgood1.WithOnline(req.Online, true),
		appgood1.WithVisible(req.Visible, true),
		appgood1.WithGoodName(req.GoodName, true),
		appgood1.WithPrice(req.Price, true),
		appgood1.WithDisplayIndex(req.DisplayIndex, true),
		appgood1.WithPurchaseLimit(req.PurchaseLimit, true),
		appgood1.WithSaleStartAt(req.SaleStartAt, true),
		appgood1.WithSaleEndAt(req.SaleEndAt, true),
		appgood1.WithServiceStartAt(req.ServiceStartAt, true),
		appgood1.WithDescriptions(req.Descriptions, true),
		appgood1.WithGoodBanner(req.GoodBanner, true),
		appgood1.WithEnablePurchase(req.EnablePurchase, true),
		appgood1.WithEnableProductPage(req.EnableProductPage, true),
		appgood1.WithCancelMode(req.CancelMode, true),
		appgood1.WithUserPurchaseLimit(req.UserPurchaseLimit, true),
		appgood1.WithDisplayColors(req.DisplayColors, true),
		appgood1.WithCancellableBeforeStart(req.CancellableBeforeStart, true),
		appgood1.WithProductPage(req.ProductPage, true),
		appgood1.WithEnableSetCommission(req.EnableSetCommission, true),
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
