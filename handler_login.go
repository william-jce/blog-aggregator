package main

import (
	"context"
	"errors"
	"fmt"
	"os"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.Args) == 0 {
		return errors.New("cmd args empty")
	}
	userName := cmd.Args[0]
	_, err := s.db.GetUser(context.Background(), userName)
	if err != nil {
		fmt.Println("user doesn't exist")
		os.Exit(1)
	}

	err = s.cfg.SetUser(userName)
	if err != nil {
		return err
	}

	fmt.Printf("User has been set to %s\n", userName)
	return nil
}
