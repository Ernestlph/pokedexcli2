package main

import (
	"errors"
	"fmt"
	"log"
)

func callbackExplore(cfg *config, args ...string) error {

	if len(args) == 0 {
		return errors.New("no location area name provided")
	}

	locationAreaName := args[0]

	resp, err := cfg.pokeapiClient.ListLocationAreaDetail(locationAreaName)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Exploring %s...", locationAreaName)
	fmt.Println("")
	fmt.Println("Pokemon found:")
	for _, area := range resp.PokemonEncounters {
		fmt.Printf(" - %s\n", area.Pokemon.Name)
	}

	return nil
}
