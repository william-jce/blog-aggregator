package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/william-jce/gator/internal/database"
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

	for _, feedItem := range feed.Channel.Item {
		description := sql.NullString{String: feedItem.Description, Valid: feedItem.Description != ""}
		var publishedAt sql.NullTime
		pubDate, err := time.Parse(time.RFC1123Z, feedItem.PubDate)
		if err == nil {
			publishedAt = sql.NullTime{Time: pubDate, Valid: true}
		}
		_, err = s.db.CreatePost(context.Background(), database.CreatePostParams{
			ID:          uuid.New(),
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
			Title:       feedItem.Title,
			Url:         feedItem.Link,
			Description: description,
			PublishedAt: publishedAt,
			FeedID:      feedToFetch.ID,
		})
		if err != nil {
			if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
				continue
			}
			log.Printf("creating post: %v\n", err)
		}
	}

	return nil
}
