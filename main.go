package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	//"unicode"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func listCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
	}
}

func cleanInput(text string) []string {
	lowerCase := strings.ToLower(text)

	result := strings.Fields(lowerCase)

	return result
}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp() error {
	fmt.Printf(`Welcome to the Pokedex!
Usage:
`)
	for _, cmd := range listCommands() {
		fmt.Printf("%v: %v\n", cmd.name, cmd.description)
	}
	return nil
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

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
		if err := cmd.callback(); err != nil {
			fmt.Printf("Error executing command '%v': %v", cmdWord, err)
		}
		//fmt.Printf("Your command was: %v\n", cleanedText[0])
	}
}
