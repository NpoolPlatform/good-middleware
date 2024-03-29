// Code generated by ent, DO NOT EDIT.

package deviceinfo

import (
	"entgo.io/ent"
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the deviceinfo type in the database.
	Label = "device_info"
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
	// FieldType holds the string denoting the type field in the database.
	FieldType = "type"
	// FieldManufacturer holds the string denoting the manufacturer field in the database.
	FieldManufacturer = "manufacturer"
	// FieldPowerConsumption holds the string denoting the power_consumption field in the database.
	FieldPowerConsumption = "power_consumption"
	// FieldShipmentAt holds the string denoting the shipment_at field in the database.
	FieldShipmentAt = "shipment_at"
	// FieldPosters holds the string denoting the posters field in the database.
	FieldPosters = "posters"
	// Table holds the table name of the deviceinfo in the database.
	Table = "device_infos"
)

// Columns holds all SQL columns for deviceinfo fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
	FieldEntID,
	FieldType,
	FieldManufacturer,
	FieldPowerConsumption,
	FieldShipmentAt,
	FieldPosters,
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
	// DefaultType holds the default value on creation for the "type" field.
	DefaultType string
	// TypeValidator is a validator for the "type" field. It is called by the builders before save.
	TypeValidator func(string) error
	// DefaultManufacturer holds the default value on creation for the "manufacturer" field.
	DefaultManufacturer string
	// ManufacturerValidator is a validator for the "manufacturer" field. It is called by the builders before save.
	ManufacturerValidator func(string) error
	// DefaultPowerConsumption holds the default value on creation for the "power_consumption" field.
	DefaultPowerConsumption uint32
	// DefaultShipmentAt holds the default value on creation for the "shipment_at" field.
	DefaultShipmentAt uint32
	// DefaultPosters holds the default value on creation for the "posters" field.
	DefaultPosters []string
)
