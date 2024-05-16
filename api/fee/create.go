package fee

import (
	"context"

	fee1 "github.com/NpoolPlatform/good-middleware/pkg/mw/fee"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/fee"
)

func (s *Server) CreateFee(ctx context.Context, in *npool.CreateFeeRequest) (*npool.CreateFeeResponse, error) {
	req := in.GetInfo()
	if req == nil {
		logger.Sugar().Errorw(
			"CreateFee",
			"In", in,
		)
		return &npool.CreateFeeResponse{}, status.Error(codes.Aborted, "invalid argument")
	}
	handler, err := fee1.NewHandler(
		ctx,
		fee1.WithEntID(req.EntID, false),
		fee1.WithGoodID(req.GoodID, false),
		fee1.WithGoodType(req.GoodType, true),
		fee1.WithName(req.Name, true),
		fee1.WithSettlementType(req.SettlementType, true),
		fee1.WithUnitValue(req.UnitValue, true),
		fee1.WithDurationDisplayType(req.DurationDisplayType, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateFee",
			"In", in,
			"Error", err,
		)
		return &npool.CreateFeeResponse{}, status.Error(codes.Aborted, err.Error())
	}

	if err = handler.CreateFee(ctx); err != nil {
		logger.Sugar().Errorw(
			"CreateFee",
			"In", in,
			"Error", err,
		)
		return &npool.CreateFeeResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.CreateFeeResponse{}, nil
}
