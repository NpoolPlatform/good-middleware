package color

import (
	"context"

	color1 "github.com/NpoolPlatform/good-middleware/pkg/mw/app/good/display/color"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/app/good/display/color"
)

func (s *Server) CreateDisplayColor(ctx context.Context, in *npool.CreateDisplayColorRequest) (*npool.CreateDisplayColorResponse, error) {
	req := in.GetInfo()
	if req == nil {
		logger.Sugar().Errorw(
			"CreateDisplayColor",
			"In", in,
		)
		return &npool.CreateDisplayColorResponse{}, status.Error(codes.Aborted, "invalid argument")
	}
	handler, err := color1.NewHandler(
		ctx,
		color1.WithEntID(req.EntID, false),
		color1.WithAppGoodID(req.AppGoodID, true),
		color1.WithColor(req.Color, false),
		color1.WithIndex(func() *uint8 {
			if req.Index == nil {
				return nil
			}
			u := uint8(*req.Index)
			return &u
		}(), false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateDisplayColor",
			"In", in,
			"Error", err,
		)
		return &npool.CreateDisplayColorResponse{}, status.Error(codes.Aborted, err.Error())
	}

	if err := handler.CreateDisplayColor(ctx); err != nil {
		logger.Sugar().Errorw(
			"CreateDisplayColor",
			"In", in,
			"Error", err,
		)
		return &npool.CreateDisplayColorResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.CreateDisplayColorResponse{}, nil
}
