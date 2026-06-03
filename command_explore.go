package main

import (
	"fmt"
)

// List Pokemon
func commandExplore(cfg *config, name string) error {
	pokemonResp, err := cfg.pokeapiClient.PokemonLocations(name)
	if err != nil {
		return err
	}

	for _, encounter := range pokemonResp.PokemonEncounters {
		fmt.Println(encounter.Pokemon.Name)
	}
	return nil
}
