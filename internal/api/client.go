package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type pokeMap struct {
	Count    int
	Next     string
	Previous string
	Results  []struct {
		Name string
		URL  string
	}
}

const BASEURL string = "https://pokeapi.co/api/v2/location-area"

func GetNextMaps(next string) (pokeMap, error) {
	url := next
	if next == "" {
		url = BASEURL
	}
	return getMaps("", url, true)
}

func GetPreviousMaps(previous string) (pokeMap, error) {
	if previous == "null" || previous == "" {
		return pokeMap{}, fmt.Errorf("you're on the first page")
	}
	return getMaps(previous, "", false)
}

func getMaps(previous, next string, getNext bool) (pokeMap, error) {
	maps := pokeMap{}
	// Next or previous
	url := previous
	if getNext {
		url = next
	}
	// Create Request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return maps, err
	}
	// Do Request and Get Response
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return maps, err
	}
	defer res.Body.Close()

	// Decode
	if err := json.NewDecoder(res.Body).Decode(&maps); err != nil {
		return maps, err
	}

	return maps, nil
}
