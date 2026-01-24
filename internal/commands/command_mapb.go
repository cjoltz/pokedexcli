package commands

import (
	"fmt"

	"github.com/cjoltz/pokedexcli/internal/api"
)

func commandMapBack(c *Config) error {
	// 1. Make Request
	mapObjects, err := api.GetPreviousMaps(c.Previous)
	if err != nil {
		return err
	}

	// 2. Update config object passed into function
	c.Previous = mapObjects.Previous
	c.Next = mapObjects.Next

	for _, res := range mapObjects.Results {
		fmt.Println(res.Name)
	}

	return nil
}
