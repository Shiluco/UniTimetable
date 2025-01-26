package schema

import (
    "time"
    "entgo.io/ent"
    "entgo.io/ent/schema/edge"
    "entgo.io/ent/schema/field"
)

// Post holds the schema definition for the Post entity.
type Post struct {
	ent.Schema
}

// Fields of the Post.
func (Post) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Unique().Immutable().StructTag(`json:"post_id"`),
		field.Int("parent_post_id").Optional().Nillable(),
		field.Int("user_id"),
		field.String("content").NotEmpty(),
		field.Int("schedule_id").Optional(),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the Post.
func (Post) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Ref("posts").Field("user_id").Unique().Required(),
		edge.From("schedule", Schedule.Type).Ref("posts").Field("schedule_id").Unique(),
		edge.From("parent", Post.Type).
            Ref("replies").
            Field("parent_post_id").
            Unique().
            Comment("親投稿（返信先）"),
        edge.To("replies", Post.Type).
            Comment("返信投稿"),
	}
}
