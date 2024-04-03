package color

import (
	"context"

	color1 "github.com/NpoolPlatform/good-middleware/pkg/mw/app/good/display/color"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/app/good/display/color"
)

func (s *Server) UpdateDisplayColor(ctx context.Context, in *npool.UpdateDisplayColorRequest) (*npool.UpdateDisplayColorResponse, error) {
	req := in.GetInfo()
	if req == nil {
		logger.Sugar().Errorw(
			"UpdateDisplayColor",
			"In", in,
		)
		return &npool.UpdateDisplayColorResponse{}, status.Error(codes.Aborted, "invalid argument")
	}
	handler, err := color1.NewHandler(
		ctx,
		color1.WithID(req.ID, false),
		color1.WithEntID(req.EntID, false),
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
			"UpdateDisplayColor",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateDisplayColorResponse{}, status.Error(codes.Aborted, err.Error())
	}

	if err := handler.UpdateDisplayColor(ctx); err != nil {
		logger.Sugar().Errorw(
			"UpdateDisplayColor",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateDisplayColorResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.UpdateDisplayColorResponse{}, nil
}
