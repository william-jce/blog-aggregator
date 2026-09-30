package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	config "github.com/william-jce/blog-aggregator/internal/config"
	"github.com/william-jce/blog-aggregator/internal/database"
)

func main() {
	godotenv.Load(".env")
	dbURL := os.Getenv("DB_URL")
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal(err)
	}

	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("unable to read config file: %v", err)
	}

	dbQueries := database.New(db)
	appState := state{db: dbQueries, cfg: &cfg}

	appCommands := commands{CommandMap: make(map[string]func(*state, command) error)}

	appCommands.register("login", handlerLogin)
	appCommands.register("register", handlerRegister)
	appCommands.register("reset", handlerReset)
	appCommands.register("users", handlerUsers)
	appCommands.register("agg", handlerAgg)
	appCommands.register("addfeed", handlerAddFeed)
	appCommands.register("feeds", handlerFeeds)
	appCommands.register("follow", handlerFollow)
	appCommands.register("following", handlerFollowing)

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
