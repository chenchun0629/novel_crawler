// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/novel_crawler/internal/data/ent/novelclue"
	"github.com/novel_crawler/internal/data/ent/predicate"
)

// NovelClueUpdate is the builder for updating NovelClue entities.
type NovelClueUpdate struct {
	config
	hooks    []Hook
	mutation *NovelClueMutation
}

// Where appends a list predicates to the NovelClueUpdate builder.
func (ncu *NovelClueUpdate) Where(ps ...predicate.NovelClue) *NovelClueUpdate {
	ncu.mutation.Where(ps...)
	return ncu
}

// SetTitle sets the "title" field.
func (ncu *NovelClueUpdate) SetTitle(s string) *NovelClueUpdate {
	ncu.mutation.SetTitle(s)
	return ncu
}

// SetDate sets the "date" field.
func (ncu *NovelClueUpdate) SetDate(t time.Time) *NovelClueUpdate {
	ncu.mutation.SetDate(t)
	return ncu
}

// SetScore sets the "score" field.
func (ncu *NovelClueUpdate) SetScore(i int) *NovelClueUpdate {
	ncu.mutation.ResetScore()
	ncu.mutation.SetScore(i)
	return ncu
}

// AddScore adds i to the "score" field.
func (ncu *NovelClueUpdate) AddScore(i int) *NovelClueUpdate {
	ncu.mutation.AddScore(i)
	return ncu
}

// SetAuthor sets the "author" field.
func (ncu *NovelClueUpdate) SetAuthor(s string) *NovelClueUpdate {
	ncu.mutation.SetAuthor(s)
	return ncu
}

// SetCategoryTitle sets the "category_title" field.
func (ncu *NovelClueUpdate) SetCategoryTitle(s string) *NovelClueUpdate {
	ncu.mutation.SetCategoryTitle(s)
	return ncu
}

// SetIntro sets the "intro" field.
func (ncu *NovelClueUpdate) SetIntro(s string) *NovelClueUpdate {
	ncu.mutation.SetIntro(s)
	return ncu
}

// SetLink sets the "link" field.
func (ncu *NovelClueUpdate) SetLink(s string) *NovelClueUpdate {
	ncu.mutation.SetLink(s)
	return ncu
}

// Mutation returns the NovelClueMutation object of the builder.
func (ncu *NovelClueUpdate) Mutation() *NovelClueMutation {
	return ncu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ncu *NovelClueUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	ncu.defaults()
	if len(ncu.hooks) == 0 {
		affected, err = ncu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*NovelClueMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ncu.mutation = mutation
			affected, err = ncu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ncu.hooks) - 1; i >= 0; i-- {
			if ncu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ncu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ncu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (ncu *NovelClueUpdate) SaveX(ctx context.Context) int {
	affected, err := ncu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ncu *NovelClueUpdate) Exec(ctx context.Context) error {
	_, err := ncu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ncu *NovelClueUpdate) ExecX(ctx context.Context) {
	if err := ncu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ncu *NovelClueUpdate) defaults() {
	if _, ok := ncu.mutation.UpdateTime(); !ok {
		v := novelclue.UpdateDefaultUpdateTime()
		ncu.mutation.SetUpdateTime(v)
	}
}

func (ncu *NovelClueUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   novelclue.Table,
			Columns: novelclue.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: novelclue.FieldID,
			},
		},
	}
	if ps := ncu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ncu.mutation.UpdateTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: novelclue.FieldUpdateTime,
		})
	}
	if value, ok := ncu.mutation.Title(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: novelclue.FieldTitle,
		})
	}
	if value, ok := ncu.mutation.Date(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: novelclue.FieldDate,
		})
	}
	if value, ok := ncu.mutation.Score(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: novelclue.FieldScore,
		})
	}
	if value, ok := ncu.mutation.AddedScore(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: novelclue.FieldScore,
		})
	}
	if value, ok := ncu.mutation.Author(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: novelclue.FieldAuthor,
		})
	}
	if value, ok := ncu.mutation.CategoryTitle(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: novelclue.FieldCategoryTitle,
		})
	}
	if value, ok := ncu.mutation.Intro(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: novelclue.FieldIntro,
		})
	}
	if value, ok := ncu.mutation.Link(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: novelclue.FieldLink,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ncu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{novelclue.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// NovelClueUpdateOne is the builder for updating a single NovelClue entity.
type NovelClueUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *NovelClueMutation
}

// SetTitle sets the "title" field.
func (ncuo *NovelClueUpdateOne) SetTitle(s string) *NovelClueUpdateOne {
	ncuo.mutation.SetTitle(s)
	return ncuo
}

// SetDate sets the "date" field.
func (ncuo *NovelClueUpdateOne) SetDate(t time.Time) *NovelClueUpdateOne {
	ncuo.mutation.SetDate(t)
	return ncuo
}

// SetScore sets the "score" field.
func (ncuo *NovelClueUpdateOne) SetScore(i int) *NovelClueUpdateOne {
	ncuo.mutation.ResetScore()
	ncuo.mutation.SetScore(i)
	return ncuo
}

// AddScore adds i to the "score" field.
func (ncuo *NovelClueUpdateOne) AddScore(i int) *NovelClueUpdateOne {
	ncuo.mutation.AddScore(i)
	return ncuo
}

// SetAuthor sets the "author" field.
func (ncuo *NovelClueUpdateOne) SetAuthor(s string) *NovelClueUpdateOne {
	ncuo.mutation.SetAuthor(s)
	return ncuo
}

// SetCategoryTitle sets the "category_title" field.
func (ncuo *NovelClueUpdateOne) SetCategoryTitle(s string) *NovelClueUpdateOne {
	ncuo.mutation.SetCategoryTitle(s)
	return ncuo
}

// SetIntro sets the "intro" field.
func (ncuo *NovelClueUpdateOne) SetIntro(s string) *NovelClueUpdateOne {
	ncuo.mutation.SetIntro(s)
	return ncuo
}

// SetLink sets the "link" field.
func (ncuo *NovelClueUpdateOne) SetLink(s string) *NovelClueUpdateOne {
	ncuo.mutation.SetLink(s)
	return ncuo
}

// Mutation returns the NovelClueMutation object of the builder.
func (ncuo *NovelClueUpdateOne) Mutation() *NovelClueMutation {
	return ncuo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ncuo *NovelClueUpdateOne) Select(field string, fields ...string) *NovelClueUpdateOne {
	ncuo.fields = append([]string{field}, fields...)
	return ncuo
}

// Save executes the query and returns the updated NovelClue entity.
func (ncuo *NovelClueUpdateOne) Save(ctx context.Context) (*NovelClue, error) {
	var (
		err  error
		node *NovelClue
	)
	ncuo.defaults()
	if len(ncuo.hooks) == 0 {
		node, err = ncuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*NovelClueMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ncuo.mutation = mutation
			node, err = ncuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(ncuo.hooks) - 1; i >= 0; i-- {
			if ncuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ncuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ncuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (ncuo *NovelClueUpdateOne) SaveX(ctx context.Context) *NovelClue {
	node, err := ncuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ncuo *NovelClueUpdateOne) Exec(ctx context.Context) error {
	_, err := ncuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ncuo *NovelClueUpdateOne) ExecX(ctx context.Context) {
	if err := ncuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ncuo *NovelClueUpdateOne) defaults() {
	if _, ok := ncuo.mutation.UpdateTime(); !ok {
		v := novelclue.UpdateDefaultUpdateTime()
		ncuo.mutation.SetUpdateTime(v)
	}
}

func (ncuo *NovelClueUpdateOne) sqlSave(ctx context.Context) (_node *NovelClue, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   novelclue.Table,
			Columns: novelclue.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: novelclue.FieldID,
			},
		},
	}
	id, ok := ncuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing NovelClue.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := ncuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, novelclue.FieldID)
		for _, f := range fields {
			if !novelclue.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != novelclue.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ncuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ncuo.mutation.UpdateTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: novelclue.FieldUpdateTime,
		})
	}
	if value, ok := ncuo.mutation.Title(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: novelclue.FieldTitle,
		})
	}
	if value, ok := ncuo.mutation.Date(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: novelclue.FieldDate,
		})
	}
	if value, ok := ncuo.mutation.Score(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: novelclue.FieldScore,
		})
	}
	if value, ok := ncuo.mutation.AddedScore(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: novelclue.FieldScore,
		})
	}
	if value, ok := ncuo.mutation.Author(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: novelclue.FieldAuthor,
		})
	}
	if value, ok := ncuo.mutation.CategoryTitle(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: novelclue.FieldCategoryTitle,
		})
	}
	if value, ok := ncuo.mutation.Intro(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: novelclue.FieldIntro,
		})
	}
	if value, ok := ncuo.mutation.Link(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: novelclue.FieldLink,
		})
	}
	_node = &NovelClue{config: ncuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ncuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{novelclue.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
