package main

import (
	"context"
	"database/sql"
	"log"
	"os"

	"github.com/Rishan-Jadva/gator/internal/config"
	"github.com/Rishan-Jadva/gator/internal/database"
	_ "github.com/lib/pq"
)

type state struct {
	db  *database.Queries
	cfg *config.Config
}

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}

	db, err := sql.Open("postgres", cfg.DBUrl)
	if err != nil {
		log.Fatalf("error connecting to db: %v", err)
	}
	defer db.Close()

	dbQueries := database.New(db)

	programState := &state{dbQueries, &cfg}
	programCommands := &commands{make(map[string]func(*state, command) error)}

	programCommands.register("login", handlerLogin)
	programCommands.register("register", handlerRegister)
	programCommands.register("reset", handlerReset)
	programCommands.register("users", handlerListUsers)
	programCommands.register("agg", handlerAgg)
	programCommands.register("addfeed", middlewareLoggedIn(handlerAddFeed))
	programCommands.register("feeds", handlerListFeeds)
	programCommands.register("follow", middlewareLoggedIn(handlerFollow))
	programCommands.register("following", middlewareLoggedIn(handlerFollowing))
	programCommands.register("unfollow", middlewareLoggedIn(handlerUnfollow))

	if len(os.Args) < 2 {
		log.Fatal("Usage: cli <command> [args...]")
	}

	commandName := os.Args[1]
	commandArgs := os.Args[2:]

	cmd := command{commandName, commandArgs}

	err = programCommands.run(programState, cmd)
	if err != nil {
		log.Fatal(err)
	}
}

func middlewareLoggedIn(handler func(s *state, cmd command, user database.User) error) func(*state, command) error {
	return func(s *state, cmd command) error {
		user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUser)
		if err != nil {
			return err
		}
		return handler(s, cmd, user)
	}
}
