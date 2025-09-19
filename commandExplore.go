package main

import (
	"encoding/json"
	"fmt"

	"github.com/Denisowiec/pokedexcli/internal/pokecache"
)

type listOfPokemonInArea struct {
	EncounterMethodRates []struct {
		EncounterMethod struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"encounter_method"`
		VersionDetails []struct {
			Rate    int `json:"rate"`
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"encounter_method_rates"`
	GameIndex int `json:"game_index"`
	ID        int `json:"id"`
	Location  struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"location"`
	Name  string `json:"name"`
	Names []struct {
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
		Name string `json:"name"`
	} `json:"names"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
		VersionDetails []struct {
			EncounterDetails []struct {
				Chance          int           `json:"chance"`
				ConditionValues []interface{} `json:"condition_values"`
				MaxLevel        int           `json:"max_level"`
				Method          struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				} `json:"method"`
				MinLevel int `json:"min_level"`
			} `json:"encounter_details"`
			MaxChance int `json:"max_chance"`
			Version   struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"pokemon_encounters"`
}

func getPokemonInArea(url string, cache *pokecache.Cache) (listOfPokemonInArea, error) {
	body, err := getData(url, cache)
	if err != nil {
		return listOfPokemonInArea{}, err
	}

	returnedPokemon := listOfPokemonInArea{}
	if err := json.Unmarshal(body, &returnedPokemon); err != nil {
		return listOfPokemonInArea{}, fmt.Errorf("error processing the data: %v", err)
	}
	return returnedPokemon, nil
}

func commandExplore(cfg *config, args []string) error {
	url := "https://pokeapi.co/api/v2/location-area/" + args[0]

	pokemonList, err := getPokemonInArea(url, &cfg.cache)
	if err != nil {
		return err
	}

	for _, entry := range pokemonList.PokemonEncounters {
		fmt.Println(entry.Pokemon.Name)
	}
	return nil
}
