package main

import (
	"fmt"
	"os"
)

func callbackExit(config *config, args ...string) error {
	fmt.Println("Exiting the Pokedex...")
	os.Exit(0)
	return nil
}
