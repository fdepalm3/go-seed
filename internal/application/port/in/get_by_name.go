package in

import (
	"context"

	"github.com/redbeestudios/go-seed/internal/application/model/pokemon"
)

//go:generate mockgen -source=./get_by_name.go -package=mocks -destination=../../../../mocks/get_pokemon_by_name.go

type GetByName interface {
	Get(ctx context.Context, name string) (*pokemon.Pokemon, error)
	Attack(ctx context.Context, name string) (*pokemon.Move, error)
	DamageRelations(ctx context.Context, name string, attack string) (string, error)
}
