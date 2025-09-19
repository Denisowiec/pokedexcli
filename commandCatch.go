package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
)

type Pokemon struct {
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
	Height int    `json:"height"`
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Stats  []struct {
		BaseStat int `json:"base_stat"`
		Effort   int `json:"effort"`
		Stat     struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"stat"`
	} `json:"stats"`
	Types []struct {
		Slot int `json:"slot"`
		Type struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"type"`
	} `json:"types"`
	Weight int `json:"weight"`
	// This one is added for the sake of this program and not part of PokeApi:
	numberCaught int
}

func commandCatch(cfg *config, args []string) error {
	pokeName := args[0]
	fmt.Printf("Throwing a Pokeball at %v...\n", pokeName)

	url := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s/", pokeName)

	body, err := getData(url, &cfg.cache)
	if err != nil {
		return fmt.Errorf("pokemon not found: %v", err)
	}

	var pokemon Pokemon

	if err = json.Unmarshal(body, &pokemon); err != nil {
		return err
	}

	// We calculate the catch chance base on the inverse of experience given
	catchChance := rand.Intn(300)

	if catchChance > pokemon.BaseExperience {
		fmt.Printf("%s was caught!\n", pokeName)
		item, ok := cfg.pokedex[pokeName]
		if !ok {
			pokemon.numberCaught = 1
			cfg.pokedex[pokeName] = pokemon
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
