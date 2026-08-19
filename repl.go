package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

func startRepl(config *config) {
	reader := bufio.NewScanner(os.Stdin)
	commands := config.commandRegistry

	for {
		fmt.Printf("Pokedex > ")
		if !reader.Scan() {
			if err := reader.Err(); err != nil {
				fmt.Fprintln(os.Stderr, err)
			}
			return
		}

		words := cleanInput(reader.Text())
		if len(words) == 0 {
			continue
		}

		command := words[0]
		use, ok := commands[command]
		if !ok {
			fmt.Println("Unknown command")
			continue
		}

		err := use.callback(config)
		if err != nil {
			fmt.Println("err: ", err)
		}
	}

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
	}

	return commands
}

func commandHelp(config *config) error {
	fmt.Printf("Welcome to the Pokedex!\n")
	fmt.Printf("Usage:\n\n")
	for _, value := range config.commandRegistry {
		fmt.Printf("%s: %s\n", value.name, value.description)
	}
	return nil
}

func commandExit(config *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func cleanInput(texts string) []string {
	newTexts := strings.ToLower(strings.TrimSpace(texts))
	words := strings.Fields(newTexts)
	return words
}
