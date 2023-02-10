package redis

import (
	"github.com/redbeestudios/go-seed/internal/application/model/pokemon"
	"github.com/redbeestudios/go-seed/pkg"
)

type pokemonDTO struct {
	Id            int     `json:"id"`
	Name          string  `json:"name"`
	PrimaryType   string  `json:"primary_type"`
	SecondaryType *string `json:"secondary_type"`
}

func (p *pokemonDTO) ToDomain() (*pokemon.Pokemon, error) {
	primaryType := pokemon.Type(p.PrimaryType)

	var secondaryType *pokemon.PokemonType
	if p.SecondaryType != nil {
		secondaryType = pokemon.Type(*p.SecondaryType)
	}

	return pokemon.NewPokemon(
		p.Id,
		p.Name,
		primaryType,
		secondaryType,
	), nil
}

func fromDomain(p *pokemon.Pokemon) *pokemonDTO {

	var secondaryType *string
	if p.SecondaryType() != nil {
		secondaryType = pkg.ToPointer(p.SecondaryType().Name)
	}

	return &pokemonDTO{
		Id:            p.Id(),
		Name:          p.Name(),
		PrimaryType:   p.PrimaryType().Name,
		SecondaryType: secondaryType,
	}
}
