package name

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	name1 "github.com/NpoolPlatform/good-middleware/pkg/mw/app/good/display/name"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/app/good/display/name"
)

func (s *Server) GetDisplayName(ctx context.Context, in *npool.GetDisplayNameRequest) (*npool.GetDisplayNameResponse, error) {
	handler, err := name1.NewHandler(
		ctx,
		name1.WithEntID(&in.EntID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetDisplayName",
			"In", in,
			"Error", err,
		)
		return &npool.GetDisplayNameResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.GetDisplayName(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetDisplayName",
			"In", in,
			"Error", err,
		)
		return &npool.GetDisplayNameResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.GetDisplayNameResponse{
		Info: info,
	}, nil
}

func (s *Server) GetDisplayNames(ctx context.Context, in *npool.GetDisplayNamesRequest) (*npool.GetDisplayNamesResponse, error) {
	handler, err := name1.NewHandler(
		ctx,
		name1.WithConds(in.GetConds()),
		name1.WithOffset(in.GetOffset()),
		name1.WithLimit(in.GetLimit()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetDisplayNames",
			"In", in,
			"Error", err,
		)
		return &npool.GetDisplayNamesResponse{}, status.Error(codes.Aborted, err.Error())
	}

	infos, total, err := handler.GetDisplayNames(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetDisplayNames",
			"In", in,
			"Error", err,
		)
		return &npool.GetDisplayNamesResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.GetDisplayNamesResponse{
		Infos: infos,
		Total: total,
	}, nil
}
