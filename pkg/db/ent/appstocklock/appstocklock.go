// Code generated by ent, DO NOT EDIT.

package appstocklock

import (
	"entgo.io/ent"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

const (
	// Label holds the string label denoting the appstocklock type in the database.
	Label = "app_stock_lock"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldDeletedAt holds the string denoting the deleted_at field in the database.
	FieldDeletedAt = "deleted_at"
	// FieldAppStockID holds the string denoting the app_stock_id field in the database.
	FieldAppStockID = "app_stock_id"
	// FieldUnits holds the string denoting the units field in the database.
	FieldUnits = "units"
	// FieldAppSpotUnits holds the string denoting the app_spot_units field in the database.
	FieldAppSpotUnits = "app_spot_units"
	// FieldLockState holds the string denoting the lock_state field in the database.
	FieldLockState = "lock_state"
	// FieldChargeBackState holds the string denoting the charge_back_state field in the database.
	FieldChargeBackState = "charge_back_state"
	// Table holds the table name of the appstocklock in the database.
	Table = "app_stock_locks"
)

// Columns holds all SQL columns for appstocklock fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
	FieldAppStockID,
	FieldUnits,
	FieldAppSpotUnits,
	FieldLockState,
	FieldChargeBackState,
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
	// DefaultAppStockID holds the default value on creation for the "app_stock_id" field.
	DefaultAppStockID func() uuid.UUID
	// DefaultUnits holds the default value on creation for the "units" field.
	DefaultUnits decimal.Decimal
	// DefaultAppSpotUnits holds the default value on creation for the "app_spot_units" field.
	DefaultAppSpotUnits decimal.Decimal
	// DefaultLockState holds the default value on creation for the "lock_state" field.
	DefaultLockState string
	// DefaultChargeBackState holds the default value on creation for the "charge_back_state" field.
	DefaultChargeBackState string
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)