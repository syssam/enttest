package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Hashtag holds the schema definition for the Hashtag entity.
type Hashtag struct {
	ent.Schema
}

// Fields of the Hashtag.
func (Hashtag) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").StorageKey("hashtag_id").NotEmpty().MaxLen(24).Immutable().DefaultFunc(func() string {
			return primitive.NewObjectID().Hex()
		}).Annotations(
			entgql.OrderField("ID"),
		),
	}
}

// Edges of the Hashtag.
func (Hashtag) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("post", Post.Type).Annotations(entgql.Skip()).Ref("hashtags").Through("hashtag_to_post", HashtagToPost.Type),
	}
}
