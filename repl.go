package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/cjoltz/pokedexcli/internal/commands"
)

func startPokedex() {
	reader := bufio.NewScanner(os.Stdin)
	cfg := commands.Config{Next: "", Previous: ""}
	for {
		fmt.Print("Pokedex > ")
		reader.Scan()

		words := cleanInput(reader.Text())
		if len(words) == 0 {
			continue
		}

		commandName := words[0]

		command, exists := commands.GetCommands()[commandName]
		if exists {
			err := command.Callback(&cfg)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}

}

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}
