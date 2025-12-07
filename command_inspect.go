package main

import (
	"fmt"
)

func commandInspect(cfg *config, words []string) error {
	if len(words) < 2 {
		return fmt.Errorf("usage: catch <pokemon-name>")
	}
	name := words[1]
	pokemon, ok := cfg.pokeDex[name]
	if !ok {
		return fmt.Errorf("you have not caught that pokemon")
	}
	fmt.Println("Name:", pokemon.Name)
	fmt.Println("Height:", pokemon.Height)
	fmt.Println("Weight:", pokemon.Weight)
	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf("  -%v: %v\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, pokeType := range pokemon.Types {
		fmt.Printf("  - %v\n", pokeType.Type.Name)
	}
	return nil
}
