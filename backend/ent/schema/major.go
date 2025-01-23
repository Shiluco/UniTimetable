package schema

import (
    "entgo.io/ent"
    "entgo.io/ent/schema/edge"
    "entgo.io/ent/schema/field"
)

// Major holds the schema definition for the Major entity.
type Major struct {
	ent.Schema
}

// Fields of the Major.
func (Major) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Unique().Immutable().StructTag(`json:"major_id"`),
		field.Int("department_id"),
		field.String("name").NotEmpty().Unique(),
	}
}

// Edges of the Major.
func (Major) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("department", Department.Type).Ref("majors").Field("department_id").Unique().Required(),
		edge.To("users", User.Type),
	}
}
