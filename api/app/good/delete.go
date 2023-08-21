//nolint:dupl
package appgood

import (
	"context"

	appgood1 "github.com/NpoolPlatform/good-middleware/pkg/mw/app/good"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/app/good"
)

func (s *Server) DeleteGood(ctx context.Context, in *npool.DeleteGoodRequest) (*npool.DeleteGoodResponse, error) {
	req := in.GetInfo()
	if req == nil {
		logger.Sugar().Errorw(
			"DeleteGood",
			"In", in,
		)
		return &npool.DeleteGoodResponse{}, status.Error(codes.Aborted, "invalid argument")
	}
	handler, err := appgood1.NewHandler(
		ctx,
		appgood1.WithID(req.ID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteGood",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteGoodResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.DeleteGood(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteGood",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteGoodResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.DeleteGoodResponse{
		Info: info,
	}, nil
}
