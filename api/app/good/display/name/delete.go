package name

import (
	"context"

	name1 "github.com/NpoolPlatform/good-middleware/pkg/mw/app/good/display/name"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/app/good/display/name"
)

func (s *Server) DeleteDisplayName(ctx context.Context, in *npool.DeleteDisplayNameRequest) (*npool.DeleteDisplayNameResponse, error) {
	req := in.GetInfo()
	if req == nil {
		logger.Sugar().Errorw(
			"DeleteDisplayName",
			"In", in,
		)
		return &npool.DeleteDisplayNameResponse{}, status.Error(codes.Aborted, "invalid argument")
	}
	handler, err := name1.NewHandler(
		ctx,
		name1.WithID(req.ID, false),
		name1.WithEntID(req.EntID, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteDisplayName",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteDisplayNameResponse{}, status.Error(codes.Aborted, err.Error())
	}

	if err := handler.DeleteDisplayName(ctx); err != nil {
		logger.Sugar().Errorw(
			"DeleteDisplayName",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteDisplayNameResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.DeleteDisplayNameResponse{}, nil
}
