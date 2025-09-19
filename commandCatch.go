package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
)

type returnedPokemon struct {
	Abilities []struct {
		Ability struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"ability"`
		IsHidden bool `json:"is_hidden"`
		Slot     int  `json:"slot"`
	} `json:"abilities"`
	BaseExperience int `json:"base_experience"`
	Cries          struct {
		Latest string `json:"latest"`
		Legacy string `json:"legacy"`
	} `json:"cries"`
	Forms []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"forms"`
	ID           int `json:"id"`
	NumberCaught int
}

func commandCatch(cfg *config, args []string) error {
	pokeName := args[0]
	fmt.Printf("Throwing a Pokeball at %v\n", pokeName)

	url := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s/", pokeName)

	body, err := getData(url, &cfg.cache)
	if err != nil {
		return fmt.Errorf("pokemon not found: %v", err)
	}

	var pokemon returnedPokemon

	if err = json.Unmarshal(body, &pokemon); err != nil {
		return err
	}

	// We calculate the catch chance base on the inverse of experience given
	invExperience := 300 - pokemon.BaseExperience
	catchChance := rand.Intn(300)

	if catchChance > invExperience {
		fmt.Printf("%s was caught!\n", pokeName)
		item, ok := cfg.pokedex[pokeName]
		if !ok {
			cfg.pokedex[pokeName] = Pokemon{
				name:         pokeName,
				numberCaught: 1,
			}
			item = cfg.pokedex[pokeName]
		} else {
			item.numberCaught += 1
		}
		fmt.Printf("Number of %s Pokemon in your pokedex: %v\n", pokeName, item.numberCaught)
	} else {
		fmt.Printf("%s excaped!\n", pokeName)
	}

	return nil
}
