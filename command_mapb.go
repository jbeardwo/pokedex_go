package main

import (
	"errors"
	"fmt"
)

func commandMapB(cfg *config, words []string) error {
	_ = words
	if cfg.Previous == nil {
		return errors.New("you're on the first page")
	}
	url := *cfg.Previous
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
