package main

import (
	"time"

	"github.com/priyanshoon/pokedex/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5*time.Second, time.Minute*5)
	config := &config{
		commandRegistry: getCommands(),
		pokemons:        make(map[string]pokeapi.Pokemon),
		pokeapiClient:   pokeClient,
	}
	startRepl(config)
}
