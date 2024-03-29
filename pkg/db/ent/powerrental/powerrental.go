// Code generated by ent, DO NOT EDIT.

package powerrental

import (
	"entgo.io/ent"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

const (
	// Label holds the string label denoting the powerrental type in the database.
	Label = "power_rental"
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
	// FieldDeviceTypeID holds the string denoting the device_type_id field in the database.
	FieldDeviceTypeID = "device_type_id"
	// FieldVendorLocationID holds the string denoting the vendor_location_id field in the database.
	FieldVendorLocationID = "vendor_location_id"
	// FieldUnitPrice holds the string denoting the unit_price field in the database.
	FieldUnitPrice = "unit_price"
	// FieldQuantityUnit holds the string denoting the quantity_unit field in the database.
	FieldQuantityUnit = "quantity_unit"
	// FieldQuantityUnitAmount holds the string denoting the quantity_unit_amount field in the database.
	FieldQuantityUnitAmount = "quantity_unit_amount"
	// FieldDeliveryAt holds the string denoting the delivery_at field in the database.
	FieldDeliveryAt = "delivery_at"
	// FieldUnitLockDeposit holds the string denoting the unit_lock_deposit field in the database.
	FieldUnitLockDeposit = "unit_lock_deposit"
	// FieldDurationType holds the string denoting the duration_type field in the database.
	FieldDurationType = "duration_type"
	// FieldStockMode holds the string denoting the stock_mode field in the database.
	FieldStockMode = "stock_mode"
	// Table holds the table name of the powerrental in the database.
	Table = "power_rentals"
)

// Columns holds all SQL columns for powerrental fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
	FieldEntID,
	FieldGoodID,
	FieldDeviceTypeID,
	FieldVendorLocationID,
	FieldUnitPrice,
	FieldQuantityUnit,
	FieldQuantityUnitAmount,
	FieldDeliveryAt,
	FieldUnitLockDeposit,
	FieldDurationType,
	FieldStockMode,
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
	// DefaultDeviceTypeID holds the default value on creation for the "device_type_id" field.
	DefaultDeviceTypeID func() uuid.UUID
	// DefaultVendorLocationID holds the default value on creation for the "vendor_location_id" field.
	DefaultVendorLocationID func() uuid.UUID
	// DefaultUnitPrice holds the default value on creation for the "unit_price" field.
	DefaultUnitPrice decimal.Decimal
	// DefaultQuantityUnit holds the default value on creation for the "quantity_unit" field.
	DefaultQuantityUnit string
	// DefaultQuantityUnitAmount holds the default value on creation for the "quantity_unit_amount" field.
	DefaultQuantityUnitAmount decimal.Decimal
	// DefaultDeliveryAt holds the default value on creation for the "delivery_at" field.
	DefaultDeliveryAt uint32
	// DefaultUnitLockDeposit holds the default value on creation for the "unit_lock_deposit" field.
	DefaultUnitLockDeposit decimal.Decimal
	// DefaultDurationType holds the default value on creation for the "duration_type" field.
	DefaultDurationType string
	// DefaultStockMode holds the default value on creation for the "stock_mode" field.
	DefaultStockMode string
)
