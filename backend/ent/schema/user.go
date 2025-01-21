package schema

import (
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
        field.String("name").NotEmpty(),
        field.String("email").Unique().NotEmpty(),
        field.String("password").Sensitive(),
    }
}
