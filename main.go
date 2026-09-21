package main

import (
	"fmt"
	"log"

	config "github.com/william-jce/blog-aggregator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Printf("Error: %v", err)
	}
	if err := cfg.SetUser("william"); err != nil {
		log.Printf("Error: %v", err)
	}

	newCfg, err := config.Read()
	if err != nil {
		log.Printf("Error: %v", err)
	}

	fmt.Print(newCfg)
}
