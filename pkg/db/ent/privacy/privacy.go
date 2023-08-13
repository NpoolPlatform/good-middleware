// Code generated by ent, DO NOT EDIT.

package privacy

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/good-middleware/pkg/db/ent"

	"entgo.io/ent/entql"
	"entgo.io/ent/privacy"
)

var (
	// Allow may be returned by rules to indicate that the policy
	// evaluation should terminate with allow decision.
	Allow = privacy.Allow

	// Deny may be returned by rules to indicate that the policy
	// evaluation should terminate with deny decision.
	Deny = privacy.Deny

	// Skip may be returned by rules to indicate that the policy
	// evaluation should continue to the next rule.
	Skip = privacy.Skip
)

// Allowf returns an formatted wrapped Allow decision.
func Allowf(format string, a ...interface{}) error {
	return fmt.Errorf(format+": %w", append(a, Allow)...)
}

// Denyf returns an formatted wrapped Deny decision.
func Denyf(format string, a ...interface{}) error {
	return fmt.Errorf(format+": %w", append(a, Deny)...)
}

// Skipf returns an formatted wrapped Skip decision.
func Skipf(format string, a ...interface{}) error {
	return fmt.Errorf(format+": %w", append(a, Skip)...)
}

// DecisionContext creates a new context from the given parent context with
// a policy decision attach to it.
func DecisionContext(parent context.Context, decision error) context.Context {
	return privacy.DecisionContext(parent, decision)
}

// DecisionFromContext retrieves the policy decision from the context.
func DecisionFromContext(ctx context.Context) (error, bool) {
	return privacy.DecisionFromContext(ctx)
}

type (
	// Policy groups query and mutation policies.
	Policy = privacy.Policy

	// QueryRule defines the interface deciding whether a
	// query is allowed and optionally modify it.
	QueryRule = privacy.QueryRule
	// QueryPolicy combines multiple query rules into a single policy.
	QueryPolicy = privacy.QueryPolicy

	// MutationRule defines the interface which decides whether a
	// mutation is allowed and optionally modifies it.
	MutationRule = privacy.MutationRule
	// MutationPolicy combines multiple mutation rules into a single policy.
	MutationPolicy = privacy.MutationPolicy
)

// QueryRuleFunc type is an adapter to allow the use of
// ordinary functions as query rules.
type QueryRuleFunc func(context.Context, ent.Query) error

// Eval returns f(ctx, q).
func (f QueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	return f(ctx, q)
}

// MutationRuleFunc type is an adapter which allows the use of
// ordinary functions as mutation rules.
type MutationRuleFunc func(context.Context, ent.Mutation) error

// EvalMutation returns f(ctx, m).
func (f MutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	return f(ctx, m)
}

// QueryMutationRule is an interface which groups query and mutation rules.
type QueryMutationRule interface {
	QueryRule
	MutationRule
}

// AlwaysAllowRule returns a rule that returns an allow decision.
func AlwaysAllowRule() QueryMutationRule {
	return fixedDecision{Allow}
}

// AlwaysDenyRule returns a rule that returns a deny decision.
func AlwaysDenyRule() QueryMutationRule {
	return fixedDecision{Deny}
}

type fixedDecision struct {
	decision error
}

func (f fixedDecision) EvalQuery(context.Context, ent.Query) error {
	return f.decision
}

func (f fixedDecision) EvalMutation(context.Context, ent.Mutation) error {
	return f.decision
}

type contextDecision struct {
	eval func(context.Context) error
}

// ContextQueryMutationRule creates a query/mutation rule from a context eval func.
func ContextQueryMutationRule(eval func(context.Context) error) QueryMutationRule {
	return contextDecision{eval}
}

func (c contextDecision) EvalQuery(ctx context.Context, _ ent.Query) error {
	return c.eval(ctx)
}

func (c contextDecision) EvalMutation(ctx context.Context, _ ent.Mutation) error {
	return c.eval(ctx)
}

// OnMutationOperation evaluates the given rule only on a given mutation operation.
func OnMutationOperation(rule MutationRule, op ent.Op) MutationRule {
	return MutationRuleFunc(func(ctx context.Context, m ent.Mutation) error {
		if m.Op().Is(op) {
			return rule.EvalMutation(ctx, m)
		}
		return Skip
	})
}

// DenyMutationOperationRule returns a rule denying specified mutation operation.
func DenyMutationOperationRule(op ent.Op) MutationRule {
	rule := MutationRuleFunc(func(_ context.Context, m ent.Mutation) error {
		return Denyf("ent/privacy: operation %s is not allowed", m.Op())
	})
	return OnMutationOperation(rule, op)
}

// The AppDefaultGoodQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type AppDefaultGoodQueryRuleFunc func(context.Context, *ent.AppDefaultGoodQuery) error

// EvalQuery return f(ctx, q).
func (f AppDefaultGoodQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.AppDefaultGoodQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.AppDefaultGoodQuery", q)
}

// The AppDefaultGoodMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type AppDefaultGoodMutationRuleFunc func(context.Context, *ent.AppDefaultGoodMutation) error

// EvalMutation calls f(ctx, m).
func (f AppDefaultGoodMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.AppDefaultGoodMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.AppDefaultGoodMutation", m)
}

// The AppGoodQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type AppGoodQueryRuleFunc func(context.Context, *ent.AppGoodQuery) error

// EvalQuery return f(ctx, q).
func (f AppGoodQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.AppGoodQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.AppGoodQuery", q)
}

// The AppGoodMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type AppGoodMutationRuleFunc func(context.Context, *ent.AppGoodMutation) error

// EvalMutation calls f(ctx, m).
func (f AppGoodMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.AppGoodMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.AppGoodMutation", m)
}

// The AppStockQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type AppStockQueryRuleFunc func(context.Context, *ent.AppStockQuery) error

// EvalQuery return f(ctx, q).
func (f AppStockQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.AppStockQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.AppStockQuery", q)
}

// The AppStockMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type AppStockMutationRuleFunc func(context.Context, *ent.AppStockMutation) error

// EvalMutation calls f(ctx, m).
func (f AppStockMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.AppStockMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.AppStockMutation", m)
}

// The CommentQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type CommentQueryRuleFunc func(context.Context, *ent.CommentQuery) error

// EvalQuery return f(ctx, q).
func (f CommentQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.CommentQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.CommentQuery", q)
}

// The CommentMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type CommentMutationRuleFunc func(context.Context, *ent.CommentMutation) error

// EvalMutation calls f(ctx, m).
func (f CommentMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.CommentMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.CommentMutation", m)
}

// The DeviceInfoQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type DeviceInfoQueryRuleFunc func(context.Context, *ent.DeviceInfoQuery) error

// EvalQuery return f(ctx, q).
func (f DeviceInfoQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.DeviceInfoQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.DeviceInfoQuery", q)
}

// The DeviceInfoMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type DeviceInfoMutationRuleFunc func(context.Context, *ent.DeviceInfoMutation) error

// EvalMutation calls f(ctx, m).
func (f DeviceInfoMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.DeviceInfoMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.DeviceInfoMutation", m)
}

// The ExtraInfoQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type ExtraInfoQueryRuleFunc func(context.Context, *ent.ExtraInfoQuery) error

// EvalQuery return f(ctx, q).
func (f ExtraInfoQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.ExtraInfoQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.ExtraInfoQuery", q)
}

// The ExtraInfoMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type ExtraInfoMutationRuleFunc func(context.Context, *ent.ExtraInfoMutation) error

// EvalMutation calls f(ctx, m).
func (f ExtraInfoMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.ExtraInfoMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.ExtraInfoMutation", m)
}

// The GoodQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type GoodQueryRuleFunc func(context.Context, *ent.GoodQuery) error

// EvalQuery return f(ctx, q).
func (f GoodQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.GoodQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.GoodQuery", q)
}

// The GoodMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type GoodMutationRuleFunc func(context.Context, *ent.GoodMutation) error

// EvalMutation calls f(ctx, m).
func (f GoodMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.GoodMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.GoodMutation", m)
}

// The GoodRewardQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type GoodRewardQueryRuleFunc func(context.Context, *ent.GoodRewardQuery) error

// EvalQuery return f(ctx, q).
func (f GoodRewardQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.GoodRewardQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.GoodRewardQuery", q)
}

// The GoodRewardMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type GoodRewardMutationRuleFunc func(context.Context, *ent.GoodRewardMutation) error

// EvalMutation calls f(ctx, m).
func (f GoodRewardMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.GoodRewardMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.GoodRewardMutation", m)
}

// The GoodRewardHistoryQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type GoodRewardHistoryQueryRuleFunc func(context.Context, *ent.GoodRewardHistoryQuery) error

// EvalQuery return f(ctx, q).
func (f GoodRewardHistoryQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.GoodRewardHistoryQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.GoodRewardHistoryQuery", q)
}

// The GoodRewardHistoryMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type GoodRewardHistoryMutationRuleFunc func(context.Context, *ent.GoodRewardHistoryMutation) error

// EvalMutation calls f(ctx, m).
func (f GoodRewardHistoryMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.GoodRewardHistoryMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.GoodRewardHistoryMutation", m)
}

// The LikeQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type LikeQueryRuleFunc func(context.Context, *ent.LikeQuery) error

// EvalQuery return f(ctx, q).
func (f LikeQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.LikeQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.LikeQuery", q)
}

// The LikeMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type LikeMutationRuleFunc func(context.Context, *ent.LikeMutation) error

// EvalMutation calls f(ctx, m).
func (f LikeMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.LikeMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.LikeMutation", m)
}

// The PromotionQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type PromotionQueryRuleFunc func(context.Context, *ent.PromotionQuery) error

// EvalQuery return f(ctx, q).
func (f PromotionQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.PromotionQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.PromotionQuery", q)
}

// The PromotionMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type PromotionMutationRuleFunc func(context.Context, *ent.PromotionMutation) error

// EvalMutation calls f(ctx, m).
func (f PromotionMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.PromotionMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.PromotionMutation", m)
}

// The RecommendQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type RecommendQueryRuleFunc func(context.Context, *ent.RecommendQuery) error

// EvalQuery return f(ctx, q).
func (f RecommendQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.RecommendQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.RecommendQuery", q)
}

// The RecommendMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type RecommendMutationRuleFunc func(context.Context, *ent.RecommendMutation) error

// EvalMutation calls f(ctx, m).
func (f RecommendMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.RecommendMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.RecommendMutation", m)
}

// The RequiredGoodQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type RequiredGoodQueryRuleFunc func(context.Context, *ent.RequiredGoodQuery) error

// EvalQuery return f(ctx, q).
func (f RequiredGoodQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.RequiredGoodQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.RequiredGoodQuery", q)
}

// The RequiredGoodMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type RequiredGoodMutationRuleFunc func(context.Context, *ent.RequiredGoodMutation) error

// EvalMutation calls f(ctx, m).
func (f RequiredGoodMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.RequiredGoodMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.RequiredGoodMutation", m)
}

// The ScoreQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type ScoreQueryRuleFunc func(context.Context, *ent.ScoreQuery) error

// EvalQuery return f(ctx, q).
func (f ScoreQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.ScoreQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.ScoreQuery", q)
}

// The ScoreMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type ScoreMutationRuleFunc func(context.Context, *ent.ScoreMutation) error

// EvalMutation calls f(ctx, m).
func (f ScoreMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.ScoreMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.ScoreMutation", m)
}

// The StockQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type StockQueryRuleFunc func(context.Context, *ent.StockQuery) error

// EvalQuery return f(ctx, q).
func (f StockQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.StockQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.StockQuery", q)
}

// The StockMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type StockMutationRuleFunc func(context.Context, *ent.StockMutation) error

// EvalMutation calls f(ctx, m).
func (f StockMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.StockMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.StockMutation", m)
}

// The VendorBrandQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type VendorBrandQueryRuleFunc func(context.Context, *ent.VendorBrandQuery) error

// EvalQuery return f(ctx, q).
func (f VendorBrandQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.VendorBrandQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.VendorBrandQuery", q)
}

// The VendorBrandMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type VendorBrandMutationRuleFunc func(context.Context, *ent.VendorBrandMutation) error

// EvalMutation calls f(ctx, m).
func (f VendorBrandMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.VendorBrandMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.VendorBrandMutation", m)
}

// The VendorLocationQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type VendorLocationQueryRuleFunc func(context.Context, *ent.VendorLocationQuery) error

// EvalQuery return f(ctx, q).
func (f VendorLocationQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.VendorLocationQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.VendorLocationQuery", q)
}

// The VendorLocationMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type VendorLocationMutationRuleFunc func(context.Context, *ent.VendorLocationMutation) error

// EvalMutation calls f(ctx, m).
func (f VendorLocationMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.VendorLocationMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.VendorLocationMutation", m)
}

type (
	// Filter is the interface that wraps the Where function
	// for filtering nodes in queries and mutations.
	Filter interface {
		// Where applies a filter on the executed query/mutation.
		Where(entql.P)
	}

	// The FilterFunc type is an adapter that allows the use of ordinary
	// functions as filters for query and mutation types.
	FilterFunc func(context.Context, Filter) error
)

// EvalQuery calls f(ctx, q) if the query implements the Filter interface, otherwise it is denied.
func (f FilterFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	fr, err := queryFilter(q)
	if err != nil {
		return err
	}
	return f(ctx, fr)
}

// EvalMutation calls f(ctx, q) if the mutation implements the Filter interface, otherwise it is denied.
func (f FilterFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	fr, err := mutationFilter(m)
	if err != nil {
		return err
	}
	return f(ctx, fr)
}

var _ QueryMutationRule = FilterFunc(nil)

func queryFilter(q ent.Query) (Filter, error) {
	switch q := q.(type) {
	case *ent.AppDefaultGoodQuery:
		return q.Filter(), nil
	case *ent.AppGoodQuery:
		return q.Filter(), nil
	case *ent.AppStockQuery:
		return q.Filter(), nil
	case *ent.CommentQuery:
		return q.Filter(), nil
	case *ent.DeviceInfoQuery:
		return q.Filter(), nil
	case *ent.ExtraInfoQuery:
		return q.Filter(), nil
	case *ent.GoodQuery:
		return q.Filter(), nil
	case *ent.GoodRewardQuery:
		return q.Filter(), nil
	case *ent.GoodRewardHistoryQuery:
		return q.Filter(), nil
	case *ent.LikeQuery:
		return q.Filter(), nil
	case *ent.PromotionQuery:
		return q.Filter(), nil
	case *ent.RecommendQuery:
		return q.Filter(), nil
	case *ent.RequiredGoodQuery:
		return q.Filter(), nil
	case *ent.ScoreQuery:
		return q.Filter(), nil
	case *ent.StockQuery:
		return q.Filter(), nil
	case *ent.VendorBrandQuery:
		return q.Filter(), nil
	case *ent.VendorLocationQuery:
		return q.Filter(), nil
	default:
		return nil, Denyf("ent/privacy: unexpected query type %T for query filter", q)
	}
}

func mutationFilter(m ent.Mutation) (Filter, error) {
	switch m := m.(type) {
	case *ent.AppDefaultGoodMutation:
		return m.Filter(), nil
	case *ent.AppGoodMutation:
		return m.Filter(), nil
	case *ent.AppStockMutation:
		return m.Filter(), nil
	case *ent.CommentMutation:
		return m.Filter(), nil
	case *ent.DeviceInfoMutation:
		return m.Filter(), nil
	case *ent.ExtraInfoMutation:
		return m.Filter(), nil
	case *ent.GoodMutation:
		return m.Filter(), nil
	case *ent.GoodRewardMutation:
		return m.Filter(), nil
	case *ent.GoodRewardHistoryMutation:
		return m.Filter(), nil
	case *ent.LikeMutation:
		return m.Filter(), nil
	case *ent.PromotionMutation:
		return m.Filter(), nil
	case *ent.RecommendMutation:
		return m.Filter(), nil
	case *ent.RequiredGoodMutation:
		return m.Filter(), nil
	case *ent.ScoreMutation:
		return m.Filter(), nil
	case *ent.StockMutation:
		return m.Filter(), nil
	case *ent.VendorBrandMutation:
		return m.Filter(), nil
	case *ent.VendorLocationMutation:
		return m.Filter(), nil
	default:
		return nil, Denyf("ent/privacy: unexpected mutation type %T for mutation filter", m)
	}
}