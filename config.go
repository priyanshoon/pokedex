package main

import "github.com/priyanshoon/pokedex/internal/pokeapi"

type config struct {
	commandRegistry      map[string]cliCommand
	pokeapiClient        pokeapi.Client
	locationName         string
	nextLocationsURL     *string
	previousLocationsURL *string
}
