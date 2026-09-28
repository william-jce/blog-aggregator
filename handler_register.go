package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/william-jce/blog-aggregator/internal/database"
)

func handlerRegister(s *state, cmd command) error {
	if len(cmd.Args) == 0 {
		return errors.New("cmd args empty")
	}
	userName := cmd.Args[0]
	_, err := s.db.GetUser(context.Background(), userName)
	if err == nil {
		return errors.New("user already exists")
	}

	user, err := s.db.CreateUser(context.Background(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      userName,
	})
	if err != nil {
		return err
	}

	err = s.cfg.SetUser(userName)
	if err != nil {
		return err
	}

	fmt.Println("user was created")
	log.Printf("user: %v", user)

	return nil
}
