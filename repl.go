package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/jbeardwo/pokedex_go/internal/pokeapi"
)

type config struct {
	pokeapiClient pokeapi.Client
	Next          *string
	Previous      *string
	pokeDex       map[string]pokeapi.Pokemon
}

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		words := cleanInput(scanner.Text())

		if len(words) == 0 {
			continue
		}

		commandName := words[0]
		command, ok := getCommands()[commandName]
		if !ok {
			fmt.Println("Unkown Command")
		} else {
			err := command.callback(cfg, words)
			if err != nil {
				fmt.Println(err)
				fmt.Println()
			}
		}
	}
}

func cleanInput(text string) []string {
	text = strings.ToLower(text)
	out := strings.Fields(text)
	return out
}

type cliCommand struct {
	name        string
	description string
	callback    func(cfg *config, words []string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "show next page of location-areas",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "show previous page of location-areas",
			callback:    commandMapB,
		},
		"explore": {
			name:        "explore",
			description: "show pokemon in specified area",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "attempt to catch specified pokemon",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "inspect a previously caught specified pokemon's stats",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "list caught pokemon names",
			callback:    commandPokedex,
		},
	}
}
