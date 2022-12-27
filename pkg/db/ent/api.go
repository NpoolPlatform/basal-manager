// Code generated by ent, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/basal-manager/pkg/db/ent/api"
	"github.com/google/uuid"
)

// API is the model entity for the API schema.
type API struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt uint32 `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt uint32 `json:"updated_at,omitempty"`
	// DeletedAt holds the value of the "deleted_at" field.
	DeletedAt uint32 `json:"deleted_at,omitempty"`
	// Protocol holds the value of the "protocol" field.
	Protocol string `json:"protocol,omitempty"`
	// ServiceName holds the value of the "service_name" field.
	ServiceName string `json:"service_name,omitempty"`
	// Method holds the value of the "method" field.
	Method string `json:"method,omitempty"`
	// MethodName holds the value of the "method_name" field.
	MethodName string `json:"method_name,omitempty"`
	// Path holds the value of the "path" field.
	Path string `json:"path,omitempty"`
	// Exported holds the value of the "exported" field.
	Exported bool `json:"exported,omitempty"`
	// PathPrefix holds the value of the "path_prefix" field.
	PathPrefix string `json:"path_prefix,omitempty"`
	// Domains holds the value of the "domains" field.
	Domains []string `json:"domains,omitempty"`
	// Depracated holds the value of the "depracated" field.
	Depracated bool `json:"depracated,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*API) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case api.FieldDomains:
			values[i] = new([]byte)
		case api.FieldExported, api.FieldDepracated:
			values[i] = new(sql.NullBool)
		case api.FieldCreatedAt, api.FieldUpdatedAt, api.FieldDeletedAt:
			values[i] = new(sql.NullInt64)
		case api.FieldProtocol, api.FieldServiceName, api.FieldMethod, api.FieldMethodName, api.FieldPath, api.FieldPathPrefix:
			values[i] = new(sql.NullString)
		case api.FieldID:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type API", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the API fields.
func (a *API) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case api.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				a.ID = *value
			}
		case api.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				a.CreatedAt = uint32(value.Int64)
			}
		case api.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				a.UpdatedAt = uint32(value.Int64)
			}
		case api.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				a.DeletedAt = uint32(value.Int64)
			}
		case api.FieldProtocol:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field protocol", values[i])
			} else if value.Valid {
				a.Protocol = value.String
			}
		case api.FieldServiceName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field service_name", values[i])
			} else if value.Valid {
				a.ServiceName = value.String
			}
		case api.FieldMethod:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field method", values[i])
			} else if value.Valid {
				a.Method = value.String
			}
		case api.FieldMethodName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field method_name", values[i])
			} else if value.Valid {
				a.MethodName = value.String
			}
		case api.FieldPath:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field path", values[i])
			} else if value.Valid {
				a.Path = value.String
			}
		case api.FieldExported:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field exported", values[i])
			} else if value.Valid {
				a.Exported = value.Bool
			}
		case api.FieldPathPrefix:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field path_prefix", values[i])
			} else if value.Valid {
				a.PathPrefix = value.String
			}
		case api.FieldDomains:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field domains", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &a.Domains); err != nil {
					return fmt.Errorf("unmarshal field domains: %w", err)
				}
			}
		case api.FieldDepracated:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field depracated", values[i])
			} else if value.Valid {
				a.Depracated = value.Bool
			}
		}
	}
	return nil
}

// Update returns a builder for updating this API.
// Note that you need to call API.Unwrap() before calling this method if this API
// was returned from a transaction, and the transaction was committed or rolled back.
func (a *API) Update() *APIUpdateOne {
	return (&APIClient{config: a.config}).UpdateOne(a)
}

// Unwrap unwraps the API entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (a *API) Unwrap() *API {
	_tx, ok := a.config.driver.(*txDriver)
	if !ok {
		panic("ent: API is not a transactional entity")
	}
	a.config.driver = _tx.drv
	return a
}

// String implements the fmt.Stringer.
func (a *API) String() string {
	var builder strings.Builder
	builder.WriteString("API(")
	builder.WriteString(fmt.Sprintf("id=%v, ", a.ID))
	builder.WriteString("created_at=")
	builder.WriteString(fmt.Sprintf("%v", a.CreatedAt))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(fmt.Sprintf("%v", a.UpdatedAt))
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(fmt.Sprintf("%v", a.DeletedAt))
	builder.WriteString(", ")
	builder.WriteString("protocol=")
	builder.WriteString(a.Protocol)
	builder.WriteString(", ")
	builder.WriteString("service_name=")
	builder.WriteString(a.ServiceName)
	builder.WriteString(", ")
	builder.WriteString("method=")
	builder.WriteString(a.Method)
	builder.WriteString(", ")
	builder.WriteString("method_name=")
	builder.WriteString(a.MethodName)
	builder.WriteString(", ")
	builder.WriteString("path=")
	builder.WriteString(a.Path)
	builder.WriteString(", ")
	builder.WriteString("exported=")
	builder.WriteString(fmt.Sprintf("%v", a.Exported))
	builder.WriteString(", ")
	builder.WriteString("path_prefix=")
	builder.WriteString(a.PathPrefix)
	builder.WriteString(", ")
	builder.WriteString("domains=")
	builder.WriteString(fmt.Sprintf("%v", a.Domains))
	builder.WriteString(", ")
	builder.WriteString("depracated=")
	builder.WriteString(fmt.Sprintf("%v", a.Depracated))
	builder.WriteByte(')')
	return builder.String()
}

// APIs is a parsable slice of API.
type APIs []*API

func (a APIs) config(cfg config) {
	for _i := range a {
		a[_i].config = cfg
	}
}