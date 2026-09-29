package main

import (
	"context"
	"fmt"
)

func handlerFeeds(s *state, cmd command) error {
	if len(cmd.Args) != 0 {
		return fmt.Errorf("usage: %s", cmd.Name)
	}

	feeds, err := s.db.GetFeeds(context.Background())
	if err != nil {
		return fmt.Errorf("fetching feeds: %w", err)
	}

	fmt.Println("Feeds:")
	for _, feed := range feeds {
		fmt.Printf("%+v\n", feed)
	}
	return nil
}
