// Copyright (c) Facebook, Inc. and its affiliates. All Rights Reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/facebookincubator/ent/dialect/gremlin"
	"github.com/facebookincubator/ent/dialect/gremlin/graph/dsl"
	"github.com/facebookincubator/ent/dialect/gremlin/graph/dsl/__"
	"github.com/facebookincubator/ent/dialect/gremlin/graph/dsl/g"
	"github.com/facebookincubator/ent/dialect/gremlin/graph/dsl/p"
	"github.com/facebookincubator/ent/entc/integration/gremlin/ent/filetype"
	"github.com/facebookincubator/ent/entc/integration/gremlin/ent/predicate"
)

// FileTypeUpdate is the builder for updating FileType entities.
type FileTypeUpdate struct {
	config
	hooks      []Hook
	mutation   *FileTypeMutation
	predicates []predicate.FileType
}

// Where adds a new predicate for the builder.
func (ftu *FileTypeUpdate) Where(ps ...predicate.FileType) *FileTypeUpdate {
	ftu.predicates = append(ftu.predicates, ps...)
	return ftu
}

// SetName sets the name field.
func (ftu *FileTypeUpdate) SetName(s string) *FileTypeUpdate {
	ftu.mutation.SetName(s)
	return ftu
}

// AddFileIDs adds the files edge to File by ids.
func (ftu *FileTypeUpdate) AddFileIDs(ids ...string) *FileTypeUpdate {
	ftu.mutation.AddFileIDs(ids...)
	return ftu
}

// AddFiles adds the files edges to File.
func (ftu *FileTypeUpdate) AddFiles(f ...*File) *FileTypeUpdate {
	ids := make([]string, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return ftu.AddFileIDs(ids...)
}

// RemoveFileIDs removes the files edge to File by ids.
func (ftu *FileTypeUpdate) RemoveFileIDs(ids ...string) *FileTypeUpdate {
	ftu.mutation.RemoveFileIDs(ids...)
	return ftu
}

// RemoveFiles removes files edges to File.
func (ftu *FileTypeUpdate) RemoveFiles(f ...*File) *FileTypeUpdate {
	ids := make([]string, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return ftu.RemoveFileIDs(ids...)
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (ftu *FileTypeUpdate) Save(ctx context.Context) (int, error) {

	var (
		err      error
		affected int
	)
	if len(ftu.hooks) == 0 {
		affected, err = ftu.gremlinSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*FileTypeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ftu.mutation = mutation
			affected, err = ftu.gremlinSave(ctx)
			return affected, err
		})
		for i := len(ftu.hooks) - 1; i >= 0; i-- {
			mut = ftu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ftu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (ftu *FileTypeUpdate) SaveX(ctx context.Context) int {
	affected, err := ftu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ftu *FileTypeUpdate) Exec(ctx context.Context) error {
	_, err := ftu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ftu *FileTypeUpdate) ExecX(ctx context.Context) {
	if err := ftu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (ftu *FileTypeUpdate) gremlinSave(ctx context.Context) (int, error) {
	res := &gremlin.Response{}
	query, bindings := ftu.gremlin().Query()
	if err := ftu.driver.Exec(ctx, query, bindings, res); err != nil {
		return 0, err
	}
	if err, ok := isConstantError(res); ok {
		return 0, err
	}
	return res.ReadInt()
}

func (ftu *FileTypeUpdate) gremlin() *dsl.Traversal {
	type constraint struct {
		pred *dsl.Traversal // constraint predicate.
		test *dsl.Traversal // test matches and its constant.
	}
	constraints := make([]*constraint, 0, 2)
	v := g.V().HasLabel(filetype.Label)
	for _, p := range ftu.predicates {
		p(v)
	}
	var (
		rv = v.Clone()
		_  = rv

		trs []*dsl.Traversal
	)
	if value, ok := ftu.mutation.Name(); ok {
		constraints = append(constraints, &constraint{
			pred: g.V().Has(filetype.Label, filetype.FieldName, value).Count(),
			test: __.Is(p.NEQ(0)).Constant(NewErrUniqueField(filetype.Label, filetype.FieldName, value)),
		})
		v.Property(dsl.Single, filetype.FieldName, value)
	}
	for _, id := range ftu.mutation.RemovedFilesIDs() {
		tr := rv.Clone().OutE(filetype.FilesLabel).Where(__.OtherV().HasID(id)).Drop().Iterate()
		trs = append(trs, tr)
	}
	for _, id := range ftu.mutation.FilesIDs() {
		v.AddE(filetype.FilesLabel).To(g.V(id)).OutV()
		constraints = append(constraints, &constraint{
			pred: g.E().HasLabel(filetype.FilesLabel).InV().HasID(id).Count(),
			test: __.Is(p.NEQ(0)).Constant(NewErrUniqueEdge(filetype.Label, filetype.FilesLabel, id)),
		})
	}
	v.Count()
	if len(constraints) > 0 {
		constraints = append(constraints, &constraint{
			pred: rv.Count(),
			test: __.Is(p.GT(1)).Constant(&ConstraintError{msg: "update traversal contains more than one vertex"}),
		})
		v = constraints[0].pred.Coalesce(constraints[0].test, v)
		for _, cr := range constraints[1:] {
			v = cr.pred.Coalesce(cr.test, v)
		}
	}
	trs = append(trs, v)
	return dsl.Join(trs...)
}

// FileTypeUpdateOne is the builder for updating a single FileType entity.
type FileTypeUpdateOne struct {
	config
	hooks    []Hook
	mutation *FileTypeMutation
}

// SetName sets the name field.
func (ftuo *FileTypeUpdateOne) SetName(s string) *FileTypeUpdateOne {
	ftuo.mutation.SetName(s)
	return ftuo
}

// AddFileIDs adds the files edge to File by ids.
func (ftuo *FileTypeUpdateOne) AddFileIDs(ids ...string) *FileTypeUpdateOne {
	ftuo.mutation.AddFileIDs(ids...)
	return ftuo
}

// AddFiles adds the files edges to File.
func (ftuo *FileTypeUpdateOne) AddFiles(f ...*File) *FileTypeUpdateOne {
	ids := make([]string, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return ftuo.AddFileIDs(ids...)
}

// RemoveFileIDs removes the files edge to File by ids.
func (ftuo *FileTypeUpdateOne) RemoveFileIDs(ids ...string) *FileTypeUpdateOne {
	ftuo.mutation.RemoveFileIDs(ids...)
	return ftuo
}

// RemoveFiles removes files edges to File.
func (ftuo *FileTypeUpdateOne) RemoveFiles(f ...*File) *FileTypeUpdateOne {
	ids := make([]string, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return ftuo.RemoveFileIDs(ids...)
}

// Save executes the query and returns the updated entity.
func (ftuo *FileTypeUpdateOne) Save(ctx context.Context) (*FileType, error) {

	var (
		err  error
		node *FileType
	)
	if len(ftuo.hooks) == 0 {
		node, err = ftuo.gremlinSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*FileTypeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ftuo.mutation = mutation
			node, err = ftuo.gremlinSave(ctx)
			return node, err
		})
		for i := len(ftuo.hooks) - 1; i >= 0; i-- {
			mut = ftuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ftuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (ftuo *FileTypeUpdateOne) SaveX(ctx context.Context) *FileType {
	ft, err := ftuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return ft
}

// Exec executes the query on the entity.
func (ftuo *FileTypeUpdateOne) Exec(ctx context.Context) error {
	_, err := ftuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ftuo *FileTypeUpdateOne) ExecX(ctx context.Context) {
	if err := ftuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (ftuo *FileTypeUpdateOne) gremlinSave(ctx context.Context) (*FileType, error) {
	res := &gremlin.Response{}
	id, ok := ftuo.mutation.ID()
	if !ok {
		return nil, fmt.Errorf("missing FileType.ID for update")
	}
	query, bindings := ftuo.gremlin(id).Query()
	if err := ftuo.driver.Exec(ctx, query, bindings, res); err != nil {
		return nil, err
	}
	if err, ok := isConstantError(res); ok {
		return nil, err
	}
	ft := &FileType{config: ftuo.config}
	if err := ft.FromResponse(res); err != nil {
		return nil, err
	}
	return ft, nil
}

func (ftuo *FileTypeUpdateOne) gremlin(id string) *dsl.Traversal {
	type constraint struct {
		pred *dsl.Traversal // constraint predicate.
		test *dsl.Traversal // test matches and its constant.
	}
	constraints := make([]*constraint, 0, 2)
	v := g.V(id)
	var (
		rv = v.Clone()
		_  = rv

		trs []*dsl.Traversal
	)
	if value, ok := ftuo.mutation.Name(); ok {
		constraints = append(constraints, &constraint{
			pred: g.V().Has(filetype.Label, filetype.FieldName, value).Count(),
			test: __.Is(p.NEQ(0)).Constant(NewErrUniqueField(filetype.Label, filetype.FieldName, value)),
		})
		v.Property(dsl.Single, filetype.FieldName, value)
	}
	for _, id := range ftuo.mutation.RemovedFilesIDs() {
		tr := rv.Clone().OutE(filetype.FilesLabel).Where(__.OtherV().HasID(id)).Drop().Iterate()
		trs = append(trs, tr)
	}
	for _, id := range ftuo.mutation.FilesIDs() {
		v.AddE(filetype.FilesLabel).To(g.V(id)).OutV()
		constraints = append(constraints, &constraint{
			pred: g.E().HasLabel(filetype.FilesLabel).InV().HasID(id).Count(),
			test: __.Is(p.NEQ(0)).Constant(NewErrUniqueEdge(filetype.Label, filetype.FilesLabel, id)),
		})
	}
	v.ValueMap(true)
	if len(constraints) > 0 {
		v = constraints[0].pred.Coalesce(constraints[0].test, v)
		for _, cr := range constraints[1:] {
			v = cr.pred.Coalesce(cr.test, v)
		}
	}
	trs = append(trs, v)
	return dsl.Join(trs...)
}