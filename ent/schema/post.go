package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Post holds the schema definition for the Post entity.
type Post struct {
	ent.Schema
}

// Fields of the Post.
func (Post) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").StorageKey("post_id").NotEmpty().MaxLen(24).Immutable().DefaultFunc(func() string {
			return primitive.NewObjectID().Hex()
		}).Annotations(
			entgql.OrderField("ID"),
		),
		field.String("title").Default("").MaxLen(200),
	}
}

func (Post) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("hashtags", Hashtag.Type).Through("hashtag_to_post", HashtagToPost.Type),
	}
}
