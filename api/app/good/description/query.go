package description

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	description1 "github.com/NpoolPlatform/good-middleware/pkg/mw/app/good/description"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/app/good/description"
)

func (s *Server) GetDescription(ctx context.Context, in *npool.GetDescriptionRequest) (*npool.GetDescriptionResponse, error) {
	handler, err := description1.NewHandler(
		ctx,
		description1.WithEntID(&in.EntID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetDescription",
			"In", in,
			"Error", err,
		)
		return &npool.GetDescriptionResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.GetDescription(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetDescription",
			"In", in,
			"Error", err,
		)
		return &npool.GetDescriptionResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.GetDescriptionResponse{
		Info: info,
	}, nil
}

func (s *Server) GetDescriptions(ctx context.Context, in *npool.GetDescriptionsRequest) (*npool.GetDescriptionsResponse, error) {
	handler, err := description1.NewHandler(
		ctx,
		description1.WithConds(in.GetConds()),
		description1.WithOffset(in.GetOffset()),
		description1.WithLimit(in.GetLimit()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetDescriptions",
			"In", in,
			"Error", err,
		)
		return &npool.GetDescriptionsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	infos, total, err := handler.GetDescriptions(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetDescriptions",
			"In", in,
			"Error", err,
		)
		return &npool.GetDescriptionsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.GetDescriptionsResponse{
		Infos: infos,
		Total: total,
	}, nil
}
