// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/hay-kot/homebox/backend/internal/data/ent/document"
	"github.com/hay-kot/homebox/backend/internal/data/ent/group"
)

// Document is the model entity for the Document schema.
type Document struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Title holds the value of the "title" field.
	Title string `json:"title,omitempty"`
	// Path holds the value of the "path" field.
	Path string `json:"path,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the DocumentQuery when eager-loading is set.
	Edges           DocumentEdges `json:"edges"`
	group_documents *uuid.UUID
}

// DocumentEdges holds the relations/edges for other nodes in the graph.
type DocumentEdges struct {
	// Group holds the value of the group edge.
	Group *Group `json:"group,omitempty"`
	// Attachments holds the value of the attachments edge.
	Attachments []*Attachment `json:"attachments,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// GroupOrErr returns the Group value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e DocumentEdges) GroupOrErr() (*Group, error) {
	if e.loadedTypes[0] {
		if e.Group == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: group.Label}
		}
		return e.Group, nil
	}
	return nil, &NotLoadedError{edge: "group"}
}

// AttachmentsOrErr returns the Attachments value or an error if the edge
// was not loaded in eager-loading.
func (e DocumentEdges) AttachmentsOrErr() ([]*Attachment, error) {
	if e.loadedTypes[1] {
		return e.Attachments, nil
	}
	return nil, &NotLoadedError{edge: "attachments"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Document) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case document.FieldTitle, document.FieldPath:
			values[i] = new(sql.NullString)
		case document.FieldCreatedAt, document.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case document.FieldID:
			values[i] = new(uuid.UUID)
		case document.ForeignKeys[0]: // group_documents
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		default:
			return nil, fmt.Errorf("unexpected column %q for type Document", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Document fields.
func (d *Document) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case document.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				d.ID = *value
			}
		case document.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				d.CreatedAt = value.Time
			}
		case document.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				d.UpdatedAt = value.Time
			}
		case document.FieldTitle:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field title", values[i])
			} else if value.Valid {
				d.Title = value.String
			}
		case document.FieldPath:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field path", values[i])
			} else if value.Valid {
				d.Path = value.String
			}
		case document.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field group_documents", values[i])
			} else if value.Valid {
				d.group_documents = new(uuid.UUID)
				*d.group_documents = *value.S.(*uuid.UUID)
			}
		}
	}
	return nil
}

// QueryGroup queries the "group" edge of the Document entity.
func (d *Document) QueryGroup() *GroupQuery {
	return (&DocumentClient{config: d.config}).QueryGroup(d)
}

// QueryAttachments queries the "attachments" edge of the Document entity.
func (d *Document) QueryAttachments() *AttachmentQuery {
	return (&DocumentClient{config: d.config}).QueryAttachments(d)
}

// Update returns a builder for updating this Document.
// Note that you need to call Document.Unwrap() before calling this method if this Document
// was returned from a transaction, and the transaction was committed or rolled back.
func (d *Document) Update() *DocumentUpdateOne {
	return (&DocumentClient{config: d.config}).UpdateOne(d)
}

// Unwrap unwraps the Document entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (d *Document) Unwrap() *Document {
	_tx, ok := d.config.driver.(*txDriver)
	if !ok {
		panic("ent: Document is not a transactional entity")
	}
	d.config.driver = _tx.drv
	return d
}

// String implements the fmt.Stringer.
func (d *Document) String() string {
	var builder strings.Builder
	builder.WriteString("Document(")
	builder.WriteString(fmt.Sprintf("id=%v, ", d.ID))
	builder.WriteString("created_at=")
	builder.WriteString(d.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(d.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("title=")
	builder.WriteString(d.Title)
	builder.WriteString(", ")
	builder.WriteString("path=")
	builder.WriteString(d.Path)
	builder.WriteByte(')')
	return builder.String()
}

// Documents is a parsable slice of Document.
type Documents []*Document

func (d Documents) config(cfg config) {
	for _i := range d {
		d[_i].config = cfg
	}
}
