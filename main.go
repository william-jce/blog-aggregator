package main

import (
	"fmt"
	"log"
	"os"

	config "github.com/william-jce/blog-aggregator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("unable to read config file: %v", err)
	}

	appState := state{Cfg: &cfg}
	appCommands := commands{CommandMap: make(map[string]func(*state, command) error)}

	appCommands.register("login", handlerLogin)

	userArgs := os.Args
	if len(userArgs) < 2 {
		fmt.Println("invalid args")
		os.Exit(1)
	}

	cmd := command{Name: userArgs[1], Args: userArgs[2:]}
	err = appCommands.run(&appState, cmd)
	if err != nil {
		fmt.Printf("unexpected issue running command: %v", err)
		os.Exit(1)
	}
}
