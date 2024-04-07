// Code generated by ent, DO NOT EDIT.

package comment

import (
	"entgo.io/ent"
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the comment type in the database.
	Label = "comment"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldDeletedAt holds the string denoting the deleted_at field in the database.
	FieldDeletedAt = "deleted_at"
	// FieldEntID holds the string denoting the ent_id field in the database.
	FieldEntID = "ent_id"
	// FieldUserID holds the string denoting the user_id field in the database.
	FieldUserID = "user_id"
	// FieldAppGoodID holds the string denoting the app_good_id field in the database.
	FieldAppGoodID = "app_good_id"
	// FieldOrderID holds the string denoting the order_id field in the database.
	FieldOrderID = "order_id"
	// FieldContent holds the string denoting the content field in the database.
	FieldContent = "content"
	// FieldReplyToID holds the string denoting the reply_to_id field in the database.
	FieldReplyToID = "reply_to_id"
	// FieldAnonymous holds the string denoting the anonymous field in the database.
	FieldAnonymous = "anonymous"
	// FieldTrialUser holds the string denoting the trial_user field in the database.
	FieldTrialUser = "trial_user"
	// FieldPurchasedUser holds the string denoting the purchased_user field in the database.
	FieldPurchasedUser = "purchased_user"
	// Table holds the table name of the comment in the database.
	Table = "comments"
)

// Columns holds all SQL columns for comment fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
	FieldEntID,
	FieldUserID,
	FieldAppGoodID,
	FieldOrderID,
	FieldContent,
	FieldReplyToID,
	FieldAnonymous,
	FieldTrialUser,
	FieldPurchasedUser,
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
	// DefaultEntID holds the default value on creation for the "ent_id" field.
	DefaultEntID func() uuid.UUID
	// DefaultUserID holds the default value on creation for the "user_id" field.
	DefaultUserID func() uuid.UUID
	// DefaultAppGoodID holds the default value on creation for the "app_good_id" field.
	DefaultAppGoodID func() uuid.UUID
	// DefaultOrderID holds the default value on creation for the "order_id" field.
	DefaultOrderID func() uuid.UUID
	// DefaultContent holds the default value on creation for the "content" field.
	DefaultContent string
	// DefaultReplyToID holds the default value on creation for the "reply_to_id" field.
	DefaultReplyToID func() uuid.UUID
	// DefaultAnonymous holds the default value on creation for the "anonymous" field.
	DefaultAnonymous bool
	// DefaultTrialUser holds the default value on creation for the "trial_user" field.
	DefaultTrialUser bool
	// DefaultPurchasedUser holds the default value on creation for the "purchased_user" field.
	DefaultPurchasedUser bool
)
