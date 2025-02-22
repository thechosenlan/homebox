// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/hay-kot/homebox/backend/internal/data/ent/group"
	"github.com/hay-kot/homebox/backend/internal/data/ent/groupinvitationtoken"
)

// GroupInvitationToken is the model entity for the GroupInvitationToken schema.
type GroupInvitationToken struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Token holds the value of the "token" field.
	Token []byte `json:"token,omitempty"`
	// ExpiresAt holds the value of the "expires_at" field.
	ExpiresAt time.Time `json:"expires_at,omitempty"`
	// Uses holds the value of the "uses" field.
	Uses int `json:"uses,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the GroupInvitationTokenQuery when eager-loading is set.
	Edges                   GroupInvitationTokenEdges `json:"edges"`
	group_invitation_tokens *uuid.UUID
}

// GroupInvitationTokenEdges holds the relations/edges for other nodes in the graph.
type GroupInvitationTokenEdges struct {
	// Group holds the value of the group edge.
	Group *Group `json:"group,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// GroupOrErr returns the Group value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e GroupInvitationTokenEdges) GroupOrErr() (*Group, error) {
	if e.loadedTypes[0] {
		if e.Group == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: group.Label}
		}
		return e.Group, nil
	}
	return nil, &NotLoadedError{edge: "group"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*GroupInvitationToken) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case groupinvitationtoken.FieldToken:
			values[i] = new([]byte)
		case groupinvitationtoken.FieldUses:
			values[i] = new(sql.NullInt64)
		case groupinvitationtoken.FieldCreatedAt, groupinvitationtoken.FieldUpdatedAt, groupinvitationtoken.FieldExpiresAt:
			values[i] = new(sql.NullTime)
		case groupinvitationtoken.FieldID:
			values[i] = new(uuid.UUID)
		case groupinvitationtoken.ForeignKeys[0]: // group_invitation_tokens
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		default:
			return nil, fmt.Errorf("unexpected column %q for type GroupInvitationToken", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the GroupInvitationToken fields.
func (git *GroupInvitationToken) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case groupinvitationtoken.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				git.ID = *value
			}
		case groupinvitationtoken.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				git.CreatedAt = value.Time
			}
		case groupinvitationtoken.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				git.UpdatedAt = value.Time
			}
		case groupinvitationtoken.FieldToken:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field token", values[i])
			} else if value != nil {
				git.Token = *value
			}
		case groupinvitationtoken.FieldExpiresAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field expires_at", values[i])
			} else if value.Valid {
				git.ExpiresAt = value.Time
			}
		case groupinvitationtoken.FieldUses:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field uses", values[i])
			} else if value.Valid {
				git.Uses = int(value.Int64)
			}
		case groupinvitationtoken.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field group_invitation_tokens", values[i])
			} else if value.Valid {
				git.group_invitation_tokens = new(uuid.UUID)
				*git.group_invitation_tokens = *value.S.(*uuid.UUID)
			}
		}
	}
	return nil
}

// QueryGroup queries the "group" edge of the GroupInvitationToken entity.
func (git *GroupInvitationToken) QueryGroup() *GroupQuery {
	return NewGroupInvitationTokenClient(git.config).QueryGroup(git)
}

// Update returns a builder for updating this GroupInvitationToken.
// Note that you need to call GroupInvitationToken.Unwrap() before calling this method if this GroupInvitationToken
// was returned from a transaction, and the transaction was committed or rolled back.
func (git *GroupInvitationToken) Update() *GroupInvitationTokenUpdateOne {
	return NewGroupInvitationTokenClient(git.config).UpdateOne(git)
}

// Unwrap unwraps the GroupInvitationToken entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (git *GroupInvitationToken) Unwrap() *GroupInvitationToken {
	_tx, ok := git.config.driver.(*txDriver)
	if !ok {
		panic("ent: GroupInvitationToken is not a transactional entity")
	}
	git.config.driver = _tx.drv
	return git
}

// String implements the fmt.Stringer.
func (git *GroupInvitationToken) String() string {
	var builder strings.Builder
	builder.WriteString("GroupInvitationToken(")
	builder.WriteString(fmt.Sprintf("id=%v, ", git.ID))
	builder.WriteString("created_at=")
	builder.WriteString(git.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(git.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("token=")
	builder.WriteString(fmt.Sprintf("%v", git.Token))
	builder.WriteString(", ")
	builder.WriteString("expires_at=")
	builder.WriteString(git.ExpiresAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("uses=")
	builder.WriteString(fmt.Sprintf("%v", git.Uses))
	builder.WriteByte(')')
	return builder.String()
}

// GroupInvitationTokens is a parsable slice of GroupInvitationToken.
type GroupInvitationTokens []*GroupInvitationToken
