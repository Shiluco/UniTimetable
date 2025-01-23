package schema

import (
    "time"
    "entgo.io/ent"
    "entgo.io/ent/schema/edge"
    "entgo.io/ent/schema/field"
)

// Schedule holds the schema definition for the Schedule entity.
type Schedule struct {
	ent.Schema
}

// Fields of the Schedule.
func (Schedule) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Unique().Immutable().StructTag(`json:"schedule_id"`),
		field.Int("user_id"),
		field.Int8("day_of_week").
			Min(0).  // 0: 日曜日
			Max(6).  // 6: 土曜日
			Comment("曜日（0:日曜日 - 6:土曜日）"),
		field.Int8("time_slot").
			Min(1).  // 1時限目
			Max(7).  // 7時限目
			Comment("時限（1-7）"),
		field.String("subject").NotEmpty().Comment("Name of the subject."),
		field.String("location").Optional().Comment("Location or classroom."),
		field.Time("created_at").Default(time.Now).Immutable().Comment("Record creation timestamp."),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now).Comment("Record update timestamp."),
	}
}

// Edges of the Schedule.
func (Schedule) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Ref("schedules").Field("user_id").Unique().Required().Comment("User who owns the schedule."),
		edge.To("posts", Post.Type).Comment("Posts associated with the schedule."),
	}
}
