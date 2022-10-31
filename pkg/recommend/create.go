package recommend

import (
	"context"

	converter "github.com/NpoolPlatform/good-manager/pkg/converter/recommend"
	mgr "github.com/NpoolPlatform/good-manager/pkg/crud/recommend"
	"github.com/NpoolPlatform/good-manager/pkg/db"
	"github.com/NpoolPlatform/good-manager/pkg/db/ent"
	constant "github.com/NpoolPlatform/good-middleware/pkg/message/const"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	"github.com/NpoolPlatform/message/npool"
	mgrpb "github.com/NpoolPlatform/message/npool/good/mgr/v1/recommend"
	"go.opentelemetry.io/otel"
	scodes "go.opentelemetry.io/otel/codes"
)

func CreateRecommend(ctx context.Context, in *mgrpb.RecommendReq) (*mgrpb.Recommend, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CreateGood")
	defer span.End()
	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	recommend, err := mgr.RowOnly(ctx, &mgrpb.Conds{
		AppID: &npool.StringVal{
			Op:    cruder.EQ,
			Value: in.GetAppID(),
		},
		GoodID: &npool.StringVal{
			Op:    cruder.EQ,
			Value: in.GetGoodID(),
		},
	})
	if err != nil && !ent.IsNotFound(err) {
		return nil, err
	}

	var info *ent.Recommend
	err = db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		if recommend != nil {
			_, err = mgr.Delete(ctx, recommend.ID.String())
			if err != nil {
				return err
			}
		}
		c := tx.Recommend.Create()
		stm, err := mgr.CreateSet(c, &mgrpb.RecommendReq{
			ID:             in.ID,
			AppID:          in.AppID,
			GoodID:         in.GoodID,
			RecommenderID:  in.RecommenderID,
			Message:        in.Message,
			RecommendIndex: in.RecommendIndex,
		})
		if err != nil {
			return err
		}

		info, err = stm.Save(_ctx)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return converter.Ent2Grpc(info), nil
}
