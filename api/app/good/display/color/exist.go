package color

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	displaycolor1 "github.com/NpoolPlatform/good-middleware/pkg/mw/app/good/display/color"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/app/good/display/color"
)

func (s *Server) ExistDisplayColorConds(ctx context.Context, in *npool.ExistDisplayColorCondsRequest) (*npool.ExistDisplayColorCondsResponse, error) {
	handler, err := displaycolor1.NewHandler(
		ctx,
		displaycolor1.WithConds(in.GetConds()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistDisplayColorConds",
			"In", in,
			"Error", err,
		)
		return &npool.ExistDisplayColorCondsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	exist, err := handler.ExistDisplayColorConds(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistDisplayColorConds",
			"In", in,
			"Error", err,
		)
		return &npool.ExistDisplayColorCondsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.ExistDisplayColorCondsResponse{
		Info: exist,
	}, nil
}
