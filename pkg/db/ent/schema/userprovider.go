package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/google/uuid"
)

// UserProvider holds the schema definition for the UserProvider entity.
type UserProvider struct {
	ent.Schema
}

// Fields of the UserProvider.
func (UserProvider) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Unique(),
		field.UUID("user_id", uuid.UUID{}),
		field.UUID("provider_id", uuid.UUID{}),
		field.String("provider_user_id").Unique(),
		field.String("user_provider_info"),
		field.Uint32("create_at").
			DefaultFunc(func() uint32 {
				return uint32(time.Now().Unix())
			}),
		field.Uint32("update_at").
			DefaultFunc(func() uint32 {
				return uint32(time.Now().Unix())
			}).
			UpdateDefault(func() uint32 {
				return uint32(time.Now().Unix())
			}),
		field.Uint32("delete_at").
			DefaultFunc(func() uint32 {
				return 0
			}),
	}
}

// Edges of the UserProvider.
func (UserProvider) Edges() []ent.Edge {
	return nil
}

func (UserProvider) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("user_id", "provider_id").Unique(),
	}
}
