package goodcoinreward

import (
	"fmt"
	"time"

	wlog "github.com/NpoolPlatform/go-service-framework/pkg/wlog"
	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
)

func (h *Handler) ConstructUpdateSQL(addTotal bool) (string, error) {
	if h.ID == nil && h.EntID == nil && (h.GoodID == nil || h.CoinTypeID == nil) {
		return "", wlog.Errorf("invalid id")
	}

	set := "set "
	now := uint32(time.Now().Unix())

	_sql := "update good_coin_rewards "
	if h.RewardTID != nil {
		_sql += fmt.Sprintf("%vreward_tid = '%v', ", set, *h.RewardTID)
		set = ""
	}
	if h.NextRewardStartAmount != nil {
		_sql += fmt.Sprintf("%vnext_reward_start_amount = '%v', ", set, *h.NextRewardStartAmount)
		set = ""
	}
	if h.LastRewardAmount != nil {
		_sql += fmt.Sprintf("%vlast_reward_amount = '%v', ", set, *h.LastRewardAmount)
		if addTotal {
			_sql += fmt.Sprintf("%vtotal_reward_amount = total_reward_amount + %v, ", set, *h.LastRewardAmount)
		}
		set = ""
	}
	if h.LastUnitRewardAmount != nil {
		_sql += fmt.Sprintf("%vlast_unit_reward_amount = '%v', ", set, *h.LastUnitRewardAmount)
		set = ""
	}
	if set != "" {
		return "", wlog.WrapError(cruder.ErrUpdateNothing)
	}
	_sql += fmt.Sprintf("updated_at = %v ", now)
	whereAnd := "where"
	if h.ID != nil {
		_sql += fmt.Sprintf("where id = %v ", *h.ID)
		whereAnd = "and"
	}
	if h.EntID != nil {
		_sql += fmt.Sprintf("%v ent_id = '%v' ", whereAnd, *h.EntID)
		whereAnd = "and"
	}
	if h.GoodID != nil && h.CoinTypeID != nil {
		_sql += fmt.Sprintf(
			"%v good_id = '%v' and coin_type_id = '%v'",
			whereAnd,
			*h.GoodID,
			*h.CoinTypeID,
		)
	}
	return _sql, nil
}
