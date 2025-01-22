package schema

import (
    "time"
    "entgo.io/ent"
    "entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
    ent.Schema
}

// Fields defines the fields of the User entity.
func (User) Fields() []ent.Field {
    return []ent.Field{
        field.Int("id").
            Positive().
            Immutable().
			Unique(),
        field.String("name").
            NotEmpty(),
        field.String("email").
            Unique().
            NotEmpty(),
        field.String("password").
            Sensitive().
            NotEmpty(),
        field.Time("created_at").
            Default(time.Now).
            Immutable(),
        field.Time("updated_at").
            Default(time.Now).
            UpdateDefault(time.Now),
    }
}

// Edges of the User.
func (User) Edges() []ent.Edge {
    return nil
}
