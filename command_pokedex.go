package main

import (
	"errors"
	"fmt"
)

func commandPokedex(cfg *config, args ...string) error {

	if len(cfg.caughtPokemon) < 1 {
		return errors.New("you have no pokemons")
	}

	fmt.Println("Your Pokedex:")
	for k := range cfg.caughtPokemon {
		fmt.Printf(" - %v\n", k)
	}
	return nil
}
