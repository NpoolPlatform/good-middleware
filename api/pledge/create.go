package pledge

import (
	"context"

	pledge1 "github.com/NpoolPlatform/good-middleware/pkg/mw/pledge"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/pledge"
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
		pledge1.WithGoodID(req.GoodID, false),
		pledge1.WithGoodType(req.GoodType, true),
		pledge1.WithName(req.Name, true),
		pledge1.WithServiceStartAt(req.ServiceStartAt, true),
		pledge1.WithStartMode(req.StartMode, true),
		pledge1.WithTestOnly(req.TestOnly, false),
		pledge1.WithBenefitIntervalHours(req.BenefitIntervalHours, true),
		pledge1.WithPurchasable(req.Purchasable, false),
		pledge1.WithOnline(req.Online, false),
		pledge1.WithContractCodeURL(req.ContractCodeURL, true),
		pledge1.WithContractCodeBranch(req.ContractCodeBranch, true),
		pledge1.WithContractState(req.ContractState, false),
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
