package main

import (
	"errors"
	"fmt"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.Args) == 0 {
		return errors.New("cmd args empty")
	}
	userName := cmd.Args[0]
	err := s.Cfg.SetUser(userName)
	if err != nil {
		return errors.New("unable to set user")
	}

	fmt.Printf("User has been set to %s\n", userName)
	return err
}
