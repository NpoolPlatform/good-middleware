package delegatedstaking

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	appdelegatedstaking1 "github.com/NpoolPlatform/good-middleware/pkg/mw/app/delegatedstaking"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/app/delegatedstaking"
)

func (s *Server) ExistDelegatedStakingConds(ctx context.Context, in *npool.ExistDelegatedStakingCondsRequest) (*npool.ExistDelegatedStakingCondsResponse, error) {
	handler, err := appdelegatedstaking1.NewHandler(
		ctx,
		appdelegatedstaking1.WithConds(in.GetConds()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistDelegatedStakingConds",
			"In", in,
			"Error", err,
		)
		return &npool.ExistDelegatedStakingCondsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.ExistDelegatedStakingConds(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistDelegatedStakingConds",
			"In", in,
			"Error", err,
		)
		return &npool.ExistDelegatedStakingCondsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.ExistDelegatedStakingCondsResponse{
		Info: info,
	}, nil
}
