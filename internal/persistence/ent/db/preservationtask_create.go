// Code generated by ent, DO NOT EDIT.

package db

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/artefactual-sdps/enduro/internal/persistence/ent/db/preservationaction"
	"github.com/artefactual-sdps/enduro/internal/persistence/ent/db/preservationtask"
	"github.com/google/uuid"
)

// PreservationTaskCreate is the builder for creating a PreservationTask entity.
type PreservationTaskCreate struct {
	config
	mutation *PreservationTaskMutation
	hooks    []Hook
}

// SetTaskID sets the "task_id" field.
func (ptc *PreservationTaskCreate) SetTaskID(u uuid.UUID) *PreservationTaskCreate {
	ptc.mutation.SetTaskID(u)
	return ptc
}

// SetName sets the "name" field.
func (ptc *PreservationTaskCreate) SetName(s string) *PreservationTaskCreate {
	ptc.mutation.SetName(s)
	return ptc
}

// SetStatus sets the "status" field.
func (ptc *PreservationTaskCreate) SetStatus(i int8) *PreservationTaskCreate {
	ptc.mutation.SetStatus(i)
	return ptc
}

// SetStartedAt sets the "started_at" field.
func (ptc *PreservationTaskCreate) SetStartedAt(t time.Time) *PreservationTaskCreate {
	ptc.mutation.SetStartedAt(t)
	return ptc
}

// SetCompletedAt sets the "completed_at" field.
func (ptc *PreservationTaskCreate) SetCompletedAt(t time.Time) *PreservationTaskCreate {
	ptc.mutation.SetCompletedAt(t)
	return ptc
}

// SetNote sets the "note" field.
func (ptc *PreservationTaskCreate) SetNote(s string) *PreservationTaskCreate {
	ptc.mutation.SetNote(s)
	return ptc
}

// SetPreservationActionID sets the "preservation_action_id" field.
func (ptc *PreservationTaskCreate) SetPreservationActionID(i int) *PreservationTaskCreate {
	ptc.mutation.SetPreservationActionID(i)
	return ptc
}

// SetActionID sets the "action" edge to the PreservationAction entity by ID.
func (ptc *PreservationTaskCreate) SetActionID(id int) *PreservationTaskCreate {
	ptc.mutation.SetActionID(id)
	return ptc
}

// SetAction sets the "action" edge to the PreservationAction entity.
func (ptc *PreservationTaskCreate) SetAction(p *PreservationAction) *PreservationTaskCreate {
	return ptc.SetActionID(p.ID)
}

// Mutation returns the PreservationTaskMutation object of the builder.
func (ptc *PreservationTaskCreate) Mutation() *PreservationTaskMutation {
	return ptc.mutation
}

// Save creates the PreservationTask in the database.
func (ptc *PreservationTaskCreate) Save(ctx context.Context) (*PreservationTask, error) {
	return withHooks(ctx, ptc.sqlSave, ptc.mutation, ptc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (ptc *PreservationTaskCreate) SaveX(ctx context.Context) *PreservationTask {
	v, err := ptc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ptc *PreservationTaskCreate) Exec(ctx context.Context) error {
	_, err := ptc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ptc *PreservationTaskCreate) ExecX(ctx context.Context) {
	if err := ptc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ptc *PreservationTaskCreate) check() error {
	if _, ok := ptc.mutation.TaskID(); !ok {
		return &ValidationError{Name: "task_id", err: errors.New(`db: missing required field "PreservationTask.task_id"`)}
	}
	if _, ok := ptc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`db: missing required field "PreservationTask.name"`)}
	}
	if _, ok := ptc.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New(`db: missing required field "PreservationTask.status"`)}
	}
	if _, ok := ptc.mutation.StartedAt(); !ok {
		return &ValidationError{Name: "started_at", err: errors.New(`db: missing required field "PreservationTask.started_at"`)}
	}
	if _, ok := ptc.mutation.CompletedAt(); !ok {
		return &ValidationError{Name: "completed_at", err: errors.New(`db: missing required field "PreservationTask.completed_at"`)}
	}
	if _, ok := ptc.mutation.Note(); !ok {
		return &ValidationError{Name: "note", err: errors.New(`db: missing required field "PreservationTask.note"`)}
	}
	if _, ok := ptc.mutation.PreservationActionID(); !ok {
		return &ValidationError{Name: "preservation_action_id", err: errors.New(`db: missing required field "PreservationTask.preservation_action_id"`)}
	}
	if v, ok := ptc.mutation.PreservationActionID(); ok {
		if err := preservationtask.PreservationActionIDValidator(v); err != nil {
			return &ValidationError{Name: "preservation_action_id", err: fmt.Errorf(`db: validator failed for field "PreservationTask.preservation_action_id": %w`, err)}
		}
	}
	if _, ok := ptc.mutation.ActionID(); !ok {
		return &ValidationError{Name: "action", err: errors.New(`db: missing required edge "PreservationTask.action"`)}
	}
	return nil
}

func (ptc *PreservationTaskCreate) sqlSave(ctx context.Context) (*PreservationTask, error) {
	if err := ptc.check(); err != nil {
		return nil, err
	}
	_node, _spec := ptc.createSpec()
	if err := sqlgraph.CreateNode(ctx, ptc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	ptc.mutation.id = &_node.ID
	ptc.mutation.done = true
	return _node, nil
}

func (ptc *PreservationTaskCreate) createSpec() (*PreservationTask, *sqlgraph.CreateSpec) {
	var (
		_node = &PreservationTask{config: ptc.config}
		_spec = sqlgraph.NewCreateSpec(preservationtask.Table, sqlgraph.NewFieldSpec(preservationtask.FieldID, field.TypeInt))
	)
	if value, ok := ptc.mutation.TaskID(); ok {
		_spec.SetField(preservationtask.FieldTaskID, field.TypeUUID, value)
		_node.TaskID = value
	}
	if value, ok := ptc.mutation.Name(); ok {
		_spec.SetField(preservationtask.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := ptc.mutation.Status(); ok {
		_spec.SetField(preservationtask.FieldStatus, field.TypeInt8, value)
		_node.Status = value
	}
	if value, ok := ptc.mutation.StartedAt(); ok {
		_spec.SetField(preservationtask.FieldStartedAt, field.TypeTime, value)
		_node.StartedAt = value
	}
	if value, ok := ptc.mutation.CompletedAt(); ok {
		_spec.SetField(preservationtask.FieldCompletedAt, field.TypeTime, value)
		_node.CompletedAt = value
	}
	if value, ok := ptc.mutation.Note(); ok {
		_spec.SetField(preservationtask.FieldNote, field.TypeString, value)
		_node.Note = value
	}
	if nodes := ptc.mutation.ActionIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   preservationtask.ActionTable,
			Columns: []string{preservationtask.ActionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(preservationaction.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.PreservationActionID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// PreservationTaskCreateBulk is the builder for creating many PreservationTask entities in bulk.
type PreservationTaskCreateBulk struct {
	config
	builders []*PreservationTaskCreate
}

// Save creates the PreservationTask entities in the database.
func (ptcb *PreservationTaskCreateBulk) Save(ctx context.Context) ([]*PreservationTask, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ptcb.builders))
	nodes := make([]*PreservationTask, len(ptcb.builders))
	mutators := make([]Mutator, len(ptcb.builders))
	for i := range ptcb.builders {
		func(i int, root context.Context) {
			builder := ptcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*PreservationTaskMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, ptcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ptcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, ptcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ptcb *PreservationTaskCreateBulk) SaveX(ctx context.Context) []*PreservationTask {
	v, err := ptcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ptcb *PreservationTaskCreateBulk) Exec(ctx context.Context) error {
	_, err := ptcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ptcb *PreservationTaskCreateBulk) ExecX(ctx context.Context) {
	if err := ptcb.Exec(ctx); err != nil {
		panic(err)
	}
}
