package main

import (
	"fmt"
	"os"
)

func commandExit(cfg *config, words []string) error {
	_ = words
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
