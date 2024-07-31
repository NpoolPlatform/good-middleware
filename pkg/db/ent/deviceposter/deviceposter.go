// Code generated by ent, DO NOT EDIT.

package deviceposter

import (
	"entgo.io/ent"
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the deviceposter type in the database.
	Label = "device_poster"
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
	// FieldDeviceTypeID holds the string denoting the device_type_id field in the database.
	FieldDeviceTypeID = "device_type_id"
	// FieldPoster holds the string denoting the poster field in the database.
	FieldPoster = "poster"
	// FieldIndex holds the string denoting the index field in the database.
	FieldIndex = "index"
	// Table holds the table name of the deviceposter in the database.
	Table = "device_posters"
)

// Columns holds all SQL columns for deviceposter fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
	FieldEntID,
	FieldDeviceTypeID,
	FieldPoster,
	FieldIndex,
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
	// DefaultDeviceTypeID holds the default value on creation for the "device_type_id" field.
	DefaultDeviceTypeID func() uuid.UUID
	// DefaultPoster holds the default value on creation for the "poster" field.
	DefaultPoster string
	// DefaultIndex holds the default value on creation for the "index" field.
	DefaultIndex uint8
)
