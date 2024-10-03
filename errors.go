package client

import (
	"errors"
	"fmt"
)

var (
	ErrFetchingPokemon = errors.New("Failed to fetch Pokemon.")
)

type PokemonFetchErr struct {
	Message    string
	StatusCode int
}

func (e PokemonFetchErr) Error() string {
	return fmt.Sprintf("Failed to fetch pokemon: %s with status code: %d", e.Message, e.StatusCode)
}
