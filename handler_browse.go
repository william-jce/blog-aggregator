package main

import (
	"context"
	"fmt"
	"strconv"

	"github.com/william-jce/gator/internal/database"
)

func handlerBrowse(s *state, cmd command, user database.User) error {
	limit := 2
	if len(cmd.Args) == 1 {
		specLimit, err := strconv.Atoi(cmd.Args[0])
		if err != nil {
			return fmt.Errorf("usage: %s <optional limit number>", cmd.Name)
		}
		limit = specLimit
	}

	posts, err := s.db.GetPostsForUser(context.Background(), database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  int32(limit),
	})
	if err != nil {
		return fmt.Errorf("getting posts for user: %w", err)
	}

	fmt.Println("Posts:")
	for _, post := range posts {
		fmt.Println("------")
		fmt.Printf("%s\n\n", post.Title)
		fmt.Printf("%s\n\n", post.Url)
		fmt.Printf("%s\n\n", post.Description.String)
	}

	return nil
}
