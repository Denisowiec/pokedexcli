package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
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

func getAreas(url string) (locAreas, error) {
	res, err := http.Get(url)
	if err != nil {
		return locAreas{}, fmt.Errorf("error downloading the data")
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return locAreas{}, fmt.Errorf("error fetching the data: %v", res.Status)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return locAreas{}, fmt.Errorf("error reading the data")
	}

	returnedAreas := locAreas{}
	if err := json.Unmarshal(body, &returnedAreas); err != nil {
		return locAreas{}, fmt.Errorf("error processing the data")
	}
	return returnedAreas, nil
}

func commandMap(cfg *config) error {
	var url string
	if cfg.next != "" {
		url = cfg.next
	} else {
		url = "https://pokeapi.co/api/v2/location-area/"
	}

	returnedAreas, err := getAreas(url)

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

func commandMapb(cfg *config) error {
	if cfg.previous == "" {
		fmt.Println("you're on the first page")
		return nil
	}
	url := cfg.previous

	returnedAreas, err := getAreas(url)

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
