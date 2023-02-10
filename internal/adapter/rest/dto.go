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

type moveResponse struct {
	Id   int             `json:"id"`
	Name string          `json:"name"`
	Type typeDescription `json:"type"`
}

func (p *pokemonResponse) ToDomain() (*pokemon.Pokemon, error) {

	primaryType := pokemon.Type(p.Types[0].Type.Name)

	var secondaryType *pokemon.PokemonType
	if len(p.Types) > 1 {
		secondaryType = pokemon.Type(p.Types[1].Type.Name)
	} else {
		secondaryType = pokemon.Type("")
	}

	return pokemon.NewPokemon(
		p.Id,
		p.Name,
		primaryType,
		secondaryType,
	), nil
}

func (p *moveResponse) ToDomain() (*pokemon.Move, error) {

	primaryType := pokemon.Type(p.Type.Name)

	return pokemon.NewMove(
		p.Id,
		p.Name,
		primaryType,
	), nil
}
