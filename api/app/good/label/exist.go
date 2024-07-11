package label

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	label1 "github.com/NpoolPlatform/good-middleware/pkg/mw/app/good/label"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/app/good/label"
)

func (s *Server) ExistLabelConds(ctx context.Context, in *npool.ExistLabelCondsRequest) (*npool.ExistLabelCondsResponse, error) {
	handler, err := label1.NewHandler(
		ctx,
		label1.WithConds(in.GetConds()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistLabelConds",
			"In", in,
			"Error", err,
		)
		return &npool.ExistLabelCondsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	exist, err := handler.ExistLabelConds(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistLabelConds",
			"In", in,
			"Error", err,
		)
		return &npool.ExistLabelCondsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.ExistLabelCondsResponse{
		Info: exist,
	}, nil
}
