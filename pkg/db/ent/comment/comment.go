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
	// FieldAppID holds the string denoting the app_id field in the database.
	FieldAppID = "app_id"
	// FieldUserID holds the string denoting the user_id field in the database.
	FieldUserID = "user_id"
	// FieldGoodID holds the string denoting the good_id field in the database.
	FieldGoodID = "good_id"
	// FieldAppGoodID holds the string denoting the app_good_id field in the database.
	FieldAppGoodID = "app_good_id"
	// FieldOrderID holds the string denoting the order_id field in the database.
	FieldOrderID = "order_id"
	// FieldContent holds the string denoting the content field in the database.
	FieldContent = "content"
	// FieldReplyToID holds the string denoting the reply_to_id field in the database.
	FieldReplyToID = "reply_to_id"
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
	FieldAppID,
	FieldUserID,
	FieldGoodID,
	FieldAppGoodID,
	FieldOrderID,
	FieldContent,
	FieldReplyToID,
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
	// DefaultAppID holds the default value on creation for the "app_id" field.
	DefaultAppID func() uuid.UUID
	// DefaultUserID holds the default value on creation for the "user_id" field.
	DefaultUserID func() uuid.UUID
	// DefaultGoodID holds the default value on creation for the "good_id" field.
	DefaultGoodID func() uuid.UUID
	// DefaultAppGoodID holds the default value on creation for the "app_good_id" field.
	DefaultAppGoodID func() uuid.UUID
	// DefaultOrderID holds the default value on creation for the "order_id" field.
	DefaultOrderID func() uuid.UUID
	// DefaultContent holds the default value on creation for the "content" field.
	DefaultContent string
	// DefaultReplyToID holds the default value on creation for the "reply_to_id" field.
	DefaultReplyToID func() uuid.UUID
)
