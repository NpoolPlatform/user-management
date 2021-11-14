package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/google/uuid"
)

// UserFrozen holds the schema definition for the UserFrozen entity.
type UserFrozen struct {
	ent.Schema
}

// Fields of the UserFrozen.
func (UserFrozen) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Unique(),
		field.UUID("user_id", uuid.UUID{}),
		field.UUID("frozen_by", uuid.UUID{}),
		field.String("frozen_cause"),
		field.Uint32("create_at").
			DefaultFunc(func() uint32 {
				return uint32(time.Now().Unix())
			}),
		field.Uint32("end_at").
			Default(0),
		field.String("status"),
		field.UUID("unfrozen_by", uuid.UUID{}),
	}
}

// Edges of the UserFrozen.
func (UserFrozen) Edges() []ent.Edge {
	return nil
}

func (UserFrozen) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("id", "status").Unique(),
	}
}
