package main

import (
	"fmt"

	"github.com/jbeardwo/pokedex_go/internal/pokeapi"
)

func commandMapB(cfg *Config) error {
	if cfg.Previous == nil {
		fmt.Println("you're on the first page")
		fmt.Println()
		return nil
	}
	url := *cfg.Previous
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
