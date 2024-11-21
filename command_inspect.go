package main

import (
	"errors"
	"fmt"
)

func callbackInspect(cfg *config, args ...string) error {

	if len(args) == 0 {
		return errors.New("no pokemon name provided")
	}

	pokemonname := args[0]

	// Check if pokemon is already caught, if not caught print "you have not caught that pokemon"
	resp, ok := cfg.caughtPokemon[pokemonname]

	if !ok {
		return errors.New("you have not caught that pokemon")
	}
	// If pokemon is caught, print pokemon details from cfg.caughtPokemon

	fmt.Printf("Name: %s\n", resp.Name)
	fmt.Printf("Height: %d\n", resp.Height)
	fmt.Printf("Weight: %d\n", resp.Weight)
	fmt.Printf("Stats:\n")
	for _, stat := range resp.Stats {
		fmt.Printf("  -%s: %d\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Printf("Types:\n")
	for _, typ := range resp.Types {
		fmt.Printf("  -%s\n", typ.Type.Name)
	}

	return nil
}
