package required

import (
	"context"

	required1 "github.com/NpoolPlatform/good-middleware/pkg/mw/good/required"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/good/required"
)

func (s *Server) CreateRequired(ctx context.Context, in *npool.CreateRequiredRequest) (*npool.CreateRequiredResponse, error) {
	req := in.GetInfo()
	if req == nil {
		logger.Sugar().Errorw(
			"CreateRequired",
			"In", in,
		)
		return &npool.CreateRequiredResponse{}, status.Error(codes.Aborted, "invalid argument")
	}
	handler, err := required1.NewHandler(
		ctx,
		required1.WithID(req.ID, false),
		required1.WithMainGoodID(req.MainGoodID, true),
		required1.WithRequiredGoodID(req.RequiredGoodID, true),
		required1.WithMust(req.Must, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateRequired",
			"In", in,
			"Error", err,
		)
		return &npool.CreateRequiredResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.CreateRequired(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateRequired",
			"In", in,
			"Error", err,
		)
		return &npool.CreateRequiredResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.CreateRequiredResponse{
		Info: info,
	}, nil
}
