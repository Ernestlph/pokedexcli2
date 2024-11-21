package main

import (
	"time"

	"github.com/Ernestlph/pokedexcli2/internal/pokeapi"
)

type config struct {
	pokeapiClient       pokeapi.Client
	nextLocationAreaURL *string
	prevLocationAreaURL *string
	caughtPokemon       map[string]pokeapi.Pokemon
}

func main() {
	cfg := config{
		pokeapiClient: pokeapi.NewClient(time.Minute * 5),
		caughtPokemon: make(map[string]pokeapi.Pokemon),
	}

	startRepl(&cfg)

}
