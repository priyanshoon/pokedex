package main

import "github.com/priyanshoon/pokedex/internal/pokeapi"

type config struct {
	commandRegistry      map[string]cliCommand
	pokemons             map[string]pokeapi.Pokemon
	pokeapiClient        pokeapi.Client
	nextLocationsURL     *string
	previousLocationsURL *string
}
