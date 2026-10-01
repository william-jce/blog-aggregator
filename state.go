package main

import (
	"github.com/william-jce/gator/internal/config"
	"github.com/william-jce/gator/internal/database"
)

type state struct {
	db  *database.Queries
	cfg *config.Config
}
