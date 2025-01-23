package schema

import (
    "time"
    "entgo.io/ent"
    "entgo.io/ent/schema/field"
	"entgo.io/ent/schema/edge"
)

// User holds the schema definition for the User entity.
type User struct {
    ent.Schema
}

// Fields defines the fields of the User entity.
func (User) Fields() []ent.Field {
    return []ent.Field{
        field.Int("id").Unique().Immutable().StructTag(`json:"user_id"`),
		field.String("name").NotEmpty(),
		field.String("email").Unique().NotEmpty(),
		field.String("password").Sensitive().NotEmpty(),
		field.Int("department_id").Optional(),
		field.Int("major_id").Optional(),
		field.String("comment").Optional(),
		field.Int8("grade").Min(1).Max(4).Optional(),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
    }
}

// Edges of the User.
func (User) Edges() []ent.Edge {
    return []ent.Edge{
		edge.From("department", Department.Type).
			Ref("users").
			Field("department_id").
			Unique(),
		edge.From("major", Major.Type).
			Ref("users").
			Field("major_id").
			Unique(),
		edge.To("posts", Post.Type),
		edge.To("schedules", Schedule.Type),
	}
}
