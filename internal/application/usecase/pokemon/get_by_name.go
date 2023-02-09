package pokemon

import (
	"context"

	"github.com/redbeestudios/go-seed/internal/application/model/pokemon"
	"github.com/redbeestudios/go-seed/internal/application/port/in"
	"github.com/redbeestudios/go-seed/internal/application/port/out"
)

var _ in.GetByName = (*GetByName)(nil)

type GetByName struct {
	pokemonRepository out.PokemonRepository
}

func (c *GetByName) Attack(ctx context.Context, name string) (*pokemon.Move, error) {
	return c.pokemonRepository.GetMoveByName(ctx, name)
}

func (c *GetByName) DamageRelations(ctx context.Context, name string, attackName string) (string, error) {

	pokemon, err := c.pokemonRepository.GetByName(ctx, name)
	if err != nil {
		return "", err
	}

	attack, err := c.pokemonRepository.GetMoveByName(ctx, attackName)
	if err != nil {
		return "", err
	}

	var multiplicador string = "x1"
	for _, v := range pokemon.PrimaryType().Weaknesses {
		if v == attack.MoveType() {
			multiplicador = "x2"
		}
	}
	return multiplicador, nil
}

func NewGetByName(pokemonRepository out.PokemonRepository) *GetByName {
	return &GetByName{
		pokemonRepository: pokemonRepository,
	}
}

func (c *GetByName) Get(ctx context.Context, name string) (*pokemon.Pokemon, error) {
	return c.pokemonRepository.GetByName(ctx, name)
}

// Si los dos tipos de un pkmn son debiles al del ataque: x4
// Si el pkmn tiene un solo tipo y  es debil o tiene dos y uno es debil y el otro no: x2
// Si ninguno de los dos es debil o uno es debil y el otro resistente: x1
// Si uno solo es resistente: x1/2
// Si los dos son resistentes: x1/4
// Si uno es inmune, no importa el otro
