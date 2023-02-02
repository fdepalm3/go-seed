package rest

import "github.com/redbeestudios/go-seed/internal/application/model/pokemon"

type typeDescription struct {
	Name string `json:"name"`
}

type responseType struct {
	Type typeDescription `json:"type"`
}

type pokemonResponse struct {
	Id    int            `json:"id"`
	Name  string         `json:"name"`
	Types []responseType `json:"types"`
}

func (p *pokemonResponse) ToDomain() (*pokemon.Pokemon, error) {

	primaryType, err := pokemon.NewPokemonType(p.Types[0].Type.Name)
	if err != nil {
		return nil, err
	}

	var secondaryType *pokemon.Type
	if len(p.Types) > 1 {
		secondaryType, err = pokemon.NewPokemonType(p.Types[1].Type.Name)
		if err != nil {
			return nil, err
		}
	} else {
		secondaryType, _ = pokemon.NewPokemonType("")
	}

	return pokemon.NewPokemon(
		p.Id,
		p.Name,
		*primaryType,
		secondaryType,
	), nil
}
