package main

import (
	"context"
	"fmt"

	"github.com/william-jce/blog-aggregator/internal/database"
)

func handlerFollowing(s *state, cmd command, user database.User) error {
	feedFollows, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return fmt.Errorf("getting feed follows: %w", err)
	}

	fmt.Printf("%s's Followed Feeds:\n", s.cfg.CurrentUserName)
	for _, feed := range feedFollows {
		fmt.Println(feed.FeedName)
	}

	return nil
}
