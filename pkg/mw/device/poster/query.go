package poster

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"

	devicepostercrud "github.com/NpoolPlatform/good-middleware/pkg/crud/device/poster"
	"github.com/NpoolPlatform/good-middleware/pkg/db"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent"
	entdevicetype "github.com/NpoolPlatform/good-middleware/pkg/db/ent/deviceinfo"
	entdeviceposter "github.com/NpoolPlatform/good-middleware/pkg/db/ent/deviceposter"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/device/poster"
)

type queryHandler struct {
	*Handler
	stmSelect *ent.DevicePosterSelect
	stmCount  *ent.DevicePosterSelect
	infos     []*npool.Poster
	total     uint32
}

func (h *queryHandler) selectPoster(stm *ent.DevicePosterQuery) *ent.DevicePosterSelect {
	return stm.Select(entdeviceposter.FieldID)
}

func (h *queryHandler) queryPoster(cli *ent.Client) error {
	if h.ID == nil && h.EntID == nil {
		return fmt.Errorf("invalid id")
	}
	stm := cli.DevicePoster.Query().Where(entdeviceposter.DeletedAt(0))
	if h.ID != nil {
		stm.Where(entdeviceposter.ID(*h.ID))
	}
	if h.EntID != nil {
		stm.Where(entdeviceposter.EntID(*h.EntID))
	}
	h.stmSelect = h.selectPoster(stm)
	return nil
}

func (h *queryHandler) queryPosters(cli *ent.Client) (*ent.DevicePosterSelect, error) {
	stm, err := devicepostercrud.SetQueryConds(cli.DevicePoster.Query(), h.PosterConds)
	if err != nil {
		return nil, err
	}
	return h.selectPoster(stm), nil
}

func (h *queryHandler) queryJoinMyself(s *sql.Selector) {
	t := sql.Table(entdeviceposter.Table)
	s.LeftJoin(t).
		On(
			s.C(entdeviceposter.FieldID),
			t.C(entdeviceposter.FieldID),
		).
		AppendSelect(
			t.C(entdeviceposter.FieldEntID),
			t.C(entdeviceposter.FieldDeviceTypeID),
			t.C(entdeviceposter.FieldPoster),
			t.C(entdeviceposter.FieldIndex),
			t.C(entdeviceposter.FieldCreatedAt),
			t.C(entdeviceposter.FieldUpdatedAt),
		)
}

func (h *queryHandler) queryJoinDeviceType(s *sql.Selector) {
	t1 := sql.Table(entdevicetype.Table)
	s.LeftJoin(t1).
		On(
			s.C(entdeviceposter.FieldDeviceTypeID),
			t1.C(entdevicetype.FieldEntID),
		).
		AppendSelect(
			sql.As(t1.C(entdevicetype.FieldType), "device_type"),
			t1.C(entdevicetype.FieldManufacturer),
		)
}

func (h *queryHandler) queryJoin() {
	h.stmSelect.Modify(func(s *sql.Selector) {
		h.queryJoinMyself(s)
		h.queryJoinDeviceType(s)
	})
	if h.stmCount == nil {
		return
	}
	h.stmCount.Modify(func(s *sql.Selector) {
		h.queryJoinDeviceType(s)
	})
}

func (h *queryHandler) scan(ctx context.Context) error {
	return h.stmSelect.Scan(ctx, &h.infos)
}

func (h *Handler) GetPoster(ctx context.Context) (*npool.Poster, error) {
	handler := &queryHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryPoster(cli); err != nil {
			return err
		}
		handler.queryJoin()
		return handler.scan(ctx)
	})
	if err != nil {
		return nil, err
	}
	if len(handler.infos) == 0 {
		return nil, fmt.Errorf("invalid deviceposter")
	}
	if len(handler.infos) > 1 {
		return nil, fmt.Errorf("too many records")
	}

	return handler.infos[0], nil
}

func (h *Handler) GetPosters(ctx context.Context) ([]*npool.Poster, uint32, error) {
	handler := &queryHandler{
		Handler: h,
	}

	var err error

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		handler.stmSelect, err = handler.queryPosters(cli)
		if err != nil {
			return err
		}
		handler.stmCount, err = handler.queryPosters(cli)
		if err != nil {
			return err
		}

		handler.queryJoin()

		total, err := handler.stmCount.Count(_ctx)
		if err != nil {
			return err
		}
		handler.total = uint32(total)

		handler.stmSelect.
			Offset(int(handler.Offset)).
			Limit(int(handler.Limit))
		return handler.scan(_ctx)
	})
	if err != nil {
		return nil, 0, err
	}

	return handler.infos, handler.total, nil
}
