package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startPokedex() {
	fmt.Println("Hello Trainer")
	scanner := bufio.NewScanner(os.Stdin)
	for true {
		fmt.Print("Pokedex > ")
		if scanner.Scan() {
			userInput := cleanInput(scanner.Text())
			fmt.Printf("Your command was: %s\n", userInput[0])
		}
	}
}

func cleanInput(text string) []string {
	sep := " "
	// fmt.Println("Hello")
	s := strings.ToLower(strings.TrimSpace(text))
	var result []string
	i := strings.Index(s, sep)
	for i > -1 {
		result = append(result, s[:i])
		nextStart := i + len(sep)
		// Empty elements not allowed
		for nextStart < len(s) {
			if string(s[nextStart]) == sep {
				nextStart++
			} else {
				break
			}
		}
		s = s[nextStart:]         // remove extracted part from s
		i = strings.Index(s, sep) // nextEnd
	}
	return append(result, s)
}
