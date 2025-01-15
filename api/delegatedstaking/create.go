package delegatedstaking

import (
	"context"

	delegatedstaking1 "github.com/NpoolPlatform/good-middleware/pkg/mw/delegatedstaking"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/delegatedstaking"
)

func (s *Server) CreateDelegatedStaking(ctx context.Context, in *npool.CreateDelegatedStakingRequest) (*npool.CreateDelegatedStakingResponse, error) {
	req := in.GetInfo()
	if req == nil {
		logger.Sugar().Errorw(
			"CreateDelegatedStaking",
			"In", in,
		)
		return &npool.CreateDelegatedStakingResponse{}, status.Error(codes.Aborted, "invalid argument")
	}
	handler, err := delegatedstaking1.NewHandler(
		ctx,
		delegatedstaking1.WithEntID(req.EntID, false),
		delegatedstaking1.WithGoodID(req.GoodID, false),
		delegatedstaking1.WithGoodType(req.GoodType, true),
		delegatedstaking1.WithName(req.Name, true),
		delegatedstaking1.WithServiceStartAt(req.ServiceStartAt, true),
		delegatedstaking1.WithStartMode(req.StartMode, true),
		delegatedstaking1.WithTestOnly(req.TestOnly, false),
		delegatedstaking1.WithBenefitIntervalHours(req.BenefitIntervalHours, true),
		delegatedstaking1.WithPurchasable(req.Purchasable, false),
		delegatedstaking1.WithOnline(req.Online, false),
		delegatedstaking1.WithContractCodeURL(req.ContractCodeURL, true),
		delegatedstaking1.WithContractCodeBranch(req.ContractCodeBranch, true),
		delegatedstaking1.WithContractState(req.ContractState, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateDelegatedStaking",
			"In", in,
			"Error", err,
		)
		return &npool.CreateDelegatedStakingResponse{}, status.Error(codes.Aborted, err.Error())
	}

	if err = handler.CreateDelegatedStaking(ctx); err != nil {
		logger.Sugar().Errorw(
			"CreateDelegatedStaking",
			"In", in,
			"Error", err,
		)
		return &npool.CreateDelegatedStakingResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.CreateDelegatedStakingResponse{}, nil
}
