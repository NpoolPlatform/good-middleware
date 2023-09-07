//nolint:dupl
package appgood

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	appgood1 "github.com/NpoolPlatform/good-middleware/pkg/mw/app/good"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/app/good"
)

func (s *Server) ExistGood(ctx context.Context, in *npool.ExistGoodRequest) (*npool.ExistGoodResponse, error) {
	handler, err := appgood1.NewHandler(
		ctx,
		appgood1.WithID(&in.ID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistGood",
			"In", in,
			"Error", err,
		)
		return &npool.ExistGoodResponse{}, status.Error(codes.Aborted, err.Error())
	}

	exist, err := handler.ExistGood(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistGood",
			"In", in,
			"Error", err,
		)
		return &npool.ExistGoodResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.ExistGoodResponse{
		Info: exist,
	}, nil
}

func (s *Server) ExistGoodConds(ctx context.Context, in *npool.ExistGoodCondsRequest) (*npool.ExistGoodCondsResponse, error) {
	handler, err := appgood1.NewHandler(
		ctx,
		appgood1.WithConds(in.GetConds()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistGoodConds",
			"In", in,
			"Error", err,
		)
		return &npool.ExistGoodCondsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	exist, err := handler.ExistGoodConds(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistGoodConds",
			"In", in,
			"Error", err,
		)
		return &npool.ExistGoodCondsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.ExistGoodCondsResponse{
		Info: exist,
	}, nil
}
