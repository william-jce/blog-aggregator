package main

import "errors"

type command struct {
	Name string
	Args []string
}

type commands struct {
	CommandMap map[string]func(*state, command) error
}

func (c *commands) run(s *state, cmd command) error {
	handler, ok := c.CommandMap[cmd.Name]
	if !ok {
		return errors.New("command doesn't exist")
	}

	return handler(s, cmd)
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.CommandMap[name] = f
}
