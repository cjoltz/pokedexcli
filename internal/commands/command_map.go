package commands

import (
	"fmt"
)

func commandMap(cfg *Config) error {
	locationsResp, err := cfg.PokeapiClient.ListLocations(cfg.nextLocationsURL)
	if err != nil {
		return err
	}

	cfg.nextLocationsURL = locationsResp.Next
	cfg.prevLocationsURL = locationsResp.Previous

	for _, loc :=  range locationsResp.Results {
		fmt.Println(loc.Name)
	}
	return nil

}
