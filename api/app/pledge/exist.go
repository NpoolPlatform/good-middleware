package pledge

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	apppledge1 "github.com/NpoolPlatform/good-middleware/pkg/mw/app/pledge"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/app/pledge"
)

func (s *Server) ExistPledgeConds(ctx context.Context, in *npool.ExistPledgeCondsRequest) (*npool.ExistPledgeCondsResponse, error) {
	handler, err := apppledge1.NewHandler(
		ctx,
		apppledge1.WithConds(in.GetConds()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistPledgeConds",
			"In", in,
			"Error", err,
		)
		return &npool.ExistPledgeCondsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.ExistPledgeConds(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistPledgeConds",
			"In", in,
			"Error", err,
		)
		return &npool.ExistPledgeCondsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.ExistPledgeCondsResponse{
		Info: info,
	}, nil
}
