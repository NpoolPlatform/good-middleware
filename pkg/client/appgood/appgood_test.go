package appgood

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"
	"time"

	"google.golang.org/grpc/credentials/insecure"

	"github.com/NpoolPlatform/go-service-framework/pkg/config"

	"bou.ke/monkey"
	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"
	"google.golang.org/grpc"

	// "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	// val "github.com/NpoolPlatform/message/npool"

	mgrpb "github.com/NpoolPlatform/message/npool/good/mgr/v1/appgood"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/appgood"

	goodmgrpb "github.com/NpoolPlatform/message/npool/good/mgr/v1/good"
	goodmwpb "github.com/NpoolPlatform/message/npool/good/mw/v1/good"

	testinit "github.com/NpoolPlatform/good-middleware/pkg/testinit"
	"github.com/google/uuid"

	"github.com/stretchr/testify/assert"

	deviceinfocrud "github.com/NpoolPlatform/good-manager/pkg/crud/deviceinfo"
	vendorlocationcrud "github.com/NpoolPlatform/good-manager/pkg/crud/vendorlocation"

	deviceinfomgrpb "github.com/NpoolPlatform/message/npool/good/mgr/v1/deviceinfo"
	vendorlocationmgrpb "github.com/NpoolPlatform/message/npool/good/mgr/v1/vendorlocation"

	goodcli "github.com/NpoolPlatform/good-middleware/pkg/client/good"
)

func init() {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}
	if err := testinit.Init(); err != nil {
		fmt.Printf("cannot init test stub: %v\n", err)
	}
}

var (
	inheritGoodID             = uuid.NewString()
	inheritGoodTitle          = "Test Inherit Good"
	inheritGoodPrice          = "12345.000000"
	inheritGoodType           = goodmgrpb.GoodType_GoodTypeClassicMining
	inheritGoodTypeStr        = goodmgrpb.GoodType_GoodTypeClassicMining.String()
	inheritGoodBenefitType    = goodmgrpb.BenefitType_BenefitTypePlatform
	inheritGoodBenefitTypeStr = goodmgrpb.BenefitType_BenefitTypePlatform.String()

	deviceInfoID     = uuid.NewString()
	vendorLocationID = uuid.NewString()

	supportCoinTypeIDs    = []string{uuid.NewString(), uuid.NewString()}
	supportCoinTypeIDsStr = fmt.Sprintf(`["%v", "%v"]`, supportCoinTypeIDs[0], supportCoinTypeIDs[1])

	posters    = []string{uuid.NewString(), uuid.NewString()}
	postersStr = fmt.Sprintf(`["%v", "%v"]`, posters[0], posters[1])

	labels    = []string{uuid.NewString(), uuid.NewString()}
	labelsStr = fmt.Sprintf(`["%v", "%v"]`, labels[0], labels[1])

	goodInfo = goodmwpb.Good{
		ID:                            uuid.NewString(),
		DeviceInfoID:                  deviceInfoID,
		DevicePostersStr:              "[]",
		DurationDays:                  365,
		CoinTypeID:                    uuid.NewString(),
		InheritFromGoodID:             &inheritGoodID,
		InheritFromGoodName:           &inheritGoodTitle,
		InheritFromGoodType:           &inheritGoodType,
		InheritFromGoodTypeStr:        &inheritGoodTypeStr,
		InheritFromGoodBenefitType:    &inheritGoodBenefitType,
		InheritFromGoodBenefitTypeStr: &inheritGoodBenefitTypeStr,
		VendorLocationID:              vendorLocationID,
		GoodType:                      goodmgrpb.GoodType_GoodTypeClassicMining,
		GoodTypeStr:                   goodmgrpb.GoodType_GoodTypeClassicMining.String(),
		BenefitType:                   goodmgrpb.BenefitType_BenefitTypePlatform,
		BenefitTypeStr:                goodmgrpb.BenefitType_BenefitTypePlatform.String(),
		Price:                         "123.000000000000000000",
		Title:                         uuid.NewString(),
		Unit:                          "TiB",
		UnitAmount:                    1,
		SupportCoinTypeIDs:            supportCoinTypeIDs,
		SupportCoinTypeIDsStr:         supportCoinTypeIDsStr,
		TestOnly:                      true,
		Posters:                       posters,
		PostersStr:                    postersStr,
		Labels:                        labels,
		LabelsStr:                     labelsStr,
		GoodTotal:                     1000,
		DeliveryAt:                    uint32(time.Now().Unix() + 1000),
		StartAt:                       uint32(time.Now().Unix() + 1000),
	}

	appGoodInfo = npool.Good{
		ID:                    uuid.NewString(),
		AppID:                 uuid.NewString(),
		GoodID:                goodInfo.ID,
		Online:                false,
		Visible:               false,
		Price:                 "123123123.000000000000000000",
		DisplayIndex:          1,
		PurchaseLimit:         3000,
		CommissionPercent:     30,
		DevicePostersStr:      goodInfo.DevicePostersStr,
		DurationDays:          goodInfo.DurationDays,
		CoinTypeID:            goodInfo.CoinTypeID,
		GoodType:              goodInfo.GoodType,
		GoodTypeStr:           goodInfo.GoodTypeStr,
		BenefitType:           goodInfo.BenefitType,
		BenefitTypeStr:        goodInfo.BenefitTypeStr,
		GoodName:              "My Test Good",
		Unit:                  goodInfo.Unit,
		UnitAmount:            goodInfo.UnitAmount,
		SupportCoinTypeIDs:    goodInfo.SupportCoinTypeIDs,
		SupportCoinTypeIDsStr: goodInfo.SupportCoinTypeIDsStr,
		TestOnly:              goodInfo.TestOnly,
		Posters:               goodInfo.Posters,
		PostersStr:            goodInfo.PostersStr,
		Labels:                goodInfo.Labels,
		LabelsStr:             goodInfo.LabelsStr,
		GoodTotal:             goodInfo.GoodTotal,
		StartAt:               goodInfo.StartAt,
		CreatedAt:             goodInfo.CreatedAt,
	}
)

var (
	goodReq = goodmwpb.GoodReq{
		ID:                 &goodInfo.ID,
		DeviceInfoID:       &goodInfo.DeviceInfoID,
		DurationDays:       &goodInfo.DurationDays,
		CoinTypeID:         &goodInfo.CoinTypeID,
		InheritFromGoodID:  goodInfo.InheritFromGoodID,
		VendorLocationID:   &goodInfo.VendorLocationID,
		Price:              &goodInfo.Price,
		BenefitType:        &goodInfo.BenefitType,
		GoodType:           &goodInfo.GoodType,
		Title:              &goodInfo.Title,
		Unit:               &goodInfo.Unit,
		UnitAmount:         &goodInfo.UnitAmount,
		SupportCoinTypeIDs: goodInfo.SupportCoinTypeIDs,
		DeliveryAt:         &goodInfo.DeliveryAt,
		StartAt:            &goodInfo.StartAt,
		TestOnly:           &goodInfo.TestOnly,
		Total:              &goodInfo.GoodTotal,
		Posters:            goodInfo.Posters,
		Labels:             goodInfo.Labels,
	}
)

var info *npool.Good

func createGood(t *testing.T) {
	var err error

	_, err = deviceinfocrud.Create(context.Background(), &deviceinfomgrpb.DeviceInfoReq{
		ID: &deviceInfoID,
	})
	assert.Nil(t, err)

	_, err = vendorlocationcrud.Create(context.Background(), &vendorlocationmgrpb.VendorLocationReq{
		ID: &vendorLocationID,
	})
	assert.Nil(t, err)

	_, err = goodcli.CreateGood(context.Background(), &goodmwpb.GoodReq{
		ID:                 &inheritGoodID,
		DeviceInfoID:       &goodInfo.DeviceInfoID,
		DurationDays:       &goodInfo.DurationDays,
		CoinTypeID:         &goodInfo.CoinTypeID,
		VendorLocationID:   &goodInfo.VendorLocationID,
		Price:              &inheritGoodPrice,
		BenefitType:        &goodInfo.BenefitType,
		GoodType:           &goodInfo.GoodType,
		Title:              &inheritGoodTitle,
		Unit:               &goodInfo.Unit,
		UnitAmount:         &goodInfo.UnitAmount,
		SupportCoinTypeIDs: goodInfo.SupportCoinTypeIDs,
		DeliveryAt:         &goodInfo.DeliveryAt,
		StartAt:            &goodInfo.StartAt,
		TestOnly:           &goodInfo.TestOnly,
		Total:              &goodInfo.GoodTotal,
		Posters:            goodInfo.Posters,
		Labels:             goodInfo.Labels,
	})
	assert.Nil(t, err)

	info1, err := goodcli.CreateGood(context.Background(), &goodReq)
	if assert.Nil(t, err) {
		goodInfo.CreatedAt = info1.CreatedAt
	}

	info, err := CreateGood(context.Background(), &mgrpb.AppGoodReq{
		ID:                &appGoodInfo.ID,
		AppID:             &appGoodInfo.AppID,
		GoodID:            &goodInfo.ID,
		Online:            &appGoodInfo.Online,
		Visible:           &appGoodInfo.Visible,
		GoodName:          &appGoodInfo.GoodName,
		Price:             &appGoodInfo.Price,
		DisplayIndex:      &appGoodInfo.DisplayIndex,
		PurchaseLimit:     &appGoodInfo.PurchaseLimit,
		CommissionPercent: &appGoodInfo.CommissionPercent,
	})
	if assert.Nil(t, err) {
		appGoodInfo.CreatedAt = info.CreatedAt
		assert.Equal(t, info, &appGoodInfo)
	}
}

func updateGood(t *testing.T) {
	/*
		var err error
		info, err = UpdateGood(context.Background(), &goodReq)
		if assert.Nil(t, err) {
			goodInfo.UpdatedAt = info.UpdatedAt
			assert.Equal(t, info, &goodInfo)
		}
	*/
}

func getGood(t *testing.T) {
	/*
		var err error
		info, err = GetGood(context.Background(), info.ID)
		if assert.Nil(t, err) {
			assert.Equal(t, info, &goodInfo)
		}
	*/
}

func getGoods(t *testing.T) {
	/*
		infos, total, err := GetGoods(context.Background(),
			&goodmgrpb.Conds{
				ID: &val.StringVal{
					Value: info.ID,
					Op:    cruder.EQ,
				},
			}, 0, 1)
		if assert.Nil(t, err) {
			assert.Equal(t, total, uint32(1))
			assert.Equal(t, infos[0], &goodInfo)
		}
	*/
}

func TestMainOrder(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	gport := config.GetIntValueWithNameSpace("", config.KeyGRPCPort)

	monkey.Patch(grpc2.GetGRPCConn, func(service string, tags ...string) (*grpc.ClientConn, error) {
		return grpc.Dial(fmt.Sprintf("localhost:%v", gport), grpc.WithTransportCredentials(insecure.NewCredentials()))
	})

	t.Run("createGood", createGood)
	t.Run("getGood", getGood)
	t.Run("getGoods", getGoods)
	t.Run("updateGood", updateGood)
}
