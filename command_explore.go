package main

import "fmt"

func commandExplore(cfg *config, args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("No location given")
	}
	exploreResp, err := cfg.pokeapiClient.LocationDetails(args[0])
	if err != nil {
		return err
	}
	fmt.Printf("Exploring %s...\n", exploreResp.Location.Name)
	fmt.Println("Found Pokemon:")
	for _, encounter := range exploreResp.PokemonEncounters {
		fmt.Println(encounter.Pokemons.Name)
	}
	return nil
}
