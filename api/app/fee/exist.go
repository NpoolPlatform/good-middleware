//nolint:dupl
package fee

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	fee1 "github.com/NpoolPlatform/good-middleware/pkg/mw/app/fee"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/app/fee"
)

func (s *Server) ExistFee(ctx context.Context, in *npool.ExistFeeRequest) (*npool.ExistFeeResponse, error) {
	handler, err := fee1.NewHandler(
		ctx,
		fee1.WithAppGoodID(&in.AppGoodID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistFee",
			"In", in,
			"Error", err,
		)
		return &npool.ExistFeeResponse{}, status.Error(codes.Aborted, err.Error())
	}

	exist, err := handler.ExistFee(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistFee",
			"In", in,
			"Error", err,
		)
		return &npool.ExistFeeResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.ExistFeeResponse{
		Info: exist,
	}, nil
}

func (s *Server) ExistFeeConds(ctx context.Context, in *npool.ExistFeeCondsRequest) (*npool.ExistFeeCondsResponse, error) {
	handler, err := fee1.NewHandler(
		ctx,
		fee1.WithConds(in.GetConds()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistFeeConds",
			"In", in,
			"Error", err,
		)
		return &npool.ExistFeeCondsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	exist, err := handler.ExistFeeConds(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistFeeConds",
			"In", in,
			"Error", err,
		)
		return &npool.ExistFeeCondsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.ExistFeeCondsResponse{
		Info: exist,
	}, nil
}
