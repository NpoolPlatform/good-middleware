package appdefaultgood

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	appdefaultgood1 "github.com/NpoolPlatform/good-middleware/pkg/mw/app/good/default"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/app/good/default"
)

func (s *Server) ExistDefaultConds(ctx context.Context, in *npool.ExistDefaultCondsRequest) (*npool.ExistDefaultCondsResponse, error) {
	handler, err := appdefaultgood1.NewHandler(
		ctx,
		appdefaultgood1.WithConds(in.GetConds()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistDefaultConds",
			"In", in,
			"Error", err,
		)
		return &npool.ExistDefaultCondsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	exist, err := handler.ExistDefaultConds(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistDefaultConds",
			"In", in,
			"Error", err,
		)
		return &npool.ExistDefaultCondsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.ExistDefaultCondsResponse{
		Info: exist,
	}, nil
}
