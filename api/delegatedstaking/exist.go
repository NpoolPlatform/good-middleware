package delegatedstaking

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	delegatedstaking1 "github.com/NpoolPlatform/good-middleware/pkg/mw/delegatedstaking"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/delegatedstaking"
)

func (s *Server) ExistDelegatedStakingConds(ctx context.Context, in *npool.ExistDelegatedStakingCondsRequest) (*npool.ExistDelegatedStakingCondsResponse, error) {
	handler, err := delegatedstaking1.NewHandler(
		ctx,
		delegatedstaking1.WithConds(in.GetConds()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistDelegatedStakingConds",
			"In", in,
			"Error", err,
		)
		return &npool.ExistDelegatedStakingCondsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	exist, err := handler.ExistDelegatedStakingConds(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistDelegatedStakingConds",
			"In", in,
			"Error", err,
		)
		return &npool.ExistDelegatedStakingCondsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.ExistDelegatedStakingCondsResponse{
		Info: exist,
	}, nil
}
