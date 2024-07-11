package fee

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	fee1 "github.com/NpoolPlatform/good-middleware/pkg/mw/fee"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/fee"
)

func (s *Server) GetFee(ctx context.Context, in *npool.GetFeeRequest) (*npool.GetFeeResponse, error) {
	handler, err := fee1.NewHandler(
		ctx,
		fee1.WithGoodID(&in.GoodID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetFee",
			"In", in,
			"Error", err,
		)
		return &npool.GetFeeResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.GetFee(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetFee",
			"In", in,
			"Error", err,
		)
		return &npool.GetFeeResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.GetFeeResponse{
		Info: info,
	}, nil
}

func (s *Server) GetFees(ctx context.Context, in *npool.GetFeesRequest) (*npool.GetFeesResponse, error) {
	handler, err := fee1.NewHandler(
		ctx,
		fee1.WithConds(in.GetConds()),
		fee1.WithOffset(in.GetOffset()),
		fee1.WithLimit(in.GetLimit()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetFees",
			"In", in,
			"Error", err,
		)
		return &npool.GetFeesResponse{}, status.Error(codes.Aborted, err.Error())
	}

	infos, total, err := handler.GetFees(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetFees",
			"In", in,
			"Error", err,
		)
		return &npool.GetFeesResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.GetFeesResponse{
		Infos: infos,
		Total: total,
	}, nil
}
