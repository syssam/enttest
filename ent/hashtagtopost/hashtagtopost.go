// Code generated by ent, DO NOT EDIT.

package hashtagtopost

const (
	// Label holds the string label denoting the hashtagtopost type in the database.
	Label = "hashtag_to_post"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "hashtag_to_post_id"
	// FieldHashtagID holds the string denoting the hashtag_id field in the database.
	FieldHashtagID = "hashtag_id"
	// FieldPostID holds the string denoting the post_id field in the database.
	FieldPostID = "post_id"
	// EdgePost holds the string denoting the post edge name in mutations.
	EdgePost = "post"
	// EdgeHashtag holds the string denoting the hashtag edge name in mutations.
	EdgeHashtag = "hashtag"
	// PostFieldID holds the string denoting the ID field of the Post.
	PostFieldID = "post_id"
	// HashtagFieldID holds the string denoting the ID field of the Hashtag.
	HashtagFieldID = "hashtag_id"
	// Table holds the table name of the hashtagtopost in the database.
	Table = "hashtag_to_posts"
	// PostTable is the table that holds the post relation/edge.
	PostTable = "hashtag_to_posts"
	// PostInverseTable is the table name for the Post entity.
	// It exists in this package in order to avoid circular dependency with the "post" package.
	PostInverseTable = "posts"
	// PostColumn is the table column denoting the post relation/edge.
	PostColumn = "post_id"
	// HashtagTable is the table that holds the hashtag relation/edge.
	HashtagTable = "hashtag_to_posts"
	// HashtagInverseTable is the table name for the Hashtag entity.
	// It exists in this package in order to avoid circular dependency with the "hashtag" package.
	HashtagInverseTable = "hashtags"
	// HashtagColumn is the table column denoting the hashtag relation/edge.
	HashtagColumn = "hashtag_id"
)

// Columns holds all SQL columns for hashtagtopost fields.
var Columns = []string{
	FieldID,
	FieldHashtagID,
	FieldPostID,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// HashtagIDValidator is a validator for the "hashtag_id" field. It is called by the builders before save.
	HashtagIDValidator func(string) error
	// PostIDValidator is a validator for the "post_id" field. It is called by the builders before save.
	PostIDValidator func(string) error
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() string
	// IDValidator is a validator for the "id" field. It is called by the builders before save.
	IDValidator func(string) error
)