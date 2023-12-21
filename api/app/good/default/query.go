package appdefaultgood

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	appdefaultgood1 "github.com/NpoolPlatform/good-middleware/pkg/mw/app/good/default"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/app/good/default"
)

func (s *Server) GetDefault(ctx context.Context, in *npool.GetDefaultRequest) (*npool.GetDefaultResponse, error) {
	handler, err := appdefaultgood1.NewHandler(
		ctx,
		appdefaultgood1.WithEntID(&in.EntID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetDefault",
			"In", in,
			"Error", err,
		)
		return &npool.GetDefaultResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.GetDefault(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetDefault",
			"In", in,
			"Error", err,
		)
		return &npool.GetDefaultResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.GetDefaultResponse{
		Info: info,
	}, nil
}

func (s *Server) GetDefaults(ctx context.Context, in *npool.GetDefaultsRequest) (*npool.GetDefaultsResponse, error) {
	handler, err := appdefaultgood1.NewHandler(
		ctx,
		appdefaultgood1.WithConds(in.GetConds()),
		appdefaultgood1.WithOffset(in.GetOffset()),
		appdefaultgood1.WithLimit(in.GetLimit()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetDefaults",
			"In", in,
			"Error", err,
		)
		return &npool.GetDefaultsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	infos, total, err := handler.GetDefaults(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetDefaults",
			"In", in,
			"Error", err,
		)
		return &npool.GetDefaultsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.GetDefaultsResponse{
		Infos: infos,
		Total: total,
	}, nil
}
