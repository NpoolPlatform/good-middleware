//nolint:dupl
package name

import (
	"context"

	name1 "github.com/NpoolPlatform/good-middleware/pkg/mw/app/good/display/name"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/app/good/display/name"
)

func (s *Server) UpdateDisplayName(ctx context.Context, in *npool.UpdateDisplayNameRequest) (*npool.UpdateDisplayNameResponse, error) {
	req := in.GetInfo()
	if req == nil {
		logger.Sugar().Errorw(
			"UpdateDisplayName",
			"In", in,
		)
		return &npool.UpdateDisplayNameResponse{}, status.Error(codes.Aborted, "invalid argument")
	}
	handler, err := name1.NewHandler(
		ctx,
		name1.WithID(req.ID, false),
		name1.WithEntID(req.EntID, false),
		name1.WithName(req.Name, false),
		name1.WithIndex(func() *uint8 {
			if req.Index == nil {
				return nil
			}
			u := uint8(*req.Index)
			return &u
		}(), false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateDisplayName",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateDisplayNameResponse{}, status.Error(codes.Aborted, err.Error())
	}

	if err := handler.UpdateDisplayName(ctx); err != nil {
		logger.Sugar().Errorw(
			"UpdateDisplayName",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateDisplayNameResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.UpdateDisplayNameResponse{}, nil
}
