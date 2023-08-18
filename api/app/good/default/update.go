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
		appdefaultgood1.WithID(req.ID, true),
		appdefaultgood1.WithAppID(req.AppID, true),
		appdefaultgood1.WithGoodID(req.GoodID, true),
		appdefaultgood1.WithAppGoodID(req.AppGoodID, true),
		appdefaultgood1.WithCoinTypeID(req.CoinTypeID, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateDefault",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateDefaultResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.UpdateDefault(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateDefault",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateDefaultResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.UpdateDefaultResponse{
		Info: info,
	}, nil
}
