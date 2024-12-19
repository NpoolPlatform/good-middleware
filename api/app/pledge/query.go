package pledge

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pledge1 "github.com/NpoolPlatform/good-middleware/pkg/mw/app/pledge"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/app/pledge"
)

func (s *Server) GetPledge(ctx context.Context, in *npool.GetPledgeRequest) (*npool.GetPledgeResponse, error) {
	handler, err := pledge1.NewHandler(
		ctx,
		pledge1.WithAppGoodID(&in.AppGoodID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetPledge",
			"In", in,
			"Error", err,
		)
		return &npool.GetPledgeResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.GetPledge(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetPledge",
			"In", in,
			"Error", err,
		)
		return &npool.GetPledgeResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.GetPledgeResponse{
		Info: info,
	}, nil
}

func (s *Server) GetPledges(ctx context.Context, in *npool.GetPledgesRequest) (*npool.GetPledgesResponse, error) {
	handler, err := pledge1.NewHandler(
		ctx,
		pledge1.WithConds(in.GetConds()),
		pledge1.WithOffset(in.GetOffset()),
		pledge1.WithLimit(in.GetLimit()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetPledges",
			"In", in,
			"Error", err,
		)
		return &npool.GetPledgesResponse{}, status.Error(codes.Aborted, err.Error())
	}

	infos, total, err := handler.GetPledges(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetPledges",
			"In", in,
			"Error", err,
		)
		return &npool.GetPledgesResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.GetPledgesResponse{
		Infos: infos,
		Total: total,
	}, nil
}
