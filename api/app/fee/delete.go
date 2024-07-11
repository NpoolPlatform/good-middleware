package fee

import (
	"context"

	fee1 "github.com/NpoolPlatform/good-middleware/pkg/mw/app/fee"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/app/fee"
)

func (s *Server) DeleteFee(ctx context.Context, in *npool.DeleteFeeRequest) (*npool.DeleteFeeResponse, error) {
	req := in.GetInfo()
	if req == nil {
		logger.Sugar().Errorw(
			"DeleteFee",
			"In", in,
		)
		return &npool.DeleteFeeResponse{}, status.Error(codes.Aborted, "invalid argument")
	}
	handler, err := fee1.NewHandler(
		ctx,
		fee1.WithID(req.ID, false),
		fee1.WithEntID(req.EntID, false),
		fee1.WithAppGoodID(req.AppGoodID, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteFee",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteFeeResponse{}, status.Error(codes.Aborted, err.Error())
	}

	if err := handler.DeleteFee(ctx); err != nil {
		logger.Sugar().Errorw(
			"DeleteFee",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteFeeResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.DeleteFeeResponse{}, nil
}
