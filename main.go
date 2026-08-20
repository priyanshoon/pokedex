package main

import (
	"time"

	"github.com/priyanshoon/pokedex/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5 * time.Second)
	config := &config{
		commandRegistry: getCommands(),
		pokeapiClient:   pokeClient,
	}
	startRepl(config)
}
