package main

import "fmt"

func commandInspect(cfg *config, args ...string) error {
	if len(args) != 1 {
		return fmt.Errorf("you must provide a pokemon's name to inspect them!")
	}

	pokemonName := args[0]
	pokeInfo, ok := cfg.pokedex[pokemonName]

	if ok {
		fmt.Printf("Name: %s\n", pokeInfo.Name)
		fmt.Printf("Height: %d\n", pokeInfo.Height)
		fmt.Printf("Weight: %d\n", pokeInfo.Weight)
		fmt.Println("Stats: ")
		for _, stat := range pokeInfo.Stats {
			fmt.Printf("  -%s: %d\n", stat.Stat.Name, stat.BaseStat)
		}
		fmt.Println("Types: ")
		for _, typeInfo := range pokeInfo.Types {
			fmt.Printf("  - %s\n", typeInfo.Type.Name)
		}
	}
	if !ok {
		fmt.Printf("Unable to inspect a wild pokemon! You haven't caught a %s yet!\n", pokemonName)
	}

	return nil
}
