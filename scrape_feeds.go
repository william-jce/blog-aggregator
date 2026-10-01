package main

import (
	"context"
	"fmt"
)

func scrapeFeeds(s *state) error {
	feedToFetch, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		return fmt.Errorf("getting next feed to fetch: %w", err)
	}

	err = s.db.MarkFeedFetched(context.Background(), feedToFetch.ID)
	if err != nil {
		return fmt.Errorf("marking feed as fetched: %w", err)
	}

	feed, err := fetchFeed(context.Background(), feedToFetch.Url)
	if err != nil {
		return fmt.Errorf("fetching feed: %w", err)
	}
	// feeds, err := s.db.GetFeeds(context.Background())
	// if err != nil {
	// 	return fmt.Errorf("getting feeds: %w", err)
	// }

	fmt.Println("Feed Titles:")
	for _, feedItem := range feed.Channel.Item {
		fmt.Printf("%s\n", feedItem.Title)
		fmt.Println("")
	}

	return nil
}
