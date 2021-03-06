// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"jerrybook/ent/book"
	"jerrybook/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// BookUpdate is the builder for updating Book entities.
type BookUpdate struct {
	config
	hooks    []Hook
	mutation *BookMutation
}

// Where appends a list predicates to the BookUpdate builder.
func (bu *BookUpdate) Where(ps ...predicate.Book) *BookUpdate {
	bu.mutation.Where(ps...)
	return bu
}

// SetUUID sets the "uuid" field.
func (bu *BookUpdate) SetUUID(u uuid.UUID) *BookUpdate {
	bu.mutation.SetUUID(u)
	return bu
}

// SetCreatedAt sets the "created_at" field.
func (bu *BookUpdate) SetCreatedAt(t time.Time) *BookUpdate {
	bu.mutation.SetCreatedAt(t)
	return bu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (bu *BookUpdate) SetNillableCreatedAt(t *time.Time) *BookUpdate {
	if t != nil {
		bu.SetCreatedAt(*t)
	}
	return bu
}

// SetUpdatedAt sets the "updated_at" field.
func (bu *BookUpdate) SetUpdatedAt(t time.Time) *BookUpdate {
	bu.mutation.SetUpdatedAt(t)
	return bu
}

// SetName sets the "name" field.
func (bu *BookUpdate) SetName(s string) *BookUpdate {
	bu.mutation.SetName(s)
	return bu
}

// SetBackgroundColor sets the "background_color" field.
func (bu *BookUpdate) SetBackgroundColor(s string) *BookUpdate {
	bu.mutation.SetBackgroundColor(s)
	return bu
}

// SetNillableBackgroundColor sets the "background_color" field if the given value is not nil.
func (bu *BookUpdate) SetNillableBackgroundColor(s *string) *BookUpdate {
	if s != nil {
		bu.SetBackgroundColor(*s)
	}
	return bu
}

// SetAuthor sets the "author" field.
func (bu *BookUpdate) SetAuthor(s string) *BookUpdate {
	bu.mutation.SetAuthor(s)
	return bu
}

// SetImageURL sets the "image_url" field.
func (bu *BookUpdate) SetImageURL(s string) *BookUpdate {
	bu.mutation.SetImageURL(s)
	return bu
}

// SetNillableImageURL sets the "image_url" field if the given value is not nil.
func (bu *BookUpdate) SetNillableImageURL(s *string) *BookUpdate {
	if s != nil {
		bu.SetImageURL(*s)
	}
	return bu
}

// ClearImageURL clears the value of the "image_url" field.
func (bu *BookUpdate) ClearImageURL() *BookUpdate {
	bu.mutation.ClearImageURL()
	return bu
}

// SetReport sets the "report" field.
func (bu *BookUpdate) SetReport(s string) *BookUpdate {
	bu.mutation.SetReport(s)
	return bu
}

// SetRating sets the "rating" field.
func (bu *BookUpdate) SetRating(f float64) *BookUpdate {
	bu.mutation.ResetRating()
	bu.mutation.SetRating(f)
	return bu
}

// AddRating adds f to the "rating" field.
func (bu *BookUpdate) AddRating(f float64) *BookUpdate {
	bu.mutation.AddRating(f)
	return bu
}

// Mutation returns the BookMutation object of the builder.
func (bu *BookUpdate) Mutation() *BookMutation {
	return bu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (bu *BookUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	bu.defaults()
	if len(bu.hooks) == 0 {
		affected, err = bu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*BookMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			bu.mutation = mutation
			affected, err = bu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(bu.hooks) - 1; i >= 0; i-- {
			if bu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = bu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, bu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (bu *BookUpdate) SaveX(ctx context.Context) int {
	affected, err := bu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (bu *BookUpdate) Exec(ctx context.Context) error {
	_, err := bu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bu *BookUpdate) ExecX(ctx context.Context) {
	if err := bu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (bu *BookUpdate) defaults() {
	if _, ok := bu.mutation.UpdatedAt(); !ok {
		v := book.UpdateDefaultUpdatedAt()
		bu.mutation.SetUpdatedAt(v)
	}
}

func (bu *BookUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   book.Table,
			Columns: book.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: book.FieldID,
			},
		},
	}
	if ps := bu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := bu.mutation.UUID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: book.FieldUUID,
		})
	}
	if value, ok := bu.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: book.FieldCreatedAt,
		})
	}
	if value, ok := bu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: book.FieldUpdatedAt,
		})
	}
	if value, ok := bu.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: book.FieldName,
		})
	}
	if value, ok := bu.mutation.BackgroundColor(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: book.FieldBackgroundColor,
		})
	}
	if value, ok := bu.mutation.Author(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: book.FieldAuthor,
		})
	}
	if value, ok := bu.mutation.ImageURL(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: book.FieldImageURL,
		})
	}
	if bu.mutation.ImageURLCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: book.FieldImageURL,
		})
	}
	if value, ok := bu.mutation.Report(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: book.FieldReport,
		})
	}
	if value, ok := bu.mutation.Rating(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: book.FieldRating,
		})
	}
	if value, ok := bu.mutation.AddedRating(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: book.FieldRating,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, bu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{book.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// BookUpdateOne is the builder for updating a single Book entity.
type BookUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *BookMutation
}

// SetUUID sets the "uuid" field.
func (buo *BookUpdateOne) SetUUID(u uuid.UUID) *BookUpdateOne {
	buo.mutation.SetUUID(u)
	return buo
}

// SetCreatedAt sets the "created_at" field.
func (buo *BookUpdateOne) SetCreatedAt(t time.Time) *BookUpdateOne {
	buo.mutation.SetCreatedAt(t)
	return buo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (buo *BookUpdateOne) SetNillableCreatedAt(t *time.Time) *BookUpdateOne {
	if t != nil {
		buo.SetCreatedAt(*t)
	}
	return buo
}

// SetUpdatedAt sets the "updated_at" field.
func (buo *BookUpdateOne) SetUpdatedAt(t time.Time) *BookUpdateOne {
	buo.mutation.SetUpdatedAt(t)
	return buo
}

// SetName sets the "name" field.
func (buo *BookUpdateOne) SetName(s string) *BookUpdateOne {
	buo.mutation.SetName(s)
	return buo
}

// SetBackgroundColor sets the "background_color" field.
func (buo *BookUpdateOne) SetBackgroundColor(s string) *BookUpdateOne {
	buo.mutation.SetBackgroundColor(s)
	return buo
}

// SetNillableBackgroundColor sets the "background_color" field if the given value is not nil.
func (buo *BookUpdateOne) SetNillableBackgroundColor(s *string) *BookUpdateOne {
	if s != nil {
		buo.SetBackgroundColor(*s)
	}
	return buo
}

// SetAuthor sets the "author" field.
func (buo *BookUpdateOne) SetAuthor(s string) *BookUpdateOne {
	buo.mutation.SetAuthor(s)
	return buo
}

// SetImageURL sets the "image_url" field.
func (buo *BookUpdateOne) SetImageURL(s string) *BookUpdateOne {
	buo.mutation.SetImageURL(s)
	return buo
}

// SetNillableImageURL sets the "image_url" field if the given value is not nil.
func (buo *BookUpdateOne) SetNillableImageURL(s *string) *BookUpdateOne {
	if s != nil {
		buo.SetImageURL(*s)
	}
	return buo
}

// ClearImageURL clears the value of the "image_url" field.
func (buo *BookUpdateOne) ClearImageURL() *BookUpdateOne {
	buo.mutation.ClearImageURL()
	return buo
}

// SetReport sets the "report" field.
func (buo *BookUpdateOne) SetReport(s string) *BookUpdateOne {
	buo.mutation.SetReport(s)
	return buo
}

// SetRating sets the "rating" field.
func (buo *BookUpdateOne) SetRating(f float64) *BookUpdateOne {
	buo.mutation.ResetRating()
	buo.mutation.SetRating(f)
	return buo
}

// AddRating adds f to the "rating" field.
func (buo *BookUpdateOne) AddRating(f float64) *BookUpdateOne {
	buo.mutation.AddRating(f)
	return buo
}

// Mutation returns the BookMutation object of the builder.
func (buo *BookUpdateOne) Mutation() *BookMutation {
	return buo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (buo *BookUpdateOne) Select(field string, fields ...string) *BookUpdateOne {
	buo.fields = append([]string{field}, fields...)
	return buo
}

// Save executes the query and returns the updated Book entity.
func (buo *BookUpdateOne) Save(ctx context.Context) (*Book, error) {
	var (
		err  error
		node *Book
	)
	buo.defaults()
	if len(buo.hooks) == 0 {
		node, err = buo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*BookMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			buo.mutation = mutation
			node, err = buo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(buo.hooks) - 1; i >= 0; i-- {
			if buo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = buo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, buo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (buo *BookUpdateOne) SaveX(ctx context.Context) *Book {
	node, err := buo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (buo *BookUpdateOne) Exec(ctx context.Context) error {
	_, err := buo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (buo *BookUpdateOne) ExecX(ctx context.Context) {
	if err := buo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (buo *BookUpdateOne) defaults() {
	if _, ok := buo.mutation.UpdatedAt(); !ok {
		v := book.UpdateDefaultUpdatedAt()
		buo.mutation.SetUpdatedAt(v)
	}
}

func (buo *BookUpdateOne) sqlSave(ctx context.Context) (_node *Book, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   book.Table,
			Columns: book.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: book.FieldID,
			},
		},
	}
	id, ok := buo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Book.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := buo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, book.FieldID)
		for _, f := range fields {
			if !book.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != book.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := buo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := buo.mutation.UUID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: book.FieldUUID,
		})
	}
	if value, ok := buo.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: book.FieldCreatedAt,
		})
	}
	if value, ok := buo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: book.FieldUpdatedAt,
		})
	}
	if value, ok := buo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: book.FieldName,
		})
	}
	if value, ok := buo.mutation.BackgroundColor(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: book.FieldBackgroundColor,
		})
	}
	if value, ok := buo.mutation.Author(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: book.FieldAuthor,
		})
	}
	if value, ok := buo.mutation.ImageURL(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: book.FieldImageURL,
		})
	}
	if buo.mutation.ImageURLCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: book.FieldImageURL,
		})
	}
	if value, ok := buo.mutation.Report(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: book.FieldReport,
		})
	}
	if value, ok := buo.mutation.Rating(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: book.FieldRating,
		})
	}
	if value, ok := buo.mutation.AddedRating(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: book.FieldRating,
		})
	}
	_node = &Book{config: buo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, buo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{book.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
