package main

import "fmt"

func commandInspect(cfg *config, args []string) error {
	pokeName := args[0]
	pokemon, ok := cfg.pokedex[pokeName]
	if !ok {
		fmt.Printf("Pokemon %s not yet caught!", pokeName)
	} else {
		fmt.Printf(`Name: %v
Number caught: %v
Height: %v
Weight: %v
Stats:
`, pokemon.Name, pokemon.numberCaught, pokemon.Height, pokemon.Weight)
		for _, stat := range pokemon.Stats {
			fmt.Printf(" - %v: %v\n", stat.Stat.Name, stat.BaseStat)
		}
		fmt.Println("Types:")
		for _, typ := range pokemon.Types {
			fmt.Printf(" - %v\n", typ.Type.Name)
		}
	}

	return nil
}
