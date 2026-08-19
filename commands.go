package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/priyanshoon/pokedex/internal/pokeapi"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

func commandMapb(cfg *config) error {
	if cfg.previous == nil {
		return errors.New("it's over for you bro, there is no previoui page")
	}

	url := cfg.previous

	locations, err := pokeapi.GetLocationAreas(*url)
	if err != nil {
		return err
	}

	cfg.previous = locations.Previous
	cfg.next = locations.Next

	for _, location := range locations.Results {
		fmt.Println(location.Name)
	}

	return nil
}

func commandMap(cfg *config) error {
	url := cfg.next

	locations, err := pokeapi.GetLocationAreas(*url)
	if err != nil {
		return err
	}

	cfg.previous = locations.Previous
	cfg.next = locations.Next

	for _, location := range locations.Results {
		fmt.Println(location.Name)
	}

	return nil
}

func commandHelp(cfg *config) error {
	fmt.Printf("Welcome to the Pokedex!\n")
	fmt.Printf("Usage:\n\n")
	for _, value := range cfg.commandRegistry {
		fmt.Printf("%s: %s\n", value.name, value.description)
	}
	return nil
}

func commandExit(cfg *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func getCommands() map[string]cliCommand {
	commands := map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Gives location areas in the pokemon world.",
			callback:    commandMap,
		},
		"mapb": {
			name:        "map",
			description: "Gives locations areas in the pokemon worlds.",
			callback:    commandMapb,
		},
	}

	return commands
}
