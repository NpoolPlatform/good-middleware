package name

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	displayname1 "github.com/NpoolPlatform/good-middleware/pkg/mw/app/good/display/name"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/app/good/display/name"
)

func (s *Server) ExistDisplayNameConds(ctx context.Context, in *npool.ExistDisplayNameCondsRequest) (*npool.ExistDisplayNameCondsResponse, error) {
	handler, err := displayname1.NewHandler(
		ctx,
		displayname1.WithConds(in.GetConds()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistDisplayNameConds",
			"In", in,
			"Error", err,
		)
		return &npool.ExistDisplayNameCondsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	exist, err := handler.ExistDisplayNameConds(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistDisplayNameConds",
			"In", in,
			"Error", err,
		)
		return &npool.ExistDisplayNameCondsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.ExistDisplayNameCondsResponse{
		Info: exist,
	}, nil
}
