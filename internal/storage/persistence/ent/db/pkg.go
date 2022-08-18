// Code generated by ent, DO NOT EDIT.

package db

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/artefactual-sdps/enduro/internal/storage/persistence/ent/db/location"
	"github.com/artefactual-sdps/enduro/internal/storage/persistence/ent/db/pkg"
	"github.com/artefactual-sdps/enduro/internal/storage/status"
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
	Status status.PackageStatus `json:"status,omitempty"`
	// ObjectKey holds the value of the "object_key" field.
	ObjectKey uuid.UUID `json:"object_key,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the PkgQuery when eager-loading is set.
	Edges PkgEdges `json:"edges"`
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
			// The edge location was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: location.Label}
		}
		return e.Location, nil
	}
	return nil, &NotLoadedError{edge: "location"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Pkg) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case pkg.FieldID, pkg.FieldLocationID:
			values[i] = new(sql.NullInt64)
		case pkg.FieldName:
			values[i] = new(sql.NullString)
		case pkg.FieldStatus:
			values[i] = new(status.PackageStatus)
		case pkg.FieldAipID, pkg.FieldObjectKey:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Pkg", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Pkg fields.
func (pk *Pkg) assignValues(columns []string, values []interface{}) error {
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
			if value, ok := values[i].(*status.PackageStatus); !ok {
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
		}
	}
	return nil
}

// QueryLocation queries the "location" edge of the Pkg entity.
func (pk *Pkg) QueryLocation() *LocationQuery {
	return (&PkgClient{config: pk.config}).QueryLocation(pk)
}

// Update returns a builder for updating this Pkg.
// Note that you need to call Pkg.Unwrap() before calling this method if this Pkg
// was returned from a transaction, and the transaction was committed or rolled back.
func (pk *Pkg) Update() *PkgUpdateOne {
	return (&PkgClient{config: pk.config}).UpdateOne(pk)
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
	builder.WriteByte(')')
	return builder.String()
}

// Pkgs is a parsable slice of Pkg.
type Pkgs []*Pkg

func (pk Pkgs) config(cfg config) {
	for _i := range pk {
		pk[_i].config = cfg
	}
}
