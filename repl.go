package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"github.com/cjoltz/pokedexcli/commands"

)

func startPokedex() {
	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		reader.Scan()

		words := cleanInput(reader.Text())
		if len(words) == 0 {
			continue
		}

		commandName := words[0]

		command, exists := commands.getCommands()[commandName]
		if exists {
			err := command.callback()
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
