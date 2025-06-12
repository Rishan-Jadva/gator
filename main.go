package main

import (
	"log"
	"os"

	"github.com/Rishan-Jadva/gator/internal/config"
)

type state struct {
	cfg *config.Config
}

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}

	programState := &state{&cfg}
	programCommands := &commands{make(map[string]func(*state, command) error)}

	programCommands.register("login", handlerLogin)

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
