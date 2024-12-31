package pledge

import (
	"context"

	pledge1 "github.com/NpoolPlatform/good-middleware/pkg/mw/app/pledge"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/app/pledge"
)

func (s *Server) CreatePledge(ctx context.Context, in *npool.CreatePledgeRequest) (*npool.CreatePledgeResponse, error) {
	req := in.GetInfo()
	if req == nil {
		logger.Sugar().Errorw(
			"CreatePledge",
			"In", in,
		)
		return &npool.CreatePledgeResponse{}, status.Error(codes.Aborted, "invalid argument")
	}
	handler, err := pledge1.NewHandler(
		ctx,
		pledge1.WithEntID(req.EntID, false),
		pledge1.WithAppID(req.AppID, true),
		pledge1.WithGoodID(req.GoodID, true),
		pledge1.WithAppGoodID(req.AppGoodID, false),
		pledge1.WithPurchasable(req.Purchasable, false),
		pledge1.WithEnableProductPage(req.EnableProductPage, false),
		pledge1.WithProductPage(req.ProductPage, false),
		pledge1.WithOnline(req.Online, false),
		pledge1.WithVisible(req.Visible, false),
		pledge1.WithName(req.Name, true),
		pledge1.WithDisplayIndex(req.DisplayIndex, false),
		pledge1.WithBanner(req.Banner, false),
		pledge1.WithServiceStartAt(req.ServiceStartAt, true),
		pledge1.WithStartMode(req.StartMode, false),
		pledge1.WithEnableSetCommission(req.EnableSetCommission, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CreatePledge",
			"In", in,
			"Error", err,
		)
		return &npool.CreatePledgeResponse{}, status.Error(codes.Aborted, err.Error())
	}

	if err = handler.CreatePledge(ctx); err != nil {
		logger.Sugar().Errorw(
			"CreatePledge",
			"In", in,
			"Error", err,
		)
		return &npool.CreatePledgeResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.CreatePledgeResponse{}, nil
}
