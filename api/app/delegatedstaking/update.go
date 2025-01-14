package delegatedstaking

import (
	"context"

	delegatedstaking1 "github.com/NpoolPlatform/good-middleware/pkg/mw/app/delegatedstaking"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/app/delegatedstaking"
)

func (s *Server) UpdateDelegatedStaking(ctx context.Context, in *npool.UpdateDelegatedStakingRequest) (*npool.UpdateDelegatedStakingResponse, error) {
	req := in.GetInfo()
	if req == nil {
		logger.Sugar().Errorw(
			"UpdateDelegatedStaking",
			"In", in,
		)
		return &npool.UpdateDelegatedStakingResponse{}, status.Error(codes.Aborted, "invalid argument")
	}
	handler, err := delegatedstaking1.NewHandler(
		ctx,
		delegatedstaking1.WithID(req.ID, false),
		delegatedstaking1.WithEntID(req.EntID, false),
		delegatedstaking1.WithAppGoodID(req.AppGoodID, false),

		delegatedstaking1.WithPurchasable(req.Purchasable, false),
		delegatedstaking1.WithEnableProductPage(req.EnableProductPage, false),
		delegatedstaking1.WithProductPage(req.ProductPage, false),
		delegatedstaking1.WithOnline(req.Online, false),
		delegatedstaking1.WithVisible(req.Visible, false),
		delegatedstaking1.WithName(req.Name, false),
		delegatedstaking1.WithDisplayIndex(req.DisplayIndex, false),
		delegatedstaking1.WithBanner(req.Banner, false),

		delegatedstaking1.WithServiceStartAt(req.ServiceStartAt, false),
		delegatedstaking1.WithStartMode(req.StartMode, false),
		delegatedstaking1.WithEnableSetCommission(req.EnableSetCommission, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateDelegatedStaking",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateDelegatedStakingResponse{}, status.Error(codes.Aborted, err.Error())
	}

	if err := handler.UpdateDelegatedStaking(ctx); err != nil {
		logger.Sugar().Errorw(
			"UpdateDelegatedStaking",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateDelegatedStakingResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.UpdateDelegatedStakingResponse{}, nil
}
