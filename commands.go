package main

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

const threshold = 40

func commandPokedex(cfg *config, args ...string) error {
	if len(cfg.pokemons) == 0 {
		return errors.New("You don't have any pokemons yet.")
	}

	fmt.Println("Your Pokedex: ")
	for pokemon := range cfg.pokemons {
		fmt.Printf(" - %s\n", pokemon)
	}
	return nil
}

func commandInspect(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a pokemon name.")
	}

	name := args[0]
	pokemon, ok := cfg.pokemons[name]
	if !ok {
		fmt.Println("you have not caught that pokemon")
		return nil
	}

	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)
	fmt.Println("Stats:")
	for _, data := range pokemon.Stats {
		fmt.Printf(" -%s: %d\n", data.Stat.Name, data.BaseStat)
	}
	fmt.Println("Types:")
	for _, data := range pokemon.Types {
		fmt.Printf(" - %s\n", data.Type.Name)
	}

	return nil
}

func commandCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a pokemon name.")
	}

	name := args[0]
	fmt.Printf("Throwing a Pokeball at %s...\n", name)

	pokemon, err := cfg.pokeapiClient.GetPokemon(name)
	if err != nil {
		return err
	}

	res := rand.Intn(pokemon.BaseExperience)

	if res > threshold {
		fmt.Printf("%s escaped!\n", name)
		return nil
	}

	cfg.pokemons[name] = pokemon
	fmt.Printf("%s was caught!\n", name)
	return nil
}

func commandExplore(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a location name.")
	}

	name := args[0]
	location, err := cfg.pokeapiClient.ListPokemons(name)
	if err != nil {
		return err
	}

	fmt.Printf("Exploring %s...\n", location.Name)
	fmt.Println("Found Pokemon: ")
	for _, enc := range location.PokemonEncounters {
		fmt.Printf(" - %s\n", enc.Pokemon.Name)
	}

	return nil
}

func commandMap(cfg *config, args ...string) error {
	locationResp, err := cfg.pokeapiClient.ListLocations(cfg.nextLocationsURL)
	if err != nil {
		return err
	}

	cfg.nextLocationsURL = locationResp.Next
	cfg.previousLocationsURL = locationResp.Previous

	for _, loc := range locationResp.Results {
		fmt.Println(loc.Name)
	}

	return nil
}

func commandMapb(cfg *config, args ...string) error {
	if cfg.previousLocationsURL == nil {
		return errors.New("you're on the first page")
	}
	locationResp, err := cfg.pokeapiClient.ListLocations(cfg.previousLocationsURL)
	if err != nil {
		return err
	}

	cfg.nextLocationsURL = locationResp.Next
	cfg.previousLocationsURL = locationResp.Previous

	for _, loc := range locationResp.Results {
		fmt.Println(loc.Name)
	}

	return nil
}

func commandHelp(cfg *config, args ...string) error {
	fmt.Printf("Welcome to the Pokedex!\n")
	fmt.Printf("Usage:\n\n")
	for _, value := range cfg.commandRegistry {
		fmt.Printf("%s: %s\n", value.name, value.description)
	}
	return nil
}

func commandExit(cfg *config, args ...string) error {
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
		"explore": {
			name:        "explore",
			description: "List all the pokemon located there.",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "catches pokemon",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "provides information about pokemons",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "list pokemons caught so far by users",
			callback:    commandPokedex,
		},
	}

	return commands
}
