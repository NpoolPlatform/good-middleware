package location

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	location1 "github.com/NpoolPlatform/good-middleware/pkg/mw/vender/location"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/vender/location"
)

func (s *Server) ExistLocationConds(ctx context.Context, in *npool.ExistLocationCondsRequest) (*npool.ExistLocationCondsResponse, error) {
	handler, err := location1.NewHandler(
		ctx,
		location1.WithConds(in.GetConds()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistLocationConds",
			"In", in,
			"Error", err,
		)
		return &npool.ExistLocationCondsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	exist, err := handler.ExistLocationConds(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistLocationConds",
			"In", in,
			"Error", err,
		)
		return &npool.ExistLocationCondsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.ExistLocationCondsResponse{
		Info: exist,
	}, nil
}
