// Code generated by ent, DO NOT EDIT.

package hook

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/good-middleware/pkg/db/ent"
)

// The AppDefaultGoodFunc type is an adapter to allow the use of ordinary
// function as AppDefaultGood mutator.
type AppDefaultGoodFunc func(context.Context, *ent.AppDefaultGoodMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f AppDefaultGoodFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.AppDefaultGoodMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.AppDefaultGoodMutation", m)
	}
	return f(ctx, mv)
}

// The AppGoodFunc type is an adapter to allow the use of ordinary
// function as AppGood mutator.
type AppGoodFunc func(context.Context, *ent.AppGoodMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f AppGoodFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.AppGoodMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.AppGoodMutation", m)
	}
	return f(ctx, mv)
}

// The AppStockFunc type is an adapter to allow the use of ordinary
// function as AppStock mutator.
type AppStockFunc func(context.Context, *ent.AppStockMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f AppStockFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.AppStockMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.AppStockMutation", m)
	}
	return f(ctx, mv)
}

// The CommentFunc type is an adapter to allow the use of ordinary
// function as Comment mutator.
type CommentFunc func(context.Context, *ent.CommentMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f CommentFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.CommentMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.CommentMutation", m)
	}
	return f(ctx, mv)
}

// The DeviceInfoFunc type is an adapter to allow the use of ordinary
// function as DeviceInfo mutator.
type DeviceInfoFunc func(context.Context, *ent.DeviceInfoMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f DeviceInfoFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.DeviceInfoMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.DeviceInfoMutation", m)
	}
	return f(ctx, mv)
}

// The ExtraInfoFunc type is an adapter to allow the use of ordinary
// function as ExtraInfo mutator.
type ExtraInfoFunc func(context.Context, *ent.ExtraInfoMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f ExtraInfoFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.ExtraInfoMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.ExtraInfoMutation", m)
	}
	return f(ctx, mv)
}

// The GoodFunc type is an adapter to allow the use of ordinary
// function as Good mutator.
type GoodFunc func(context.Context, *ent.GoodMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f GoodFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.GoodMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.GoodMutation", m)
	}
	return f(ctx, mv)
}

// The GoodRewardFunc type is an adapter to allow the use of ordinary
// function as GoodReward mutator.
type GoodRewardFunc func(context.Context, *ent.GoodRewardMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f GoodRewardFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.GoodRewardMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.GoodRewardMutation", m)
	}
	return f(ctx, mv)
}

// The GoodRewardHistoryFunc type is an adapter to allow the use of ordinary
// function as GoodRewardHistory mutator.
type GoodRewardHistoryFunc func(context.Context, *ent.GoodRewardHistoryMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f GoodRewardHistoryFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.GoodRewardHistoryMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.GoodRewardHistoryMutation", m)
	}
	return f(ctx, mv)
}

// The LikeFunc type is an adapter to allow the use of ordinary
// function as Like mutator.
type LikeFunc func(context.Context, *ent.LikeMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f LikeFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.LikeMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.LikeMutation", m)
	}
	return f(ctx, mv)
}

// The PromotionFunc type is an adapter to allow the use of ordinary
// function as Promotion mutator.
type PromotionFunc func(context.Context, *ent.PromotionMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f PromotionFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.PromotionMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.PromotionMutation", m)
	}
	return f(ctx, mv)
}

// The RecommendFunc type is an adapter to allow the use of ordinary
// function as Recommend mutator.
type RecommendFunc func(context.Context, *ent.RecommendMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f RecommendFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.RecommendMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.RecommendMutation", m)
	}
	return f(ctx, mv)
}

// The RequiredGoodFunc type is an adapter to allow the use of ordinary
// function as RequiredGood mutator.
type RequiredGoodFunc func(context.Context, *ent.RequiredGoodMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f RequiredGoodFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.RequiredGoodMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.RequiredGoodMutation", m)
	}
	return f(ctx, mv)
}

// The ScoreFunc type is an adapter to allow the use of ordinary
// function as Score mutator.
type ScoreFunc func(context.Context, *ent.ScoreMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f ScoreFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.ScoreMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.ScoreMutation", m)
	}
	return f(ctx, mv)
}

// The StockFunc type is an adapter to allow the use of ordinary
// function as Stock mutator.
type StockFunc func(context.Context, *ent.StockMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f StockFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.StockMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.StockMutation", m)
	}
	return f(ctx, mv)
}

// The VendorLocationFunc type is an adapter to allow the use of ordinary
// function as VendorLocation mutator.
type VendorLocationFunc func(context.Context, *ent.VendorLocationMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f VendorLocationFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.VendorLocationMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.VendorLocationMutation", m)
	}
	return f(ctx, mv)
}

// Condition is a hook condition function.
type Condition func(context.Context, ent.Mutation) bool

// And groups conditions with the AND operator.
func And(first, second Condition, rest ...Condition) Condition {
	return func(ctx context.Context, m ent.Mutation) bool {
		if !first(ctx, m) || !second(ctx, m) {
			return false
		}
		for _, cond := range rest {
			if !cond(ctx, m) {
				return false
			}
		}
		return true
	}
}

// Or groups conditions with the OR operator.
func Or(first, second Condition, rest ...Condition) Condition {
	return func(ctx context.Context, m ent.Mutation) bool {
		if first(ctx, m) || second(ctx, m) {
			return true
		}
		for _, cond := range rest {
			if cond(ctx, m) {
				return true
			}
		}
		return false
	}
}

// Not negates a given condition.
func Not(cond Condition) Condition {
	return func(ctx context.Context, m ent.Mutation) bool {
		return !cond(ctx, m)
	}
}

// HasOp is a condition testing mutation operation.
func HasOp(op ent.Op) Condition {
	return func(_ context.Context, m ent.Mutation) bool {
		return m.Op().Is(op)
	}
}

// HasAddedFields is a condition validating `.AddedField` on fields.
func HasAddedFields(field string, fields ...string) Condition {
	return func(_ context.Context, m ent.Mutation) bool {
		if _, exists := m.AddedField(field); !exists {
			return false
		}
		for _, field := range fields {
			if _, exists := m.AddedField(field); !exists {
				return false
			}
		}
		return true
	}
}

// HasClearedFields is a condition validating `.FieldCleared` on fields.
func HasClearedFields(field string, fields ...string) Condition {
	return func(_ context.Context, m ent.Mutation) bool {
		if exists := m.FieldCleared(field); !exists {
			return false
		}
		for _, field := range fields {
			if exists := m.FieldCleared(field); !exists {
				return false
			}
		}
		return true
	}
}

// HasFields is a condition validating `.Field` on fields.
func HasFields(field string, fields ...string) Condition {
	return func(_ context.Context, m ent.Mutation) bool {
		if _, exists := m.Field(field); !exists {
			return false
		}
		for _, field := range fields {
			if _, exists := m.Field(field); !exists {
				return false
			}
		}
		return true
	}
}

// If executes the given hook under condition.
//
//	hook.If(ComputeAverage, And(HasFields(...), HasAddedFields(...)))
//
func If(hk ent.Hook, cond Condition) ent.Hook {
	return func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			if cond(ctx, m) {
				return hk(next).Mutate(ctx, m)
			}
			return next.Mutate(ctx, m)
		})
	}
}

// On executes the given hook only for the given operation.
//
//	hook.On(Log, ent.Delete|ent.Create)
//
func On(hk ent.Hook, op ent.Op) ent.Hook {
	return If(hk, HasOp(op))
}

// Unless skips the given hook only for the given operation.
//
//	hook.Unless(Log, ent.Update|ent.UpdateOne)
//
func Unless(hk ent.Hook, op ent.Op) ent.Hook {
	return If(hk, Not(HasOp(op)))
}

// FixedError is a hook returning a fixed error.
func FixedError(err error) ent.Hook {
	return func(ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(context.Context, ent.Mutation) (ent.Value, error) {
			return nil, err
		})
	}
}

// Reject returns a hook that rejects all operations that match op.
//
//	func (T) Hooks() []ent.Hook {
//		return []ent.Hook{
//			Reject(ent.Delete|ent.Update),
//		}
//	}
//
func Reject(op ent.Op) ent.Hook {
	hk := FixedError(fmt.Errorf("%s operation is not allowed", op))
	return On(hk, op)
}

// Chain acts as a list of hooks and is effectively immutable.
// Once created, it will always hold the same set of hooks in the same order.
type Chain struct {
	hooks []ent.Hook
}

// NewChain creates a new chain of hooks.
func NewChain(hooks ...ent.Hook) Chain {
	return Chain{append([]ent.Hook(nil), hooks...)}
}

// Hook chains the list of hooks and returns the final hook.
func (c Chain) Hook() ent.Hook {
	return func(mutator ent.Mutator) ent.Mutator {
		for i := len(c.hooks) - 1; i >= 0; i-- {
			mutator = c.hooks[i](mutator)
		}
		return mutator
	}
}

// Append extends a chain, adding the specified hook
// as the last ones in the mutation flow.
func (c Chain) Append(hooks ...ent.Hook) Chain {
	newHooks := make([]ent.Hook, 0, len(c.hooks)+len(hooks))
	newHooks = append(newHooks, c.hooks...)
	newHooks = append(newHooks, hooks...)
	return Chain{newHooks}
}

// Extend extends a chain, adding the specified chain
// as the last ones in the mutation flow.
func (c Chain) Extend(chain Chain) Chain {
	return c.Append(chain.hooks...)
}
