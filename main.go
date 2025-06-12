package main

import (
	"fmt"

	"github.com/Rishan-Jadva/gator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	err = cfg.SetUser("rndmsns")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	updatedCfg, err := config.Read()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	fmt.Printf("Config loaded:\nDB URL: %s\nCurrent User: %s\n", updatedCfg.DBUrl, updatedCfg.CurrentUser)
}
