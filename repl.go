package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Config struct {
	Next     *string
	Previous *string
}

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	cfg := &Config{}

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
			err := command.callback(cfg)
			if err != nil {
				fmt.Println(err)
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
	callback    func(cfg *Config) error
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
	}
}
