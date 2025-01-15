package delegatedstaking

import (
	"context"

	delegatedstaking1 "github.com/NpoolPlatform/good-middleware/pkg/mw/delegatedstaking"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/delegatedstaking"
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
		delegatedstaking1.WithGoodID(req.GoodID, false),

		delegatedstaking1.WithGoodType(req.GoodType, false),
		delegatedstaking1.WithName(req.Name, false),
		delegatedstaking1.WithServiceStartAt(req.ServiceStartAt, false),
		delegatedstaking1.WithStartMode(req.StartMode, false),
		delegatedstaking1.WithTestOnly(req.TestOnly, false),
		delegatedstaking1.WithBenefitIntervalHours(req.BenefitIntervalHours, false),
		delegatedstaking1.WithPurchasable(req.Purchasable, false),
		delegatedstaking1.WithOnline(req.Online, false),

		delegatedstaking1.WithRewardState(req.RewardState, false),
		delegatedstaking1.WithRewardAt(req.RewardAt, false),
		delegatedstaking1.WithRewards(req.Rewards, false),
		delegatedstaking1.WithContractCodeURL(req.ContractCodeURL, false),
		delegatedstaking1.WithContractCodeBranch(req.ContractCodeBranch, false),
		delegatedstaking1.WithContractState(req.ContractState, false),
		delegatedstaking1.WithRollback(req.Rollback, false),
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
