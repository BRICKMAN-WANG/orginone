// Code generated by entc, DO NOT EDIT.

package schema

import (
	"context"
	"fmt"
	"orginone/common/schema/asredeploydata"
	"orginone/common/schema/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// AsRedeployDataDelete is the builder for deleting a AsRedeployData entity.
type AsRedeployDataDelete struct {
	config
	hooks    []Hook
	mutation *AsRedeployDataMutation
}

// Where appends a list predicates to the AsRedeployDataDelete builder.
func (ardd *AsRedeployDataDelete) Where(ps ...predicate.AsRedeployData) *AsRedeployDataDelete {
	ardd.mutation.Where(ps...)
	return ardd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (ardd *AsRedeployDataDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(ardd.hooks) == 0 {
		affected, err = ardd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AsRedeployDataMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ardd.mutation = mutation
			affected, err = ardd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ardd.hooks) - 1; i >= 0; i-- {
			if ardd.hooks[i] == nil {
				return 0, fmt.Errorf("schema: uninitialized hook (forgotten import schema/runtime?)")
			}
			mut = ardd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ardd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (ardd *AsRedeployDataDelete) ExecX(ctx context.Context) int {
	n, err := ardd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (ardd *AsRedeployDataDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: asredeploydata.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: asredeploydata.FieldID,
			},
		},
	}
	if ps := ardd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, ardd.driver, _spec)
}

// AsRedeployDataDeleteOne is the builder for deleting a single AsRedeployData entity.
type AsRedeployDataDeleteOne struct {
	ardd *AsRedeployDataDelete
}

// Exec executes the deletion query.
func (arddo *AsRedeployDataDeleteOne) Exec(ctx context.Context) error {
	n, err := arddo.ardd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{asredeploydata.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (arddo *AsRedeployDataDeleteOne) ExecX(ctx context.Context) {
	arddo.ardd.ExecX(ctx)
}
