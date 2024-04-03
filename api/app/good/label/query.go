package label

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	label1 "github.com/NpoolPlatform/good-middleware/pkg/mw/app/good/label"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/app/good/label"
)

func (s *Server) GetLabel(ctx context.Context, in *npool.GetLabelRequest) (*npool.GetLabelResponse, error) {
	handler, err := label1.NewHandler(
		ctx,
		label1.WithEntID(&in.EntID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetLabel",
			"In", in,
			"Error", err,
		)
		return &npool.GetLabelResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.GetLabel(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetLabel",
			"In", in,
			"Error", err,
		)
		return &npool.GetLabelResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.GetLabelResponse{
		Info: info,
	}, nil
}

func (s *Server) GetLabels(ctx context.Context, in *npool.GetLabelsRequest) (*npool.GetLabelsResponse, error) {
	handler, err := label1.NewHandler(
		ctx,
		label1.WithConds(in.GetConds()),
		label1.WithOffset(in.GetOffset()),
		label1.WithLimit(in.GetLimit()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetLabels",
			"In", in,
			"Error", err,
		)
		return &npool.GetLabelsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	infos, total, err := handler.GetLabels(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetLabels",
			"In", in,
			"Error", err,
		)
		return &npool.GetLabelsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.GetLabelsResponse{
		Infos: infos,
		Total: total,
	}, nil
}
