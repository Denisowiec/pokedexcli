package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

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
	var body []byte

	// We check if the url requested is present in the cache
	val, ok := cache.Get(url)
	if ok {
		body = val
	} else {
		// Not in cache, so we proceed to download it off the Internet, and save the result in cache
		res, err := http.Get(url)
		if err != nil {
			return locAreas{}, fmt.Errorf("error downloading the data: %v", err)
		}
		defer res.Body.Close()

		if res.StatusCode != http.StatusOK {
			return locAreas{}, fmt.Errorf("error fetching the data: %v", res.Status)
		}

		body, err = io.ReadAll(res.Body)
		if err != nil {
			return locAreas{}, fmt.Errorf("error reading the data: %v", err)
		}

		cache.Add(url, body)
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
