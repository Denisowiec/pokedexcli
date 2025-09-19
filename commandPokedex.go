package main

import "fmt"

func commandPokedex(cfg *config, args []string) error {
	fmt.Println("Your pokedex:")
	if len(cfg.pokedex) == 0 {
		fmt.Println("No pokemon caught, yet")
		return nil
	}
	for _, item := range cfg.pokedex {
		fmt.Printf(" - %v, number: %v\n", item.Name, item.numberCaught)
	}
	return nil
}
