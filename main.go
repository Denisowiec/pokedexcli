package main

import (
	"fmt"
	"strings"
	//"unicode"
)

func cleanInput(text string) []string {
	lowerCase := strings.ToLower(text)
	var result []string

	result = strings.Fields(lowerCase)

	return result
}

func main() {
	example := "Hello, World!"

	exampleOutput := cleanInput(example)

	for _, val := range exampleOutput {
		fmt.Println(val)
	}

	//fmt.Println("Hello, World!")
}
