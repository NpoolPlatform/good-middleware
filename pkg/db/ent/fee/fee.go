// Code generated by ent, DO NOT EDIT.

package fee

import (
	"entgo.io/ent"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

const (
	// Label holds the string label denoting the fee type in the database.
	Label = "fee"
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
	// FieldGoodID holds the string denoting the good_id field in the database.
	FieldGoodID = "good_id"
	// FieldSettlementType holds the string denoting the settlement_type field in the database.
	FieldSettlementType = "settlement_type"
	// FieldUnitValue holds the string denoting the unit_value field in the database.
	FieldUnitValue = "unit_value"
	// FieldDurationDisplayType holds the string denoting the duration_display_type field in the database.
	FieldDurationDisplayType = "duration_display_type"
	// Table holds the table name of the fee in the database.
	Table = "fees"
)

// Columns holds all SQL columns for fee fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
	FieldEntID,
	FieldGoodID,
	FieldSettlementType,
	FieldUnitValue,
	FieldDurationDisplayType,
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
	// DefaultGoodID holds the default value on creation for the "good_id" field.
	DefaultGoodID func() uuid.UUID
	// DefaultSettlementType holds the default value on creation for the "settlement_type" field.
	DefaultSettlementType string
	// DefaultUnitValue holds the default value on creation for the "unit_value" field.
	DefaultUnitValue decimal.Decimal
	// DefaultDurationDisplayType holds the default value on creation for the "duration_display_type" field.
	DefaultDurationDisplayType string
)
