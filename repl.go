package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/linkaiden/pokedexcli/internal/pokeapi"
)

type config struct {
	pokeapiClient pokeapi.Client
	pokedex       map[string]pokeapi.Pokemon
	next          *string
	previous      *string
}

func startREPL(cfg *config) {
	reader := bufio.NewScanner(os.Stdin)
	commands := getCommands()

	for {
		if reader.Err() != nil {
			break
		}

		fmt.Print("Pokedex > ")
		reader.Scan()

		words := cleanInput(reader.Text())
		if len(words) == 0 {
			continue
		}

		cmdName := words[0]
		args := []string{}
		if len(words) > 1 {
			args = words[1:]
		}

		cmd, exists := commands[cmdName]

		if exists {
			err := cmd.callback(cfg, args...)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}
}

func cleanInput(text string) []string {
	lowered := strings.ToLower(text)
	words := strings.Fields(lowered)
	return words
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Get the next (or first if not listed yet) page of locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Display the previous page of locations",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore <location_name>",
			description: "Explore a location",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch <pokemon_name>",
			description: "Attempt to catch a pokemon",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect <pokemon_name>",
			description: "Inspect attributes of a caught pokemon",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "View a list of what you've caught",
			callback:    commandPokedex,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
}
