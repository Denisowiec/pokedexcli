package main

import (
	"strings"
	//"unicode"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

type config struct {
	previous string
	next     string
}

func listCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "List all location areas in the world of Pokemon, in batches of 20",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "List the previous batch of 20 location areas in the worl of Pokemon",
			callback:    commandMapb,
		},
	}
}

func cleanInput(text string) []string {
	lowerCase := strings.ToLower(text)

	result := strings.Fields(lowerCase)

	return result
}

func main() {
	startRepl()
}
