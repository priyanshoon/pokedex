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
	callback    func() error
}

func startRepl() {
	reader := bufio.NewScanner(os.Stdin)
	commands := getCommands()

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

		use.callback()
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

func commandHelp() error {
	fmt.Printf("Welcome to the Pokedex!\n")
	fmt.Printf("Usage:\n\n")
	for key, value := range getCommands() {
		fmt.Printf("%s: %s\n", key, value.description)
	}
	return nil
}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func cleanInput(texts string) []string {
	newTexts := strings.ToLower(strings.TrimSpace(texts))
	words := strings.Fields(newTexts)
	return words
}
