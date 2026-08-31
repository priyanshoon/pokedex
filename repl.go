package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

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

		args := []string{}
		if len(words) > 1 {
			args = words[1:]
		}

		err := use.callback(config, args...)
		if err != nil {
			fmt.Println("err: ", err)
		}
	}

}

func cleanInput(texts string) []string {
	newTexts := strings.ToLower(strings.TrimSpace(texts))
	words := strings.Fields(newTexts)
	return words
}
