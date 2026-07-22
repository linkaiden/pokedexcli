package main

import (
	"github.com/linkaiden/pokedexcli/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(pokeapi.DefaultClientTimeout, pokeapi.DefaultCacheReapInterval)
	cfg := &config{
		pokeapiClient: pokeClient,
		pokedex:       make(map[string]pokeapi.Pokemon),
	}

	startREPL(cfg)
}
