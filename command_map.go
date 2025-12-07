package main

import (
	"fmt"
)

func commandMap(cfg *config, words []string) error {
	_ = words
	url := "https://pokeapi.co/api/v2/location-area/"
	if cfg.Next != nil {
		url = *cfg.Next
	}
	locationAreaResponse, err := cfg.pokeapiClient.GetLocationAreas(url)
	if err != nil {
		return fmt.Errorf("error: %v", err)
	}
	for _, locationArea := range locationAreaResponse.Results {
		fmt.Println(locationArea.Name)
	}
	fmt.Println()

	cfg.Next = locationAreaResponse.Next
	cfg.Previous = locationAreaResponse.Previous

	return nil
}
