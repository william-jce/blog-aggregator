package main

import (
	"fmt"
	"log"
	"time"
)

func handlerAgg(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <time_between_reqs e.g. 1s, 1m, 1h>\n", cmd.Name)
	}

	timeBetweenReqs, err := time.ParseDuration(cmd.Args[0])
	if err != nil {
		return fmt.Errorf("parsing time between requests: %w", err)
	}

	fmt.Printf("Collecting feeds every %s\n", cmd.Args[0])
	ticker := time.NewTicker(timeBetweenReqs)
	defer ticker.Stop()
	for ; ; <-ticker.C {
		err := scrapeFeeds(s)
		if err != nil {
			log.Printf("scraping feeds: %v", err)
		}
	}
	return nil
}
