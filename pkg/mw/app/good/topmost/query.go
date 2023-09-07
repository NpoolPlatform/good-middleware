package topmost

import (
	"context"
	"encoding/json"
	"fmt"

	"entgo.io/ent/dialect/sql"

	topmostcrud "github.com/NpoolPlatform/good-middleware/pkg/crud/app/good/topmost"
	"github.com/NpoolPlatform/good-middleware/pkg/db"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent"
	enttopmost "github.com/NpoolPlatform/good-middleware/pkg/db/ent/topmost"
	types "github.com/NpoolPlatform/message/npool/basetypes/good/v1"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/app/good/topmost"

	"github.com/shopspring/decimal"
)

type queryHandler struct {
	*Handler
	stmSelect *ent.TopMostSelect
	stmCount  *ent.TopMostSelect
	infos     []*npool.TopMost
	total     uint32
}

func (h *queryHandler) selectTopMost(stm *ent.TopMostQuery) *ent.TopMostSelect {
	return stm.Select(enttopmost.FieldID)
}

func (h *queryHandler) queryTopMost(cli *ent.Client) {
	h.stmSelect = h.selectTopMost(
		cli.TopMost.
			Query().
			Where(
				enttopmost.ID(*h.ID),
				enttopmost.DeletedAt(0),
			),
	)
}

func (h *queryHandler) queryTopMosts(cli *ent.Client) (*ent.TopMostSelect, error) {
	stm, err := topmostcrud.SetQueryConds(cli.TopMost.Query(), h.Conds)
	if err != nil {
		return nil, err
	}
	return h.selectTopMost(stm), nil
}

func (h *queryHandler) queryJoinMyself(s *sql.Selector) {
	t := sql.Table(enttopmost.Table)
	s.LeftJoin(t).
		On(
			s.C(enttopmost.FieldID),
			t.C(enttopmost.FieldID),
		).
		AppendSelect(
			sql.As(t.C(enttopmost.FieldAppID), "app_id"),
			sql.As(t.C(enttopmost.FieldTopMostType), "top_most_type"),
			sql.As(t.C(enttopmost.FieldTitle), "title"),
			sql.As(t.C(enttopmost.FieldMessage), "message"),
			sql.As(t.C(enttopmost.FieldPosters), "posters"),
			sql.As(t.C(enttopmost.FieldStartAt), "start_at"),
			sql.As(t.C(enttopmost.FieldEndAt), "end_at"),
			sql.As(t.C(enttopmost.FieldThresholdCredits), "threshold_credits"),
			sql.As(t.C(enttopmost.FieldRegisterElapsedSeconds), "register_elapsed_seconds"),
			sql.As(t.C(enttopmost.FieldThresholdPurchases), "threshold_purchases"),
			sql.As(t.C(enttopmost.FieldThresholdPaymentAmount), "threshold_payment_amount"),
			sql.As(t.C(enttopmost.FieldKycMust), "kyc_must"),
			sql.As(t.C(enttopmost.FieldCreatedAt), "created_at"),
			sql.As(t.C(enttopmost.FieldUpdatedAt), "updated_at"),
		)
}

func (h *queryHandler) queryJoin() {
	h.stmSelect.Modify(func(s *sql.Selector) {
		h.queryJoinMyself(s)
	})
	if h.stmCount == nil {
		return
	}
	h.stmCount.Modify(func(s *sql.Selector) {
	})
}

func (h *queryHandler) scan(ctx context.Context) error {
	return h.stmSelect.Scan(ctx, &h.infos)
}

func (h *queryHandler) formalize() {
	for _, info := range h.infos {
		info.TopMostType = types.GoodTopMostType(types.GoodTopMostType_value[info.TopMostTypeStr])
		_ = json.Unmarshal([]byte(info.PostersStr), &info.Posters)
		amount, err := decimal.NewFromString(info.ThresholdCredits)
		if err != nil {
			info.ThresholdCredits = decimal.NewFromInt(0).String()
		} else {
			info.ThresholdCredits = amount.String()
		}
		amount, err = decimal.NewFromString(info.ThresholdPaymentAmount)
		if err != nil {
			info.ThresholdPaymentAmount = decimal.NewFromInt(0).String()
		} else {
			info.ThresholdPaymentAmount = amount.String()
		}
	}
}

func (h *Handler) GetTopMost(ctx context.Context) (*npool.TopMost, error) {
	handler := &queryHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		handler.queryTopMost(cli)
		handler.queryJoin()
		return handler.scan(ctx)
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

	handler.formalize()

	return handler.infos[0], nil
}

func (h *Handler) GetTopMosts(ctx context.Context) ([]*npool.TopMost, uint32, error) {
	handler := &queryHandler{
		Handler: h,
	}

	var err error

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		handler.stmSelect, err = handler.queryTopMosts(cli)
		if err != nil {
			return err
		}
		handler.stmCount, err = handler.queryTopMosts(cli)
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

	handler.formalize()

	return handler.infos, handler.total, nil
}
