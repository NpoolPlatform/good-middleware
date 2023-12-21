// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent/deviceinfo"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent/predicate"
	"github.com/google/uuid"
)

// DeviceInfoUpdate is the builder for updating DeviceInfo entities.
type DeviceInfoUpdate struct {
	config
	hooks     []Hook
	mutation  *DeviceInfoMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the DeviceInfoUpdate builder.
func (diu *DeviceInfoUpdate) Where(ps ...predicate.DeviceInfo) *DeviceInfoUpdate {
	diu.mutation.Where(ps...)
	return diu
}

// SetCreatedAt sets the "created_at" field.
func (diu *DeviceInfoUpdate) SetCreatedAt(u uint32) *DeviceInfoUpdate {
	diu.mutation.ResetCreatedAt()
	diu.mutation.SetCreatedAt(u)
	return diu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (diu *DeviceInfoUpdate) SetNillableCreatedAt(u *uint32) *DeviceInfoUpdate {
	if u != nil {
		diu.SetCreatedAt(*u)
	}
	return diu
}

// AddCreatedAt adds u to the "created_at" field.
func (diu *DeviceInfoUpdate) AddCreatedAt(u int32) *DeviceInfoUpdate {
	diu.mutation.AddCreatedAt(u)
	return diu
}

// SetUpdatedAt sets the "updated_at" field.
func (diu *DeviceInfoUpdate) SetUpdatedAt(u uint32) *DeviceInfoUpdate {
	diu.mutation.ResetUpdatedAt()
	diu.mutation.SetUpdatedAt(u)
	return diu
}

// AddUpdatedAt adds u to the "updated_at" field.
func (diu *DeviceInfoUpdate) AddUpdatedAt(u int32) *DeviceInfoUpdate {
	diu.mutation.AddUpdatedAt(u)
	return diu
}

// SetDeletedAt sets the "deleted_at" field.
func (diu *DeviceInfoUpdate) SetDeletedAt(u uint32) *DeviceInfoUpdate {
	diu.mutation.ResetDeletedAt()
	diu.mutation.SetDeletedAt(u)
	return diu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (diu *DeviceInfoUpdate) SetNillableDeletedAt(u *uint32) *DeviceInfoUpdate {
	if u != nil {
		diu.SetDeletedAt(*u)
	}
	return diu
}

// AddDeletedAt adds u to the "deleted_at" field.
func (diu *DeviceInfoUpdate) AddDeletedAt(u int32) *DeviceInfoUpdate {
	diu.mutation.AddDeletedAt(u)
	return diu
}

// SetEntID sets the "ent_id" field.
func (diu *DeviceInfoUpdate) SetEntID(u uuid.UUID) *DeviceInfoUpdate {
	diu.mutation.SetEntID(u)
	return diu
}

// SetNillableEntID sets the "ent_id" field if the given value is not nil.
func (diu *DeviceInfoUpdate) SetNillableEntID(u *uuid.UUID) *DeviceInfoUpdate {
	if u != nil {
		diu.SetEntID(*u)
	}
	return diu
}

// SetType sets the "type" field.
func (diu *DeviceInfoUpdate) SetType(s string) *DeviceInfoUpdate {
	diu.mutation.SetType(s)
	return diu
}

// SetNillableType sets the "type" field if the given value is not nil.
func (diu *DeviceInfoUpdate) SetNillableType(s *string) *DeviceInfoUpdate {
	if s != nil {
		diu.SetType(*s)
	}
	return diu
}

// ClearType clears the value of the "type" field.
func (diu *DeviceInfoUpdate) ClearType() *DeviceInfoUpdate {
	diu.mutation.ClearType()
	return diu
}

// SetManufacturer sets the "manufacturer" field.
func (diu *DeviceInfoUpdate) SetManufacturer(s string) *DeviceInfoUpdate {
	diu.mutation.SetManufacturer(s)
	return diu
}

// SetNillableManufacturer sets the "manufacturer" field if the given value is not nil.
func (diu *DeviceInfoUpdate) SetNillableManufacturer(s *string) *DeviceInfoUpdate {
	if s != nil {
		diu.SetManufacturer(*s)
	}
	return diu
}

// ClearManufacturer clears the value of the "manufacturer" field.
func (diu *DeviceInfoUpdate) ClearManufacturer() *DeviceInfoUpdate {
	diu.mutation.ClearManufacturer()
	return diu
}

// SetPowerConsumption sets the "power_consumption" field.
func (diu *DeviceInfoUpdate) SetPowerConsumption(u uint32) *DeviceInfoUpdate {
	diu.mutation.ResetPowerConsumption()
	diu.mutation.SetPowerConsumption(u)
	return diu
}

// SetNillablePowerConsumption sets the "power_consumption" field if the given value is not nil.
func (diu *DeviceInfoUpdate) SetNillablePowerConsumption(u *uint32) *DeviceInfoUpdate {
	if u != nil {
		diu.SetPowerConsumption(*u)
	}
	return diu
}

// AddPowerConsumption adds u to the "power_consumption" field.
func (diu *DeviceInfoUpdate) AddPowerConsumption(u int32) *DeviceInfoUpdate {
	diu.mutation.AddPowerConsumption(u)
	return diu
}

// ClearPowerConsumption clears the value of the "power_consumption" field.
func (diu *DeviceInfoUpdate) ClearPowerConsumption() *DeviceInfoUpdate {
	diu.mutation.ClearPowerConsumption()
	return diu
}

// SetShipmentAt sets the "shipment_at" field.
func (diu *DeviceInfoUpdate) SetShipmentAt(u uint32) *DeviceInfoUpdate {
	diu.mutation.ResetShipmentAt()
	diu.mutation.SetShipmentAt(u)
	return diu
}

// SetNillableShipmentAt sets the "shipment_at" field if the given value is not nil.
func (diu *DeviceInfoUpdate) SetNillableShipmentAt(u *uint32) *DeviceInfoUpdate {
	if u != nil {
		diu.SetShipmentAt(*u)
	}
	return diu
}

// AddShipmentAt adds u to the "shipment_at" field.
func (diu *DeviceInfoUpdate) AddShipmentAt(u int32) *DeviceInfoUpdate {
	diu.mutation.AddShipmentAt(u)
	return diu
}

// ClearShipmentAt clears the value of the "shipment_at" field.
func (diu *DeviceInfoUpdate) ClearShipmentAt() *DeviceInfoUpdate {
	diu.mutation.ClearShipmentAt()
	return diu
}

// SetPosters sets the "posters" field.
func (diu *DeviceInfoUpdate) SetPosters(s []string) *DeviceInfoUpdate {
	diu.mutation.SetPosters(s)
	return diu
}

// ClearPosters clears the value of the "posters" field.
func (diu *DeviceInfoUpdate) ClearPosters() *DeviceInfoUpdate {
	diu.mutation.ClearPosters()
	return diu
}

// Mutation returns the DeviceInfoMutation object of the builder.
func (diu *DeviceInfoUpdate) Mutation() *DeviceInfoMutation {
	return diu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (diu *DeviceInfoUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if err := diu.defaults(); err != nil {
		return 0, err
	}
	if len(diu.hooks) == 0 {
		if err = diu.check(); err != nil {
			return 0, err
		}
		affected, err = diu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*DeviceInfoMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = diu.check(); err != nil {
				return 0, err
			}
			diu.mutation = mutation
			affected, err = diu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(diu.hooks) - 1; i >= 0; i-- {
			if diu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = diu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, diu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (diu *DeviceInfoUpdate) SaveX(ctx context.Context) int {
	affected, err := diu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (diu *DeviceInfoUpdate) Exec(ctx context.Context) error {
	_, err := diu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (diu *DeviceInfoUpdate) ExecX(ctx context.Context) {
	if err := diu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (diu *DeviceInfoUpdate) defaults() error {
	if _, ok := diu.mutation.UpdatedAt(); !ok {
		if deviceinfo.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized deviceinfo.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := deviceinfo.UpdateDefaultUpdatedAt()
		diu.mutation.SetUpdatedAt(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (diu *DeviceInfoUpdate) check() error {
	if v, ok := diu.mutation.GetType(); ok {
		if err := deviceinfo.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`ent: validator failed for field "DeviceInfo.type": %w`, err)}
		}
	}
	if v, ok := diu.mutation.Manufacturer(); ok {
		if err := deviceinfo.ManufacturerValidator(v); err != nil {
			return &ValidationError{Name: "manufacturer", err: fmt.Errorf(`ent: validator failed for field "DeviceInfo.manufacturer": %w`, err)}
		}
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (diu *DeviceInfoUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *DeviceInfoUpdate {
	diu.modifiers = append(diu.modifiers, modifiers...)
	return diu
}

func (diu *DeviceInfoUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   deviceinfo.Table,
			Columns: deviceinfo.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: deviceinfo.FieldID,
			},
		},
	}
	if ps := diu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := diu.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: deviceinfo.FieldCreatedAt,
		})
	}
	if value, ok := diu.mutation.AddedCreatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: deviceinfo.FieldCreatedAt,
		})
	}
	if value, ok := diu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: deviceinfo.FieldUpdatedAt,
		})
	}
	if value, ok := diu.mutation.AddedUpdatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: deviceinfo.FieldUpdatedAt,
		})
	}
	if value, ok := diu.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: deviceinfo.FieldDeletedAt,
		})
	}
	if value, ok := diu.mutation.AddedDeletedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: deviceinfo.FieldDeletedAt,
		})
	}
	if value, ok := diu.mutation.EntID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: deviceinfo.FieldEntID,
		})
	}
	if value, ok := diu.mutation.GetType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: deviceinfo.FieldType,
		})
	}
	if diu.mutation.TypeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: deviceinfo.FieldType,
		})
	}
	if value, ok := diu.mutation.Manufacturer(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: deviceinfo.FieldManufacturer,
		})
	}
	if diu.mutation.ManufacturerCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: deviceinfo.FieldManufacturer,
		})
	}
	if value, ok := diu.mutation.PowerConsumption(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: deviceinfo.FieldPowerConsumption,
		})
	}
	if value, ok := diu.mutation.AddedPowerConsumption(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: deviceinfo.FieldPowerConsumption,
		})
	}
	if diu.mutation.PowerConsumptionCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Column: deviceinfo.FieldPowerConsumption,
		})
	}
	if value, ok := diu.mutation.ShipmentAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: deviceinfo.FieldShipmentAt,
		})
	}
	if value, ok := diu.mutation.AddedShipmentAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: deviceinfo.FieldShipmentAt,
		})
	}
	if diu.mutation.ShipmentAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Column: deviceinfo.FieldShipmentAt,
		})
	}
	if value, ok := diu.mutation.Posters(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: deviceinfo.FieldPosters,
		})
	}
	if diu.mutation.PostersCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Column: deviceinfo.FieldPosters,
		})
	}
	_spec.Modifiers = diu.modifiers
	if n, err = sqlgraph.UpdateNodes(ctx, diu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{deviceinfo.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// DeviceInfoUpdateOne is the builder for updating a single DeviceInfo entity.
type DeviceInfoUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *DeviceInfoMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetCreatedAt sets the "created_at" field.
func (diuo *DeviceInfoUpdateOne) SetCreatedAt(u uint32) *DeviceInfoUpdateOne {
	diuo.mutation.ResetCreatedAt()
	diuo.mutation.SetCreatedAt(u)
	return diuo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (diuo *DeviceInfoUpdateOne) SetNillableCreatedAt(u *uint32) *DeviceInfoUpdateOne {
	if u != nil {
		diuo.SetCreatedAt(*u)
	}
	return diuo
}

// AddCreatedAt adds u to the "created_at" field.
func (diuo *DeviceInfoUpdateOne) AddCreatedAt(u int32) *DeviceInfoUpdateOne {
	diuo.mutation.AddCreatedAt(u)
	return diuo
}

// SetUpdatedAt sets the "updated_at" field.
func (diuo *DeviceInfoUpdateOne) SetUpdatedAt(u uint32) *DeviceInfoUpdateOne {
	diuo.mutation.ResetUpdatedAt()
	diuo.mutation.SetUpdatedAt(u)
	return diuo
}

// AddUpdatedAt adds u to the "updated_at" field.
func (diuo *DeviceInfoUpdateOne) AddUpdatedAt(u int32) *DeviceInfoUpdateOne {
	diuo.mutation.AddUpdatedAt(u)
	return diuo
}

// SetDeletedAt sets the "deleted_at" field.
func (diuo *DeviceInfoUpdateOne) SetDeletedAt(u uint32) *DeviceInfoUpdateOne {
	diuo.mutation.ResetDeletedAt()
	diuo.mutation.SetDeletedAt(u)
	return diuo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (diuo *DeviceInfoUpdateOne) SetNillableDeletedAt(u *uint32) *DeviceInfoUpdateOne {
	if u != nil {
		diuo.SetDeletedAt(*u)
	}
	return diuo
}

// AddDeletedAt adds u to the "deleted_at" field.
func (diuo *DeviceInfoUpdateOne) AddDeletedAt(u int32) *DeviceInfoUpdateOne {
	diuo.mutation.AddDeletedAt(u)
	return diuo
}

// SetEntID sets the "ent_id" field.
func (diuo *DeviceInfoUpdateOne) SetEntID(u uuid.UUID) *DeviceInfoUpdateOne {
	diuo.mutation.SetEntID(u)
	return diuo
}

// SetNillableEntID sets the "ent_id" field if the given value is not nil.
func (diuo *DeviceInfoUpdateOne) SetNillableEntID(u *uuid.UUID) *DeviceInfoUpdateOne {
	if u != nil {
		diuo.SetEntID(*u)
	}
	return diuo
}

// SetType sets the "type" field.
func (diuo *DeviceInfoUpdateOne) SetType(s string) *DeviceInfoUpdateOne {
	diuo.mutation.SetType(s)
	return diuo
}

// SetNillableType sets the "type" field if the given value is not nil.
func (diuo *DeviceInfoUpdateOne) SetNillableType(s *string) *DeviceInfoUpdateOne {
	if s != nil {
		diuo.SetType(*s)
	}
	return diuo
}

// ClearType clears the value of the "type" field.
func (diuo *DeviceInfoUpdateOne) ClearType() *DeviceInfoUpdateOne {
	diuo.mutation.ClearType()
	return diuo
}

// SetManufacturer sets the "manufacturer" field.
func (diuo *DeviceInfoUpdateOne) SetManufacturer(s string) *DeviceInfoUpdateOne {
	diuo.mutation.SetManufacturer(s)
	return diuo
}

// SetNillableManufacturer sets the "manufacturer" field if the given value is not nil.
func (diuo *DeviceInfoUpdateOne) SetNillableManufacturer(s *string) *DeviceInfoUpdateOne {
	if s != nil {
		diuo.SetManufacturer(*s)
	}
	return diuo
}

// ClearManufacturer clears the value of the "manufacturer" field.
func (diuo *DeviceInfoUpdateOne) ClearManufacturer() *DeviceInfoUpdateOne {
	diuo.mutation.ClearManufacturer()
	return diuo
}

// SetPowerConsumption sets the "power_consumption" field.
func (diuo *DeviceInfoUpdateOne) SetPowerConsumption(u uint32) *DeviceInfoUpdateOne {
	diuo.mutation.ResetPowerConsumption()
	diuo.mutation.SetPowerConsumption(u)
	return diuo
}

// SetNillablePowerConsumption sets the "power_consumption" field if the given value is not nil.
func (diuo *DeviceInfoUpdateOne) SetNillablePowerConsumption(u *uint32) *DeviceInfoUpdateOne {
	if u != nil {
		diuo.SetPowerConsumption(*u)
	}
	return diuo
}

// AddPowerConsumption adds u to the "power_consumption" field.
func (diuo *DeviceInfoUpdateOne) AddPowerConsumption(u int32) *DeviceInfoUpdateOne {
	diuo.mutation.AddPowerConsumption(u)
	return diuo
}

// ClearPowerConsumption clears the value of the "power_consumption" field.
func (diuo *DeviceInfoUpdateOne) ClearPowerConsumption() *DeviceInfoUpdateOne {
	diuo.mutation.ClearPowerConsumption()
	return diuo
}

// SetShipmentAt sets the "shipment_at" field.
func (diuo *DeviceInfoUpdateOne) SetShipmentAt(u uint32) *DeviceInfoUpdateOne {
	diuo.mutation.ResetShipmentAt()
	diuo.mutation.SetShipmentAt(u)
	return diuo
}

// SetNillableShipmentAt sets the "shipment_at" field if the given value is not nil.
func (diuo *DeviceInfoUpdateOne) SetNillableShipmentAt(u *uint32) *DeviceInfoUpdateOne {
	if u != nil {
		diuo.SetShipmentAt(*u)
	}
	return diuo
}

// AddShipmentAt adds u to the "shipment_at" field.
func (diuo *DeviceInfoUpdateOne) AddShipmentAt(u int32) *DeviceInfoUpdateOne {
	diuo.mutation.AddShipmentAt(u)
	return diuo
}

// ClearShipmentAt clears the value of the "shipment_at" field.
func (diuo *DeviceInfoUpdateOne) ClearShipmentAt() *DeviceInfoUpdateOne {
	diuo.mutation.ClearShipmentAt()
	return diuo
}

// SetPosters sets the "posters" field.
func (diuo *DeviceInfoUpdateOne) SetPosters(s []string) *DeviceInfoUpdateOne {
	diuo.mutation.SetPosters(s)
	return diuo
}

// ClearPosters clears the value of the "posters" field.
func (diuo *DeviceInfoUpdateOne) ClearPosters() *DeviceInfoUpdateOne {
	diuo.mutation.ClearPosters()
	return diuo
}

// Mutation returns the DeviceInfoMutation object of the builder.
func (diuo *DeviceInfoUpdateOne) Mutation() *DeviceInfoMutation {
	return diuo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (diuo *DeviceInfoUpdateOne) Select(field string, fields ...string) *DeviceInfoUpdateOne {
	diuo.fields = append([]string{field}, fields...)
	return diuo
}

// Save executes the query and returns the updated DeviceInfo entity.
func (diuo *DeviceInfoUpdateOne) Save(ctx context.Context) (*DeviceInfo, error) {
	var (
		err  error
		node *DeviceInfo
	)
	if err := diuo.defaults(); err != nil {
		return nil, err
	}
	if len(diuo.hooks) == 0 {
		if err = diuo.check(); err != nil {
			return nil, err
		}
		node, err = diuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*DeviceInfoMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = diuo.check(); err != nil {
				return nil, err
			}
			diuo.mutation = mutation
			node, err = diuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(diuo.hooks) - 1; i >= 0; i-- {
			if diuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = diuo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, diuo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*DeviceInfo)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from DeviceInfoMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (diuo *DeviceInfoUpdateOne) SaveX(ctx context.Context) *DeviceInfo {
	node, err := diuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (diuo *DeviceInfoUpdateOne) Exec(ctx context.Context) error {
	_, err := diuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (diuo *DeviceInfoUpdateOne) ExecX(ctx context.Context) {
	if err := diuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (diuo *DeviceInfoUpdateOne) defaults() error {
	if _, ok := diuo.mutation.UpdatedAt(); !ok {
		if deviceinfo.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized deviceinfo.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := deviceinfo.UpdateDefaultUpdatedAt()
		diuo.mutation.SetUpdatedAt(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (diuo *DeviceInfoUpdateOne) check() error {
	if v, ok := diuo.mutation.GetType(); ok {
		if err := deviceinfo.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`ent: validator failed for field "DeviceInfo.type": %w`, err)}
		}
	}
	if v, ok := diuo.mutation.Manufacturer(); ok {
		if err := deviceinfo.ManufacturerValidator(v); err != nil {
			return &ValidationError{Name: "manufacturer", err: fmt.Errorf(`ent: validator failed for field "DeviceInfo.manufacturer": %w`, err)}
		}
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (diuo *DeviceInfoUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *DeviceInfoUpdateOne {
	diuo.modifiers = append(diuo.modifiers, modifiers...)
	return diuo
}

func (diuo *DeviceInfoUpdateOne) sqlSave(ctx context.Context) (_node *DeviceInfo, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   deviceinfo.Table,
			Columns: deviceinfo.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: deviceinfo.FieldID,
			},
		},
	}
	id, ok := diuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "DeviceInfo.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := diuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, deviceinfo.FieldID)
		for _, f := range fields {
			if !deviceinfo.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != deviceinfo.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := diuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := diuo.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: deviceinfo.FieldCreatedAt,
		})
	}
	if value, ok := diuo.mutation.AddedCreatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: deviceinfo.FieldCreatedAt,
		})
	}
	if value, ok := diuo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: deviceinfo.FieldUpdatedAt,
		})
	}
	if value, ok := diuo.mutation.AddedUpdatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: deviceinfo.FieldUpdatedAt,
		})
	}
	if value, ok := diuo.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: deviceinfo.FieldDeletedAt,
		})
	}
	if value, ok := diuo.mutation.AddedDeletedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: deviceinfo.FieldDeletedAt,
		})
	}
	if value, ok := diuo.mutation.EntID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: deviceinfo.FieldEntID,
		})
	}
	if value, ok := diuo.mutation.GetType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: deviceinfo.FieldType,
		})
	}
	if diuo.mutation.TypeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: deviceinfo.FieldType,
		})
	}
	if value, ok := diuo.mutation.Manufacturer(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: deviceinfo.FieldManufacturer,
		})
	}
	if diuo.mutation.ManufacturerCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: deviceinfo.FieldManufacturer,
		})
	}
	if value, ok := diuo.mutation.PowerConsumption(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: deviceinfo.FieldPowerConsumption,
		})
	}
	if value, ok := diuo.mutation.AddedPowerConsumption(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: deviceinfo.FieldPowerConsumption,
		})
	}
	if diuo.mutation.PowerConsumptionCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Column: deviceinfo.FieldPowerConsumption,
		})
	}
	if value, ok := diuo.mutation.ShipmentAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: deviceinfo.FieldShipmentAt,
		})
	}
	if value, ok := diuo.mutation.AddedShipmentAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: deviceinfo.FieldShipmentAt,
		})
	}
	if diuo.mutation.ShipmentAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Column: deviceinfo.FieldShipmentAt,
		})
	}
	if value, ok := diuo.mutation.Posters(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: deviceinfo.FieldPosters,
		})
	}
	if diuo.mutation.PostersCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Column: deviceinfo.FieldPosters,
		})
	}
	_spec.Modifiers = diuo.modifiers
	_node = &DeviceInfo{config: diuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, diuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{deviceinfo.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
