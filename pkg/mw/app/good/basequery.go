package good

import (
	"entgo.io/ent/dialect/sql"

	wlog "github.com/NpoolPlatform/go-service-framework/pkg/wlog"
	goodbasecrud "github.com/NpoolPlatform/good-middleware/pkg/crud/app/good/goodbase"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent"
	entappgoodbase "github.com/NpoolPlatform/good-middleware/pkg/db/ent/appgoodbase"
	entgoodbase "github.com/NpoolPlatform/good-middleware/pkg/db/ent/goodbase"
)

type baseQueryHandler struct {
	*Handler
	stmSelect *ent.AppGoodBaseSelect
}

func (h *baseQueryHandler) selectGood(stm *ent.AppGoodBaseQuery) *ent.AppGoodBaseSelect {
	return stm.Select(entgoodbase.FieldID)
}

func (h *baseQueryHandler) queryGood(cli *ent.Client) {
	h.stmSelect = h.selectGood(
		cli.AppGoodBase.
			Query().
			Where(
				entappgoodbase.EntID(*h.EntID),
				entappgoodbase.DeletedAt(0),
			),
	)
}

func (h *baseQueryHandler) queryGoods(cli *ent.Client) (*ent.AppGoodBaseSelect, error) {
	stm, err := goodbasecrud.SetQueryConds(cli.AppGoodBase.Query(), h.AppGoodConds)
	if err != nil {
		return nil, wlog.WrapError(err)
	}
	return h.selectGood(stm), nil
}

func (h *baseQueryHandler) queryJoinMyself(s *sql.Selector) {
	t1 := sql.Table(entappgoodbase.Table)
	s.Join(t1).
		On(
			s.C(entappgoodbase.FieldID),
			t1.C(entappgoodbase.FieldID),
		).
		AppendSelect(
			t1.C(entappgoodbase.FieldEntID),
			t1.C(entappgoodbase.FieldGoodID),
			sql.As(t1.C(entappgoodbase.FieldPurchasable), "app_good_purchasable"),
			t1.C(entappgoodbase.FieldEnableProductPage),
			t1.C(entappgoodbase.FieldProductPage),
			sql.As(t1.C(entappgoodbase.FieldOnline), "app_good_online"),
			t1.C(entappgoodbase.FieldVisible),
			sql.As(t1.C(entappgoodbase.FieldName), "app_good_name"),
			t1.C(entappgoodbase.FieldDisplayIndex),
			t1.C(entappgoodbase.FieldBanner),
		)
}

func (h *baseQueryHandler) queryJoinGoodBase(s *sql.Selector) {
	t1 := sql.Table(entgoodbase.Table)
	s.Join(t1).
		On(
			s.C(entappgoodbase.FieldGoodID),
			t1.C(entgoodbase.FieldEntID),
		)
	s.AppendSelect(
		t1.C(entgoodbase.FieldGoodType),
		t1.C(entgoodbase.FieldBenefitType),
		sql.As(t1.C(entgoodbase.FieldName), "good_name"),
		t1.C(entgoodbase.FieldServiceStartAt),
		t1.C(entgoodbase.FieldStartMode),
		t1.C(entgoodbase.FieldTestOnly),
		t1.C(entgoodbase.FieldBenefitIntervalHours),
		sql.As(t1.C(entgoodbase.FieldPurchasable), "good_purchasable"),
		sql.As(t1.C(entgoodbase.FieldOnline), "good_online"),
	)
}

func (h *baseQueryHandler) queryJoin() {
	h.stmSelect.Modify(func(s *sql.Selector) {
		h.queryJoinMyself(s)
		h.queryJoinGoodBase(s)
	})
}
