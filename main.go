package main

func main() {
	url := "https://pokeapi.co/api/v2/location-area"
	config := &config{
		commandRegistry: getCommands(),
		next:            &url,
		previous:        nil,
	}
	startRepl(config)
}
