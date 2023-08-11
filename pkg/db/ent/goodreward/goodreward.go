// Code generated by ent, DO NOT EDIT.

package goodreward

import (
	"entgo.io/ent"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

const (
	// Label holds the string label denoting the goodreward type in the database.
	Label = "good_reward"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldDeletedAt holds the string denoting the deleted_at field in the database.
	FieldDeletedAt = "deleted_at"
	// FieldGoodID holds the string denoting the good_id field in the database.
	FieldGoodID = "good_id"
	// FieldRewardState holds the string denoting the reward_state field in the database.
	FieldRewardState = "reward_state"
	// FieldLastRewardAt holds the string denoting the last_reward_at field in the database.
	FieldLastRewardAt = "last_reward_at"
	// FieldRewardTid holds the string denoting the reward_tid field in the database.
	FieldRewardTid = "reward_tid"
	// FieldNextRewardStartAmount holds the string denoting the next_reward_start_amount field in the database.
	FieldNextRewardStartAmount = "next_reward_start_amount"
	// FieldLastRewardAmount holds the string denoting the last_reward_amount field in the database.
	FieldLastRewardAmount = "last_reward_amount"
	// Table holds the table name of the goodreward in the database.
	Table = "good_rewards"
)

// Columns holds all SQL columns for goodreward fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
	FieldGoodID,
	FieldRewardState,
	FieldLastRewardAt,
	FieldRewardTid,
	FieldNextRewardStartAmount,
	FieldLastRewardAmount,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

// Note that the variables below are initialized by the runtime
// package on the initialization of the application. Therefore,
// it should be imported in the main as follows:
//
//	import _ "github.com/NpoolPlatform/good-middleware/pkg/db/ent/runtime"
//
var (
	Hooks  [1]ent.Hook
	Policy ent.Policy
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() uint32
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() uint32
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() uint32
	// DefaultDeletedAt holds the default value on creation for the "deleted_at" field.
	DefaultDeletedAt func() uint32
	// DefaultRewardState holds the default value on creation for the "reward_state" field.
	DefaultRewardState string
	// DefaultLastRewardAt holds the default value on creation for the "last_reward_at" field.
	DefaultLastRewardAt uint32
	// DefaultRewardTid holds the default value on creation for the "reward_tid" field.
	DefaultRewardTid func() uuid.UUID
	// DefaultNextRewardStartAmount holds the default value on creation for the "next_reward_start_amount" field.
	DefaultNextRewardStartAmount decimal.Decimal
	// DefaultLastRewardAmount holds the default value on creation for the "last_reward_amount" field.
	DefaultLastRewardAmount decimal.Decimal
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)
