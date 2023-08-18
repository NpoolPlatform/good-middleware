package like

import (
	"context"

	like1 "github.com/NpoolPlatform/good-middleware/pkg/mw/good/like"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/good/like"
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
		like1.WithID(req.ID, true),
		like1.WithAppID(req.AppID, true),
		like1.WithGoodID(req.GoodID, true),
		like1.WithUserID(req.UserID, true),
		like1.WithLike(req.Like, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateLike",
			"In", in,
			"Error", err,
		)
		return &npool.CreateLikeResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.CreateLike(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateLike",
			"In", in,
			"Error", err,
		)
		return &npool.CreateLikeResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.CreateLikeResponse{
		Info: info,
	}, nil
}
