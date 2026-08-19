package main

func main() {
	config := &config{
		commandRegistry: map[string]cliCommand{
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
		},
	}
	startRepl(config)
}
