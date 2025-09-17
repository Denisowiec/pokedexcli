package main

import (
	"bufio"
	"fmt"
	"os"
)

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	cfg := config{}

	for {
		fmt.Printf("Pokedex > ")
		scanner.Scan()
		cleanedText := cleanInput(scanner.Text())
		commands := listCommands()

		// Check if command exists in the registry
		cmdWord := cleanedText[0]
		cmd, ok := commands[cmdWord]

		if !ok {
			fmt.Println("Unknown command")
			continue
		}
		if err := cmd.callback(&cfg); err != nil {
			fmt.Printf("Error executing command '%v': %v", cmdWord, err)
		}
		//fmt.Printf("Your command was: %v\n", cleanedText[0])
	}
}
