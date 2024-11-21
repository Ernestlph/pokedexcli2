package main

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
)

func callbackCatch(cfg *config, args ...string) error {

	if len(args) == 0 {
		return errors.New("no pokemon name provided")
	}

	pokemonname := args[0]

	resp, err := cfg.pokeapiClient.GetPokemon(pokemonname)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Throwing a pokeball at %s...", pokemonname)
	fmt.Println("")

	// Create random number from 0 to 450
	// If random number is less than or equal to base experience, pokemon is caught
	// If random number is greater than base experience, pokemon escapes
	randomNumber := rand.Intn(450)

	if resp.BaseExperience <= randomNumber {
		fmt.Printf("%s was caught!\n", pokemonname)
		cfg.caughtPokemon[pokemonname] = resp
		fmt.Println("You may now inspect it with the inspect command.")
	} else {
		fmt.Printf("%s escaped!\n", pokemonname)
	}

	return nil
}
