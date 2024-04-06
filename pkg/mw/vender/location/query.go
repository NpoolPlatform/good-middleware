package location

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"

	locationcrud "github.com/NpoolPlatform/good-middleware/pkg/crud/vender/location"
	"github.com/NpoolPlatform/good-middleware/pkg/db"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent"
	entvendorbrand "github.com/NpoolPlatform/good-middleware/pkg/db/ent/vendorbrand"
	entvendorlocation "github.com/NpoolPlatform/good-middleware/pkg/db/ent/vendorlocation"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/vender/location"
)

type queryHandler struct {
	*Handler
	stmSelect *ent.VendorLocationSelect
	stmCount  *ent.VendorLocationSelect
	infos     []*npool.Location
	total     uint32
}

func (h *queryHandler) selectVendorLocation(stm *ent.VendorLocationQuery) *ent.VendorLocationSelect {
	return stm.Select(entvendorlocation.FieldID)
}

func (h *queryHandler) queryVendorLocation(cli *ent.Client) error {
	if h.ID == nil && h.EntID == nil {
		return fmt.Errorf("invalid id")
	}
	stm := cli.VendorLocation.Query().Where(entvendorlocation.DeletedAt(0))
	if h.ID != nil {
		stm.Where(entvendorlocation.ID(*h.ID))
	}
	if h.EntID != nil {
		stm.Where(entvendorlocation.EntID(*h.EntID))
	}
	h.stmSelect = h.selectVendorLocation(stm)
	return nil
}

func (h *queryHandler) queryVendorLocations(ctx context.Context, cli *ent.Client) (*ent.VendorLocationSelect, error) {
	stm, err := locationcrud.SetQueryConds(cli.VendorLocation.Query(), h.Conds)
	if err != nil {
		return nil, err
	}
	return h.selectVendorLocation(stm), nil
}

func (h *queryHandler) queryJoinMyself(s *sql.Selector) {
	t1 := sql.Table(entvendorlocation.Table)
	s.LeftJoin(t1).
		On(
			s.C(entvendorlocation.FieldEntID),
			t1.C(entvendorlocation.FieldEntID),
		).
		AppendSelect(
			t1.C(entvendorlocation.FieldEntID),
			t1.C(entvendorlocation.FieldCountry),
			t1.C(entvendorlocation.FieldProvince),
			t1.C(entvendorlocation.FieldCity),
			t1.C(entvendorlocation.FieldAddress),
			t1.C(entvendorlocation.FieldBrandID),
			t1.C(entvendorlocation.FieldCreatedAt),
			t1.C(entvendorlocation.FieldUpdatedAt),
		)
}

func (h *queryHandler) queryJoinBrand(s *sql.Selector) {
	t1 := sql.Table(entvendorbrand.Table)
	s.LeftJoin(t1).
		On(
			s.C(entvendorlocation.FieldBrandID),
			t1.C(entvendorbrand.FieldEntID),
		).
		AppendSelect(
			sql.As(t1.C(entvendorbrand.FieldName), "brand_name"),
			sql.As(t1.C(entvendorbrand.FieldLogo), "brand_logo"),
		)
}

func (h *queryHandler) queryJoin() {
	h.stmSelect.Modify(func(s *sql.Selector) {
		h.queryJoinMyself(s)
		h.queryJoinBrand(s)
	})
	if h.stmCount == nil {
		return
	}
	h.stmCount.Modify(func(s *sql.Selector) {
		h.queryJoinBrand(s)
	})
}

func (h *queryHandler) scan(ctx context.Context) error {
	return h.stmSelect.Scan(ctx, &h.infos)
}

func (h *Handler) GetLocation(ctx context.Context) (*npool.Location, error) {
	handler := &queryHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryVendorLocation(cli); err != nil {
			return err
		}
		handler.queryJoin()
		handler.stmSelect.
			Offset(0).
			Limit(2)
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

func (h *Handler) GetLocations(ctx context.Context) (infos []*npool.Location, total uint32, err error) {
	handler := &queryHandler{
		Handler: h,
	}
	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if handler.stmSelect, err = handler.queryVendorLocations(_ctx, cli); err != nil {
			return err
		}
		if handler.stmCount, err = handler.queryVendorLocations(_ctx, cli); err != nil {
			return err
		}

		handler.queryJoin()

		_total, err := handler.stmCount.Count(_ctx)
		if err != nil {
			return err
		}
		handler.total = uint32(_total)

		handler.stmSelect.
			Offset(int(h.Offset)).
			Limit(int(h.Limit))
		return handler.scan(_ctx)
	})
	if err != nil {
		return nil, 0, err
	}
	return handler.infos, handler.total, nil
}
