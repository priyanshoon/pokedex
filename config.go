package main

type config struct {
	commandRegistry map[string]cliCommand
	next            *string
	previous        *string
}
