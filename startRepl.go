package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/Denisowiec/pokedexcli/internal/pokecache"
)

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	cfg := config{
		next:     "",
		previous: "",
		cache:    pokecache.NewCache(5 * time.Second),
		pokedex:  map[string]Pokemon{},
	}

	for {
		fmt.Printf("Pokedex > ")
		scanner.Scan()
		cleanedText := cleanInput(scanner.Text())
		commands := listCommands()

		// Check if command exists in the registry
		cmdWord := cleanedText[0]
		args := cleanedText[1:]
		cmd, ok := commands[cmdWord]

		if !ok {
			fmt.Println("Unknown command")
			continue
		}
		if err := cmd.callback(&cfg, args); err != nil {
			fmt.Printf("Error executing command '%v': %v", cmdWord, err)
		}
		//fmt.Printf("Your command was: %v\n", cleanedText[0])
	}
}
