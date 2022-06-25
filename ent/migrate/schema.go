// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// HashtagsColumns holds the columns for the "hashtags" table.
	HashtagsColumns = []*schema.Column{
		{Name: "hashtag_id", Type: field.TypeString, Size: 24},
	}
	// HashtagsTable holds the schema information for the "hashtags" table.
	HashtagsTable = &schema.Table{
		Name:       "hashtags",
		Columns:    HashtagsColumns,
		PrimaryKey: []*schema.Column{HashtagsColumns[0]},
	}
	// HashtagToPostsColumns holds the columns for the "hashtag_to_posts" table.
	HashtagToPostsColumns = []*schema.Column{
		{Name: "hashtag_to_post_id", Type: field.TypeString, Size: 24},
		{Name: "post_id", Type: field.TypeString, Size: 24},
		{Name: "hashtag_id", Type: field.TypeString, Size: 24},
	}
	// HashtagToPostsTable holds the schema information for the "hashtag_to_posts" table.
	HashtagToPostsTable = &schema.Table{
		Name:       "hashtag_to_posts",
		Columns:    HashtagToPostsColumns,
		PrimaryKey: []*schema.Column{HashtagToPostsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "hashtag_to_posts_posts_post",
				Columns:    []*schema.Column{HashtagToPostsColumns[1]},
				RefColumns: []*schema.Column{PostsColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "hashtag_to_posts_hashtags_hashtag",
				Columns:    []*schema.Column{HashtagToPostsColumns[2]},
				RefColumns: []*schema.Column{HashtagsColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "hashtagtopost_post_id_hashtag_id",
				Unique:  true,
				Columns: []*schema.Column{HashtagToPostsColumns[1], HashtagToPostsColumns[2]},
			},
		},
	}
	// PostsColumns holds the columns for the "posts" table.
	PostsColumns = []*schema.Column{
		{Name: "post_id", Type: field.TypeString, Size: 24},
		{Name: "title", Type: field.TypeString, Size: 200, Default: ""},
	}
	// PostsTable holds the schema information for the "posts" table.
	PostsTable = &schema.Table{
		Name:       "posts",
		Columns:    PostsColumns,
		PrimaryKey: []*schema.Column{PostsColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		HashtagsTable,
		HashtagToPostsTable,
		PostsTable,
	}
)

func init() {
	HashtagToPostsTable.ForeignKeys[0].RefTable = PostsTable
	HashtagToPostsTable.ForeignKeys[1].RefTable = HashtagsTable
}