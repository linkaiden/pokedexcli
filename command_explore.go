package main

import (
	"fmt"
)

func commandExplore(cfg *config, args ...string) error {
	if len(args) != 1 {
		return fmt.Errorf("you must provide an area name to explore")
	}

	areaName := args[0]
	locationResp, err := cfg.pokeapiClient.ListLocationInfo(areaName)
	if err != nil {
		return err
	}

	fmt.Printf("Exploring %s...", areaName)
	fmt.Println("Found Pokemon:")
	for _, location := range locationResp.PokemonEncounters {
		fmt.Printf("- %s\n", location.Pokemon.Name)
	}

	return nil
}
