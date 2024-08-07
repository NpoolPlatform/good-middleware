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

func (s *Server) UpdateDefault(ctx context.Context, in *npool.UpdateDefaultRequest) (*npool.UpdateDefaultResponse, error) {
	req := in.GetInfo()
	if req == nil {
		logger.Sugar().Errorw(
			"UpdateDefault",
			"In", in,
		)
		return &npool.UpdateDefaultResponse{}, status.Error(codes.Aborted, "invalid argument")
	}
	handler, err := appdefaultgood1.NewHandler(
		ctx,
		appdefaultgood1.WithID(req.ID, false),
		appdefaultgood1.WithEntID(req.EntID, false),
		appdefaultgood1.WithAppGoodID(req.AppGoodID, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateDefault",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateDefaultResponse{}, status.Error(codes.Aborted, err.Error())
	}

	if err := handler.UpdateDefault(ctx); err != nil {
		logger.Sugar().Errorw(
			"UpdateDefault",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateDefaultResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.UpdateDefaultResponse{}, nil
}
