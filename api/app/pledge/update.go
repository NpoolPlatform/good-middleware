package pledge

import (
	"context"

	pledge1 "github.com/NpoolPlatform/good-middleware/pkg/mw/app/pledge"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/app/pledge"
)

func (s *Server) UpdatePledge(ctx context.Context, in *npool.UpdatePledgeRequest) (*npool.UpdatePledgeResponse, error) {
	req := in.GetInfo()
	if req == nil {
		logger.Sugar().Errorw(
			"UpdatePledge",
			"In", in,
		)
		return &npool.UpdatePledgeResponse{}, status.Error(codes.Aborted, "invalid argument")
	}
	handler, err := pledge1.NewHandler(
		ctx,
		pledge1.WithID(req.ID, false),
		pledge1.WithEntID(req.EntID, false),
		pledge1.WithAppGoodID(req.AppGoodID, false),

		pledge1.WithPurchasable(req.Purchasable, false),
		pledge1.WithEnableProductPage(req.EnableProductPage, false),
		pledge1.WithProductPage(req.ProductPage, false),
		pledge1.WithOnline(req.Online, false),
		pledge1.WithVisible(req.Visible, false),
		pledge1.WithName(req.Name, false),
		pledge1.WithDisplayIndex(req.DisplayIndex, false),
		pledge1.WithBanner(req.Banner, false),

		pledge1.WithServiceStartAt(req.ServiceStartAt, false),
		pledge1.WithStartMode(req.StartMode, false),
		pledge1.WithEnableSetCommission(req.EnableSetCommission, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdatePledge",
			"In", in,
			"Error", err,
		)
		return &npool.UpdatePledgeResponse{}, status.Error(codes.Aborted, err.Error())
	}

	if err := handler.UpdatePledge(ctx); err != nil {
		logger.Sugar().Errorw(
			"UpdatePledge",
			"In", in,
			"Error", err,
		)
		return &npool.UpdatePledgeResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.UpdatePledgeResponse{}, nil
}
