//nolint:dupl
package malfunction

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	malfunction1 "github.com/NpoolPlatform/good-middleware/pkg/mw/good/malfunction"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/good/malfunction"
)

func (s *Server) ExistMalfunction(ctx context.Context, in *npool.ExistMalfunctionRequest) (*npool.ExistMalfunctionResponse, error) {
	handler, err := malfunction1.NewHandler(
		ctx,
		malfunction1.WithEntID(&in.EntID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistMalfunction",
			"In", in,
			"Error", err,
		)
		return &npool.ExistMalfunctionResponse{}, status.Error(codes.Aborted, err.Error())
	}

	exist, err := handler.ExistMalfunction(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistMalfunction",
			"In", in,
			"Error", err,
		)
		return &npool.ExistMalfunctionResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.ExistMalfunctionResponse{
		Info: exist,
	}, nil
}

func (s *Server) ExistMalfunctionConds(ctx context.Context, in *npool.ExistMalfunctionCondsRequest) (*npool.ExistMalfunctionCondsResponse, error) {
	handler, err := malfunction1.NewHandler(
		ctx,
		malfunction1.WithConds(in.GetConds()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistMalfunctionConds",
			"In", in,
			"Error", err,
		)
		return &npool.ExistMalfunctionCondsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	exist, err := handler.ExistMalfunctionConds(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistMalfunctionConds",
			"In", in,
			"Error", err,
		)
		return &npool.ExistMalfunctionCondsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.ExistMalfunctionCondsResponse{
		Info: exist,
	}, nil
}
