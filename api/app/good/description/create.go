//nolint:dupl
package description

import (
	"context"

	description1 "github.com/NpoolPlatform/good-middleware/pkg/mw/app/good/description"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/app/good/description"
)

func (s *Server) CreateDescription(ctx context.Context, in *npool.CreateDescriptionRequest) (*npool.CreateDescriptionResponse, error) {
	req := in.GetInfo()
	if req == nil {
		logger.Sugar().Errorw(
			"CreateDescription",
			"In", in,
		)
		return &npool.CreateDescriptionResponse{}, status.Error(codes.Aborted, "invalid argument")
	}
	handler, err := description1.NewHandler(
		ctx,
		description1.WithEntID(req.EntID, false),
		description1.WithAppGoodID(req.AppGoodID, true),
		description1.WithDescription(req.Description, false),
		description1.WithIndex(func() *uint8 {
			if req.Index == nil {
				return nil
			}
			u := uint8(*req.Index)
			return &u
		}(), false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateDescription",
			"In", in,
			"Error", err,
		)
		return &npool.CreateDescriptionResponse{}, status.Error(codes.Aborted, err.Error())
	}

	if err := handler.CreateDescription(ctx); err != nil {
		logger.Sugar().Errorw(
			"CreateDescription",
			"In", in,
			"Error", err,
		)
		return &npool.CreateDescriptionResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.CreateDescriptionResponse{}, nil
}
