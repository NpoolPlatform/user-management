package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/google/uuid"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Unique(),
		field.String("username").Unique(),
		field.String("password"),
		field.String("salt"),
		field.String("display_name"),
		field.String("phone_number").Unique(),
		field.String("email_address").Unique(),
		field.Int32("login_times").
			Default(0),
		field.Bool("kyc_verify").
			Default(false),
		field.Bool("ga_verify").
			Default(false),
		field.String("signup_method"),
		field.Int64("create_at").
			DefaultFunc(func() int64 {
				return time.Now().Unix()
			}),
		field.Int64("update_at").
			DefaultFunc(func() int64 {
				return time.Now().Unix()
			}).
			UpdateDefault(func() int64 {
				return time.Now().Unix()
			}),
		field.Int64("delete_at").
			DefaultFunc(func() int64 {
				return 0
			}),
		field.String("avatar"),
		field.String("region"),
		field.Int32("age"),
		field.String("gender"),
		field.String("birthday"),
		field.String("country"),
		field.String("province"),
		field.String("city"),
		field.String("career"),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}

func (User) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("username"),
		index.Fields("phone_number"),
		index.Fields("email_address"),
	}
}
