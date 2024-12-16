// Code generated by ent, DO NOT EDIT.

package apppledge

import (
	"entgo.io/ent"
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the apppledge type in the database.
	Label = "app_pledge"
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
	// FieldAppGoodID holds the string denoting the app_good_id field in the database.
	FieldAppGoodID = "app_good_id"
	// FieldServiceStartAt holds the string denoting the service_start_at field in the database.
	FieldServiceStartAt = "service_start_at"
	// FieldStartMode holds the string denoting the start_mode field in the database.
	FieldStartMode = "start_mode"
	// FieldEnableSetCommission holds the string denoting the enable_set_commission field in the database.
	FieldEnableSetCommission = "enable_set_commission"
	// Table holds the table name of the apppledge in the database.
	Table = "app_pledges"
)

// Columns holds all SQL columns for apppledge fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
	FieldEntID,
	FieldAppGoodID,
	FieldServiceStartAt,
	FieldStartMode,
	FieldEnableSetCommission,
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
	// DefaultAppGoodID holds the default value on creation for the "app_good_id" field.
	DefaultAppGoodID func() uuid.UUID
	// DefaultServiceStartAt holds the default value on creation for the "service_start_at" field.
	DefaultServiceStartAt uint32
	// DefaultStartMode holds the default value on creation for the "start_mode" field.
	DefaultStartMode string
	// DefaultEnableSetCommission holds the default value on creation for the "enable_set_commission" field.
	DefaultEnableSetCommission bool
)
