package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// HashtagToPost holds the schema definition for the HashtagToPost entity.
type HashtagToPost struct {
	ent.Schema
}

// Fields of the HashtagToPost.
func (HashtagToPost) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").StorageKey("hashtag_to_post_id").NotEmpty().MaxLen(24).Immutable().DefaultFunc(func() string {
			return primitive.NewObjectID().Hex()
		}).Annotations(
			entgql.OrderField("ID"),
		),
		field.String("hashtag_id").NotEmpty().MaxLen(24).Annotations(
			entsql.Annotation{
				Default: "",
			},
			entgql.Type("ID"),
		),
		field.String("post_id").NotEmpty().MaxLen(24).Annotations(
			entsql.Annotation{
				Default: "",
			},
			entgql.Type("ID"),
		),
	}
}

// Edges of the HashtagToPost.
func (HashtagToPost) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("post", Post.Type).Annotations(entgql.Skip()).
			Required().
			Unique().
			Field("post_id"),
		edge.To("hashtag", Hashtag.Type).Annotations(entgql.Skip()).
			Required().
			Unique().
			Field("hashtag_id"),
	}
}
