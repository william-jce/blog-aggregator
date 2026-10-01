package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
	config "github.com/william-jce/gator/internal/config"
	"github.com/william-jce/gator/internal/database"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("unable to read config file: %v", err)
	}

	db, err := sql.Open("postgres", cfg.DbUrl)
	if err != nil {
		log.Fatal(err)
	}

	dbQueries := database.New(db)
	appState := state{db: dbQueries, cfg: &cfg}

	appCommands := commands{CommandMap: make(map[string]func(*state, command) error)}

	appCommands.register("login", handlerLogin)
	appCommands.register("register", handlerRegister)
	appCommands.register("reset", handlerReset)
	appCommands.register("users", handlerUsers)
	appCommands.register("agg", handlerAgg)
	appCommands.register("addfeed", middlewareLoggedIn(handlerAddFeed))
	appCommands.register("feeds", handlerFeeds)
	appCommands.register("follow", middlewareLoggedIn(handlerFollow))
	appCommands.register("following", middlewareLoggedIn(handlerFollowing))
	appCommands.register("unfollow", middlewareLoggedIn(handlerUnfollow))
	appCommands.register("browse", middlewareLoggedIn(handlerBrowse))

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
