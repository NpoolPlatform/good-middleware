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

// The AppDelegatedStakingFunc type is an adapter to allow the use of ordinary
// function as AppDelegatedStaking mutator.
type AppDelegatedStakingFunc func(context.Context, *ent.AppDelegatedStakingMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f AppDelegatedStakingFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.AppDelegatedStakingMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.AppDelegatedStakingMutation", m)
	}
	return f(ctx, mv)
}

// The AppFeeFunc type is an adapter to allow the use of ordinary
// function as AppFee mutator.
type AppFeeFunc func(context.Context, *ent.AppFeeMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f AppFeeFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.AppFeeMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.AppFeeMutation", m)
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

// The AppGoodBaseFunc type is an adapter to allow the use of ordinary
// function as AppGoodBase mutator.
type AppGoodBaseFunc func(context.Context, *ent.AppGoodBaseMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f AppGoodBaseFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.AppGoodBaseMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.AppGoodBaseMutation", m)
	}
	return f(ctx, mv)
}

// The AppGoodDescriptionFunc type is an adapter to allow the use of ordinary
// function as AppGoodDescription mutator.
type AppGoodDescriptionFunc func(context.Context, *ent.AppGoodDescriptionMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f AppGoodDescriptionFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.AppGoodDescriptionMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.AppGoodDescriptionMutation", m)
	}
	return f(ctx, mv)
}

// The AppGoodDisplayColorFunc type is an adapter to allow the use of ordinary
// function as AppGoodDisplayColor mutator.
type AppGoodDisplayColorFunc func(context.Context, *ent.AppGoodDisplayColorMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f AppGoodDisplayColorFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.AppGoodDisplayColorMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.AppGoodDisplayColorMutation", m)
	}
	return f(ctx, mv)
}

// The AppGoodDisplayNameFunc type is an adapter to allow the use of ordinary
// function as AppGoodDisplayName mutator.
type AppGoodDisplayNameFunc func(context.Context, *ent.AppGoodDisplayNameMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f AppGoodDisplayNameFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.AppGoodDisplayNameMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.AppGoodDisplayNameMutation", m)
	}
	return f(ctx, mv)
}

// The AppGoodLabelFunc type is an adapter to allow the use of ordinary
// function as AppGoodLabel mutator.
type AppGoodLabelFunc func(context.Context, *ent.AppGoodLabelMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f AppGoodLabelFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.AppGoodLabelMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.AppGoodLabelMutation", m)
	}
	return f(ctx, mv)
}

// The AppGoodPosterFunc type is an adapter to allow the use of ordinary
// function as AppGoodPoster mutator.
type AppGoodPosterFunc func(context.Context, *ent.AppGoodPosterMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f AppGoodPosterFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.AppGoodPosterMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.AppGoodPosterMutation", m)
	}
	return f(ctx, mv)
}

// The AppLegacyPowerRentalFunc type is an adapter to allow the use of ordinary
// function as AppLegacyPowerRental mutator.
type AppLegacyPowerRentalFunc func(context.Context, *ent.AppLegacyPowerRentalMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f AppLegacyPowerRentalFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.AppLegacyPowerRentalMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.AppLegacyPowerRentalMutation", m)
	}
	return f(ctx, mv)
}

// The AppMiningGoodStockFunc type is an adapter to allow the use of ordinary
// function as AppMiningGoodStock mutator.
type AppMiningGoodStockFunc func(context.Context, *ent.AppMiningGoodStockMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f AppMiningGoodStockFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.AppMiningGoodStockMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.AppMiningGoodStockMutation", m)
	}
	return f(ctx, mv)
}

// The AppPowerRentalFunc type is an adapter to allow the use of ordinary
// function as AppPowerRental mutator.
type AppPowerRentalFunc func(context.Context, *ent.AppPowerRentalMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f AppPowerRentalFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.AppPowerRentalMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.AppPowerRentalMutation", m)
	}
	return f(ctx, mv)
}

// The AppSimulatePowerRentalFunc type is an adapter to allow the use of ordinary
// function as AppSimulatePowerRental mutator.
type AppSimulatePowerRentalFunc func(context.Context, *ent.AppSimulatePowerRentalMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f AppSimulatePowerRentalFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.AppSimulatePowerRentalMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.AppSimulatePowerRentalMutation", m)
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

// The AppStockLockFunc type is an adapter to allow the use of ordinary
// function as AppStockLock mutator.
type AppStockLockFunc func(context.Context, *ent.AppStockLockMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f AppStockLockFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.AppStockLockMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.AppStockLockMutation", m)
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

// The DelegatedStakingFunc type is an adapter to allow the use of ordinary
// function as DelegatedStaking mutator.
type DelegatedStakingFunc func(context.Context, *ent.DelegatedStakingMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f DelegatedStakingFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.DelegatedStakingMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.DelegatedStakingMutation", m)
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

// The DeviceManufacturerFunc type is an adapter to allow the use of ordinary
// function as DeviceManufacturer mutator.
type DeviceManufacturerFunc func(context.Context, *ent.DeviceManufacturerMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f DeviceManufacturerFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.DeviceManufacturerMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.DeviceManufacturerMutation", m)
	}
	return f(ctx, mv)
}

// The DevicePosterFunc type is an adapter to allow the use of ordinary
// function as DevicePoster mutator.
type DevicePosterFunc func(context.Context, *ent.DevicePosterMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f DevicePosterFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.DevicePosterMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.DevicePosterMutation", m)
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

// The FbmCrowdFundingFunc type is an adapter to allow the use of ordinary
// function as FbmCrowdFunding mutator.
type FbmCrowdFundingFunc func(context.Context, *ent.FbmCrowdFundingMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f FbmCrowdFundingFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.FbmCrowdFundingMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.FbmCrowdFundingMutation", m)
	}
	return f(ctx, mv)
}

// The FeeFunc type is an adapter to allow the use of ordinary
// function as Fee mutator.
type FeeFunc func(context.Context, *ent.FeeMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f FeeFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.FeeMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.FeeMutation", m)
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

// The GoodBaseFunc type is an adapter to allow the use of ordinary
// function as GoodBase mutator.
type GoodBaseFunc func(context.Context, *ent.GoodBaseMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f GoodBaseFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.GoodBaseMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.GoodBaseMutation", m)
	}
	return f(ctx, mv)
}

// The GoodCoinFunc type is an adapter to allow the use of ordinary
// function as GoodCoin mutator.
type GoodCoinFunc func(context.Context, *ent.GoodCoinMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f GoodCoinFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.GoodCoinMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.GoodCoinMutation", m)
	}
	return f(ctx, mv)
}

// The GoodCoinRewardFunc type is an adapter to allow the use of ordinary
// function as GoodCoinReward mutator.
type GoodCoinRewardFunc func(context.Context, *ent.GoodCoinRewardMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f GoodCoinRewardFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.GoodCoinRewardMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.GoodCoinRewardMutation", m)
	}
	return f(ctx, mv)
}

// The GoodMalfunctionFunc type is an adapter to allow the use of ordinary
// function as GoodMalfunction mutator.
type GoodMalfunctionFunc func(context.Context, *ent.GoodMalfunctionMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f GoodMalfunctionFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.GoodMalfunctionMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.GoodMalfunctionMutation", m)
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

// The MiningGoodStockFunc type is an adapter to allow the use of ordinary
// function as MiningGoodStock mutator.
type MiningGoodStockFunc func(context.Context, *ent.MiningGoodStockMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f MiningGoodStockFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.MiningGoodStockMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.MiningGoodStockMutation", m)
	}
	return f(ctx, mv)
}

// The PowerRentalFunc type is an adapter to allow the use of ordinary
// function as PowerRental mutator.
type PowerRentalFunc func(context.Context, *ent.PowerRentalMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f PowerRentalFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.PowerRentalMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.PowerRentalMutation", m)
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

// The RequiredAppGoodFunc type is an adapter to allow the use of ordinary
// function as RequiredAppGood mutator.
type RequiredAppGoodFunc func(context.Context, *ent.RequiredAppGoodMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f RequiredAppGoodFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.RequiredAppGoodMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.RequiredAppGoodMutation", m)
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

// The TopMostFunc type is an adapter to allow the use of ordinary
// function as TopMost mutator.
type TopMostFunc func(context.Context, *ent.TopMostMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f TopMostFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.TopMostMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.TopMostMutation", m)
	}
	return f(ctx, mv)
}

// The TopMostConstraintFunc type is an adapter to allow the use of ordinary
// function as TopMostConstraint mutator.
type TopMostConstraintFunc func(context.Context, *ent.TopMostConstraintMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f TopMostConstraintFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.TopMostConstraintMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.TopMostConstraintMutation", m)
	}
	return f(ctx, mv)
}

// The TopMostGoodFunc type is an adapter to allow the use of ordinary
// function as TopMostGood mutator.
type TopMostGoodFunc func(context.Context, *ent.TopMostGoodMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f TopMostGoodFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.TopMostGoodMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.TopMostGoodMutation", m)
	}
	return f(ctx, mv)
}

// The TopMostGoodConstraintFunc type is an adapter to allow the use of ordinary
// function as TopMostGoodConstraint mutator.
type TopMostGoodConstraintFunc func(context.Context, *ent.TopMostGoodConstraintMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f TopMostGoodConstraintFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.TopMostGoodConstraintMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.TopMostGoodConstraintMutation", m)
	}
	return f(ctx, mv)
}

// The TopMostGoodPosterFunc type is an adapter to allow the use of ordinary
// function as TopMostGoodPoster mutator.
type TopMostGoodPosterFunc func(context.Context, *ent.TopMostGoodPosterMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f TopMostGoodPosterFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.TopMostGoodPosterMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.TopMostGoodPosterMutation", m)
	}
	return f(ctx, mv)
}

// The TopMostPosterFunc type is an adapter to allow the use of ordinary
// function as TopMostPoster mutator.
type TopMostPosterFunc func(context.Context, *ent.TopMostPosterMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f TopMostPosterFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.TopMostPosterMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.TopMostPosterMutation", m)
	}
	return f(ctx, mv)
}

// The VendorBrandFunc type is an adapter to allow the use of ordinary
// function as VendorBrand mutator.
type VendorBrandFunc func(context.Context, *ent.VendorBrandMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f VendorBrandFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.VendorBrandMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.VendorBrandMutation", m)
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
