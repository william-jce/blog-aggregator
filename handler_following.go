package main

import (
	"context"
	"fmt"
)

func handlerFollowing(s *state, cmd command) error {
	user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return fmt.Errorf("getting user: %w", err)
	}

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
