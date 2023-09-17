// Code generated by ent, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/hexennacht/signme/services/sm-user-api/core/entity"
	"github.com/hexennacht/signme/services/sm-user-api/ent/credential"
)

// Credential is the model entity for the Credential schema.
type Credential struct {
	config `json:"-"`
	// ID of the ent.
	ID int64 `json:"id,omitempty"`
	// CredentialType holds the value of the "credential_type" field.
	CredentialType credential.CredentialType `json:"credential_type,omitempty"`
	// Credential holds the value of the "credential" field.
	Credential entity.Credential `json:"credential,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	// DeletedAt holds the value of the "deleted_at" field.
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the CredentialQuery when eager-loading is set.
	Edges        CredentialEdges `json:"edges"`
	selectValues sql.SelectValues
}

// CredentialEdges holds the relations/edges for other nodes in the graph.
type CredentialEdges struct {
	// Users holds the value of the users edge.
	Users []*User `json:"users,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// UsersOrErr returns the Users value or an error if the edge
// was not loaded in eager-loading.
func (e CredentialEdges) UsersOrErr() ([]*User, error) {
	if e.loadedTypes[0] {
		return e.Users, nil
	}
	return nil, &NotLoadedError{edge: "users"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Credential) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case credential.FieldCredential:
			values[i] = new([]byte)
		case credential.FieldID:
			values[i] = new(sql.NullInt64)
		case credential.FieldCredentialType:
			values[i] = new(sql.NullString)
		case credential.FieldCreatedAt, credential.FieldUpdatedAt, credential.FieldDeletedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Credential fields.
func (c *Credential) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case credential.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			c.ID = int64(value.Int64)
		case credential.FieldCredentialType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field credential_type", values[i])
			} else if value.Valid {
				c.CredentialType = credential.CredentialType(value.String)
			}
		case credential.FieldCredential:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field credential", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &c.Credential); err != nil {
					return fmt.Errorf("unmarshal field credential: %w", err)
				}
			}
		case credential.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				c.CreatedAt = value.Time
			}
		case credential.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				c.UpdatedAt = new(time.Time)
				*c.UpdatedAt = value.Time
			}
		case credential.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				c.DeletedAt = new(time.Time)
				*c.DeletedAt = value.Time
			}
		default:
			c.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Credential.
// This includes values selected through modifiers, order, etc.
func (c *Credential) Value(name string) (ent.Value, error) {
	return c.selectValues.Get(name)
}

// QueryUsers queries the "users" edge of the Credential entity.
func (c *Credential) QueryUsers() *UserQuery {
	return NewCredentialClient(c.config).QueryUsers(c)
}

// Update returns a builder for updating this Credential.
// Note that you need to call Credential.Unwrap() before calling this method if this Credential
// was returned from a transaction, and the transaction was committed or rolled back.
func (c *Credential) Update() *CredentialUpdateOne {
	return NewCredentialClient(c.config).UpdateOne(c)
}

// Unwrap unwraps the Credential entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (c *Credential) Unwrap() *Credential {
	_tx, ok := c.config.driver.(*txDriver)
	if !ok {
		panic("ent: Credential is not a transactional entity")
	}
	c.config.driver = _tx.drv
	return c
}

// String implements the fmt.Stringer.
func (c *Credential) String() string {
	var builder strings.Builder
	builder.WriteString("Credential(")
	builder.WriteString(fmt.Sprintf("id=%v, ", c.ID))
	builder.WriteString("credential_type=")
	builder.WriteString(fmt.Sprintf("%v", c.CredentialType))
	builder.WriteString(", ")
	builder.WriteString("credential=")
	builder.WriteString(fmt.Sprintf("%v", c.Credential))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(c.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	if v := c.UpdatedAt; v != nil {
		builder.WriteString("updated_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	if v := c.DeletedAt; v != nil {
		builder.WriteString("deleted_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteByte(')')
	return builder.String()
}

// Credentials is a parsable slice of Credential.
type Credentials []*Credential
