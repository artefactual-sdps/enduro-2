// Code generated by ent, DO NOT EDIT.

package db

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/artefactual-sdps/enduro/internal/storage/persistence/ent/db/location"
	"github.com/artefactual-sdps/enduro/internal/storage/persistence/ent/db/pkg"
	"github.com/artefactual-sdps/enduro/internal/storage/types"
	"github.com/google/uuid"
)

// Pkg is the model entity for the Pkg schema.
type Pkg struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// AipID holds the value of the "aip_id" field.
	AipID uuid.UUID `json:"aip_id,omitempty"`
	// LocationID holds the value of the "location_id" field.
	LocationID int `json:"location_id,omitempty"`
	// Status holds the value of the "status" field.
	Status types.PackageStatus `json:"status,omitempty"`
	// ObjectKey holds the value of the "object_key" field.
	ObjectKey uuid.UUID `json:"object_key,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the PkgQuery when eager-loading is set.
	Edges        PkgEdges `json:"edges"`
	selectValues sql.SelectValues
}

// PkgEdges holds the relations/edges for other nodes in the graph.
type PkgEdges struct {
	// Location holds the value of the location edge.
	Location *Location `json:"location,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// LocationOrErr returns the Location value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e PkgEdges) LocationOrErr() (*Location, error) {
	if e.loadedTypes[0] {
		if e.Location == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: location.Label}
		}
		return e.Location, nil
	}
	return nil, &NotLoadedError{edge: "location"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Pkg) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case pkg.FieldID, pkg.FieldLocationID:
			values[i] = new(sql.NullInt64)
		case pkg.FieldName:
			values[i] = new(sql.NullString)
		case pkg.FieldCreatedAt:
			values[i] = new(sql.NullTime)
		case pkg.FieldStatus:
			values[i] = new(types.PackageStatus)
		case pkg.FieldAipID, pkg.FieldObjectKey:
			values[i] = new(uuid.UUID)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Pkg fields.
func (pk *Pkg) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case pkg.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			pk.ID = int(value.Int64)
		case pkg.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				pk.Name = value.String
			}
		case pkg.FieldAipID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field aip_id", values[i])
			} else if value != nil {
				pk.AipID = *value
			}
		case pkg.FieldLocationID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field location_id", values[i])
			} else if value.Valid {
				pk.LocationID = int(value.Int64)
			}
		case pkg.FieldStatus:
			if value, ok := values[i].(*types.PackageStatus); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value != nil {
				pk.Status = *value
			}
		case pkg.FieldObjectKey:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field object_key", values[i])
			} else if value != nil {
				pk.ObjectKey = *value
			}
		case pkg.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				pk.CreatedAt = value.Time
			}
		default:
			pk.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Pkg.
// This includes values selected through modifiers, order, etc.
func (pk *Pkg) Value(name string) (ent.Value, error) {
	return pk.selectValues.Get(name)
}

// QueryLocation queries the "location" edge of the Pkg entity.
func (pk *Pkg) QueryLocation() *LocationQuery {
	return NewPkgClient(pk.config).QueryLocation(pk)
}

// Update returns a builder for updating this Pkg.
// Note that you need to call Pkg.Unwrap() before calling this method if this Pkg
// was returned from a transaction, and the transaction was committed or rolled back.
func (pk *Pkg) Update() *PkgUpdateOne {
	return NewPkgClient(pk.config).UpdateOne(pk)
}

// Unwrap unwraps the Pkg entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (pk *Pkg) Unwrap() *Pkg {
	_tx, ok := pk.config.driver.(*txDriver)
	if !ok {
		panic("db: Pkg is not a transactional entity")
	}
	pk.config.driver = _tx.drv
	return pk
}

// String implements the fmt.Stringer.
func (pk *Pkg) String() string {
	var builder strings.Builder
	builder.WriteString("Pkg(")
	builder.WriteString(fmt.Sprintf("id=%v, ", pk.ID))
	builder.WriteString("name=")
	builder.WriteString(pk.Name)
	builder.WriteString(", ")
	builder.WriteString("aip_id=")
	builder.WriteString(fmt.Sprintf("%v", pk.AipID))
	builder.WriteString(", ")
	builder.WriteString("location_id=")
	builder.WriteString(fmt.Sprintf("%v", pk.LocationID))
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(fmt.Sprintf("%v", pk.Status))
	builder.WriteString(", ")
	builder.WriteString("object_key=")
	builder.WriteString(fmt.Sprintf("%v", pk.ObjectKey))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(pk.CreatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Pkgs is a parsable slice of Pkg.
type Pkgs []*Pkg
