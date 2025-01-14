package delegatedstaking

import (
	"context"

	delegatedstaking1 "github.com/NpoolPlatform/good-middleware/pkg/mw/delegatedstaking"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/delegatedstaking"
)

func (s *Server) DeleteDelegatedStaking(ctx context.Context, in *npool.DeleteDelegatedStakingRequest) (*npool.DeleteDelegatedStakingResponse, error) {
	req := in.GetInfo()
	if req == nil {
		logger.Sugar().Errorw(
			"DeleteDelegatedStaking",
			"In", in,
		)
		return &npool.DeleteDelegatedStakingResponse{}, status.Error(codes.Aborted, "invalid argument")
	}
	handler, err := delegatedstaking1.NewHandler(
		ctx,
		delegatedstaking1.WithID(req.ID, false),
		delegatedstaking1.WithEntID(req.EntID, false),
		delegatedstaking1.WithGoodID(req.GoodID, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteDelegatedStaking",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteDelegatedStakingResponse{}, status.Error(codes.Aborted, err.Error())
	}

	if err := handler.DeleteDelegatedStaking(ctx); err != nil {
		logger.Sugar().Errorw(
			"DeleteDelegatedStaking",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteDelegatedStakingResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.DeleteDelegatedStakingResponse{}, nil
}
