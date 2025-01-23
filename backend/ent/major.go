// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/Shiluco/UniTimetable/backend/ent/department"
	"github.com/Shiluco/UniTimetable/backend/ent/major"
)

// Major is the model entity for the Major schema.
type Major struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"major_id"`
	// DepartmentID holds the value of the "department_id" field.
	DepartmentID int `json:"department_id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the MajorQuery when eager-loading is set.
	Edges        MajorEdges `json:"edges"`
	selectValues sql.SelectValues
}

// MajorEdges holds the relations/edges for other nodes in the graph.
type MajorEdges struct {
	// Department holds the value of the department edge.
	Department *Department `json:"department,omitempty"`
	// Users holds the value of the users edge.
	Users []*User `json:"users,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// DepartmentOrErr returns the Department value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e MajorEdges) DepartmentOrErr() (*Department, error) {
	if e.Department != nil {
		return e.Department, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: department.Label}
	}
	return nil, &NotLoadedError{edge: "department"}
}

// UsersOrErr returns the Users value or an error if the edge
// was not loaded in eager-loading.
func (e MajorEdges) UsersOrErr() ([]*User, error) {
	if e.loadedTypes[1] {
		return e.Users, nil
	}
	return nil, &NotLoadedError{edge: "users"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Major) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case major.FieldID, major.FieldDepartmentID:
			values[i] = new(sql.NullInt64)
		case major.FieldName:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Major fields.
func (m *Major) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case major.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			m.ID = int(value.Int64)
		case major.FieldDepartmentID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field department_id", values[i])
			} else if value.Valid {
				m.DepartmentID = int(value.Int64)
			}
		case major.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				m.Name = value.String
			}
		default:
			m.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Major.
// This includes values selected through modifiers, order, etc.
func (m *Major) Value(name string) (ent.Value, error) {
	return m.selectValues.Get(name)
}

// QueryDepartment queries the "department" edge of the Major entity.
func (m *Major) QueryDepartment() *DepartmentQuery {
	return NewMajorClient(m.config).QueryDepartment(m)
}

// QueryUsers queries the "users" edge of the Major entity.
func (m *Major) QueryUsers() *UserQuery {
	return NewMajorClient(m.config).QueryUsers(m)
}

// Update returns a builder for updating this Major.
// Note that you need to call Major.Unwrap() before calling this method if this Major
// was returned from a transaction, and the transaction was committed or rolled back.
func (m *Major) Update() *MajorUpdateOne {
	return NewMajorClient(m.config).UpdateOne(m)
}

// Unwrap unwraps the Major entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (m *Major) Unwrap() *Major {
	_tx, ok := m.config.driver.(*txDriver)
	if !ok {
		panic("ent: Major is not a transactional entity")
	}
	m.config.driver = _tx.drv
	return m
}

// String implements the fmt.Stringer.
func (m *Major) String() string {
	var builder strings.Builder
	builder.WriteString("Major(")
	builder.WriteString(fmt.Sprintf("id=%v, ", m.ID))
	builder.WriteString("department_id=")
	builder.WriteString(fmt.Sprintf("%v", m.DepartmentID))
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(m.Name)
	builder.WriteByte(')')
	return builder.String()
}

// Majors is a parsable slice of Major.
type Majors []*Major
