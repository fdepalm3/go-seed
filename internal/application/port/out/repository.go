package out

import (
	"context"

	"github.com/redbeestudios/go-seed/internal/application/model/pokemon"
)

//go:generate mockgen -source=./repository.go -package=mocks -destination=../../../../mocks/pokemon_repository.go

type PokemonRepository interface {
	GetByName(ctx context.Context, name string) (*pokemon.Pokemon, error)
	SavePokemon(ctx context.Context, pokemon *pokemon.Pokemon) error
	GetMoveByName(ctx context.Context, name string) (*pokemon.Move, error)
	SendPokemonMessage(ctx context.Context, message string) error
}
