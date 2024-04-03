package color

import (
	"context"

	color1 "github.com/NpoolPlatform/good-middleware/pkg/mw/app/good/display/color"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/app/good/display/color"
)

func (s *Server) DeleteDisplayColor(ctx context.Context, in *npool.DeleteDisplayColorRequest) (*npool.DeleteDisplayColorResponse, error) {
	req := in.GetInfo()
	if req == nil {
		logger.Sugar().Errorw(
			"DeleteDisplayColor",
			"In", in,
		)
		return &npool.DeleteDisplayColorResponse{}, status.Error(codes.Aborted, "invalid argument")
	}
	handler, err := color1.NewHandler(
		ctx,
		color1.WithID(req.ID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteDisplayColor",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteDisplayColorResponse{}, status.Error(codes.Aborted, err.Error())
	}

	if err := handler.DeleteDisplayColor(ctx); err != nil {
		logger.Sugar().Errorw(
			"DeleteDisplayColor",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteDisplayColorResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.DeleteDisplayColorResponse{}, nil
}
