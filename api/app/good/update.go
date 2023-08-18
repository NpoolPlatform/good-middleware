//nolint:dupl
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
		appgood1.WithAppID(req.AppID, true),
		appgood1.WithGoodID(req.GoodID, true),
		appgood1.WithOnline(req.Online, false),
		appgood1.WithVisible(req.Visible, false),
		appgood1.WithGoodName(req.GoodName, false),
		appgood1.WithPrice(req.Price, false),
		appgood1.WithDisplayIndex(req.DisplayIndex, false),
		appgood1.WithPurchaseLimit(req.PurchaseLimit, false),
		appgood1.WithSaleStartAt(req.SaleStartAt, false),
		appgood1.WithSaleEndAt(req.SaleEndAt, false),
		appgood1.WithServiceStartAt(req.ServiceStartAt, false),
		appgood1.WithDescriptions(req.Descriptions, false),
		appgood1.WithGoodBanner(req.GoodBanner, false),
		appgood1.WithEnablePurchase(req.EnablePurchase, false),
		appgood1.WithEnableProductPage(req.EnableProductPage, false),
		appgood1.WithCancelMode(req.CancelMode, false),
		appgood1.WithUserPurchaseLimit(req.UserPurchaseLimit, false),
		appgood1.WithDisplayColors(req.DisplayColors, false),
		appgood1.WithCancellableBeforeStart(req.CancellableBeforeStart, false),
		appgood1.WithProductPage(req.ProductPage, false),
		appgood1.WithEnableSetCommission(req.EnableSetCommission, false),
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
