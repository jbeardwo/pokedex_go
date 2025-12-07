package main

import (
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, words []string) error {
	if len(words) < 2 {
		return fmt.Errorf("usage: catch <pokemon-name>")
	}
	url := "https://pokeapi.co/api/v2/pokemon/"
	name := words[1]
	url += name

	pokemon, err := cfg.pokeapiClient.GetPokemon(url)
	if err != nil {
		return fmt.Errorf("error: %v", err)
	}
	fmt.Printf("Throwing a Pokeball at %v...\n", pokemon.Name)
	catchProbability := getCatchProbability(float64(pokemon.BaseExperience))
	catchRoll := rand.Float64()

	if catchRoll < catchProbability {
		fmt.Printf("%v was caught!\n", pokemon.Name)
		cfg.pokeDex[pokemon.Name] = *pokemon
	} else {
		fmt.Printf("%v escaped!\n", pokemon.Name)
	}

	return nil
}

func getCatchProbability(baseExp float64) float64 {
	minExp, maxExp := 36.0, 395.0
	minProb, maxProb := .1, .9
	normalized := (baseExp - minExp) / (maxExp - minExp)
	return maxProb - (maxProb-minProb)*normalized
}
