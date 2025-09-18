package main

import "fmt"

func commandHelp(cfg *config, args []string) error {
	fmt.Printf(`Welcome to the Pokedex!
Usage:
`)
	for _, cmd := range listCommands() {
		fmt.Printf("%v: %v\n", cmd.name, cmd.description)
	}
	return nil
}
