package appdefaultgood

import (
	"context"

	appdefaultgood1 "github.com/NpoolPlatform/good-middleware/pkg/mw/app/good/default"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/app/good/default"
)

func (s *Server) DeleteDefault(ctx context.Context, in *npool.DeleteDefaultRequest) (*npool.DeleteDefaultResponse, error) {
	req := in.GetInfo()
	if req == nil {
		logger.Sugar().Errorw(
			"DeleteDefault",
			"In", in,
		)
		return &npool.DeleteDefaultResponse{}, status.Error(codes.Aborted, "invalid argument")
	}
	handler, err := appdefaultgood1.NewHandler(
		ctx,
		appdefaultgood1.WithID(req.ID, false),
		appdefaultgood1.WithEntID(req.EntID, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteDefault",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteDefaultResponse{}, status.Error(codes.Aborted, err.Error())
	}

	if err := handler.DeleteDefault(ctx); err != nil {
		logger.Sugar().Errorw(
			"DeleteDefault",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteDefaultResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.DeleteDefaultResponse{}, nil
}
