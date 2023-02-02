package in

import (
	"context"

	"github.com/redbeestudios/go-seed/internal/application/model/pokemon"
)

//go:generate mockgen -source=./save_pokemon.go -package=mocks -destination=../../../../mocks/save_pokemon.go

type SavePokemon interface {
	Save(ctx context.Context, pokemon *pokemon.Pokemon) error
}
