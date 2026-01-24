package commands

import (
	"fmt"
	"github.com/cjoltz/pokedexcli/internal/api"
)

type Config struct {
	Next     string
	Previous string
}

func commandMap(c *Config) error {
	// 1. Make Request
	mapObjects, err := api.GetNextMaps(c.Next)
	if err != nil {
		return err
	}

	// 2. Update config object passed into function
	c.Previous = mapObjects.Previous
	c.Next = mapObjects.Next

	// 3. Display Names 
	for _, res := range mapObjects.Results {
		fmt.Println(res.Name)
	}
	return nil
}
