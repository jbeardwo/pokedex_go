package main

import (
	"fmt"

	"github.com/jbeardwo/pokedex_go/pokeapi"
)

func commandMap(cfg *Config) error {
	url := "https://pokeapi.co/api/v2/location-area/"
	if cfg.Next != nil {
		url = *cfg.Next
	}
	locationAreaResponse, err := pokeapi.GetLocationAreas(url)
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
