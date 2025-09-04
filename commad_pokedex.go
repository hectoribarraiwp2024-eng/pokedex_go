package main

import (
	"fmt"
	"errors"
)

func commandPokedex(cfg *config, args ...string) error {
	if len(cfg.caughtPokemon) < 1 {
		return errors.New("you have caught no pokemon")
	}

	for key := range cfg.caughtPokemon {
		fmt.Printf(" - %s\n", key)
	}

	return nil
}