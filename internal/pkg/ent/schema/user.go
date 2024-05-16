package schema

import (
	"time"

	"hiyoko-fiber/internal/domain/entities/users"
	"hiyoko-fiber/internal/pkg/ent/util"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Mixin of the User.
// updateやcreateをmixinするとカラムの順序を制御出来ない
func (User) Mixin() []ent.Mixin {
	return []ent.Mixin{}
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").
			GoType(util.ULID("")).
			DefaultFunc(func() util.ULID {
				return util.NewULID()
			}).
			Immutable().
			Unique(),
		field.Int8("status").
			GoType(users.Status(0)).
			Default(users.Status(0).Default().ToInt8()),
		field.String("original_id").
			MaxLen(255).
			Optional(),
		field.String("email").
			MaxLen(255).
			Unique().
			NotEmpty(),
		field.String("password").
			MaxLen(255).
			Sensitive().
			NotEmpty(),
		field.Time("created_at").
			Immutable().
			Default(time.Now).
			SchemaType(map[string]string{
				dialect.MySQL: "datetime",
			}),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now).
			SchemaType(map[string]string{
				dialect.MySQL: "datetime",
			}),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{}
}

func (User) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("original_id").Unique(),
	}
}
