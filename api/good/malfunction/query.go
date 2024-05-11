package malfunction

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	malfunction1 "github.com/NpoolPlatform/good-middleware/pkg/mw/good/malfunction"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/good/malfunction"
)

func (s *Server) GetMalfunction(ctx context.Context, in *npool.GetMalfunctionRequest) (*npool.GetMalfunctionResponse, error) {
	handler, err := malfunction1.NewHandler(
		ctx,
		malfunction1.WithEntID(&in.EntID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetMalfunction",
			"In", in,
			"Error", err,
		)
		return &npool.GetMalfunctionResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.GetMalfunction(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetMalfunction",
			"In", in,
			"Error", err,
		)
		return &npool.GetMalfunctionResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.GetMalfunctionResponse{
		Info: info,
	}, nil
}

func (s *Server) GetMalfunctions(ctx context.Context, in *npool.GetMalfunctionsRequest) (*npool.GetMalfunctionsResponse, error) {
	handler, err := malfunction1.NewHandler(
		ctx,
		malfunction1.WithConds(in.GetConds()),
		malfunction1.WithOffset(in.GetOffset()),
		malfunction1.WithLimit(in.GetLimit()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetMalfunctions",
			"In", in,
			"Error", err,
		)
		return &npool.GetMalfunctionsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	infos, total, err := handler.GetMalfunctions(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetMalfunctions",
			"In", in,
			"Error", err,
		)
		return &npool.GetMalfunctionsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.GetMalfunctionsResponse{
		Infos: infos,
		Total: total,
	}, nil
}
