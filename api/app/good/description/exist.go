package description

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	description1 "github.com/NpoolPlatform/good-middleware/pkg/mw/app/good/description"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/app/good/description"
)

func (s *Server) ExistDescriptionConds(ctx context.Context, in *npool.ExistDescriptionCondsRequest) (*npool.ExistDescriptionCondsResponse, error) {
	handler, err := description1.NewHandler(
		ctx,
		description1.WithConds(in.GetConds()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistDescriptionConds",
			"In", in,
			"Error", err,
		)
		return &npool.ExistDescriptionCondsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	exist, err := handler.ExistDescriptionConds(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistDescriptionConds",
			"In", in,
			"Error", err,
		)
		return &npool.ExistDescriptionCondsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.ExistDescriptionCondsResponse{
		Info: exist,
	}, nil
}
