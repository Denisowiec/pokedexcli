package main

import (
	"strings"

	"github.com/Denisowiec/pokedexcli/internal/pokecache"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config, []string) error
}

type config struct {
	previous string
	next     string
	cache    pokecache.Cache
	pokedex  map[string]Pokemon
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
		"explore": {
			name:        "explore",
			description: "List all Pokemon that can be found in a given location area",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Attempt to catch a pokemon and add it to your pokedex",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "Inspect a pokemon in your pokedex",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "List all pokemon in your pokedex",
			callback:    commandPokedex,
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
