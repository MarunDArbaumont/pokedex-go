package main

import (
	"fmt"
)

func commandPokedex(cfg *config, args []string) error {
	pokedex := cfg.pokeapiClient.Pokedex
	if len(pokedex) > 0 {
		fmt.Println("Your Pokedex:")
		for pokemon, _ := range pokedex {
			fmt.Printf("- %v\n", pokemon)
		}
		return nil
	}
	fmt.Println("No pokemon in your pokedex")
	return nil
}
