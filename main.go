package main

import (
	"time"

	"github.com/cjoltz/pokedexcli/internal/commands"
	"github.com/cjoltz/pokedexcli/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5 * time.Second, time.Minute * 5)
	cfg := &commands.Config{
		PokeapiClient: pokeClient,
	}
	startPokedex(cfg)
}
