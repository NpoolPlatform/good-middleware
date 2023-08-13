package location

import (
	"context"
	"fmt"

	// deviceinfocrud "github.com/NpoolPlatform/good-middleware/pkg/crud/deviceinfo"
	"github.com/NpoolPlatform/good-middleware/pkg/db"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent"
	entlocation "github.com/NpoolPlatform/good-middleware/pkg/db/ent/vendorlocation"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/vender/location"
)

type queryHandler struct {
	*Handler
	stmSelect *ent.VendorLocationSelect
	infos     []*npool.Location
	total     uint32
}

func (h *queryHandler) selectVendorLocation(stm *ent.VendorLocationQuery) {
	h.stmSelect = stm.Select(
		entlocation.FieldID,
		entlocation.FieldCountry,
		entlocation.FieldProvince,
		entlocation.FieldCity,
		entlocation.FieldAddress,
		entlocation.FieldBrandID,
		entlocation.FieldCreatedAt,
		entlocation.FieldUpdatedAt,
	)
}

func (h *queryHandler) queryVendorLocation(cli *ent.Client) {
	h.selectVendorLocation(
		cli.VendorLocation.
			Query().
			Where(
				entlocation.ID(*h.ID),
				entlocation.DeletedAt(0),
			),
	)
}

func (h *queryHandler) scan(ctx context.Context) error {
	return h.stmSelect.Scan(ctx, &h.infos)
}

func (h *Handler) GetLocation(ctx context.Context) (*npool.Location, error) {
	if h.ID == nil {
		return nil, fmt.Errorf("invalid id")
	}

	handler := &queryHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		handler.queryVendorLocation(cli)
		return handler.scan(_ctx)
	})
	if err != nil {
		return nil, err
	}
	if len(handler.infos) == 0 {
		return nil, nil
	}
	if len(handler.infos) > 1 {
		return nil, fmt.Errorf("too many records")
	}
	return handler.infos[0], nil
}