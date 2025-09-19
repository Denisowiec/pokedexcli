package main

import (
	"encoding/json"
	"fmt"

	"github.com/Denisowiec/pokedexcli/internal/pokecache"
)

type locAreas struct {
	Count    int
	Next     string
	Previous string
	Results  []struct {
		Name string
		Url  string
	}
}

func getAreas(url string, cache *pokecache.Cache) (locAreas, error) {
	body, err := getData(url, cache)
	if err != nil {
		return locAreas{}, err
	}

	returnedAreas := locAreas{}
	if err := json.Unmarshal(body, &returnedAreas); err != nil {
		return locAreas{}, fmt.Errorf("error processing the data")
	}
	return returnedAreas, nil
}

func commandMap(cfg *config, args []string) error {
	var url string
	if cfg.next != "" {
		url = cfg.next
	} else {
		url = "https://pokeapi.co/api/v2/location-area/"
	}

	returnedAreas, err := getAreas(url, &cfg.cache)
	if err != nil {
		return err
	}

	cfg.next = returnedAreas.Next
	cfg.previous = returnedAreas.Previous

	for _, loc := range returnedAreas.Results {
		fmt.Println(loc.Name)
	}
	return nil
}

func commandMapb(cfg *config, args []string) error {
	if cfg.previous == "" {
		fmt.Println("you're on the first page")
		return nil
	}
	url := cfg.previous

	returnedAreas, err := getAreas(url, &cfg.cache)

	if err != nil {
		return err
	}

	cfg.next = returnedAreas.Next
	cfg.previous = returnedAreas.Previous

	for _, loc := range returnedAreas.Results {
		fmt.Println(loc.Name)
	}
	return nil
}
