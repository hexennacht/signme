package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"

	"github.com/hexennacht/signme/services/sm-user-api/core/entity"
)

// Credential holds the schema definition for the Credential entity.
type Credential struct {
	ent.Schema
}

// Fields of the Credential.
func (Credential) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").Positive().Immutable(),
		field.Enum("credential_type").Values("private", "public").Default("private").Immutable(),
		field.JSON("credential", entity.Credential{}).Immutable(),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now).Nillable(),
		field.Time("deleted_at").Default(nil).Nillable(),
	}
}

// Edges of the Credential.
func (Credential) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("users", User.Type).
			Ref("credentials"),
	}
}
