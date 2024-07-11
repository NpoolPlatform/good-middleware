package color

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	color1 "github.com/NpoolPlatform/good-middleware/pkg/mw/app/good/display/color"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/app/good/display/color"
)

func (s *Server) GetDisplayColor(ctx context.Context, in *npool.GetDisplayColorRequest) (*npool.GetDisplayColorResponse, error) {
	handler, err := color1.NewHandler(
		ctx,
		color1.WithEntID(&in.EntID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetDisplayColor",
			"In", in,
			"Error", err,
		)
		return &npool.GetDisplayColorResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.GetDisplayColor(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetDisplayColor",
			"In", in,
			"Error", err,
		)
		return &npool.GetDisplayColorResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.GetDisplayColorResponse{
		Info: info,
	}, nil
}

func (s *Server) GetDisplayColors(ctx context.Context, in *npool.GetDisplayColorsRequest) (*npool.GetDisplayColorsResponse, error) {
	handler, err := color1.NewHandler(
		ctx,
		color1.WithConds(in.GetConds()),
		color1.WithOffset(in.GetOffset()),
		color1.WithLimit(in.GetLimit()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetDisplayColors",
			"In", in,
			"Error", err,
		)
		return &npool.GetDisplayColorsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	infos, total, err := handler.GetDisplayColors(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetDisplayColors",
			"In", in,
			"Error", err,
		)
		return &npool.GetDisplayColorsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.GetDisplayColorsResponse{
		Infos: infos,
		Total: total,
	}, nil
}
