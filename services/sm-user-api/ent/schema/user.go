package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").Positive().Immutable(),
		field.String("name").NotEmpty(),
		field.String("username").NotEmpty().Unique(),
		field.String("password").NotEmpty(),
		field.String("profile_picture").Default("https://picsum.photos/300/300.jpg"),
		field.Enum("status").Values("new", "active", "banned", "deleted").Default("new"),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now).Nillable(),
		field.Time("deleted_at").Nillable(),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}

func (User) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("status", "created_at"),
	}
}
