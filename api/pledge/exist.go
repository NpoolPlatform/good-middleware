package pledge

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pledge1 "github.com/NpoolPlatform/good-middleware/pkg/mw/pledge"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/pledge"
)

func (s *Server) ExistPledgeConds(ctx context.Context, in *npool.ExistPledgeCondsRequest) (*npool.ExistPledgeCondsResponse, error) {
	handler, err := pledge1.NewHandler(
		ctx,
		pledge1.WithConds(in.GetConds()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistPledgeConds",
			"In", in,
			"Error", err,
		)
		return &npool.ExistPledgeCondsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	exist, err := handler.ExistPledgeConds(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistPledgeConds",
			"In", in,
			"Error", err,
		)
		return &npool.ExistPledgeCondsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.ExistPledgeCondsResponse{
		Info: exist,
	}, nil
}
