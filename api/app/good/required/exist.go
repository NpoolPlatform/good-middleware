package required

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	required1 "github.com/NpoolPlatform/good-middleware/pkg/mw/app/good/required"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/app/good/required"
)

func (s *Server) ExistRequiredConds(ctx context.Context, in *npool.ExistRequiredCondsRequest) (*npool.ExistRequiredCondsResponse, error) {
	handler, err := required1.NewHandler(
		ctx,
		required1.WithConds(in.GetConds()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistRequiredConds",
			"In", in,
			"Error", err,
		)
		return &npool.ExistRequiredCondsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.ExistRequiredConds(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistRequiredConds",
			"In", in,
			"Error", err,
		)
		return &npool.ExistRequiredCondsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.ExistRequiredCondsResponse{
		Info: info,
	}, nil
}
