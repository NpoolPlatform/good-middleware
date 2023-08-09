// Code generated by ent, DO NOT EDIT.

package deviceinfo

import (
	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent/predicate"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.DeviceInfo {
	return predicate.DeviceInfo(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.DeviceInfo {
	return predicate.DeviceInfo(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.DeviceInfo {
	return predicate.DeviceInfo(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.DeviceInfo {
	return predicate.DeviceInfo(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.DeviceInfo {
	return predicate.DeviceInfo(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.DeviceInfo {
	return predicate.DeviceInfo(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.DeviceInfo {
	return predicate.DeviceInfo(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.DeviceInfo {
	return predicate.DeviceInfo(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.DeviceInfo {
	return predicate.DeviceInfo(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v uint32) predicate.DeviceInfo {
	return predicate.DeviceInfo(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v uint32) predicate.DeviceInfo {
	return predicate.DeviceInfo(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v uint32) predicate.DeviceInfo {
	return predicate.DeviceInfo(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeletedAt), v))
	})
}

// Type applies equality check predicate on the "type" field. It's identical to TypeEQ.
func Type(v string) predicate.DeviceInfo {
	return predicate.DeviceInfo(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldType), v))
	})
}

// Manufacturer applies equality check predicate on the "manufacturer" field. It's identical to ManufacturerEQ.
func Manufacturer(v string) predicate.DeviceInfo {
	return predicate.DeviceInfo(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldManufacturer), v))
	})
}

// PowerComsuption applies equality check predicate on the "power_comsuption" field. It's identical to PowerComsuptionEQ.
func PowerComsuption(v uint32) predicate.DeviceInfo {
	return predicate.DeviceInfo(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPowerComsuption), v))
	})
}

// ShipmentAt applies equality check predicate on the "shipment_at" field. It's identical to ShipmentAtEQ.
func ShipmentAt(v uint32) predicate.DeviceInfo {
	return predicate.DeviceInfo(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldShipmentAt), v))
	})
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v uint32) predicate.DeviceInfo {
	return predicate.DeviceInfo(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v uint32) predicate.DeviceInfo {
	return predicate.DeviceInfo(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...uint32) predicate.DeviceInfo {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.DeviceInfo(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...uint32) predicate.DeviceInfo {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.DeviceInfo(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v uint32) predicate.DeviceInfo {
	return predicate.DeviceInfo(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v uint32) predicate.DeviceInfo {
	return predicate.DeviceInfo(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v uint32) predicate.DeviceInfo {
	return predicate.DeviceInfo(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v uint32) predicate.DeviceInfo {
	return predicate.DeviceInfo(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v uint32) predicate.DeviceInfo {
	return predicate.DeviceInfo(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v uint32) predicate.DeviceInfo {
	return predicate.DeviceInfo(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...uint32) predicate.DeviceInfo {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.DeviceInfo(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...uint32) predicate.DeviceInfo {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.DeviceInfo(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v uint32) predicate.DeviceInfo {
	return predicate.DeviceInfo(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v uint32) predicate.DeviceInfo {
	return predicate.DeviceInfo(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v uint32) predicate.DeviceInfo {
	return predicate.DeviceInfo(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v uint32) predicate.DeviceInfo {
	return predicate.DeviceInfo(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdatedAt), v))
	})
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v uint32) predicate.DeviceInfo {
	return predicate.DeviceInfo(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v uint32) predicate.DeviceInfo {
	return predicate.DeviceInfo(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...uint32) predicate.DeviceInfo {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.DeviceInfo(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldDeletedAt), v...))
	})
}

// DeletedAtNotIn applies the NotIn predicate on the "deleted_at" field.
func DeletedAtNotIn(vs ...uint32) predicate.DeviceInfo {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.DeviceInfo(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldDeletedAt), v...))
	})
}

// DeletedAtGT applies the GT predicate on the "deleted_at" field.
func DeletedAtGT(v uint32) predicate.DeviceInfo {
	return predicate.DeviceInfo(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v uint32) predicate.DeviceInfo {
	return predicate.DeviceInfo(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v uint32) predicate.DeviceInfo {
	return predicate.DeviceInfo(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v uint32) predicate.DeviceInfo {
	return predicate.DeviceInfo(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDeletedAt), v))
	})
}

// TypeEQ applies the EQ predicate on the "type" field.
func TypeEQ(v string) predicate.DeviceInfo {
	return predicate.DeviceInfo(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldType), v))
	})
}

// TypeNEQ applies the NEQ predicate on the "type" field.
func TypeNEQ(v string) predicate.DeviceInfo {
	return predicate.DeviceInfo(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldType), v))
	})
}

// TypeIn applies the In predicate on the "type" field.
func TypeIn(vs ...string) predicate.DeviceInfo {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.DeviceInfo(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldType), v...))
	})
}

// TypeNotIn applies the NotIn predicate on the "type" field.
func TypeNotIn(vs ...string) predicate.DeviceInfo {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.DeviceInfo(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldType), v...))
	})
}

// TypeGT applies the GT predicate on the "type" field.
func TypeGT(v string) predicate.DeviceInfo {
	return predicate.DeviceInfo(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldType), v))
	})
}

// TypeGTE applies the GTE predicate on the "type" field.
func TypeGTE(v string) predicate.DeviceInfo {
	return predicate.DeviceInfo(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldType), v))
	})
}

// TypeLT applies the LT predicate on the "type" field.
func TypeLT(v string) predicate.DeviceInfo {
	return predicate.DeviceInfo(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldType), v))
	})
}

// TypeLTE applies the LTE predicate on the "type" field.
func TypeLTE(v string) predicate.DeviceInfo {
	return predicate.DeviceInfo(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldType), v))
	})
}

// TypeContains applies the Contains predicate on the "type" field.
func TypeContains(v string) predicate.DeviceInfo {
	return predicate.DeviceInfo(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldType), v))
	})
}

// TypeHasPrefix applies the HasPrefix predicate on the "type" field.
func TypeHasPrefix(v string) predicate.DeviceInfo {
	return predicate.DeviceInfo(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldType), v))
	})
}

// TypeHasSuffix applies the HasSuffix predicate on the "type" field.
func TypeHasSuffix(v string) predicate.DeviceInfo {
	return predicate.DeviceInfo(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldType), v))
	})
}

// TypeIsNil applies the IsNil predicate on the "type" field.
func TypeIsNil() predicate.DeviceInfo {
	return predicate.DeviceInfo(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldType)))
	})
}

// TypeNotNil applies the NotNil predicate on the "type" field.
func TypeNotNil() predicate.DeviceInfo {
	return predicate.DeviceInfo(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldType)))
	})
}

// TypeEqualFold applies the EqualFold predicate on the "type" field.
func TypeEqualFold(v string) predicate.DeviceInfo {
	return predicate.DeviceInfo(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldType), v))
	})
}

// TypeContainsFold applies the ContainsFold predicate on the "type" field.
func TypeContainsFold(v string) predicate.DeviceInfo {
	return predicate.DeviceInfo(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldType), v))
	})
}

// ManufacturerEQ applies the EQ predicate on the "manufacturer" field.
func ManufacturerEQ(v string) predicate.DeviceInfo {
	return predicate.DeviceInfo(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldManufacturer), v))
	})
}

// ManufacturerNEQ applies the NEQ predicate on the "manufacturer" field.
func ManufacturerNEQ(v string) predicate.DeviceInfo {
	return predicate.DeviceInfo(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldManufacturer), v))
	})
}

// ManufacturerIn applies the In predicate on the "manufacturer" field.
func ManufacturerIn(vs ...string) predicate.DeviceInfo {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.DeviceInfo(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldManufacturer), v...))
	})
}

// ManufacturerNotIn applies the NotIn predicate on the "manufacturer" field.
func ManufacturerNotIn(vs ...string) predicate.DeviceInfo {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.DeviceInfo(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldManufacturer), v...))
	})
}

// ManufacturerGT applies the GT predicate on the "manufacturer" field.
func ManufacturerGT(v string) predicate.DeviceInfo {
	return predicate.DeviceInfo(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldManufacturer), v))
	})
}

// ManufacturerGTE applies the GTE predicate on the "manufacturer" field.
func ManufacturerGTE(v string) predicate.DeviceInfo {
	return predicate.DeviceInfo(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldManufacturer), v))
	})
}

// ManufacturerLT applies the LT predicate on the "manufacturer" field.
func ManufacturerLT(v string) predicate.DeviceInfo {
	return predicate.DeviceInfo(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldManufacturer), v))
	})
}

// ManufacturerLTE applies the LTE predicate on the "manufacturer" field.
func ManufacturerLTE(v string) predicate.DeviceInfo {
	return predicate.DeviceInfo(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldManufacturer), v))
	})
}

// ManufacturerContains applies the Contains predicate on the "manufacturer" field.
func ManufacturerContains(v string) predicate.DeviceInfo {
	return predicate.DeviceInfo(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldManufacturer), v))
	})
}

// ManufacturerHasPrefix applies the HasPrefix predicate on the "manufacturer" field.
func ManufacturerHasPrefix(v string) predicate.DeviceInfo {
	return predicate.DeviceInfo(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldManufacturer), v))
	})
}

// ManufacturerHasSuffix applies the HasSuffix predicate on the "manufacturer" field.
func ManufacturerHasSuffix(v string) predicate.DeviceInfo {
	return predicate.DeviceInfo(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldManufacturer), v))
	})
}

// ManufacturerIsNil applies the IsNil predicate on the "manufacturer" field.
func ManufacturerIsNil() predicate.DeviceInfo {
	return predicate.DeviceInfo(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldManufacturer)))
	})
}

// ManufacturerNotNil applies the NotNil predicate on the "manufacturer" field.
func ManufacturerNotNil() predicate.DeviceInfo {
	return predicate.DeviceInfo(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldManufacturer)))
	})
}

// ManufacturerEqualFold applies the EqualFold predicate on the "manufacturer" field.
func ManufacturerEqualFold(v string) predicate.DeviceInfo {
	return predicate.DeviceInfo(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldManufacturer), v))
	})
}

// ManufacturerContainsFold applies the ContainsFold predicate on the "manufacturer" field.
func ManufacturerContainsFold(v string) predicate.DeviceInfo {
	return predicate.DeviceInfo(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldManufacturer), v))
	})
}

// PowerComsuptionEQ applies the EQ predicate on the "power_comsuption" field.
func PowerComsuptionEQ(v uint32) predicate.DeviceInfo {
	return predicate.DeviceInfo(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPowerComsuption), v))
	})
}

// PowerComsuptionNEQ applies the NEQ predicate on the "power_comsuption" field.
func PowerComsuptionNEQ(v uint32) predicate.DeviceInfo {
	return predicate.DeviceInfo(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldPowerComsuption), v))
	})
}

// PowerComsuptionIn applies the In predicate on the "power_comsuption" field.
func PowerComsuptionIn(vs ...uint32) predicate.DeviceInfo {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.DeviceInfo(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldPowerComsuption), v...))
	})
}

// PowerComsuptionNotIn applies the NotIn predicate on the "power_comsuption" field.
func PowerComsuptionNotIn(vs ...uint32) predicate.DeviceInfo {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.DeviceInfo(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldPowerComsuption), v...))
	})
}

// PowerComsuptionGT applies the GT predicate on the "power_comsuption" field.
func PowerComsuptionGT(v uint32) predicate.DeviceInfo {
	return predicate.DeviceInfo(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldPowerComsuption), v))
	})
}

// PowerComsuptionGTE applies the GTE predicate on the "power_comsuption" field.
func PowerComsuptionGTE(v uint32) predicate.DeviceInfo {
	return predicate.DeviceInfo(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldPowerComsuption), v))
	})
}

// PowerComsuptionLT applies the LT predicate on the "power_comsuption" field.
func PowerComsuptionLT(v uint32) predicate.DeviceInfo {
	return predicate.DeviceInfo(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldPowerComsuption), v))
	})
}

// PowerComsuptionLTE applies the LTE predicate on the "power_comsuption" field.
func PowerComsuptionLTE(v uint32) predicate.DeviceInfo {
	return predicate.DeviceInfo(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldPowerComsuption), v))
	})
}

// PowerComsuptionIsNil applies the IsNil predicate on the "power_comsuption" field.
func PowerComsuptionIsNil() predicate.DeviceInfo {
	return predicate.DeviceInfo(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldPowerComsuption)))
	})
}

// PowerComsuptionNotNil applies the NotNil predicate on the "power_comsuption" field.
func PowerComsuptionNotNil() predicate.DeviceInfo {
	return predicate.DeviceInfo(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldPowerComsuption)))
	})
}

// ShipmentAtEQ applies the EQ predicate on the "shipment_at" field.
func ShipmentAtEQ(v uint32) predicate.DeviceInfo {
	return predicate.DeviceInfo(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldShipmentAt), v))
	})
}

// ShipmentAtNEQ applies the NEQ predicate on the "shipment_at" field.
func ShipmentAtNEQ(v uint32) predicate.DeviceInfo {
	return predicate.DeviceInfo(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldShipmentAt), v))
	})
}

// ShipmentAtIn applies the In predicate on the "shipment_at" field.
func ShipmentAtIn(vs ...uint32) predicate.DeviceInfo {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.DeviceInfo(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldShipmentAt), v...))
	})
}

// ShipmentAtNotIn applies the NotIn predicate on the "shipment_at" field.
func ShipmentAtNotIn(vs ...uint32) predicate.DeviceInfo {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.DeviceInfo(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldShipmentAt), v...))
	})
}

// ShipmentAtGT applies the GT predicate on the "shipment_at" field.
func ShipmentAtGT(v uint32) predicate.DeviceInfo {
	return predicate.DeviceInfo(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldShipmentAt), v))
	})
}

// ShipmentAtGTE applies the GTE predicate on the "shipment_at" field.
func ShipmentAtGTE(v uint32) predicate.DeviceInfo {
	return predicate.DeviceInfo(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldShipmentAt), v))
	})
}

// ShipmentAtLT applies the LT predicate on the "shipment_at" field.
func ShipmentAtLT(v uint32) predicate.DeviceInfo {
	return predicate.DeviceInfo(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldShipmentAt), v))
	})
}

// ShipmentAtLTE applies the LTE predicate on the "shipment_at" field.
func ShipmentAtLTE(v uint32) predicate.DeviceInfo {
	return predicate.DeviceInfo(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldShipmentAt), v))
	})
}

// ShipmentAtIsNil applies the IsNil predicate on the "shipment_at" field.
func ShipmentAtIsNil() predicate.DeviceInfo {
	return predicate.DeviceInfo(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldShipmentAt)))
	})
}

// ShipmentAtNotNil applies the NotNil predicate on the "shipment_at" field.
func ShipmentAtNotNil() predicate.DeviceInfo {
	return predicate.DeviceInfo(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldShipmentAt)))
	})
}

// PostersIsNil applies the IsNil predicate on the "posters" field.
func PostersIsNil() predicate.DeviceInfo {
	return predicate.DeviceInfo(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldPosters)))
	})
}

// PostersNotNil applies the NotNil predicate on the "posters" field.
func PostersNotNil() predicate.DeviceInfo {
	return predicate.DeviceInfo(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldPosters)))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.DeviceInfo) predicate.DeviceInfo {
	return predicate.DeviceInfo(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.DeviceInfo) predicate.DeviceInfo {
	return predicate.DeviceInfo(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.DeviceInfo) predicate.DeviceInfo {
	return predicate.DeviceInfo(func(s *sql.Selector) {
		p(s.Not())
	})
}
