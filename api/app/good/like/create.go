package like

import (
	"context"

	like1 "github.com/NpoolPlatform/good-middleware/pkg/mw/app/good/like"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/app/good/like"
)

func (s *Server) CreateLike(ctx context.Context, in *npool.CreateLikeRequest) (*npool.CreateLikeResponse, error) {
	req := in.GetInfo()
	if req == nil {
		logger.Sugar().Errorw(
			"CreateLike",
			"In", in,
		)
		return &npool.CreateLikeResponse{}, status.Error(codes.Aborted, "invalid argument")
	}
	handler, err := like1.NewHandler(
		ctx,
		like1.WithEntID(req.EntID, false),
		like1.WithAppGoodID(req.AppGoodID, true),
		like1.WithLike(req.Like, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateLike",
			"In", in,
			"Error", err,
		)
		return &npool.CreateLikeResponse{}, status.Error(codes.Aborted, err.Error())
	}

	if err := handler.CreateLike(ctx); err != nil {
		logger.Sugar().Errorw(
			"CreateLike",
			"In", in,
			"Error", err,
		)
		return &npool.CreateLikeResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.CreateLikeResponse{}, nil
}
