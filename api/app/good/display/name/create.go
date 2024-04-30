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

func (s *Server) CreateDisplayName(ctx context.Context, in *npool.CreateDisplayNameRequest) (*npool.CreateDisplayNameResponse, error) {
	req := in.GetInfo()
	if req == nil {
		logger.Sugar().Errorw(
			"CreateDisplayName",
			"In", in,
		)
		return &npool.CreateDisplayNameResponse{}, status.Error(codes.Aborted, "invalid argument")
	}
	handler, err := name1.NewHandler(
		ctx,
		name1.WithEntID(req.EntID, false),
		name1.WithAppGoodID(req.AppGoodID, true),
		name1.WithName(req.Name, true),
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
			"CreateDisplayName",
			"In", in,
			"Error", err,
		)
		return &npool.CreateDisplayNameResponse{}, status.Error(codes.Aborted, err.Error())
	}

	if err := handler.CreateDisplayName(ctx); err != nil {
		logger.Sugar().Errorw(
			"CreateDisplayName",
			"In", in,
			"Error", err,
		)
		return &npool.CreateDisplayNameResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.CreateDisplayNameResponse{}, nil
}
