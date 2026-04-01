package main

import (
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("No pokemon given")
	}
	pokemonResp, err := cfg.pokeapiClient.PokemonDetails(args[0])
	if err != nil {
		return err
	}
	fmt.Printf("Throwing a Pokeball at %s...\n", args[0])
	if rand.Intn(pokemonResp.BaseExperience / 10) == 1 {
		fmt.Printf("%s was caught!\n", args[0])
		cfg.pokeapiClient.Pokedex[pokemonResp.Name] = pokemonResp
		fmt.Println("You may now inspect it with the inspect command.")
		return nil
	}
	fmt.Printf("%s escaped!\n", args[0])
	return nil
}
