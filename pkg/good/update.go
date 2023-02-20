package good

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	extramgrpb "github.com/NpoolPlatform/message/npool/good/mgr/v1/extrainfo"
	goodmgrpb "github.com/NpoolPlatform/message/npool/good/mgr/v1/good"
	stockmgrpb "github.com/NpoolPlatform/message/npool/good/mgr/v1/stock"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/good"

	extracrud "github.com/NpoolPlatform/good-manager/pkg/crud/extrainfo"
	goodcrud "github.com/NpoolPlatform/good-manager/pkg/crud/good"
	stockcrud "github.com/NpoolPlatform/good-manager/pkg/crud/stock"

	"github.com/NpoolPlatform/good-manager/pkg/db"
	"github.com/NpoolPlatform/good-manager/pkg/db/ent"

	entappgood "github.com/NpoolPlatform/good-manager/pkg/db/ent/appgood"
	entextra "github.com/NpoolPlatform/good-manager/pkg/db/ent/extrainfo"
	entgood "github.com/NpoolPlatform/good-manager/pkg/db/ent/good"
	entstock "github.com/NpoolPlatform/good-manager/pkg/db/ent/stock"

	"github.com/google/uuid"

	npoolpb "github.com/NpoolPlatform/message/npool"
)

// nolint
func UpdateGood(ctx context.Context, in *npool.GoodReq) (*npool.Good, error) {
	var err error

	err = checkGoodExist(ctx, in)
	if err != nil {
		return nil, err
	}

	err = db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		info, err := tx.Good.Query().Where(entgood.ID(uuid.MustParse(in.GetID()))).ForUpdate().Only(_ctx)
		if err != nil {
			return err
		}

		u, err := goodcrud.UpdateSet(
			info,
			&goodmgrpb.GoodReq{
				DeviceInfoID:           in.DeviceInfoID,
				DurationDays:           in.DurationDays,
				CoinTypeID:             in.CoinTypeID,
				InheritFromGoodID:      in.InheritFromGoodID,
				VendorLocationID:       in.VendorLocationID,
				Price:                  in.Price,
				Title:                  in.Title,
				Unit:                   in.Unit,
				UnitAmount:             in.UnitAmount,
				SupportCoinTypeIDs:     in.SupportCoinTypeIDs,
				DeliveryAt:             in.DeliveryAt,
				StartAt:                in.StartAt,
				TestOnly:               in.TestOnly,
				BenefitIntervalHours:   in.BenefitIntervalHours,
				BenefitState:           in.BenefitState,
				BenefitTIDs:            in.BenefitTIDs,
				NextBenefitStartAmount: in.NextBenefitStartAmount,
				LastBenefitAmount:      in.LastBenefitAmount,
			},
		)
		if err != nil {
			return err
		}

		_, err = u.Save(_ctx)
		if err != nil {
			return err
		}

		extra, err := tx.
			ExtraInfo.
			Query().
			Where(
				entextra.GoodID(uuid.MustParse(in.GetID())),
			).
			ForUpdate().
			Only(ctx)
		if err != nil {
			return err
		}

		u1, err := extracrud.UpdateSet(
			extra.Update(),
			&extramgrpb.ExtraInfoReq{
				Posters: in.Posters,
				Labels:  in.Labels,
			},
		)
		if err != nil {
			return err
		}

		_, err = u1.Save(_ctx)
		if err != nil {
			return err
		}

		stock, err := tx.
			Stock.
			Query().
			Where(
				entstock.GoodID(uuid.MustParse(in.GetID())),
			).
			ForUpdate().
			Only(ctx)
		if err != nil {
			return err
		}

		u2, err := stockcrud.UpdateSet(
			stock.Update(),
			&stockmgrpb.StockReq{
				Total: in.Total,
			},
		)
		if err != nil {
			return err
		}

		_, err = u2.Save(_ctx)
		if err != nil {
			return err
		}

		u3, err := stockcrud.AddFieldSet(
			stock,
			&stockmgrpb.StockReq{
				Locked:    in.Locked,
				InService: in.InService,
				WaitStart: in.WaitStart,
				Sold:      in.Sold,
			},
		)
		if err != nil {
			return err
		}

		_, err = u3.Save(_ctx)
		if err != nil {
			return err
		}

		if in.StartAt == nil {
			return nil
		}

		appGoods, err := tx.
			AppGood.
			Query().
			Where(
				entappgood.GoodID(uuid.MustParse(in.GetID())),
			).
			All(_ctx)
		if err != nil {
			return err
		}

		for _, ag := range appGoods {
			if in.GetStartAt() > ag.ServiceStartAt {
				_, err := tx.
					AppGood.
					UpdateOneID(ag.ID).
					SetServiceStartAt(in.GetStartAt()).
					Save(_ctx)
				if err != nil {
					return err
				}
			}
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return GetGood(ctx, in.GetID())
}

func checkGoodExist(ctx context.Context, in *npool.GoodReq) error {
	goodExist, err := goodcrud.Exist(ctx, uuid.MustParse(in.GetID()))
	if err != nil {
		return err
	}
	if !goodExist {
		return fmt.Errorf("good not exsit")
	}

	extraExist, err := extracrud.ExistConds(ctx, &extramgrpb.Conds{
		GoodID: &npoolpb.StringVal{
			Op:    cruder.EQ,
			Value: in.GetID(),
		},
	})
	if err != nil {
		return err
	}
	if !extraExist {
		_, err = extracrud.Create(ctx, &extramgrpb.ExtraInfoReq{
			GoodID:  in.ID,
			Posters: in.Posters,
			Labels:  in.Labels,
		})
		if err != nil {
			return err
		}
	}

	stockExist, err := stockcrud.ExistConds(ctx, &stockmgrpb.Conds{
		GoodID: &npoolpb.StringVal{
			Op:    cruder.EQ,
			Value: in.GetID(),
		},
	})
	if err != nil {
		return err
	}

	if !stockExist {
		_, err = stockcrud.Create(ctx, &stockmgrpb.StockReq{
			GoodID: in.ID,
			Total:  in.Total,
		})
		if err != nil {
			return err
		}
	}

	return nil
}
