// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"enttest/ent/hashtag"
	"enttest/ent/hashtagtopost"
	"enttest/ent/post"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// HashtagToPostCreate is the builder for creating a HashtagToPost entity.
type HashtagToPostCreate struct {
	config
	mutation *HashtagToPostMutation
	hooks    []Hook
}

// SetHashtagID sets the "hashtag_id" field.
func (htpc *HashtagToPostCreate) SetHashtagID(s string) *HashtagToPostCreate {
	htpc.mutation.SetHashtagID(s)
	return htpc
}

// SetPostID sets the "post_id" field.
func (htpc *HashtagToPostCreate) SetPostID(s string) *HashtagToPostCreate {
	htpc.mutation.SetPostID(s)
	return htpc
}

// SetID sets the "id" field.
func (htpc *HashtagToPostCreate) SetID(s string) *HashtagToPostCreate {
	htpc.mutation.SetID(s)
	return htpc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (htpc *HashtagToPostCreate) SetNillableID(s *string) *HashtagToPostCreate {
	if s != nil {
		htpc.SetID(*s)
	}
	return htpc
}

// SetPost sets the "post" edge to the Post entity.
func (htpc *HashtagToPostCreate) SetPost(p *Post) *HashtagToPostCreate {
	return htpc.SetPostID(p.ID)
}

// SetHashtag sets the "hashtag" edge to the Hashtag entity.
func (htpc *HashtagToPostCreate) SetHashtag(h *Hashtag) *HashtagToPostCreate {
	return htpc.SetHashtagID(h.ID)
}

// Mutation returns the HashtagToPostMutation object of the builder.
func (htpc *HashtagToPostCreate) Mutation() *HashtagToPostMutation {
	return htpc.mutation
}

// Save creates the HashtagToPost in the database.
func (htpc *HashtagToPostCreate) Save(ctx context.Context) (*HashtagToPost, error) {
	var (
		err  error
		node *HashtagToPost
	)
	htpc.defaults()
	if len(htpc.hooks) == 0 {
		if err = htpc.check(); err != nil {
			return nil, err
		}
		node, err = htpc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*HashtagToPostMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = htpc.check(); err != nil {
				return nil, err
			}
			htpc.mutation = mutation
			if node, err = htpc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(htpc.hooks) - 1; i >= 0; i-- {
			if htpc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = htpc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, htpc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*HashtagToPost)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from HashtagToPostMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (htpc *HashtagToPostCreate) SaveX(ctx context.Context) *HashtagToPost {
	v, err := htpc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (htpc *HashtagToPostCreate) Exec(ctx context.Context) error {
	_, err := htpc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (htpc *HashtagToPostCreate) ExecX(ctx context.Context) {
	if err := htpc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (htpc *HashtagToPostCreate) defaults() {
	if _, ok := htpc.mutation.ID(); !ok {
		v := hashtagtopost.DefaultID()
		htpc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (htpc *HashtagToPostCreate) check() error {
	if _, ok := htpc.mutation.HashtagID(); !ok {
		return &ValidationError{Name: "hashtag_id", err: errors.New(`ent: missing required field "HashtagToPost.hashtag_id"`)}
	}
	if v, ok := htpc.mutation.HashtagID(); ok {
		if err := hashtagtopost.HashtagIDValidator(v); err != nil {
			return &ValidationError{Name: "hashtag_id", err: fmt.Errorf(`ent: validator failed for field "HashtagToPost.hashtag_id": %w`, err)}
		}
	}
	if _, ok := htpc.mutation.PostID(); !ok {
		return &ValidationError{Name: "post_id", err: errors.New(`ent: missing required field "HashtagToPost.post_id"`)}
	}
	if v, ok := htpc.mutation.PostID(); ok {
		if err := hashtagtopost.PostIDValidator(v); err != nil {
			return &ValidationError{Name: "post_id", err: fmt.Errorf(`ent: validator failed for field "HashtagToPost.post_id": %w`, err)}
		}
	}
	if v, ok := htpc.mutation.ID(); ok {
		if err := hashtagtopost.IDValidator(v); err != nil {
			return &ValidationError{Name: "id", err: fmt.Errorf(`ent: validator failed for field "HashtagToPost.id": %w`, err)}
		}
	}
	if _, ok := htpc.mutation.PostID(); !ok {
		return &ValidationError{Name: "post", err: errors.New(`ent: missing required edge "HashtagToPost.post"`)}
	}
	if _, ok := htpc.mutation.HashtagID(); !ok {
		return &ValidationError{Name: "hashtag", err: errors.New(`ent: missing required edge "HashtagToPost.hashtag"`)}
	}
	return nil
}

func (htpc *HashtagToPostCreate) sqlSave(ctx context.Context) (*HashtagToPost, error) {
	_node, _spec := htpc.createSpec()
	if err := sqlgraph.CreateNode(ctx, htpc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected HashtagToPost.ID type: %T", _spec.ID.Value)
		}
	}
	return _node, nil
}

func (htpc *HashtagToPostCreate) createSpec() (*HashtagToPost, *sqlgraph.CreateSpec) {
	var (
		_node = &HashtagToPost{config: htpc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: hashtagtopost.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: hashtagtopost.FieldID,
			},
		}
	)
	if id, ok := htpc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if nodes := htpc.mutation.PostIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   hashtagtopost.PostTable,
			Columns: []string{hashtagtopost.PostColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: post.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.PostID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := htpc.mutation.HashtagIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   hashtagtopost.HashtagTable,
			Columns: []string{hashtagtopost.HashtagColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: hashtag.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.HashtagID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// HashtagToPostCreateBulk is the builder for creating many HashtagToPost entities in bulk.
type HashtagToPostCreateBulk struct {
	config
	builders []*HashtagToPostCreate
}

// Save creates the HashtagToPost entities in the database.
func (htpcb *HashtagToPostCreateBulk) Save(ctx context.Context) ([]*HashtagToPost, error) {
	specs := make([]*sqlgraph.CreateSpec, len(htpcb.builders))
	nodes := make([]*HashtagToPost, len(htpcb.builders))
	mutators := make([]Mutator, len(htpcb.builders))
	for i := range htpcb.builders {
		func(i int, root context.Context) {
			builder := htpcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*HashtagToPostMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, htpcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, htpcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
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
		if _, err := mutators[0].Mutate(ctx, htpcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (htpcb *HashtagToPostCreateBulk) SaveX(ctx context.Context) []*HashtagToPost {
	v, err := htpcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (htpcb *HashtagToPostCreateBulk) Exec(ctx context.Context) error {
	_, err := htpcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (htpcb *HashtagToPostCreateBulk) ExecX(ctx context.Context) {
	if err := htpcb.Exec(ctx); err != nil {
		panic(err)
	}
}