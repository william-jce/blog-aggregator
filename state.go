package main

import (
	"github.com/william-jce/blog-aggregator/internal/config"
	"github.com/william-jce/blog-aggregator/internal/database"
)

type state struct {
	db  *database.Queries
	cfg *config.Config
}
