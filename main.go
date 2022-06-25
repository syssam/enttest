package main

import (
	"context"
	"log"

	"enttest/ent"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	client, err := ent.Open("mysql", "root:@tcp(localhost:3306)/enttest?parseTime=True")
	if err != nil {
		log.Fatal(err)
	}
	client = client.Debug()
	defer client.Close()

	ctx := context.Background()
	// Run the auto migration tool.
	if err := client.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	if err := Do(ctx, client); err != nil {
		log.Fatal(err)
	}
}

func Do(ctx context.Context, client *ent.Client) error {
	hashtag, err := client.Hashtag.Create().Save(ctx)
	if err != nil {
		return err
	}
	_, err = client.Post.Create().SetTitle("text").AddHashtags(hashtag).Save(ctx)
	if err != nil {
		return err
	}

	/*
		currently
		INSERT INTO `hashtags` (`hashtag_id`) VALUES (?) args=[62b6fef98b5abe4e0b97fc46]
		INSERT INTO `posts` (`title`, `post_id`) VALUES (?, ?) args=[text 62b6fef98b5abe4e0b97fc47]
		INSERT INTO `hashtag_to_posts` (`post_id`, `hashtag_id`) VALUES (?, ?) args=[62b6fef98b5abe4e0b97fc47 62b6fef98b5abe4e0b97fc46]

		expected
		INSERT INTO `hashtags` (`hashtag_id`) VALUES (?) args=[62b6fef98b5abe4e0b97fc46]
		INSERT INTO `posts` (`title`, `post_id`) VALUES (?, ?) args=[text 62b6fef98b5abe4e0b97fc47]
		INSERT INTO `hashtag_to_posts` (`hashtag_to_post_id`, `post_id`, `hashtag_id`) VALUES (?, ?, ?) args=[xxx 62b6fef98b5abe4e0b97fc47 62b6fef98b5abe4e0b97fc46]
	*/

	return nil
}
