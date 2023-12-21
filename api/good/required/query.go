package required

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	required1 "github.com/NpoolPlatform/good-middleware/pkg/mw/good/required"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/good/required"
)

func (s *Server) GetRequired(ctx context.Context, in *npool.GetRequiredRequest) (*npool.GetRequiredResponse, error) {
	handler, err := required1.NewHandler(
		ctx,
		required1.WithEntID(&in.EntID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetRequired",
			"In", in,
			"Error", err,
		)
		return &npool.GetRequiredResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.GetRequired(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetRequired",
			"In", in,
			"Error", err,
		)
		return &npool.GetRequiredResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.GetRequiredResponse{
		Info: info,
	}, nil
}

func (s *Server) GetRequireds(ctx context.Context, in *npool.GetRequiredsRequest) (*npool.GetRequiredsResponse, error) {
	handler, err := required1.NewHandler(
		ctx,
		required1.WithConds(in.GetConds()),
		required1.WithOffset(in.GetOffset()),
		required1.WithLimit(in.GetLimit()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetRequireds",
			"In", in,
			"Error", err,
		)
		return &npool.GetRequiredsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	infos, total, err := handler.GetRequireds(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetRequireds",
			"In", in,
			"Error", err,
		)
		return &npool.GetRequiredsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.GetRequiredsResponse{
		Infos: infos,
		Total: total,
	}, nil
}
