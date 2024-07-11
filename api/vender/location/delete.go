package location

import (
	"context"

	location1 "github.com/NpoolPlatform/good-middleware/pkg/mw/vender/location"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/vender/location"
)

func (s *Server) DeleteLocation(ctx context.Context, in *npool.DeleteLocationRequest) (*npool.DeleteLocationResponse, error) {
	req := in.GetInfo()
	if req == nil {
		logger.Sugar().Errorw(
			"DeleteLocation",
			"In", in,
		)
		return &npool.DeleteLocationResponse{}, status.Error(codes.Aborted, "invalid argument")
	}
	handler, err := location1.NewHandler(
		ctx,
		location1.WithID(req.ID, false),
		location1.WithEntID(req.EntID, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteLocation",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteLocationResponse{}, status.Error(codes.Aborted, err.Error())
	}

	if err := handler.DeleteLocation(ctx); err != nil {
		logger.Sugar().Errorw(
			"DeleteLocation",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteLocationResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.DeleteLocationResponse{}, nil
}
