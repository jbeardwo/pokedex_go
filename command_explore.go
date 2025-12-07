package main

import "fmt"

func commandExplore(cfg *config, words []string) error {
	if len(words) < 2 {
		return fmt.Errorf("usage: explore <area-name>")
	}
	url := "https://pokeapi.co/api/v2/location-area/"
	area := words[1]
	url += area

	fmt.Printf("Exploring %v...", area)
	ExploreAreaResponse, err := cfg.pokeapiClient.GetExploreArea(url)
	if err != nil {
		return fmt.Errorf("error: %v", err)
	}
	fmt.Println("Found Pokemon:")
	for _, encounter := range ExploreAreaResponse.PokemonEncounters {
		fmt.Println("-", encounter.Pokemon.Name)
	}
	fmt.Println()

	return nil
}
