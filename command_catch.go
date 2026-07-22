package main

import (
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return fmt.Errorf("you must provide a pokemon name to attempt to catch one")
	}

	pokemonName := args[0]
	pokemon, err := cfg.pokeapiClient.GetPokemonInfo(pokemonName)
	if err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonName)
	threshold := 50
	roll := rand.Intn(pokemon.BaseExperience)

	if roll < threshold {
		fmt.Printf("%s was caught!\n", pokemonName)
		fmt.Printf("You may now inspect your %s with the inspect command!\n", pokemonName)
		cfg.pokedex[pokemonName] = pokemon
	} else {
		fmt.Printf("%s has escaped!\n", pokemonName)
	}

	return nil
}
