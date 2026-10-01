package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/william-jce/gator/internal/database"
)

func handlerFollow(s *state, cmd command, user database.User) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <url>", cmd.Name)
	}

	url := cmd.Args[0]

	feed, err := s.db.GetFeed(context.Background(), url)
	if err != nil {
		return fmt.Errorf("getting feed: %w", err)
	}

	feedFollow, err := s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	})
	if err != nil {
		return fmt.Errorf("creating feed follow record: %w", err)
	}

	fmt.Printf("%s's Feed:\n", s.cfg.CurrentUserName)
	for _, feedItem := range feedFollow {
		fmt.Println(feedItem)
	}
	return nil
}
