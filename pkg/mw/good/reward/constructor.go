package goodreward

import (
	"fmt"
	"time"

	wlog "github.com/NpoolPlatform/go-service-framework/pkg/wlog"
	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
)

func (h *Handler) ConstructUpdateSQL() (string, error) {
	if h.GoodID == nil && h.ID == nil && h.EntID == nil {
		return "", wlog.Errorf("invalid id")
	}

	set := "set "
	now := uint32(time.Now().Unix())

	_sql := "update good_rewards "
	if h.RewardState != nil {
		_sql += fmt.Sprintf("%vreward_state = '%v', ", set, h.RewardState.String())
		set = ""
	}
	if h.LastRewardAt != nil {
		_sql += fmt.Sprintf("%vlast_reward_at = %v, ", set, *h.LastRewardAt)
		set = ""
	}
	if set != "" {
		return "", wlog.WrapError(cruder.ErrUpdateNothing)
	}
	_sql += fmt.Sprintf("updated_at = %v ", now)
	_sql += "where deleted_at = 0 "
	if h.ID != nil {
		_sql += fmt.Sprintf("and id = %v ", *h.ID)
	}
	if h.GoodID != nil {
		_sql += fmt.Sprintf("and good_id = '%v' ", *h.GoodID)
	}
	if h.EntID != nil {
		_sql += fmt.Sprintf("and ent_id = '%v' ", *h.EntID)
	}
	if h.LastRewardAt != nil {
		_sql += "and not exists ("
		_sql += "select 1 from good_reward_histories "
		_sql += fmt.Sprintf(
			"where good_id = '%v' and deleted_at = 0 and reward_date = %v",
			*h.GoodID,
			*h.LastRewardAt,
		)
		_sql += " limit 1)"
	}
	return _sql, nil
}
