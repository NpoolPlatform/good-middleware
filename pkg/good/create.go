package good

import (
	"context"

	"go.opentelemetry.io/otel"
	scodes "go.opentelemetry.io/otel/codes"

	constant "github.com/NpoolPlatform/good-middleware/pkg/message/const"
	commontracer "github.com/NpoolPlatform/good-middleware/pkg/tracer"

	extrainfomgrpb "github.com/NpoolPlatform/message/npool/good/mgr/v1/extrainfo"
	goodmgrpb "github.com/NpoolPlatform/message/npool/good/mgr/v1/good"
	stockmgrpb "github.com/NpoolPlatform/message/npool/good/mgr/v1/stock"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/good"

	"github.com/NpoolPlatform/good-manager/pkg/db"
	"github.com/NpoolPlatform/good-manager/pkg/db/ent"

	extrainfocrud "github.com/NpoolPlatform/good-manager/pkg/crud/extrainfo"
	goodcrud "github.com/NpoolPlatform/good-manager/pkg/crud/good"
	stockcrud "github.com/NpoolPlatform/good-manager/pkg/crud/stock"
)

func CreateGood(ctx context.Context, in *npool.GoodReq) (*npool.Good, error) {
	var id string
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CreateGood")
	defer span.End()
	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = commontracer.TraceInvoker(span, "good", "middleware", "CRUD")

	err = db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		c := tx.Good.Create()
		stm, err := goodcrud.CreateSet(c, &goodmgrpb.GoodReq{
			ID:                 in.ID,
			DeviceInfoID:       in.DeviceInfoID,
			DurationDays:       in.DurationDays,
			CoinTypeID:         in.CoinTypeID,
			InheritFromGoodID:  in.InheritFromGoodID,
			VendorLocationID:   in.VendorLocationID,
			BenefitType:        in.BenefitType,
			GoodType:           in.GoodType,
			Price:              in.Price,
			Title:              in.Title,
			Unit:               in.Unit,
			UnitAmount:         in.UnitAmount,
			SupportCoinTypeIDs: in.SupportCoinTypeIDs,
			DeliveryAt:         in.DeliveryAt,
			StartAt:            in.StartAt,
			TestOnly:           in.TestOnly,
		})
		if err != nil {
			return err
		}

		info, err := stm.Save(_ctx)
		if err != nil {
			return err
		}

		id = info.ID.String()

		c1 := tx.Stock.Create()
		stm1, err := stockcrud.CreateSet(c1, &stockmgrpb.StockReq{
			GoodID: &id,
			Total:  in.Total,
		})
		if err != nil {
			return err
		}

		_, err = stm1.Save(_ctx)
		if err != nil {
			return err
		}

		c2 := tx.ExtraInfo.Create()
		stm2, err := extrainfocrud.CreateSet(c2, &extrainfomgrpb.ExtraInfoReq{
			GoodID:  &id,
			Posters: in.Posters,
			Labels:  in.Labels,
		})
		if err != nil {
			return err
		}

		_, err = stm2.Save(_ctx)
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return GetGood(ctx, id)
}
