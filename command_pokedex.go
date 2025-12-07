package main

import (
	"fmt"
)

func commandPokedex(cfg *config, words []string) error {
	fmt.Println("Your Pokedex:")
	for _, pokemon := range cfg.pokeDex {
		fmt.Printf(" - %v\n", pokemon.Name)
	}
	return nil
}
