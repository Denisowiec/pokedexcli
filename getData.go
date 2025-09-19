package main

import (
	"fmt"
	"io"
	"net/http"

	"github.com/Denisowiec/pokedexcli/internal/pokecache"
)

func getData(url string, cache *pokecache.Cache) ([]byte, error) {
	var body []byte

	// We check if the url requested is present in the cache
	val, ok := cache.Get(url)
	if ok {
		body = val
	} else {
		// Not in cache, so we proceed to download it off the Internet, and save the result in cache
		res, err := http.Get(url)
		if err != nil {
			return []byte{}, fmt.Errorf("error downloading the data: %v", err)
		}
		defer res.Body.Close()

		if res.StatusCode != http.StatusOK {
			return []byte{}, fmt.Errorf("error fetching the data: %v", res.Status)
		}

		body, err = io.ReadAll(res.Body)
		if err != nil {
			return []byte{}, fmt.Errorf("error reading the data: %v", err)
		}

		cache.Add(url, body)
	}
	return body, nil
}
