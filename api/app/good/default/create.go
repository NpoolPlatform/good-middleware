//nolint:dupl
package appdefaultgood

import (
	"context"

	appdefaultgood1 "github.com/NpoolPlatform/good-middleware/pkg/mw/app/good/default"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/app/good/default"
)

func (s *Server) CreateDefault(ctx context.Context, in *npool.CreateDefaultRequest) (*npool.CreateDefaultResponse, error) {
	req := in.GetInfo()
	if req == nil {
		logger.Sugar().Errorw(
			"CreateDefault",
			"In", in,
		)
		return &npool.CreateDefaultResponse{}, status.Error(codes.Aborted, "invalid argument")
	}
	handler, err := appdefaultgood1.NewHandler(
		ctx,
		appdefaultgood1.WithID(req.ID, false),
		appdefaultgood1.WithAppGoodID(req.AppGoodID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateDefault",
			"In", in,
			"Error", err,
		)
		return &npool.CreateDefaultResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.CreateDefault(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateDefault",
			"In", in,
			"Error", err,
		)
		return &npool.CreateDefaultResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.CreateDefaultResponse{
		Info: info,
	}, nil
}
