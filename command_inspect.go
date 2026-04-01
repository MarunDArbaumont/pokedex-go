package main

import (
	"fmt"
)

func commandInspect(cfg *config, args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("No pokemon given")
	}
	pokemon, exists := cfg.pokeapiClient.Pokedex[args[0]]
	if exists {
		fmt.Printf("Name: %v\n", pokemon.Name)
		fmt.Printf("Name: %v\n", pokemon.Height)
		fmt.Printf("Name: %v\n", pokemon.Weight)
		fmt.Printf("Stats:\n")
		for _, stat := range pokemon.Stats {
			fmt.Printf("- %v: %v\n", stat.Stat.Name, stat.BaseStat)
		}
		fmt.Printf("Types:\n")
		for _, pokemonType := range pokemon.Types {
			fmt.Printf("- %v\n", pokemonType.Type.Name)
		}
		return nil
	}
	fmt.Printf("you have not caught %v yet\n", args[0])
	return nil
}
