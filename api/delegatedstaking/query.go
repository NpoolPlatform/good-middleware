package delegatedstaking

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	delegatedstaking1 "github.com/NpoolPlatform/good-middleware/pkg/mw/delegatedstaking"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/delegatedstaking"
)

func (s *Server) GetDelegatedStaking(ctx context.Context, in *npool.GetDelegatedStakingRequest) (*npool.GetDelegatedStakingResponse, error) {
	handler, err := delegatedstaking1.NewHandler(
		ctx,
		delegatedstaking1.WithGoodID(&in.GoodID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetDelegatedStaking",
			"In", in,
			"Error", err,
		)
		return &npool.GetDelegatedStakingResponse{}, status.Error(codes.Aborted, err.Error())
	}
	info, err := handler.GetDelegatedStaking(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetDelegatedStaking",
			"In", in,
			"Error", err,
		)
		return &npool.GetDelegatedStakingResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.GetDelegatedStakingResponse{
		Info: info,
	}, nil
}

func (s *Server) GetDelegatedStakings(ctx context.Context, in *npool.GetDelegatedStakingsRequest) (*npool.GetDelegatedStakingsResponse, error) {
	handler, err := delegatedstaking1.NewHandler(
		ctx,
		delegatedstaking1.WithConds(in.GetConds()),
		delegatedstaking1.WithOffset(in.GetOffset()),
		delegatedstaking1.WithLimit(in.GetLimit()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetDelegatedStakings",
			"In", in,
			"Error", err,
		)
		return &npool.GetDelegatedStakingsResponse{}, status.Error(codes.Aborted, err.Error())
	}
	infos, total, err := handler.GetDelegatedStakings(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetDelegatedStakings",
			"In", in,
			"Error", err,
		)
		return &npool.GetDelegatedStakingsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.GetDelegatedStakingsResponse{
		Infos: infos,
		Total: total,
	}, nil
}
